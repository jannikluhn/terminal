package oracle

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

func (_Oracle *OracleCaller) RequestExists(opts *bind.CallOpts, transferIdentifier OracleTransferIdentifier) (bool, error) {
	oracleRequest, err := _Oracle.GetRequest(opts, transferIdentifier)
	if err != nil {
		return false, err
	}
	return oracleRequest.LastChallengeTime != 0, nil
}
