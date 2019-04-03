// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethash

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

// VotingContractABI is the input ABI used to generate the binding from.
const VotingContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"voterCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"canCreateBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governanceAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"canVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockMakerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddVoter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"address\"}],\"name\":\"RemovedVoter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddBlockMaker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"address\"}],\"name\":\"RemovedBlockMaker\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setVoteThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"getCanonHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addBlockMaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeBlockMaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isBlockMaker\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\"},{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VotingContract is an auto generated Go binding around an Ethereum contract.
type VotingContract struct {
	VotingContractCaller     // Read-only binding to the contract
	VotingContractTransactor // Write-only binding to the contract
	VotingContractFilterer   // Log filterer for contract events
}

// VotingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingContractSession struct {
	Contract     *VotingContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingContractCallerSession struct {
	Contract *VotingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VotingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingContractTransactorSession struct {
	Contract     *VotingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VotingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingContractRaw struct {
	Contract *VotingContract // Generic contract binding to access the raw methods on
}

// VotingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingContractCallerRaw struct {
	Contract *VotingContractCaller // Generic read-only contract binding to access the raw methods on
}

// VotingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingContractTransactorRaw struct {
	Contract *VotingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotingContract creates a new instance of VotingContract, bound to a specific deployed contract.
func NewVotingContract(address common.Address, backend bind.ContractBackend) (*VotingContract, error) {
	contract, err := bindVotingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotingContract{VotingContractCaller: VotingContractCaller{contract: contract}, VotingContractTransactor: VotingContractTransactor{contract: contract}, VotingContractFilterer: VotingContractFilterer{contract: contract}}, nil
}

// NewVotingContractCaller creates a new read-only instance of VotingContract, bound to a specific deployed contract.
func NewVotingContractCaller(address common.Address, caller bind.ContractCaller) (*VotingContractCaller, error) {
	contract, err := bindVotingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingContractCaller{contract: contract}, nil
}

// NewVotingContractTransactor creates a new write-only instance of VotingContract, bound to a specific deployed contract.
func NewVotingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingContractTransactor, error) {
	contract, err := bindVotingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingContractTransactor{contract: contract}, nil
}

// NewVotingContractFilterer creates a new log filterer instance of VotingContract, bound to a specific deployed contract.
func NewVotingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingContractFilterer, error) {
	contract, err := bindVotingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingContractFilterer{contract: contract}, nil
}

// bindVotingContract binds a generic wrapper to an already deployed contract.
func bindVotingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingContract *VotingContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VotingContract.Contract.VotingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingContract *VotingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingContract.Contract.VotingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingContract *VotingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingContract.Contract.VotingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingContract *VotingContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VotingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingContract *VotingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingContract *VotingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingContract.Contract.contract.Transact(opts, method, params...)
}

// BlockMakerCount is a free data retrieval call binding the contract method 0xcf528985.
//
// Solidity: function blockMakerCount() constant returns(uint256)
func (_VotingContract *VotingContractCaller) BlockMakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "blockMakerCount")
	return *ret0, err
}

// BlockMakerCount is a free data retrieval call binding the contract method 0xcf528985.
//
// Solidity: function blockMakerCount() constant returns(uint256)
func (_VotingContract *VotingContractSession) BlockMakerCount() (*big.Int, error) {
	return _VotingContract.Contract.BlockMakerCount(&_VotingContract.CallOpts)
}

// BlockMakerCount is a free data retrieval call binding the contract method 0xcf528985.
//
// Solidity: function blockMakerCount() constant returns(uint256)
func (_VotingContract *VotingContractCallerSession) BlockMakerCount() (*big.Int, error) {
	return _VotingContract.Contract.BlockMakerCount(&_VotingContract.CallOpts)
}

// CanCreateBlocks is a free data retrieval call binding the contract method 0x488099a6.
//
// Solidity: function canCreateBlocks( address) constant returns(bool)
func (_VotingContract *VotingContractCaller) CanCreateBlocks(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "canCreateBlocks", arg0)
	return *ret0, err
}

// CanCreateBlocks is a free data retrieval call binding the contract method 0x488099a6.
//
// Solidity: function canCreateBlocks( address) constant returns(bool)
func (_VotingContract *VotingContractSession) CanCreateBlocks(arg0 common.Address) (bool, error) {
	return _VotingContract.Contract.CanCreateBlocks(&_VotingContract.CallOpts, arg0)
}

// CanCreateBlocks is a free data retrieval call binding the contract method 0x488099a6.
//
// Solidity: function canCreateBlocks( address) constant returns(bool)
func (_VotingContract *VotingContractCallerSession) CanCreateBlocks(arg0 common.Address) (bool, error) {
	return _VotingContract.Contract.CanCreateBlocks(&_VotingContract.CallOpts, arg0)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote( address) constant returns(bool)
func (_VotingContract *VotingContractCaller) CanVote(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "canVote", arg0)
	return *ret0, err
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote( address) constant returns(bool)
func (_VotingContract *VotingContractSession) CanVote(arg0 common.Address) (bool, error) {
	return _VotingContract.Contract.CanVote(&_VotingContract.CallOpts, arg0)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote( address) constant returns(bool)
func (_VotingContract *VotingContractCallerSession) CanVote(arg0 common.Address) (bool, error) {
	return _VotingContract.Contract.CanVote(&_VotingContract.CallOpts, arg0)
}

// GetCanonHash is a free data retrieval call binding the contract method 0x559c390c.
//
// Solidity: function getCanonHash(height uint256) constant returns(bytes32)
func (_VotingContract *VotingContractCaller) GetCanonHash(opts *bind.CallOpts, height *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "getCanonHash", height)
	return *ret0, err
}

// GetCanonHash is a free data retrieval call binding the contract method 0x559c390c.
//
// Solidity: function getCanonHash(height uint256) constant returns(bytes32)
func (_VotingContract *VotingContractSession) GetCanonHash(height *big.Int) ([32]byte, error) {
	return _VotingContract.Contract.GetCanonHash(&_VotingContract.CallOpts, height)
}

// GetCanonHash is a free data retrieval call binding the contract method 0x559c390c.
//
// Solidity: function getCanonHash(height uint256) constant returns(bytes32)
func (_VotingContract *VotingContractCallerSession) GetCanonHash(height *big.Int) ([32]byte, error) {
	return _VotingContract.Contract.GetCanonHash(&_VotingContract.CallOpts, height)
}

// GetEntry is a free data retrieval call binding the contract method 0x98ba676d.
//
// Solidity: function getEntry(height uint256, n uint256) constant returns(bytes32)
func (_VotingContract *VotingContractCaller) GetEntry(opts *bind.CallOpts, height *big.Int, n *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "getEntry", height, n)
	return *ret0, err
}

// GetEntry is a free data retrieval call binding the contract method 0x98ba676d.
//
// Solidity: function getEntry(height uint256, n uint256) constant returns(bytes32)
func (_VotingContract *VotingContractSession) GetEntry(height *big.Int, n *big.Int) ([32]byte, error) {
	return _VotingContract.Contract.GetEntry(&_VotingContract.CallOpts, height, n)
}

// GetEntry is a free data retrieval call binding the contract method 0x98ba676d.
//
// Solidity: function getEntry(height uint256, n uint256) constant returns(bytes32)
func (_VotingContract *VotingContractCallerSession) GetEntry(height *big.Int, n *big.Int) ([32]byte, error) {
	return _VotingContract.Contract.GetEntry(&_VotingContract.CallOpts, height, n)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_VotingContract *VotingContractCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_VotingContract *VotingContractSession) GetSize() (*big.Int, error) {
	return _VotingContract.Contract.GetSize(&_VotingContract.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_VotingContract *VotingContractCallerSession) GetSize() (*big.Int, error) {
	return _VotingContract.Contract.GetSize(&_VotingContract.CallOpts)
}

// GovernanceAddress is a free data retrieval call binding the contract method 0x795053d3.
//
// Solidity: function governanceAddress() constant returns(address)
func (_VotingContract *VotingContractCaller) GovernanceAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "governanceAddress")
	return *ret0, err
}

// GovernanceAddress is a free data retrieval call binding the contract method 0x795053d3.
//
// Solidity: function governanceAddress() constant returns(address)
func (_VotingContract *VotingContractSession) GovernanceAddress() (common.Address, error) {
	return _VotingContract.Contract.GovernanceAddress(&_VotingContract.CallOpts)
}

// GovernanceAddress is a free data retrieval call binding the contract method 0x795053d3.
//
// Solidity: function governanceAddress() constant returns(address)
func (_VotingContract *VotingContractCallerSession) GovernanceAddress() (common.Address, error) {
	return _VotingContract.Contract.GovernanceAddress(&_VotingContract.CallOpts)
}

// IsBlockMaker is a free data retrieval call binding the contract method 0xe814d1c7.
//
// Solidity: function isBlockMaker(addr address) constant returns(bool)
func (_VotingContract *VotingContractCaller) IsBlockMaker(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "isBlockMaker", addr)
	return *ret0, err
}

// IsBlockMaker is a free data retrieval call binding the contract method 0xe814d1c7.
//
// Solidity: function isBlockMaker(addr address) constant returns(bool)
func (_VotingContract *VotingContractSession) IsBlockMaker(addr common.Address) (bool, error) {
	return _VotingContract.Contract.IsBlockMaker(&_VotingContract.CallOpts, addr)
}

// IsBlockMaker is a free data retrieval call binding the contract method 0xe814d1c7.
//
// Solidity: function isBlockMaker(addr address) constant returns(bool)
func (_VotingContract *VotingContractCallerSession) IsBlockMaker(addr common.Address) (bool, error) {
	return _VotingContract.Contract.IsBlockMaker(&_VotingContract.CallOpts, addr)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(bool)
func (_VotingContract *VotingContractCaller) IsVoter(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "isVoter", addr)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(bool)
func (_VotingContract *VotingContractSession) IsVoter(addr common.Address) (bool, error) {
	return _VotingContract.Contract.IsVoter(&_VotingContract.CallOpts, addr)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(bool)
func (_VotingContract *VotingContractCallerSession) IsVoter(addr common.Address) (bool, error) {
	return _VotingContract.Contract.IsVoter(&_VotingContract.CallOpts, addr)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() constant returns(uint256)
func (_VotingContract *VotingContractCaller) VoteThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "voteThreshold")
	return *ret0, err
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() constant returns(uint256)
func (_VotingContract *VotingContractSession) VoteThreshold() (*big.Int, error) {
	return _VotingContract.Contract.VoteThreshold(&_VotingContract.CallOpts)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() constant returns(uint256)
func (_VotingContract *VotingContractCallerSession) VoteThreshold() (*big.Int, error) {
	return _VotingContract.Contract.VoteThreshold(&_VotingContract.CallOpts)
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() constant returns(uint256)
func (_VotingContract *VotingContractCaller) VoterCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "voterCount")
	return *ret0, err
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() constant returns(uint256)
func (_VotingContract *VotingContractSession) VoterCount() (*big.Int, error) {
	return _VotingContract.Contract.VoterCount(&_VotingContract.CallOpts)
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() constant returns(uint256)
func (_VotingContract *VotingContractCallerSession) VoterCount() (*big.Int, error) {
	return _VotingContract.Contract.VoterCount(&_VotingContract.CallOpts)
}

// AddBlockMaker is a paid mutator transaction binding the contract method 0x72a571fc.
//
// Solidity: function addBlockMaker(addr address) returns()
func (_VotingContract *VotingContractTransactor) AddBlockMaker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "addBlockMaker", addr)
}

// AddBlockMaker is a paid mutator transaction binding the contract method 0x72a571fc.
//
// Solidity: function addBlockMaker(addr address) returns()
func (_VotingContract *VotingContractSession) AddBlockMaker(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.AddBlockMaker(&_VotingContract.TransactOpts, addr)
}

// AddBlockMaker is a paid mutator transaction binding the contract method 0x72a571fc.
//
// Solidity: function addBlockMaker(addr address) returns()
func (_VotingContract *VotingContractTransactorSession) AddBlockMaker(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.AddBlockMaker(&_VotingContract.TransactOpts, addr)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(addr address) returns()
func (_VotingContract *VotingContractTransactor) AddVoter(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "addVoter", addr)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(addr address) returns()
func (_VotingContract *VotingContractSession) AddVoter(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.AddVoter(&_VotingContract.TransactOpts, addr)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(addr address) returns()
func (_VotingContract *VotingContractTransactorSession) AddVoter(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.AddVoter(&_VotingContract.TransactOpts, addr)
}

// RemoveBlockMaker is a paid mutator transaction binding the contract method 0x284d163c.
//
// Solidity: function removeBlockMaker(addr address) returns()
func (_VotingContract *VotingContractTransactor) RemoveBlockMaker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "removeBlockMaker", addr)
}

// RemoveBlockMaker is a paid mutator transaction binding the contract method 0x284d163c.
//
// Solidity: function removeBlockMaker(addr address) returns()
func (_VotingContract *VotingContractSession) RemoveBlockMaker(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.RemoveBlockMaker(&_VotingContract.TransactOpts, addr)
}

// RemoveBlockMaker is a paid mutator transaction binding the contract method 0x284d163c.
//
// Solidity: function removeBlockMaker(addr address) returns()
func (_VotingContract *VotingContractTransactorSession) RemoveBlockMaker(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.RemoveBlockMaker(&_VotingContract.TransactOpts, addr)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(addr address) returns()
func (_VotingContract *VotingContractTransactor) RemoveVoter(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "removeVoter", addr)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(addr address) returns()
func (_VotingContract *VotingContractSession) RemoveVoter(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.RemoveVoter(&_VotingContract.TransactOpts, addr)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(addr address) returns()
func (_VotingContract *VotingContractTransactorSession) RemoveVoter(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.RemoveVoter(&_VotingContract.TransactOpts, addr)
}

// SetVoteThreshold is a paid mutator transaction binding the contract method 0x12909485.
//
// Solidity: function setVoteThreshold(threshold uint256) returns()
func (_VotingContract *VotingContractTransactor) SetVoteThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "setVoteThreshold", threshold)
}

// SetVoteThreshold is a paid mutator transaction binding the contract method 0x12909485.
//
// Solidity: function setVoteThreshold(threshold uint256) returns()
func (_VotingContract *VotingContractSession) SetVoteThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _VotingContract.Contract.SetVoteThreshold(&_VotingContract.TransactOpts, threshold)
}

// SetVoteThreshold is a paid mutator transaction binding the contract method 0x12909485.
//
// Solidity: function setVoteThreshold(threshold uint256) returns()
func (_VotingContract *VotingContractTransactorSession) SetVoteThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _VotingContract.Contract.SetVoteThreshold(&_VotingContract.TransactOpts, threshold)
}

// Vote is a paid mutator transaction binding the contract method 0x68bb8bb6.
//
// Solidity: function vote(height uint256, hash bytes32) returns()
func (_VotingContract *VotingContractTransactor) Vote(opts *bind.TransactOpts, height *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "vote", height, hash)
}

// Vote is a paid mutator transaction binding the contract method 0x68bb8bb6.
//
// Solidity: function vote(height uint256, hash bytes32) returns()
func (_VotingContract *VotingContractSession) Vote(height *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _VotingContract.Contract.Vote(&_VotingContract.TransactOpts, height, hash)
}

// Vote is a paid mutator transaction binding the contract method 0x68bb8bb6.
//
// Solidity: function vote(height uint256, hash bytes32) returns()
func (_VotingContract *VotingContractTransactorSession) Vote(height *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _VotingContract.Contract.Vote(&_VotingContract.TransactOpts, height, hash)
}

// VotingContractAddBlockMakerIterator is returned from FilterAddBlockMaker and is used to iterate over the raw logs and unpacked data for AddBlockMaker events raised by the VotingContract contract.
type VotingContractAddBlockMakerIterator struct {
	Event *VotingContractAddBlockMaker // Event containing the contract specifics and raw log

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
func (it *VotingContractAddBlockMakerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingContractAddBlockMaker)
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
		it.Event = new(VotingContractAddBlockMaker)
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
func (it *VotingContractAddBlockMakerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingContractAddBlockMakerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingContractAddBlockMaker represents a AddBlockMaker event raised by the VotingContract contract.
type VotingContractAddBlockMaker struct {
	common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddBlockMaker is a free log retrieval operation binding the contract event 0x1a4ce6942f7aa91856332e618fc90159f13a340611a308f5d7327ba0707e5685.
//
// Solidity: e AddBlockMaker( address)
func (_VotingContract *VotingContractFilterer) FilterAddBlockMaker(opts *bind.FilterOpts) (*VotingContractAddBlockMakerIterator, error) {

	logs, sub, err := _VotingContract.contract.FilterLogs(opts, "AddBlockMaker")
	if err != nil {
		return nil, err
	}
	return &VotingContractAddBlockMakerIterator{contract: _VotingContract.contract, event: "AddBlockMaker", logs: logs, sub: sub}, nil
}

// WatchAddBlockMaker is a free log subscription operation binding the contract event 0x1a4ce6942f7aa91856332e618fc90159f13a340611a308f5d7327ba0707e5685.
//
// Solidity: e AddBlockMaker( address)
func (_VotingContract *VotingContractFilterer) WatchAddBlockMaker(opts *bind.WatchOpts, sink chan<- *VotingContractAddBlockMaker) (event.Subscription, error) {

	logs, sub, err := _VotingContract.contract.WatchLogs(opts, "AddBlockMaker")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingContractAddBlockMaker)
				if err := _VotingContract.contract.UnpackLog(event, "AddBlockMaker", log); err != nil {
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

// VotingContractAddVoterIterator is returned from FilterAddVoter and is used to iterate over the raw logs and unpacked data for AddVoter events raised by the VotingContract contract.
type VotingContractAddVoterIterator struct {
	Event *VotingContractAddVoter // Event containing the contract specifics and raw log

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
func (it *VotingContractAddVoterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingContractAddVoter)
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
		it.Event = new(VotingContractAddVoter)
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
func (it *VotingContractAddVoterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingContractAddVoterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingContractAddVoter represents a AddVoter event raised by the VotingContract contract.
type VotingContractAddVoter struct {
	common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddVoter is a free log retrieval operation binding the contract event 0x0ad2eca75347acd5160276fe4b5dad46987e4ff4af9e574195e3e9bc15d7e0ff.
//
// Solidity: e AddVoter( address)
func (_VotingContract *VotingContractFilterer) FilterAddVoter(opts *bind.FilterOpts) (*VotingContractAddVoterIterator, error) {

	logs, sub, err := _VotingContract.contract.FilterLogs(opts, "AddVoter")
	if err != nil {
		return nil, err
	}
	return &VotingContractAddVoterIterator{contract: _VotingContract.contract, event: "AddVoter", logs: logs, sub: sub}, nil
}

// WatchAddVoter is a free log subscription operation binding the contract event 0x0ad2eca75347acd5160276fe4b5dad46987e4ff4af9e574195e3e9bc15d7e0ff.
//
// Solidity: e AddVoter( address)
func (_VotingContract *VotingContractFilterer) WatchAddVoter(opts *bind.WatchOpts, sink chan<- *VotingContractAddVoter) (event.Subscription, error) {

	logs, sub, err := _VotingContract.contract.WatchLogs(opts, "AddVoter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingContractAddVoter)
				if err := _VotingContract.contract.UnpackLog(event, "AddVoter", log); err != nil {
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

// VotingContractRemovedBlockMakerIterator is returned from FilterRemovedBlockMaker and is used to iterate over the raw logs and unpacked data for RemovedBlockMaker events raised by the VotingContract contract.
type VotingContractRemovedBlockMakerIterator struct {
	Event *VotingContractRemovedBlockMaker // Event containing the contract specifics and raw log

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
func (it *VotingContractRemovedBlockMakerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingContractRemovedBlockMaker)
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
		it.Event = new(VotingContractRemovedBlockMaker)
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
func (it *VotingContractRemovedBlockMakerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingContractRemovedBlockMakerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingContractRemovedBlockMaker represents a RemovedBlockMaker event raised by the VotingContract contract.
type VotingContractRemovedBlockMaker struct {
	common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlockMaker is a free log retrieval operation binding the contract event 0x8cee3054364d6799f1c8962580ad61273d9d38ca1ff26516bd1ad23c099a6022.
//
// Solidity: e RemovedBlockMaker( address)
func (_VotingContract *VotingContractFilterer) FilterRemovedBlockMaker(opts *bind.FilterOpts) (*VotingContractRemovedBlockMakerIterator, error) {

	logs, sub, err := _VotingContract.contract.FilterLogs(opts, "RemovedBlockMaker")
	if err != nil {
		return nil, err
	}
	return &VotingContractRemovedBlockMakerIterator{contract: _VotingContract.contract, event: "RemovedBlockMaker", logs: logs, sub: sub}, nil
}

// WatchRemovedBlockMaker is a free log subscription operation binding the contract event 0x8cee3054364d6799f1c8962580ad61273d9d38ca1ff26516bd1ad23c099a6022.
//
// Solidity: e RemovedBlockMaker( address)
func (_VotingContract *VotingContractFilterer) WatchRemovedBlockMaker(opts *bind.WatchOpts, sink chan<- *VotingContractRemovedBlockMaker) (event.Subscription, error) {

	logs, sub, err := _VotingContract.contract.WatchLogs(opts, "RemovedBlockMaker")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingContractRemovedBlockMaker)
				if err := _VotingContract.contract.UnpackLog(event, "RemovedBlockMaker", log); err != nil {
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

// VotingContractRemovedVoterIterator is returned from FilterRemovedVoter and is used to iterate over the raw logs and unpacked data for RemovedVoter events raised by the VotingContract contract.
type VotingContractRemovedVoterIterator struct {
	Event *VotingContractRemovedVoter // Event containing the contract specifics and raw log

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
func (it *VotingContractRemovedVoterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingContractRemovedVoter)
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
		it.Event = new(VotingContractRemovedVoter)
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
func (it *VotingContractRemovedVoterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingContractRemovedVoterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingContractRemovedVoter represents a RemovedVoter event raised by the VotingContract contract.
type VotingContractRemovedVoter struct {
	common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemovedVoter is a free log retrieval operation binding the contract event 0x183393fc5cffbfc7d03d623966b85f76b9430f42d3aada2ac3f3deabc78899e8.
//
// Solidity: e RemovedVoter( address)
func (_VotingContract *VotingContractFilterer) FilterRemovedVoter(opts *bind.FilterOpts) (*VotingContractRemovedVoterIterator, error) {

	logs, sub, err := _VotingContract.contract.FilterLogs(opts, "RemovedVoter")
	if err != nil {
		return nil, err
	}
	return &VotingContractRemovedVoterIterator{contract: _VotingContract.contract, event: "RemovedVoter", logs: logs, sub: sub}, nil
}

// WatchRemovedVoter is a free log subscription operation binding the contract event 0x183393fc5cffbfc7d03d623966b85f76b9430f42d3aada2ac3f3deabc78899e8.
//
// Solidity: e RemovedVoter( address)
func (_VotingContract *VotingContractFilterer) WatchRemovedVoter(opts *bind.WatchOpts, sink chan<- *VotingContractRemovedVoter) (event.Subscription, error) {

	logs, sub, err := _VotingContract.contract.WatchLogs(opts, "RemovedVoter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingContractRemovedVoter)
				if err := _VotingContract.contract.UnpackLog(event, "RemovedVoter", log); err != nil {
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

// VotingContractVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the VotingContract contract.
type VotingContractVoteIterator struct {
	Event *VotingContractVote // Event containing the contract specifics and raw log

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
func (it *VotingContractVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingContractVote)
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
		it.Event = new(VotingContractVote)
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
func (it *VotingContractVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingContractVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingContractVote represents a Vote event raised by the VotingContract contract.
type VotingContractVote struct {
	Sender      common.Address
	BlockNumber *big.Int
	BlockHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x3d03ba7f4b5227cdb385f2610906e5bcee147171603ec40005b30915ad20e258.
//
// Solidity: e Vote(sender indexed address, blockNumber uint256, blockHash bytes32)
func (_VotingContract *VotingContractFilterer) FilterVote(opts *bind.FilterOpts, sender []common.Address) (*VotingContractVoteIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VotingContract.contract.FilterLogs(opts, "Vote", senderRule)
	if err != nil {
		return nil, err
	}
	return &VotingContractVoteIterator{contract: _VotingContract.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x3d03ba7f4b5227cdb385f2610906e5bcee147171603ec40005b30915ad20e258.
//
// Solidity: e Vote(sender indexed address, blockNumber uint256, blockHash bytes32)
func (_VotingContract *VotingContractFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *VotingContractVote, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VotingContract.contract.WatchLogs(opts, "Vote", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingContractVote)
				if err := _VotingContract.contract.UnpackLog(event, "Vote", log); err != nil {
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
