#! /usr/bin/env bash

set -eu

DATADIR=${SHROOT}/geth-chain-sink
PKFILE=${SHROOT}/privkey_sink.txt
GENESISFILE=${SHROOT}/genesis_sink.json

PK=b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773
ADDR=0x1dF62f291b2E969fB0849d99D9Ce41e2F137006e

rm -rf ${DATADIR}

geth init --datadir ${DATADIR} ${GENESISFILE}

echo ${PK} >${PKFILE}
yes "" | geth --datadir ${DATADIR} account import ${PKFILE}
rm ${PKFILE}

yes "" | geth --datadir ${DATADIR} --miner.etherbase=${ADDR} --http --http.port 8545 --http.api personal,eth,net,web3,debug --http.corsdomain "*" --ws --ws.port 8546 --ws.api personal,eth,net,web3,debug --gcmode archive --rpc.allow-unprotected-txs --nodiscover --port 30303 --mine --unlock ${ADDR} --allow-insecure-unlock
