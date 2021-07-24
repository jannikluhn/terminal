#! /usr/bin/env bash

set -eu

SHROOT=sourcechain
DATADIR=${SHROOT}/geth-chain
PKFILE=${SHROOT}/privkey.txt

PK=b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3774
ADDR=0x4578FCAFBA09e095C405348a5829850B9B19510B

rm -rf ${SHROOT}
mkdir ${SHROOT}
echo ${PK} >${PKFILE}
yes "" | geth --datadir ${DATADIR} account import ${PKFILE}
rm ${PKFILE}
exec geth --dev --dev.period=3 --datadir ${DATADIR} --miner.etherbase=${ADDR} --http --http.port 8555 --http.api personal,eth,net,web3,debug --http.corsdomain "*" --ws --ws.port 8556 --ws.api personal,eth,net,web3,debug --gcmode archive --rpc.allow-unprotected-txs
