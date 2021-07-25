package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/jannikluhn/terminal/node/contracts/endpointcontract"
	"github.com/jannikluhn/terminal/node/contracts/oracle"
)

const sinkAddressHex = "0xBe0B0f08A599F07699E98A9D001084e97b9a900A"
const sinkOracleAddressHex = "0xFA33c8EF8b5c4f3003361c876a298D1DB61ccA4e"
const sinkPrivateKeyHex = "b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773"

const sourceAddressHex = "0x4fd74C68a8c68610Bd40BdC550f195bD367Ff238"
const sourceOracleAddressHex = "0x14b47DBb4b0F7bff4205F4D1f0c1BEEBCE47b33C"
const sourcePrivateKeyHex = "b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3774"

const sinkRpcUrl = "ws://127.0.0.1:8546/"
const sourceRpcUrl = "ws://127.0.0.1:8556/"

type Connection struct {
	Client           *ethclient.Client
	ChainID          uint64
	EndpointContract *endpointcontract.Endpointcontract
	OracleContract   *oracle.Oracle
	PrivateKey       *ecdsa.PrivateKey
}

func NewConnection(ctx context.Context, rpcUrl string, endpointContractAddress common.Address, oracleContractAddress common.Address, privateKey *ecdsa.PrivateKey) (Connection, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return Connection{}, errors.Wrapf(err, "failed to connect to Ethereum node at %s", rpcUrl)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return Connection{}, errors.Wrap(err, "failed to query chain id")
	}

	endpointContract, err := endpointcontract.NewEndpointcontract(endpointContractAddress, client)
	if err != nil {
		return Connection{}, errors.Wrap(err, "faild to create endpoint contract instance")
	}
	oracleContract, err := oracle.NewOracle(oracleContractAddress, client)
	if err != nil {
		return Connection{}, errors.Wrap(err, "faild to create oracle contract instance")
	}

	return Connection{
		Client:           client,
		ChainID:          chainID.Uint64(),
		EndpointContract: endpointContract,
		OracleContract:   oracleContract,
		PrivateKey:       privateKey,
	}, nil
}

func (conn *Connection) TransactOpts() (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(conn.PrivateKey, new(big.Int).SetUint64(conn.ChainID))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create transactor")
	}
	return opts, nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Fatal error:", err)
	}
}

type Transfer struct {
	FromEndpoint    endpointcontract.Endpoint
	TransactionHash common.Hash
	BlockNumber     uint64
	LogIndex        uint64

	ToEndpoint endpointcontract.Endpoint
	Receiver   common.Address
	Amount     *big.Int
}

func (t *Transfer) TransferIdentifier() oracle.OracleTransferIdentifier {
	return oracle.OracleTransferIdentifier{
		ChainID:     t.FromEndpoint.ChainID,
		BlockNumber: t.BlockNumber,
		TxHash:      t.TransactionHash,
		LogIndex:    uint32(t.LogIndex),
	}
}

func NewTransferFromRequestEvent(chainID uint64, ev *endpointcontract.EndpointcontractTransferRequested) Transfer {
	return Transfer{
		FromEndpoint: endpointcontract.Endpoint{
			ChainID:         chainID,
			ContractAddress: ev.Raw.Address,
		},
		TransactionHash: ev.Raw.TxHash,
		BlockNumber:     ev.Raw.BlockNumber,
		LogIndex:        uint64(ev.Raw.Index),

		ToEndpoint: ev.Endpoint,
		Receiver:   ev.Receiver,
		Amount:     ev.Amount,
	}
}

func listenForRequests(ctx context.Context, conn Connection) (chan Transfer, chan error) {
	requestChannel := make(chan Transfer)
	errorChannel := make(chan error)

	go func() {
		fail := func(err error) {
			errorChannel <- err
			close(errorChannel)
			close(requestChannel)
		}

		watchOpts := &bind.WatchOpts{
			Context: ctx,
			Start:   nil,
		}
		transferRequestedChannel := make(chan *endpointcontract.EndpointcontractTransferRequested)
		sub, err := conn.EndpointContract.EndpointcontractFilterer.WatchTransferRequested(watchOpts, transferRequestedChannel)
		if err != nil {
			fail(errors.Wrap(err, "faild to subscribe to TransferRequested events"))
			return
		}

		for {
			select {
			case err := <-sub.Err():
				fail(errors.Wrap(err, "event subscription error"))
				return
			case ev := <-transferRequestedChannel:
				transfer := NewTransferFromRequestEvent(conn.ChainID, ev)
				requestChannel <- transfer
			case <-ctx.Done():
				close(errorChannel)
				close(requestChannel)
				return
			}
		}
	}()

	return requestChannel, errorChannel
}

func processTransfer(ctx context.Context, connections map[uint64]Connection, transfer Transfer) error {
	toConn, ok := connections[transfer.ToEndpoint.ChainID]
	if !ok {
		return errors.Errorf("received transfer request to unknown chain id %d: %+v", transfer.ToEndpoint.ChainID, transfer)
	}

	// check if request is already handled by someone else
	callOpts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	exists, err := toConn.OracleContract.RequestExists(callOpts, transfer.TransferIdentifier())
	if err != nil {
		return errors.Wrapf(err, "failed to check if transfer %+v already exists", transfer)
	}
	if exists {
		log.Printf("Transfer %+v is already handled", transfer)
		return nil
	}

	transactOpts, err := toConn.TransactOpts()
	if err != nil {
		return err
	}
	tx, err := toConn.OracleContract.StartRequest(transactOpts, transfer.TransferIdentifier(), transfer.Amount, transfer.Receiver)
	if err != nil {
		return errors.Wrapf(err, "error starting request for transfer %+v", transfer)
	}

	log.Printf("Sent StartRequest tx with hash %s for transfer %+v", tx.Hash(), transfer)

	return nil
}

func run() error {
	ctx := context.Background()

	sinkPrivateKey, err := crypto.HexToECDSA(sinkPrivateKeyHex)
	if err != nil {
		return errors.Wrapf(err, "got invalid private key")
	}
	sourcePrivateKey, err := crypto.HexToECDSA(sourcePrivateKeyHex)
	if err != nil {
		return errors.Wrapf(err, "got invalid private key")
	}

	sinkConnection, err := NewConnection(
		ctx,
		sinkRpcUrl,
		common.HexToAddress(sinkAddressHex),
		common.HexToAddress(sinkOracleAddressHex),
		sinkPrivateKey,
	)
	if err != nil {
		return err
	}
	sourceConnection, err := NewConnection(
		ctx,
		sourceRpcUrl,
		common.HexToAddress(sourceAddressHex),
		common.HexToAddress(sourceOracleAddressHex),
		sourcePrivateKey,
	)
	if err != nil {
		return err
	}
	if sinkConnection.ChainID == sourceConnection.ChainID {
		return errors.Errorf("got two endpoint connections with same chain id %d", sinkConnection.ChainID)
	}

	connections := make(map[uint64]Connection)
	connections[sinkConnection.ChainID] = sinkConnection
	connections[sourceConnection.ChainID] = sourceConnection

	cancelCtx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()
	sinkRequestChannel, sinkErrorChannel := listenForRequests(cancelCtx, sinkConnection)
	sourceRequestChannel, sourceErrorChannel := listenForRequests(cancelCtx, sourceConnection)

	for {
		select {
		case err := <-sinkErrorChannel:
			return errors.Wrap(err, "error listening to sink transfer requests")
		case err := <-sourceErrorChannel:
			return errors.Wrap(err, "error listening to source transfer requests")
		case transfer := <-sinkRequestChannel:
			err := processTransfer(cancelCtx, connections, transfer)
			if err != nil {
				log.Printf("failed to process transfer: %s", err)
			}
		case transfer := <-sourceRequestChannel:
			err := processTransfer(cancelCtx, connections, transfer)
			if err != nil {
				log.Printf("failed to process transfer: %s", err)
			}
		case <-time.After(5 * time.Second):
			log.Println("listening for events")
		}
	}
}
