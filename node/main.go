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

	endpointcontract "github.com/jannikluhn/terminal/node/contracts"
)

type EndpointLocation struct {
	ChainID         uint64
	ContractAddress common.Address
}

const sinkAddressHex = "0xFA33c8EF8b5c4f3003361c876a298D1DB61ccA4e"
const sourceAddressHex = "0x14b47DBb4b0F7bff4205F4D1f0c1BEEBCE47b33C"
const sinkRpcUrl = "ws://127.0.0.1:8546/"
const sourceRpcUrl = "ws://127.0.0.1:8556/"

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

func listenForRequests(ctx context.Context, client *ethclient.Client, endpoint endpointcontract.Endpoint) (chan Transfer, chan error) {
	requestChannel := make(chan Transfer)
	errorChannel := make(chan error)

	go func() {
		fail := func(err error) {
			errorChannel <- err
			close(errorChannel)
			close(requestChannel)
		}
		contract, err := endpointcontract.NewEndpointcontract(endpoint.ContractAddress, client)
		if err != nil {
			fail(errors.Wrap(err, "faild to create endpoint contract instance"))
			return
		}

		chainID, err := client.ChainID(ctx)
		if err != nil {
			fail(errors.Wrap(err, "faild to query chain id"))
			return
		}

		watchOpts := &bind.WatchOpts{
			Context: ctx,
			Start:   nil,
		}
		transferRequestedChannel := make(chan *endpointcontract.EndpointcontractTransferRequested)
		sub, err := contract.EndpointcontractFilterer.WatchTransferRequested(watchOpts, transferRequestedChannel)
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
				transfer := NewTransferFromRequestEvent(chainID.Uint64(), ev)
				requestChannel <- transfer
			}
		}
	}()

	return requestChannel, errorChannel
}

func run() error {
	ctx := context.Background()

	sinkClient, err := ethclient.Dial(sinkRpcUrl)
	if err != nil {
		return errors.Wrap(err, "failed to connect to sink Ethereum node")
	}
	sourceClient, err := ethclient.Dial(sourceRpcUrl)
	if err != nil {
		return errors.Wrap(err, "failed to connect to sink Ethereum node")
	}

	sinkChainID, err := sinkClient.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query sink chain id")
	}
	sourceChainID, err := sourceClient.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query source chain id")
	}

	sinkEndpoint := endpointcontract.Endpoint{
		ChainID:         sinkChainID.Uint64(),
		ContractAddress: common.HexToAddress(sinkAddressHex),
	}
	sourceEndpoint := endpointcontract.Endpoint{
		ChainID:         sourceChainID.Uint64(),
		ContractAddress: common.HexToAddress(sourceAddressHex),
	}

	sinkRequestChannel, sinkErrorChannel := listenForRequests(ctx, sinkClient, sinkEndpoint)
	sourceRequestChannel, sourceErrorChannel := listenForRequests(ctx, sourceClient, sourceEndpoint)

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
