package main

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/jannikluhn/terminal/node/contracts/endpointcontract"
)

const sinkAddressHex = "0xFA33c8EF8b5c4f3003361c876a298D1DB61ccA4e"
const sourceAddressHex = "0x14b47DBb4b0F7bff4205F4D1f0c1BEEBCE47b33C"
const sinkRpcUrl = "ws://127.0.0.1:8546/"
const sourceRpcUrl = "ws://127.0.0.1:8556/"

type EndpointConnection struct {
	Client   *ethclient.Client
	Contract *endpointcontract.Endpointcontract
	ChainID  uint64
}

func NewEndpointConnection(ctx context.Context, rpcUrl string, endpointContractAddress common.Address) (EndpointConnection, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return EndpointConnection{}, errors.Wrapf(err, "failed to connect to Ethereum node at %s", rpcUrl)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return EndpointConnection{}, errors.Wrap(err, "failed to query chain id")
	}

	contract, err := endpointcontract.NewEndpointcontract(endpointContractAddress, client)
	if err != nil {
		return EndpointConnection{}, errors.Wrap(err, "faild to create endpoint contract instance")
	}

	return EndpointConnection{
		Client:   client,
		Contract: contract,
		ChainID:  chainID.Uint64(),
	}, nil
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

func listenForRequests(ctx context.Context, conn EndpointConnection) (chan Transfer, chan error) {
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
		sub, err := conn.Contract.EndpointcontractFilterer.WatchTransferRequested(watchOpts, transferRequestedChannel)
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

func processTransferRequests(requests chan Transfer) error {
	for {

	}
}

func run() error {
	ctx := context.Background()

	sinkConnection, err := NewEndpointConnection(ctx, sinkRpcUrl, common.HexToAddress(sinkAddressHex))
	if err != nil {
		return err
	}
	sourceConnection, err := NewEndpointConnection(ctx, sourceRpcUrl, common.HexToAddress(sourceAddressHex))
	if err != nil {
		return err
	}

	sinkRequestChannel, sinkErrorChannel := listenForRequests(ctx, sinkConnection)
	sourceRequestChannel, sourceErrorChannel := listenForRequests(ctx, sourceConnection)

	for {
		select {
		case err := <-sinkErrorChannel:
			return errors.Wrap(err, "error listening to sink transfer requests")
		case err := <-sourceErrorChannel:
			return errors.Wrap(err, "error listening to source transfer requests")
		case transfer := <-sinkRequestChannel:
			log.Printf("Received sink transfer request %+v", transfer)
		case transfer := <-sourceRequestChannel:
			log.Printf("Received source transfer request %+v", transfer)
		case <-time.After(5 * time.Second):
			log.Println("listening for events")
		}
	}
}
