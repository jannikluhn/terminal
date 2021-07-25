#! /usr/bin/env bash

set -eu

DATADIR=${SHROOT}/geth-chain-source
PKFILE=${SHROOT}/privkey_source.txt
GENESISFILE=${SHROOT}/genesis_source.json

PK=b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3774
ADDR=0x4578FCAFBA09e095C405348a5829850B9B19510B

rm -rf ${DATADIR}

geth init --datadir ${DATADIR} ${GENESISFILE}

echo ${PK} >${PKFILE}
yes "" | geth --datadir ${DATADIR} account import ${PKFILE}
rm ${PKFILE}

yes "" | geth --datadir ${DATADIR} --miner.etherbase=${ADDR} --http --http.port 8555 --http.api personal,eth,net,web3,debug --http.corsdomain "*" --ws --ws.port 8556 --ws.api personal,eth,net,web3,debug --gcmode archive --rpc.allow-unprotected-txs --nodiscover --port 30304 --mine --unlock ${ADDR} --allow-insecure-unlock
