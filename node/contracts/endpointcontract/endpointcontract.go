// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package endpointcontract

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

// Endpoint is an auto generated low-level Go binding around an user-defined struct.
type Endpoint struct {
	ChainID         uint64
	ContractAddress common.Address
}

// EndpointcontractABI is the input ABI used to generate the binding from.
const EndpointcontractABI = "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEndpoint\",\"name\":\"endpoint\",\"type\":\"tuple\"}],\"name\":\"ConnectedToEndpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEndpoint\",\"name\":\"endpoint\",\"type\":\"tuple\"}],\"name\":\"DisconnectedFromEndpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEndpoint\",\"name\":\"endpoint\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structEndpoint\",\"name\":\"endpoint\",\"type\":\"tuple\"}],\"name\":\"connectToEndpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"connectedEndpoints\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structEndpoint\",\"name\":\"endpoint\",\"type\":\"tuple\"}],\"name\":\"disconnectFromEndpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structEndpoint\",\"name\":\"endpoint\",\"type\":\"tuple\"}],\"name\":\"endpointHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structEndpoint\",\"name\":\"endpoint\",\"type\":\"tuple\"}],\"name\":\"isConnectedTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structEndpoint\",\"name\":\"endpoint\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structEndpoint\",\"name\":\"endpoint\",\"type\":\"tuple\"}],\"name\":\"revertIfNotConnectedTo\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Endpointcontract is an auto generated Go binding around an Ethereum contract.
type Endpointcontract struct {
	EndpointcontractCaller     // Read-only binding to the contract
	EndpointcontractTransactor // Write-only binding to the contract
	EndpointcontractFilterer   // Log filterer for contract events
}

// EndpointcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EndpointcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EndpointcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EndpointcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EndpointcontractSession struct {
	Contract     *Endpointcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EndpointcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EndpointcontractCallerSession struct {
	Contract *EndpointcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EndpointcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EndpointcontractTransactorSession struct {
	Contract     *EndpointcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EndpointcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EndpointcontractRaw struct {
	Contract *Endpointcontract // Generic contract binding to access the raw methods on
}

// EndpointcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EndpointcontractCallerRaw struct {
	Contract *EndpointcontractCaller // Generic read-only contract binding to access the raw methods on
}

// EndpointcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EndpointcontractTransactorRaw struct {
	Contract *EndpointcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEndpointcontract creates a new instance of Endpointcontract, bound to a specific deployed contract.
func NewEndpointcontract(address common.Address, backend bind.ContractBackend) (*Endpointcontract, error) {
	contract, err := bindEndpointcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Endpointcontract{EndpointcontractCaller: EndpointcontractCaller{contract: contract}, EndpointcontractTransactor: EndpointcontractTransactor{contract: contract}, EndpointcontractFilterer: EndpointcontractFilterer{contract: contract}}, nil
}

// NewEndpointcontractCaller creates a new read-only instance of Endpointcontract, bound to a specific deployed contract.
func NewEndpointcontractCaller(address common.Address, caller bind.ContractCaller) (*EndpointcontractCaller, error) {
	contract, err := bindEndpointcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EndpointcontractCaller{contract: contract}, nil
}

// NewEndpointcontractTransactor creates a new write-only instance of Endpointcontract, bound to a specific deployed contract.
func NewEndpointcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*EndpointcontractTransactor, error) {
	contract, err := bindEndpointcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EndpointcontractTransactor{contract: contract}, nil
}

// NewEndpointcontractFilterer creates a new log filterer instance of Endpointcontract, bound to a specific deployed contract.
func NewEndpointcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*EndpointcontractFilterer, error) {
	contract, err := bindEndpointcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EndpointcontractFilterer{contract: contract}, nil
}

// bindEndpointcontract binds a generic wrapper to an already deployed contract.
func bindEndpointcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EndpointcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Endpointcontract *EndpointcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Endpointcontract.Contract.EndpointcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Endpointcontract *EndpointcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Endpointcontract.Contract.EndpointcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Endpointcontract *EndpointcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Endpointcontract.Contract.EndpointcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Endpointcontract *EndpointcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Endpointcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Endpointcontract *EndpointcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Endpointcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Endpointcontract *EndpointcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Endpointcontract.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Endpointcontract *EndpointcontractCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Endpointcontract *EndpointcontractSession) ADMINROLE() ([32]byte, error) {
	return _Endpointcontract.Contract.ADMINROLE(&_Endpointcontract.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Endpointcontract *EndpointcontractCallerSession) ADMINROLE() ([32]byte, error) {
	return _Endpointcontract.Contract.ADMINROLE(&_Endpointcontract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Endpointcontract *EndpointcontractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Endpointcontract *EndpointcontractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Endpointcontract.Contract.DEFAULTADMINROLE(&_Endpointcontract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Endpointcontract *EndpointcontractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Endpointcontract.Contract.DEFAULTADMINROLE(&_Endpointcontract.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_Endpointcontract *EndpointcontractCaller) ORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_Endpointcontract *EndpointcontractSession) ORACLEROLE() ([32]byte, error) {
	return _Endpointcontract.Contract.ORACLEROLE(&_Endpointcontract.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_Endpointcontract *EndpointcontractCallerSession) ORACLEROLE() ([32]byte, error) {
	return _Endpointcontract.Contract.ORACLEROLE(&_Endpointcontract.CallOpts)
}

// ConnectedEndpoints is a free data retrieval call binding the contract method 0x4d970931.
//
// Solidity: function connectedEndpoints(bytes32 ) view returns(bool)
func (_Endpointcontract *EndpointcontractCaller) ConnectedEndpoints(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "connectedEndpoints", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConnectedEndpoints is a free data retrieval call binding the contract method 0x4d970931.
//
// Solidity: function connectedEndpoints(bytes32 ) view returns(bool)
func (_Endpointcontract *EndpointcontractSession) ConnectedEndpoints(arg0 [32]byte) (bool, error) {
	return _Endpointcontract.Contract.ConnectedEndpoints(&_Endpointcontract.CallOpts, arg0)
}

// ConnectedEndpoints is a free data retrieval call binding the contract method 0x4d970931.
//
// Solidity: function connectedEndpoints(bytes32 ) view returns(bool)
func (_Endpointcontract *EndpointcontractCallerSession) ConnectedEndpoints(arg0 [32]byte) (bool, error) {
	return _Endpointcontract.Contract.ConnectedEndpoints(&_Endpointcontract.CallOpts, arg0)
}

// EndpointHash is a free data retrieval call binding the contract method 0x7e5b908a.
//
// Solidity: function endpointHash((uint64,address) endpoint) pure returns(bytes32)
func (_Endpointcontract *EndpointcontractCaller) EndpointHash(opts *bind.CallOpts, endpoint Endpoint) ([32]byte, error) {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "endpointHash", endpoint)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EndpointHash is a free data retrieval call binding the contract method 0x7e5b908a.
//
// Solidity: function endpointHash((uint64,address) endpoint) pure returns(bytes32)
func (_Endpointcontract *EndpointcontractSession) EndpointHash(endpoint Endpoint) ([32]byte, error) {
	return _Endpointcontract.Contract.EndpointHash(&_Endpointcontract.CallOpts, endpoint)
}

// EndpointHash is a free data retrieval call binding the contract method 0x7e5b908a.
//
// Solidity: function endpointHash((uint64,address) endpoint) pure returns(bytes32)
func (_Endpointcontract *EndpointcontractCallerSession) EndpointHash(endpoint Endpoint) ([32]byte, error) {
	return _Endpointcontract.Contract.EndpointHash(&_Endpointcontract.CallOpts, endpoint)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Endpointcontract *EndpointcontractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Endpointcontract *EndpointcontractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Endpointcontract.Contract.GetRoleAdmin(&_Endpointcontract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Endpointcontract *EndpointcontractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Endpointcontract.Contract.GetRoleAdmin(&_Endpointcontract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Endpointcontract *EndpointcontractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Endpointcontract *EndpointcontractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Endpointcontract.Contract.HasRole(&_Endpointcontract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Endpointcontract *EndpointcontractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Endpointcontract.Contract.HasRole(&_Endpointcontract.CallOpts, role, account)
}

// IsConnectedTo is a free data retrieval call binding the contract method 0x90f2408f.
//
// Solidity: function isConnectedTo((uint64,address) endpoint) view returns(bool)
func (_Endpointcontract *EndpointcontractCaller) IsConnectedTo(opts *bind.CallOpts, endpoint Endpoint) (bool, error) {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "isConnectedTo", endpoint)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConnectedTo is a free data retrieval call binding the contract method 0x90f2408f.
//
// Solidity: function isConnectedTo((uint64,address) endpoint) view returns(bool)
func (_Endpointcontract *EndpointcontractSession) IsConnectedTo(endpoint Endpoint) (bool, error) {
	return _Endpointcontract.Contract.IsConnectedTo(&_Endpointcontract.CallOpts, endpoint)
}

// IsConnectedTo is a free data retrieval call binding the contract method 0x90f2408f.
//
// Solidity: function isConnectedTo((uint64,address) endpoint) view returns(bool)
func (_Endpointcontract *EndpointcontractCallerSession) IsConnectedTo(endpoint Endpoint) (bool, error) {
	return _Endpointcontract.Contract.IsConnectedTo(&_Endpointcontract.CallOpts, endpoint)
}

// RevertIfNotConnectedTo is a free data retrieval call binding the contract method 0x07608e76.
//
// Solidity: function revertIfNotConnectedTo((uint64,address) endpoint) view returns()
func (_Endpointcontract *EndpointcontractCaller) RevertIfNotConnectedTo(opts *bind.CallOpts, endpoint Endpoint) error {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "revertIfNotConnectedTo", endpoint)

	if err != nil {
		return err
	}

	return err

}

// RevertIfNotConnectedTo is a free data retrieval call binding the contract method 0x07608e76.
//
// Solidity: function revertIfNotConnectedTo((uint64,address) endpoint) view returns()
func (_Endpointcontract *EndpointcontractSession) RevertIfNotConnectedTo(endpoint Endpoint) error {
	return _Endpointcontract.Contract.RevertIfNotConnectedTo(&_Endpointcontract.CallOpts, endpoint)
}

// RevertIfNotConnectedTo is a free data retrieval call binding the contract method 0x07608e76.
//
// Solidity: function revertIfNotConnectedTo((uint64,address) endpoint) view returns()
func (_Endpointcontract *EndpointcontractCallerSession) RevertIfNotConnectedTo(endpoint Endpoint) error {
	return _Endpointcontract.Contract.RevertIfNotConnectedTo(&_Endpointcontract.CallOpts, endpoint)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Endpointcontract *EndpointcontractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Endpointcontract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Endpointcontract *EndpointcontractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Endpointcontract.Contract.SupportsInterface(&_Endpointcontract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Endpointcontract *EndpointcontractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Endpointcontract.Contract.SupportsInterface(&_Endpointcontract.CallOpts, interfaceId)
}

// ConnectToEndpoint is a paid mutator transaction binding the contract method 0xb2e5f7e9.
//
// Solidity: function connectToEndpoint((uint64,address) endpoint) returns()
func (_Endpointcontract *EndpointcontractTransactor) ConnectToEndpoint(opts *bind.TransactOpts, endpoint Endpoint) (*types.Transaction, error) {
	return _Endpointcontract.contract.Transact(opts, "connectToEndpoint", endpoint)
}

// ConnectToEndpoint is a paid mutator transaction binding the contract method 0xb2e5f7e9.
//
// Solidity: function connectToEndpoint((uint64,address) endpoint) returns()
func (_Endpointcontract *EndpointcontractSession) ConnectToEndpoint(endpoint Endpoint) (*types.Transaction, error) {
	return _Endpointcontract.Contract.ConnectToEndpoint(&_Endpointcontract.TransactOpts, endpoint)
}

// ConnectToEndpoint is a paid mutator transaction binding the contract method 0xb2e5f7e9.
//
// Solidity: function connectToEndpoint((uint64,address) endpoint) returns()
func (_Endpointcontract *EndpointcontractTransactorSession) ConnectToEndpoint(endpoint Endpoint) (*types.Transaction, error) {
	return _Endpointcontract.Contract.ConnectToEndpoint(&_Endpointcontract.TransactOpts, endpoint)
}

// DisconnectFromEndpoint is a paid mutator transaction binding the contract method 0x131e69f8.
//
// Solidity: function disconnectFromEndpoint((uint64,address) endpoint) returns()
func (_Endpointcontract *EndpointcontractTransactor) DisconnectFromEndpoint(opts *bind.TransactOpts, endpoint Endpoint) (*types.Transaction, error) {
	return _Endpointcontract.contract.Transact(opts, "disconnectFromEndpoint", endpoint)
}

// DisconnectFromEndpoint is a paid mutator transaction binding the contract method 0x131e69f8.
//
// Solidity: function disconnectFromEndpoint((uint64,address) endpoint) returns()
func (_Endpointcontract *EndpointcontractSession) DisconnectFromEndpoint(endpoint Endpoint) (*types.Transaction, error) {
	return _Endpointcontract.Contract.DisconnectFromEndpoint(&_Endpointcontract.TransactOpts, endpoint)
}

// DisconnectFromEndpoint is a paid mutator transaction binding the contract method 0x131e69f8.
//
// Solidity: function disconnectFromEndpoint((uint64,address) endpoint) returns()
func (_Endpointcontract *EndpointcontractTransactorSession) DisconnectFromEndpoint(endpoint Endpoint) (*types.Transaction, error) {
	return _Endpointcontract.Contract.DisconnectFromEndpoint(&_Endpointcontract.TransactOpts, endpoint)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Endpointcontract *EndpointcontractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Endpointcontract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Endpointcontract *EndpointcontractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Endpointcontract.Contract.GrantRole(&_Endpointcontract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Endpointcontract *EndpointcontractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Endpointcontract.Contract.GrantRole(&_Endpointcontract.TransactOpts, role, account)
}

// ReleaseTransfer is a paid mutator transaction binding the contract method 0xaf56cd59.
//
// Solidity: function releaseTransfer(address receiver, uint256 amount) returns()
func (_Endpointcontract *EndpointcontractTransactor) ReleaseTransfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Endpointcontract.contract.Transact(opts, "releaseTransfer", receiver, amount)
}

// ReleaseTransfer is a paid mutator transaction binding the contract method 0xaf56cd59.
//
// Solidity: function releaseTransfer(address receiver, uint256 amount) returns()
func (_Endpointcontract *EndpointcontractSession) ReleaseTransfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Endpointcontract.Contract.ReleaseTransfer(&_Endpointcontract.TransactOpts, receiver, amount)
}

// ReleaseTransfer is a paid mutator transaction binding the contract method 0xaf56cd59.
//
// Solidity: function releaseTransfer(address receiver, uint256 amount) returns()
func (_Endpointcontract *EndpointcontractTransactorSession) ReleaseTransfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Endpointcontract.Contract.ReleaseTransfer(&_Endpointcontract.TransactOpts, receiver, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Endpointcontract *EndpointcontractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Endpointcontract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Endpointcontract *EndpointcontractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Endpointcontract.Contract.RenounceRole(&_Endpointcontract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Endpointcontract *EndpointcontractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Endpointcontract.Contract.RenounceRole(&_Endpointcontract.TransactOpts, role, account)
}

// RequestTransfer is a paid mutator transaction binding the contract method 0x65190401.
//
// Solidity: function requestTransfer((uint64,address) endpoint, address receiver, uint256 amount) returns()
func (_Endpointcontract *EndpointcontractTransactor) RequestTransfer(opts *bind.TransactOpts, endpoint Endpoint, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Endpointcontract.contract.Transact(opts, "requestTransfer", endpoint, receiver, amount)
}

// RequestTransfer is a paid mutator transaction binding the contract method 0x65190401.
//
// Solidity: function requestTransfer((uint64,address) endpoint, address receiver, uint256 amount) returns()
func (_Endpointcontract *EndpointcontractSession) RequestTransfer(endpoint Endpoint, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Endpointcontract.Contract.RequestTransfer(&_Endpointcontract.TransactOpts, endpoint, receiver, amount)
}

// RequestTransfer is a paid mutator transaction binding the contract method 0x65190401.
//
// Solidity: function requestTransfer((uint64,address) endpoint, address receiver, uint256 amount) returns()
func (_Endpointcontract *EndpointcontractTransactorSession) RequestTransfer(endpoint Endpoint, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Endpointcontract.Contract.RequestTransfer(&_Endpointcontract.TransactOpts, endpoint, receiver, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Endpointcontract *EndpointcontractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Endpointcontract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Endpointcontract *EndpointcontractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Endpointcontract.Contract.RevokeRole(&_Endpointcontract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Endpointcontract *EndpointcontractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Endpointcontract.Contract.RevokeRole(&_Endpointcontract.TransactOpts, role, account)
}

// EndpointcontractConnectedToEndpointIterator is returned from FilterConnectedToEndpoint and is used to iterate over the raw logs and unpacked data for ConnectedToEndpoint events raised by the Endpointcontract contract.
type EndpointcontractConnectedToEndpointIterator struct {
	Event *EndpointcontractConnectedToEndpoint // Event containing the contract specifics and raw log

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
func (it *EndpointcontractConnectedToEndpointIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointcontractConnectedToEndpoint)
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
		it.Event = new(EndpointcontractConnectedToEndpoint)
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
func (it *EndpointcontractConnectedToEndpointIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointcontractConnectedToEndpointIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointcontractConnectedToEndpoint represents a ConnectedToEndpoint event raised by the Endpointcontract contract.
type EndpointcontractConnectedToEndpoint struct {
	Endpoint Endpoint
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConnectedToEndpoint is a free log retrieval operation binding the contract event 0xf419a17f113f5c02b2f88e9d18774fd5c364592ba353748cbe656138d786d6b8.
//
// Solidity: event ConnectedToEndpoint((uint64,address) endpoint)
func (_Endpointcontract *EndpointcontractFilterer) FilterConnectedToEndpoint(opts *bind.FilterOpts) (*EndpointcontractConnectedToEndpointIterator, error) {

	logs, sub, err := _Endpointcontract.contract.FilterLogs(opts, "ConnectedToEndpoint")
	if err != nil {
		return nil, err
	}
	return &EndpointcontractConnectedToEndpointIterator{contract: _Endpointcontract.contract, event: "ConnectedToEndpoint", logs: logs, sub: sub}, nil
}

// WatchConnectedToEndpoint is a free log subscription operation binding the contract event 0xf419a17f113f5c02b2f88e9d18774fd5c364592ba353748cbe656138d786d6b8.
//
// Solidity: event ConnectedToEndpoint((uint64,address) endpoint)
func (_Endpointcontract *EndpointcontractFilterer) WatchConnectedToEndpoint(opts *bind.WatchOpts, sink chan<- *EndpointcontractConnectedToEndpoint) (event.Subscription, error) {

	logs, sub, err := _Endpointcontract.contract.WatchLogs(opts, "ConnectedToEndpoint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointcontractConnectedToEndpoint)
				if err := _Endpointcontract.contract.UnpackLog(event, "ConnectedToEndpoint", log); err != nil {
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

// ParseConnectedToEndpoint is a log parse operation binding the contract event 0xf419a17f113f5c02b2f88e9d18774fd5c364592ba353748cbe656138d786d6b8.
//
// Solidity: event ConnectedToEndpoint((uint64,address) endpoint)
func (_Endpointcontract *EndpointcontractFilterer) ParseConnectedToEndpoint(log types.Log) (*EndpointcontractConnectedToEndpoint, error) {
	event := new(EndpointcontractConnectedToEndpoint)
	if err := _Endpointcontract.contract.UnpackLog(event, "ConnectedToEndpoint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointcontractDisconnectedFromEndpointIterator is returned from FilterDisconnectedFromEndpoint and is used to iterate over the raw logs and unpacked data for DisconnectedFromEndpoint events raised by the Endpointcontract contract.
type EndpointcontractDisconnectedFromEndpointIterator struct {
	Event *EndpointcontractDisconnectedFromEndpoint // Event containing the contract specifics and raw log

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
func (it *EndpointcontractDisconnectedFromEndpointIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointcontractDisconnectedFromEndpoint)
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
		it.Event = new(EndpointcontractDisconnectedFromEndpoint)
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
func (it *EndpointcontractDisconnectedFromEndpointIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointcontractDisconnectedFromEndpointIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointcontractDisconnectedFromEndpoint represents a DisconnectedFromEndpoint event raised by the Endpointcontract contract.
type EndpointcontractDisconnectedFromEndpoint struct {
	Endpoint Endpoint
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDisconnectedFromEndpoint is a free log retrieval operation binding the contract event 0x389ae7698956d1e58be75146dd0c4f53b3b7448a874100828b1d07f2f38c46d8.
//
// Solidity: event DisconnectedFromEndpoint((uint64,address) endpoint)
func (_Endpointcontract *EndpointcontractFilterer) FilterDisconnectedFromEndpoint(opts *bind.FilterOpts) (*EndpointcontractDisconnectedFromEndpointIterator, error) {

	logs, sub, err := _Endpointcontract.contract.FilterLogs(opts, "DisconnectedFromEndpoint")
	if err != nil {
		return nil, err
	}
	return &EndpointcontractDisconnectedFromEndpointIterator{contract: _Endpointcontract.contract, event: "DisconnectedFromEndpoint", logs: logs, sub: sub}, nil
}

// WatchDisconnectedFromEndpoint is a free log subscription operation binding the contract event 0x389ae7698956d1e58be75146dd0c4f53b3b7448a874100828b1d07f2f38c46d8.
//
// Solidity: event DisconnectedFromEndpoint((uint64,address) endpoint)
func (_Endpointcontract *EndpointcontractFilterer) WatchDisconnectedFromEndpoint(opts *bind.WatchOpts, sink chan<- *EndpointcontractDisconnectedFromEndpoint) (event.Subscription, error) {

	logs, sub, err := _Endpointcontract.contract.WatchLogs(opts, "DisconnectedFromEndpoint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointcontractDisconnectedFromEndpoint)
				if err := _Endpointcontract.contract.UnpackLog(event, "DisconnectedFromEndpoint", log); err != nil {
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

// ParseDisconnectedFromEndpoint is a log parse operation binding the contract event 0x389ae7698956d1e58be75146dd0c4f53b3b7448a874100828b1d07f2f38c46d8.
//
// Solidity: event DisconnectedFromEndpoint((uint64,address) endpoint)
func (_Endpointcontract *EndpointcontractFilterer) ParseDisconnectedFromEndpoint(log types.Log) (*EndpointcontractDisconnectedFromEndpoint, error) {
	event := new(EndpointcontractDisconnectedFromEndpoint)
	if err := _Endpointcontract.contract.UnpackLog(event, "DisconnectedFromEndpoint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointcontractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Endpointcontract contract.
type EndpointcontractRoleAdminChangedIterator struct {
	Event *EndpointcontractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EndpointcontractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointcontractRoleAdminChanged)
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
		it.Event = new(EndpointcontractRoleAdminChanged)
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
func (it *EndpointcontractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointcontractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointcontractRoleAdminChanged represents a RoleAdminChanged event raised by the Endpointcontract contract.
type EndpointcontractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Endpointcontract *EndpointcontractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EndpointcontractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Endpointcontract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EndpointcontractRoleAdminChangedIterator{contract: _Endpointcontract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Endpointcontract *EndpointcontractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EndpointcontractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Endpointcontract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointcontractRoleAdminChanged)
				if err := _Endpointcontract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Endpointcontract *EndpointcontractFilterer) ParseRoleAdminChanged(log types.Log) (*EndpointcontractRoleAdminChanged, error) {
	event := new(EndpointcontractRoleAdminChanged)
	if err := _Endpointcontract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointcontractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Endpointcontract contract.
type EndpointcontractRoleGrantedIterator struct {
	Event *EndpointcontractRoleGranted // Event containing the contract specifics and raw log

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
func (it *EndpointcontractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointcontractRoleGranted)
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
		it.Event = new(EndpointcontractRoleGranted)
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
func (it *EndpointcontractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointcontractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointcontractRoleGranted represents a RoleGranted event raised by the Endpointcontract contract.
type EndpointcontractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Endpointcontract *EndpointcontractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EndpointcontractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Endpointcontract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EndpointcontractRoleGrantedIterator{contract: _Endpointcontract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Endpointcontract *EndpointcontractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EndpointcontractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Endpointcontract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointcontractRoleGranted)
				if err := _Endpointcontract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Endpointcontract *EndpointcontractFilterer) ParseRoleGranted(log types.Log) (*EndpointcontractRoleGranted, error) {
	event := new(EndpointcontractRoleGranted)
	if err := _Endpointcontract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointcontractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Endpointcontract contract.
type EndpointcontractRoleRevokedIterator struct {
	Event *EndpointcontractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EndpointcontractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointcontractRoleRevoked)
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
		it.Event = new(EndpointcontractRoleRevoked)
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
func (it *EndpointcontractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointcontractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointcontractRoleRevoked represents a RoleRevoked event raised by the Endpointcontract contract.
type EndpointcontractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Endpointcontract *EndpointcontractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EndpointcontractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Endpointcontract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EndpointcontractRoleRevokedIterator{contract: _Endpointcontract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Endpointcontract *EndpointcontractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EndpointcontractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Endpointcontract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointcontractRoleRevoked)
				if err := _Endpointcontract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Endpointcontract *EndpointcontractFilterer) ParseRoleRevoked(log types.Log) (*EndpointcontractRoleRevoked, error) {
	event := new(EndpointcontractRoleRevoked)
	if err := _Endpointcontract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointcontractTransferReleasedIterator is returned from FilterTransferReleased and is used to iterate over the raw logs and unpacked data for TransferReleased events raised by the Endpointcontract contract.
type EndpointcontractTransferReleasedIterator struct {
	Event *EndpointcontractTransferReleased // Event containing the contract specifics and raw log

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
func (it *EndpointcontractTransferReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointcontractTransferReleased)
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
		it.Event = new(EndpointcontractTransferReleased)
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
func (it *EndpointcontractTransferReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointcontractTransferReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointcontractTransferReleased represents a TransferReleased event raised by the Endpointcontract contract.
type EndpointcontractTransferReleased struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferReleased is a free log retrieval operation binding the contract event 0x0b3206d324e5269b2561359f447fd104c22f9cf582bb3886656a6bfd736c4e15.
//
// Solidity: event TransferReleased(address receiver, uint256 amount)
func (_Endpointcontract *EndpointcontractFilterer) FilterTransferReleased(opts *bind.FilterOpts) (*EndpointcontractTransferReleasedIterator, error) {

	logs, sub, err := _Endpointcontract.contract.FilterLogs(opts, "TransferReleased")
	if err != nil {
		return nil, err
	}
	return &EndpointcontractTransferReleasedIterator{contract: _Endpointcontract.contract, event: "TransferReleased", logs: logs, sub: sub}, nil
}

// WatchTransferReleased is a free log subscription operation binding the contract event 0x0b3206d324e5269b2561359f447fd104c22f9cf582bb3886656a6bfd736c4e15.
//
// Solidity: event TransferReleased(address receiver, uint256 amount)
func (_Endpointcontract *EndpointcontractFilterer) WatchTransferReleased(opts *bind.WatchOpts, sink chan<- *EndpointcontractTransferReleased) (event.Subscription, error) {

	logs, sub, err := _Endpointcontract.contract.WatchLogs(opts, "TransferReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointcontractTransferReleased)
				if err := _Endpointcontract.contract.UnpackLog(event, "TransferReleased", log); err != nil {
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

// ParseTransferReleased is a log parse operation binding the contract event 0x0b3206d324e5269b2561359f447fd104c22f9cf582bb3886656a6bfd736c4e15.
//
// Solidity: event TransferReleased(address receiver, uint256 amount)
func (_Endpointcontract *EndpointcontractFilterer) ParseTransferReleased(log types.Log) (*EndpointcontractTransferReleased, error) {
	event := new(EndpointcontractTransferReleased)
	if err := _Endpointcontract.contract.UnpackLog(event, "TransferReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointcontractTransferRequestedIterator is returned from FilterTransferRequested and is used to iterate over the raw logs and unpacked data for TransferRequested events raised by the Endpointcontract contract.
type EndpointcontractTransferRequestedIterator struct {
	Event *EndpointcontractTransferRequested // Event containing the contract specifics and raw log

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
func (it *EndpointcontractTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointcontractTransferRequested)
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
		it.Event = new(EndpointcontractTransferRequested)
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
func (it *EndpointcontractTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointcontractTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointcontractTransferRequested represents a TransferRequested event raised by the Endpointcontract contract.
type EndpointcontractTransferRequested struct {
	Endpoint Endpoint
	Sender   common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferRequested is a free log retrieval operation binding the contract event 0xb88260dad45db6f5739be2ac520f6360d808a7b5ec21b2b3367e76e39af006bc.
//
// Solidity: event TransferRequested((uint64,address) endpoint, address sender, address receiver, uint256 amount)
func (_Endpointcontract *EndpointcontractFilterer) FilterTransferRequested(opts *bind.FilterOpts) (*EndpointcontractTransferRequestedIterator, error) {

	logs, sub, err := _Endpointcontract.contract.FilterLogs(opts, "TransferRequested")
	if err != nil {
		return nil, err
	}
	return &EndpointcontractTransferRequestedIterator{contract: _Endpointcontract.contract, event: "TransferRequested", logs: logs, sub: sub}, nil
}

// WatchTransferRequested is a free log subscription operation binding the contract event 0xb88260dad45db6f5739be2ac520f6360d808a7b5ec21b2b3367e76e39af006bc.
//
// Solidity: event TransferRequested((uint64,address) endpoint, address sender, address receiver, uint256 amount)
func (_Endpointcontract *EndpointcontractFilterer) WatchTransferRequested(opts *bind.WatchOpts, sink chan<- *EndpointcontractTransferRequested) (event.Subscription, error) {

	logs, sub, err := _Endpointcontract.contract.WatchLogs(opts, "TransferRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointcontractTransferRequested)
				if err := _Endpointcontract.contract.UnpackLog(event, "TransferRequested", log); err != nil {
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

// ParseTransferRequested is a log parse operation binding the contract event 0xb88260dad45db6f5739be2ac520f6360d808a7b5ec21b2b3367e76e39af006bc.
//
// Solidity: event TransferRequested((uint64,address) endpoint, address sender, address receiver, uint256 amount)
func (_Endpointcontract *EndpointcontractFilterer) ParseTransferRequested(log types.Log) (*EndpointcontractTransferRequested, error) {
	event := new(EndpointcontractTransferRequested)
	if err := _Endpointcontract.contract.UnpackLog(event, "TransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
