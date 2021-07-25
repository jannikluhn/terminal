// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OracleRequest is an auto generated low-level Go binding around an user-defined struct.
type OracleRequest struct {
	Legit             bool
	Value             *big.Int
	LastStake         *big.Int
	LastChallengeTime uint32
	Receiver          common.Address
	Claimed           bool
	Transferred       bool
	Stakers           []common.Address
}

// OracleTransferIdentifier is an auto generated low-level Go binding around an user-defined struct.
type OracleTransferIdentifier struct {
	ChainID     uint64
	BlockNumber uint64
	LogIndex    uint32
	TxHash      [32]byte
}

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chaindID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chaindID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"StartRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chaindID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"TransferSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chaindID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"legit\",\"type\":\"bool\"}],\"name\":\"UpdatedRequest\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structOracle.TransferIdentifier\",\"name\":\"transferIdentifier\",\"type\":\"tuple\"}],\"name\":\"TransferIdentifierHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structOracle.TransferIdentifier\",\"name\":\"transferIdentifier\",\"type\":\"tuple\"}],\"name\":\"claimRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structOracle.TransferIdentifier\",\"name\":\"transferIdentifier\",\"type\":\"tuple\"}],\"name\":\"closeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractEndpointContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structOracle.TransferIdentifier\",\"name\":\"transferIdentifier\",\"type\":\"tuple\"}],\"name\":\"getRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"legit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastStake\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"lastChallengeTime\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferred\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"internalType\":\"structOracle.Request\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_startingStake\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_transferToken\",\"type\":\"address\"},{\"internalType\":\"contractEndpointContract\",\"name\":\"_endpoint\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structOracle.TransferIdentifier\",\"name\":\"transferIdentifier\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"isSuccessfulRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structOracle.TransferIdentifier\",\"name\":\"transferIdentifier\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"startRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"logIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"internalType\":\"structOracle.TransferIdentifier\",\"name\":\"transferIdentifier\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"expectedNewStake\",\"type\":\"uint256\"}],\"name\":\"updateRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// TransferIdentifierHash is a free data retrieval call binding the contract method 0x88828d44.
//
// Solidity: function TransferIdentifierHash((uint64,uint64,uint32,bytes32) transferIdentifier) pure returns(bytes32)
func (_Oracle *OracleCaller) TransferIdentifierHash(opts *bind.CallOpts, transferIdentifier OracleTransferIdentifier) ([32]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "TransferIdentifierHash", transferIdentifier)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TransferIdentifierHash is a free data retrieval call binding the contract method 0x88828d44.
//
// Solidity: function TransferIdentifierHash((uint64,uint64,uint32,bytes32) transferIdentifier) pure returns(bytes32)
func (_Oracle *OracleSession) TransferIdentifierHash(transferIdentifier OracleTransferIdentifier) ([32]byte, error) {
	return _Oracle.Contract.TransferIdentifierHash(&_Oracle.CallOpts, transferIdentifier)
}

// TransferIdentifierHash is a free data retrieval call binding the contract method 0x88828d44.
//
// Solidity: function TransferIdentifierHash((uint64,uint64,uint32,bytes32) transferIdentifier) pure returns(bytes32)
func (_Oracle *OracleCallerSession) TransferIdentifierHash(transferIdentifier OracleTransferIdentifier) ([32]byte, error) {
	return _Oracle.Contract.TransferIdentifierHash(&_Oracle.CallOpts, transferIdentifier)
}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() view returns(uint32)
func (_Oracle *OracleCaller) ChallengePeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "challengePeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() view returns(uint32)
func (_Oracle *OracleSession) ChallengePeriod() (uint32, error) {
	return _Oracle.Contract.ChallengePeriod(&_Oracle.CallOpts)
}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() view returns(uint32)
func (_Oracle *OracleCallerSession) ChallengePeriod() (uint32, error) {
	return _Oracle.Contract.ChallengePeriod(&_Oracle.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Oracle *OracleCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Oracle *OracleSession) Endpoint() (common.Address, error) {
	return _Oracle.Contract.Endpoint(&_Oracle.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Oracle *OracleCallerSession) Endpoint() (common.Address, error) {
	return _Oracle.Contract.Endpoint(&_Oracle.CallOpts)
}

// GetRequest is a free data retrieval call binding the contract method 0xcfb55cef.
//
// Solidity: function getRequest((uint64,uint64,uint32,bytes32) transferIdentifier) view returns((bool,uint256,uint256,uint32,address,bool,bool,address[]))
func (_Oracle *OracleCaller) GetRequest(opts *bind.CallOpts, transferIdentifier OracleTransferIdentifier) (OracleRequest, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRequest", transferIdentifier)

	if err != nil {
		return *new(OracleRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleRequest)).(*OracleRequest)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0xcfb55cef.
//
// Solidity: function getRequest((uint64,uint64,uint32,bytes32) transferIdentifier) view returns((bool,uint256,uint256,uint32,address,bool,bool,address[]))
func (_Oracle *OracleSession) GetRequest(transferIdentifier OracleTransferIdentifier) (OracleRequest, error) {
	return _Oracle.Contract.GetRequest(&_Oracle.CallOpts, transferIdentifier)
}

// GetRequest is a free data retrieval call binding the contract method 0xcfb55cef.
//
// Solidity: function getRequest((uint64,uint64,uint32,bytes32) transferIdentifier) view returns((bool,uint256,uint256,uint32,address,bool,bool,address[]))
func (_Oracle *OracleCallerSession) GetRequest(transferIdentifier OracleTransferIdentifier) (OracleRequest, error) {
	return _Oracle.Contract.GetRequest(&_Oracle.CallOpts, transferIdentifier)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Oracle *OracleCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Oracle *OracleSession) Initialized() (bool, error) {
	return _Oracle.Contract.Initialized(&_Oracle.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Oracle *OracleCallerSession) Initialized() (bool, error) {
	return _Oracle.Contract.Initialized(&_Oracle.CallOpts)
}

// IsSuccessfulRequest is a free data retrieval call binding the contract method 0x4f5e50b2.
//
// Solidity: function isSuccessfulRequest((uint64,uint64,uint32,bytes32) transferIdentifier, uint256 value, address receiver) view returns(bool)
func (_Oracle *OracleCaller) IsSuccessfulRequest(opts *bind.CallOpts, transferIdentifier OracleTransferIdentifier, value *big.Int, receiver common.Address) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isSuccessfulRequest", transferIdentifier, value, receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSuccessfulRequest is a free data retrieval call binding the contract method 0x4f5e50b2.
//
// Solidity: function isSuccessfulRequest((uint64,uint64,uint32,bytes32) transferIdentifier, uint256 value, address receiver) view returns(bool)
func (_Oracle *OracleSession) IsSuccessfulRequest(transferIdentifier OracleTransferIdentifier, value *big.Int, receiver common.Address) (bool, error) {
	return _Oracle.Contract.IsSuccessfulRequest(&_Oracle.CallOpts, transferIdentifier, value, receiver)
}

// IsSuccessfulRequest is a free data retrieval call binding the contract method 0x4f5e50b2.
//
// Solidity: function isSuccessfulRequest((uint64,uint64,uint32,bytes32) transferIdentifier, uint256 value, address receiver) view returns(bool)
func (_Oracle *OracleCallerSession) IsSuccessfulRequest(transferIdentifier OracleTransferIdentifier, value *big.Int, receiver common.Address) (bool, error) {
	return _Oracle.Contract.IsSuccessfulRequest(&_Oracle.CallOpts, transferIdentifier, value, receiver)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_Oracle *OracleCaller) StakedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "stakedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_Oracle *OracleSession) StakedToken() (common.Address, error) {
	return _Oracle.Contract.StakedToken(&_Oracle.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_Oracle *OracleCallerSession) StakedToken() (common.Address, error) {
	return _Oracle.Contract.StakedToken(&_Oracle.CallOpts)
}

// StartingStake is a free data retrieval call binding the contract method 0x727035df.
//
// Solidity: function startingStake() view returns(uint256)
func (_Oracle *OracleCaller) StartingStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "startingStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingStake is a free data retrieval call binding the contract method 0x727035df.
//
// Solidity: function startingStake() view returns(uint256)
func (_Oracle *OracleSession) StartingStake() (*big.Int, error) {
	return _Oracle.Contract.StartingStake(&_Oracle.CallOpts)
}

// StartingStake is a free data retrieval call binding the contract method 0x727035df.
//
// Solidity: function startingStake() view returns(uint256)
func (_Oracle *OracleCallerSession) StartingStake() (*big.Int, error) {
	return _Oracle.Contract.StartingStake(&_Oracle.CallOpts)
}

// TransferToken is a free data retrieval call binding the contract method 0x799a5359.
//
// Solidity: function transferToken() view returns(address)
func (_Oracle *OracleCaller) TransferToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "transferToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransferToken is a free data retrieval call binding the contract method 0x799a5359.
//
// Solidity: function transferToken() view returns(address)
func (_Oracle *OracleSession) TransferToken() (common.Address, error) {
	return _Oracle.Contract.TransferToken(&_Oracle.CallOpts)
}

// TransferToken is a free data retrieval call binding the contract method 0x799a5359.
//
// Solidity: function transferToken() view returns(address)
func (_Oracle *OracleCallerSession) TransferToken() (common.Address, error) {
	return _Oracle.Contract.TransferToken(&_Oracle.CallOpts)
}

// ClaimRequest is a paid mutator transaction binding the contract method 0x5811dd2e.
//
// Solidity: function claimRequest((uint64,uint64,uint32,bytes32) transferIdentifier) returns()
func (_Oracle *OracleTransactor) ClaimRequest(opts *bind.TransactOpts, transferIdentifier OracleTransferIdentifier) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "claimRequest", transferIdentifier)
}

// ClaimRequest is a paid mutator transaction binding the contract method 0x5811dd2e.
//
// Solidity: function claimRequest((uint64,uint64,uint32,bytes32) transferIdentifier) returns()
func (_Oracle *OracleSession) ClaimRequest(transferIdentifier OracleTransferIdentifier) (*types.Transaction, error) {
	return _Oracle.Contract.ClaimRequest(&_Oracle.TransactOpts, transferIdentifier)
}

// ClaimRequest is a paid mutator transaction binding the contract method 0x5811dd2e.
//
// Solidity: function claimRequest((uint64,uint64,uint32,bytes32) transferIdentifier) returns()
func (_Oracle *OracleTransactorSession) ClaimRequest(transferIdentifier OracleTransferIdentifier) (*types.Transaction, error) {
	return _Oracle.Contract.ClaimRequest(&_Oracle.TransactOpts, transferIdentifier)
}

// CloseRequest is a paid mutator transaction binding the contract method 0x7ce2f854.
//
// Solidity: function closeRequest((uint64,uint64,uint32,bytes32) transferIdentifier) returns()
func (_Oracle *OracleTransactor) CloseRequest(opts *bind.TransactOpts, transferIdentifier OracleTransferIdentifier) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "closeRequest", transferIdentifier)
}

// CloseRequest is a paid mutator transaction binding the contract method 0x7ce2f854.
//
// Solidity: function closeRequest((uint64,uint64,uint32,bytes32) transferIdentifier) returns()
func (_Oracle *OracleSession) CloseRequest(transferIdentifier OracleTransferIdentifier) (*types.Transaction, error) {
	return _Oracle.Contract.CloseRequest(&_Oracle.TransactOpts, transferIdentifier)
}

// CloseRequest is a paid mutator transaction binding the contract method 0x7ce2f854.
//
// Solidity: function closeRequest((uint64,uint64,uint32,bytes32) transferIdentifier) returns()
func (_Oracle *OracleTransactorSession) CloseRequest(transferIdentifier OracleTransferIdentifier) (*types.Transaction, error) {
	return _Oracle.Contract.CloseRequest(&_Oracle.TransactOpts, transferIdentifier)
}

// Init is a paid mutator transaction binding the contract method 0xb4fb44a4.
//
// Solidity: function init(uint32 _challengePeriod, uint256 _startingStake, address _stakedToken, address _transferToken, address _endpoint) returns()
func (_Oracle *OracleTransactor) Init(opts *bind.TransactOpts, _challengePeriod uint32, _startingStake *big.Int, _stakedToken common.Address, _transferToken common.Address, _endpoint common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "init", _challengePeriod, _startingStake, _stakedToken, _transferToken, _endpoint)
}

// Init is a paid mutator transaction binding the contract method 0xb4fb44a4.
//
// Solidity: function init(uint32 _challengePeriod, uint256 _startingStake, address _stakedToken, address _transferToken, address _endpoint) returns()
func (_Oracle *OracleSession) Init(_challengePeriod uint32, _startingStake *big.Int, _stakedToken common.Address, _transferToken common.Address, _endpoint common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Init(&_Oracle.TransactOpts, _challengePeriod, _startingStake, _stakedToken, _transferToken, _endpoint)
}

// Init is a paid mutator transaction binding the contract method 0xb4fb44a4.
//
// Solidity: function init(uint32 _challengePeriod, uint256 _startingStake, address _stakedToken, address _transferToken, address _endpoint) returns()
func (_Oracle *OracleTransactorSession) Init(_challengePeriod uint32, _startingStake *big.Int, _stakedToken common.Address, _transferToken common.Address, _endpoint common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Init(&_Oracle.TransactOpts, _challengePeriod, _startingStake, _stakedToken, _transferToken, _endpoint)
}

// StartRequest is a paid mutator transaction binding the contract method 0xf2ecf2cf.
//
// Solidity: function startRequest((uint64,uint64,uint32,bytes32) transferIdentifier, uint256 value, address receiver) returns()
func (_Oracle *OracleTransactor) StartRequest(opts *bind.TransactOpts, transferIdentifier OracleTransferIdentifier, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "startRequest", transferIdentifier, value, receiver)
}

// StartRequest is a paid mutator transaction binding the contract method 0xf2ecf2cf.
//
// Solidity: function startRequest((uint64,uint64,uint32,bytes32) transferIdentifier, uint256 value, address receiver) returns()
func (_Oracle *OracleSession) StartRequest(transferIdentifier OracleTransferIdentifier, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.StartRequest(&_Oracle.TransactOpts, transferIdentifier, value, receiver)
}

// StartRequest is a paid mutator transaction binding the contract method 0xf2ecf2cf.
//
// Solidity: function startRequest((uint64,uint64,uint32,bytes32) transferIdentifier, uint256 value, address receiver) returns()
func (_Oracle *OracleTransactorSession) StartRequest(transferIdentifier OracleTransferIdentifier, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.StartRequest(&_Oracle.TransactOpts, transferIdentifier, value, receiver)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0x5b555135.
//
// Solidity: function updateRequest((uint64,uint64,uint32,bytes32) transferIdentifier, uint256 expectedNewStake) returns()
func (_Oracle *OracleTransactor) UpdateRequest(opts *bind.TransactOpts, transferIdentifier OracleTransferIdentifier, expectedNewStake *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "updateRequest", transferIdentifier, expectedNewStake)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0x5b555135.
//
// Solidity: function updateRequest((uint64,uint64,uint32,bytes32) transferIdentifier, uint256 expectedNewStake) returns()
func (_Oracle *OracleSession) UpdateRequest(transferIdentifier OracleTransferIdentifier, expectedNewStake *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateRequest(&_Oracle.TransactOpts, transferIdentifier, expectedNewStake)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0x5b555135.
//
// Solidity: function updateRequest((uint64,uint64,uint32,bytes32) transferIdentifier, uint256 expectedNewStake) returns()
func (_Oracle *OracleTransactorSession) UpdateRequest(transferIdentifier OracleTransferIdentifier, expectedNewStake *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateRequest(&_Oracle.TransactOpts, transferIdentifier, expectedNewStake)
}

// OracleClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Oracle contract.
type OracleClaimIterator struct {
	Event *OracleClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleClaim represents a Claim event raised by the Oracle contract.
type OracleClaim struct {
	ChaindID    uint64
	BlockNumber uint64
	TxHash      [32]byte
	LogIndex    uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x58f3005c82600a42850ce31d94e79a7a06a4c1d69b602651a74acf3b34fd797f.
//
// Solidity: event Claim(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex)
func (_Oracle *OracleFilterer) FilterClaim(opts *bind.FilterOpts) (*OracleClaimIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &OracleClaimIterator{contract: _Oracle.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x58f3005c82600a42850ce31d94e79a7a06a4c1d69b602651a74acf3b34fd797f.
//
// Solidity: event Claim(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex)
func (_Oracle *OracleFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *OracleClaim) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleClaim)
				if err := _Oracle.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x58f3005c82600a42850ce31d94e79a7a06a4c1d69b602651a74acf3b34fd797f.
//
// Solidity: event Claim(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex)
func (_Oracle *OracleFilterer) ParseClaim(log types.Log) (*OracleClaim, error) {
	event := new(OracleClaim)
	if err := _Oracle.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleStartRequestIterator is returned from FilterStartRequest and is used to iterate over the raw logs and unpacked data for StartRequest events raised by the Oracle contract.
type OracleStartRequestIterator struct {
	Event *OracleStartRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleStartRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleStartRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleStartRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleStartRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleStartRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleStartRequest represents a StartRequest event raised by the Oracle contract.
type OracleStartRequest struct {
	ChaindID    uint64
	BlockNumber uint64
	TxHash      [32]byte
	LogIndex    uint32
	Value       *big.Int
	Receiver    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStartRequest is a free log retrieval operation binding the contract event 0x73ead83b5a1e8b1cbfdff4d7195ff52534782991ae2f95db880dba6c4526cea6.
//
// Solidity: event StartRequest(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint256 value, address receiver)
func (_Oracle *OracleFilterer) FilterStartRequest(opts *bind.FilterOpts) (*OracleStartRequestIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "StartRequest")
	if err != nil {
		return nil, err
	}
	return &OracleStartRequestIterator{contract: _Oracle.contract, event: "StartRequest", logs: logs, sub: sub}, nil
}

// WatchStartRequest is a free log subscription operation binding the contract event 0x73ead83b5a1e8b1cbfdff4d7195ff52534782991ae2f95db880dba6c4526cea6.
//
// Solidity: event StartRequest(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint256 value, address receiver)
func (_Oracle *OracleFilterer) WatchStartRequest(opts *bind.WatchOpts, sink chan<- *OracleStartRequest) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "StartRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleStartRequest)
				if err := _Oracle.contract.UnpackLog(event, "StartRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartRequest is a log parse operation binding the contract event 0x73ead83b5a1e8b1cbfdff4d7195ff52534782991ae2f95db880dba6c4526cea6.
//
// Solidity: event StartRequest(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint256 value, address receiver)
func (_Oracle *OracleFilterer) ParseStartRequest(log types.Log) (*OracleStartRequest, error) {
	event := new(OracleStartRequest)
	if err := _Oracle.contract.UnpackLog(event, "StartRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTransferSuccessfulIterator is returned from FilterTransferSuccessful and is used to iterate over the raw logs and unpacked data for TransferSuccessful events raised by the Oracle contract.
type OracleTransferSuccessfulIterator struct {
	Event *OracleTransferSuccessful // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleTransferSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTransferSuccessful)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleTransferSuccessful)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleTransferSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTransferSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTransferSuccessful represents a TransferSuccessful event raised by the Oracle contract.
type OracleTransferSuccessful struct {
	ChaindID    uint64
	BlockNumber uint64
	TxHash      [32]byte
	LogIndex    uint32
	Value       *big.Int
	Receiver    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferSuccessful is a free log retrieval operation binding the contract event 0x24f7d010bbb7e472f65d80503685e78e29e3c441f27700a4904033ff2d1efda0.
//
// Solidity: event TransferSuccessful(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint256 value, address receiver)
func (_Oracle *OracleFilterer) FilterTransferSuccessful(opts *bind.FilterOpts) (*OracleTransferSuccessfulIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "TransferSuccessful")
	if err != nil {
		return nil, err
	}
	return &OracleTransferSuccessfulIterator{contract: _Oracle.contract, event: "TransferSuccessful", logs: logs, sub: sub}, nil
}

// WatchTransferSuccessful is a free log subscription operation binding the contract event 0x24f7d010bbb7e472f65d80503685e78e29e3c441f27700a4904033ff2d1efda0.
//
// Solidity: event TransferSuccessful(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint256 value, address receiver)
func (_Oracle *OracleFilterer) WatchTransferSuccessful(opts *bind.WatchOpts, sink chan<- *OracleTransferSuccessful) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "TransferSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTransferSuccessful)
				if err := _Oracle.contract.UnpackLog(event, "TransferSuccessful", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSuccessful is a log parse operation binding the contract event 0x24f7d010bbb7e472f65d80503685e78e29e3c441f27700a4904033ff2d1efda0.
//
// Solidity: event TransferSuccessful(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint256 value, address receiver)
func (_Oracle *OracleFilterer) ParseTransferSuccessful(log types.Log) (*OracleTransferSuccessful, error) {
	event := new(OracleTransferSuccessful)
	if err := _Oracle.contract.UnpackLog(event, "TransferSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleUpdatedRequestIterator is returned from FilterUpdatedRequest and is used to iterate over the raw logs and unpacked data for UpdatedRequest events raised by the Oracle contract.
type OracleUpdatedRequestIterator struct {
	Event *OracleUpdatedRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleUpdatedRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleUpdatedRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleUpdatedRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleUpdatedRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleUpdatedRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleUpdatedRequest represents a UpdatedRequest event raised by the Oracle contract.
type OracleUpdatedRequest struct {
	ChaindID    uint64
	BlockNumber uint64
	TxHash      [32]byte
	LogIndex    uint32
	NewStake    *big.Int
	Legit       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRequest is a free log retrieval operation binding the contract event 0xd0c30a42aa79c547b734320baa819e21f63b84c2038527a07462df42489de15c.
//
// Solidity: event UpdatedRequest(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint256 newStake, bool legit)
func (_Oracle *OracleFilterer) FilterUpdatedRequest(opts *bind.FilterOpts) (*OracleUpdatedRequestIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "UpdatedRequest")
	if err != nil {
		return nil, err
	}
	return &OracleUpdatedRequestIterator{contract: _Oracle.contract, event: "UpdatedRequest", logs: logs, sub: sub}, nil
}

// WatchUpdatedRequest is a free log subscription operation binding the contract event 0xd0c30a42aa79c547b734320baa819e21f63b84c2038527a07462df42489de15c.
//
// Solidity: event UpdatedRequest(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint256 newStake, bool legit)
func (_Oracle *OracleFilterer) WatchUpdatedRequest(opts *bind.WatchOpts, sink chan<- *OracleUpdatedRequest) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "UpdatedRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleUpdatedRequest)
				if err := _Oracle.contract.UnpackLog(event, "UpdatedRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedRequest is a log parse operation binding the contract event 0xd0c30a42aa79c547b734320baa819e21f63b84c2038527a07462df42489de15c.
//
// Solidity: event UpdatedRequest(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint256 newStake, bool legit)
func (_Oracle *OracleFilterer) ParseUpdatedRequest(log types.Log) (*OracleUpdatedRequest, error) {
	event := new(OracleUpdatedRequest)
	if err := _Oracle.contract.UnpackLog(event, "UpdatedRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
