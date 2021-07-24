package main

import (
	"context"
	"log"
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

const sinkChainID = 5
const sinkAddressHex = "0xFA33c8EF8b5c4f3003361c876a298D1DB61ccA4e"

const sinkRpcUrl = "ws://127.0.0.1:8546/"

func main() {
	err := listen()
	if err != nil {
		log.Fatalf("Error listening to events:", err)
	}
}

func listen() error {
	ctx := context.Background()

	client, err := ethclient.Dial(sinkRpcUrl)
	if err != nil {
		return errors.Wrap(err, "failed to connect to Ethereum node")
	}

	sinkLocation := EndpointLocation{
		ChainID:         sinkChainID,
		ContractAddress: common.HexToAddress(sinkAddressHex),
	}
	sinkContract, err := endpointcontract.NewEndpointcontract(sinkLocation.ContractAddress, client)
	if err != nil {
		return errors.Wrap(err, "faild to create endpoint contract instance")
	}

	watchOpts := &bind.WatchOpts{
		Context: ctx,
		Start:   nil,
	}
	transferRequestedChannel := make(chan *endpointcontract.EndpointcontractTransferRequested)
	sub, err := sinkContract.EndpointcontractFilterer.WatchTransferRequested(watchOpts, transferRequestedChannel)
	if err != nil {
		return errors.Wrap(err, "faild to subscribe to TransferRequested events")
	}

	for {
		select {
		case err := <-sub.Err():
			return errors.Wrap(err, "event subscription error")
		case ev := <-transferRequestedChannel:
			log.Println("Received EndpointRequested event in block %d, tx %s, log index %d", ev.Raw.BlockNumber, ev.Raw.TxHash, ev.Raw.Index)
		case <-time.After(5 * time.Second):
			log.Println("listening for events")
		}
	}
}
