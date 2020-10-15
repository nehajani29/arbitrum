// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChallengeFactoryABI is the input ABI used to generate the binding from.
const ChallengeFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inboxTopChallengeTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executionChallengeTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oneStepProofAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oneStepProofAddress2\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_TYPE_STR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"executionChallengeTemplate\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"generateCloneAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inboxTopChallengeTemplate\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oneStepProofAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oneStepProofAddress2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeFactoryFuncSigs = map[string]string{
	"e6fcd194": "INVALID_TYPE_STR()",
	"865da1c2": "createChallenge(address,address,uint256,bytes32,uint256)",
	"9b5f4dfc": "executionChallengeTemplate()",
	"729406c8": "generateCloneAddress(address,address,uint256)",
	"f089f1c2": "inboxTopChallengeTemplate()",
	"52ddf4a5": "oneStepProofAddress()",
	"4930fdc5": "oneStepProofAddress2()",
}

// ChallengeFactoryBin is the compiled bytecode used for deploying new contracts.
var ChallengeFactoryBin = "0x608060405234801561001057600080fd5b506040516107623803806107628339818101604052608081101561003357600080fd5b50805160208201516040830151606090930151600080546001600160a01b039485166001600160a01b03199182161790915560018054938516938216939093179092556002805494841694831694909417909355600380549290931691161790556106bf806100a36000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063865da1c21161005b578063865da1c2146100e45780639b5f4dfc14610126578063e6fcd1941461012e578063f089f1c2146101ab5761007d565b80634930fdc51461008257806352ddf4a5146100a6578063729406c8146100ae575b600080fd5b61008a6101b3565b604080516001600160a01b039092168252519081900360200190f35b61008a6101c2565b61008a600480360360608110156100c457600080fd5b506001600160a01b038135811691602081013590911690604001356101d1565b61008a600480360360a08110156100fa57600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060800135610268565b61008a61038e565b61013661039d565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610170578181015183820152602001610158565b50505050905090810190601f16801561019d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61008a6103c5565b6003546001600160a01b031681565b6002546001600160a01b031681565b60006001600160f81b0319306101e786866103d4565b6101f86101f386610421565b610511565b60405160200180856001600160f81b0319166001600160f81b0319168152600101846001600160a01b03166001600160a01b031660601b81526014018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c90509392505050565b60008061027483610421565b9050600061028182610558565b604080516301568f2760e11b81523360048201526001600160a01b038b811660248301528a81166044830152606482018a9052608482018990529151929350908316916302ad1e4e9160a48082019260009290919082900301818387803b1580156102eb57600080fd5b505af11580156102ff573d6000803e3d6000fd5b5050505061030b610680565b84141561038357600254600354604080516373ee855360e11b81526001600160a01b0393841660048201529183166024830152519183169163e7dd0aa69160448082019260009290919082900301818387803b15801561036a57600080fd5b505af115801561037e573d6000803e3d6000fd5b505050505b979650505050505050565b6001546001600160a01b031681565b6040518060400160405280600c81526020016b494e56414c49445f5459504560a01b81525081565b6000546001600160a01b031681565b604080516bffffffffffffffffffffffff19606094851b811660208084019190915293851b1660348201523390931b60488401528051603c818503018152605c9093019052815191012090565b600061042b610685565b82141561044457506000546001600160a01b031661050c565b61044c610680565b82141561046557506001546001600160a01b031661050c565b60408051808201909152600c81526b494e56414c49445f5459504560a01b602082015260405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156104d15781810151838201526020016104b9565b50505050905090810190601f1680156104fe5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b919050565b604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b815260609190911b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037902090565b6000816001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b15801561059357600080fd5b505afa1580156105a7573d6000803e3d6000fd5b505050506040513d60208110156105bd57600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b60208201529061062e5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104d15781810151838201526020016104b9565b5060008260601b9050604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528160148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f0949350505050565b600190565b60009056fea265627a7a723158203f480ba88e50ab2bfde7100501ce755a073995f218abbc0b6711b3cfd0052cb164736f6c63430005110032"

// DeployChallengeFactory deploys a new Ethereum contract, binding an instance of ChallengeFactory to it.
func DeployChallengeFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _inboxTopChallengeTemplate common.Address, _executionChallengeTemplate common.Address, _oneStepProofAddress common.Address, _oneStepProofAddress2 common.Address) (common.Address, *types.Transaction, *ChallengeFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeFactoryBin), backend, _inboxTopChallengeTemplate, _executionChallengeTemplate, _oneStepProofAddress, _oneStepProofAddress2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeFactory{ChallengeFactoryCaller: ChallengeFactoryCaller{contract: contract}, ChallengeFactoryTransactor: ChallengeFactoryTransactor{contract: contract}, ChallengeFactoryFilterer: ChallengeFactoryFilterer{contract: contract}}, nil
}

// ChallengeFactory is an auto generated Go binding around an Ethereum contract.
type ChallengeFactory struct {
	ChallengeFactoryCaller     // Read-only binding to the contract
	ChallengeFactoryTransactor // Write-only binding to the contract
	ChallengeFactoryFilterer   // Log filterer for contract events
}

// ChallengeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeFactorySession struct {
	Contract     *ChallengeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeFactoryCallerSession struct {
	Contract *ChallengeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChallengeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeFactoryTransactorSession struct {
	Contract     *ChallengeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChallengeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeFactoryRaw struct {
	Contract *ChallengeFactory // Generic contract binding to access the raw methods on
}

// ChallengeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeFactoryCallerRaw struct {
	Contract *ChallengeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeFactoryTransactorRaw struct {
	Contract *ChallengeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeFactory creates a new instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactory(address common.Address, backend bind.ContractBackend) (*ChallengeFactory, error) {
	contract, err := bindChallengeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactory{ChallengeFactoryCaller: ChallengeFactoryCaller{contract: contract}, ChallengeFactoryTransactor: ChallengeFactoryTransactor{contract: contract}, ChallengeFactoryFilterer: ChallengeFactoryFilterer{contract: contract}}, nil
}

// NewChallengeFactoryCaller creates a new read-only instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactoryCaller(address common.Address, caller bind.ContractCaller) (*ChallengeFactoryCaller, error) {
	contract, err := bindChallengeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactoryCaller{contract: contract}, nil
}

// NewChallengeFactoryTransactor creates a new write-only instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeFactoryTransactor, error) {
	contract, err := bindChallengeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactoryTransactor{contract: contract}, nil
}

// NewChallengeFactoryFilterer creates a new log filterer instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFactoryFilterer, error) {
	contract, err := bindChallengeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactoryFilterer{contract: contract}, nil
}

// bindChallengeFactory binds a generic wrapper to an already deployed contract.
func bindChallengeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseChallengeFactoryABI parses the ABI
func ParseChallengeFactoryABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeFactoryABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeFactory *ChallengeFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeFactory.Contract.ChallengeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeFactory *ChallengeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.ChallengeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeFactory *ChallengeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.ChallengeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeFactory *ChallengeFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeFactory *ChallengeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeFactory *ChallengeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.contract.Transact(opts, method, params...)
}

// INVALIDTYPESTR is a free data retrieval call binding the contract method 0xe6fcd194.
//
// Solidity: function INVALID_TYPE_STR() constant returns(string)
func (_ChallengeFactory *ChallengeFactoryCaller) INVALIDTYPESTR(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "INVALID_TYPE_STR")
	return *ret0, err
}

// INVALIDTYPESTR is a free data retrieval call binding the contract method 0xe6fcd194.
//
// Solidity: function INVALID_TYPE_STR() constant returns(string)
func (_ChallengeFactory *ChallengeFactorySession) INVALIDTYPESTR() (string, error) {
	return _ChallengeFactory.Contract.INVALIDTYPESTR(&_ChallengeFactory.CallOpts)
}

// INVALIDTYPESTR is a free data retrieval call binding the contract method 0xe6fcd194.
//
// Solidity: function INVALID_TYPE_STR() constant returns(string)
func (_ChallengeFactory *ChallengeFactoryCallerSession) INVALIDTYPESTR() (string, error) {
	return _ChallengeFactory.Contract.INVALIDTYPESTR(&_ChallengeFactory.CallOpts)
}

// ExecutionChallengeTemplate is a free data retrieval call binding the contract method 0x9b5f4dfc.
//
// Solidity: function executionChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) ExecutionChallengeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "executionChallengeTemplate")
	return *ret0, err
}

// ExecutionChallengeTemplate is a free data retrieval call binding the contract method 0x9b5f4dfc.
//
// Solidity: function executionChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) ExecutionChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.ExecutionChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// ExecutionChallengeTemplate is a free data retrieval call binding the contract method 0x9b5f4dfc.
//
// Solidity: function executionChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) ExecutionChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.ExecutionChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) GenerateCloneAddress(opts *bind.CallOpts, asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "generateCloneAddress", asserter, challenger, challengeType)
	return *ret0, err
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) GenerateCloneAddress(asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	return _ChallengeFactory.Contract.GenerateCloneAddress(&_ChallengeFactory.CallOpts, asserter, challenger, challengeType)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) GenerateCloneAddress(asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	return _ChallengeFactory.Contract.GenerateCloneAddress(&_ChallengeFactory.CallOpts, asserter, challenger, challengeType)
}

// InboxTopChallengeTemplate is a free data retrieval call binding the contract method 0xf089f1c2.
//
// Solidity: function inboxTopChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) InboxTopChallengeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "inboxTopChallengeTemplate")
	return *ret0, err
}

// InboxTopChallengeTemplate is a free data retrieval call binding the contract method 0xf089f1c2.
//
// Solidity: function inboxTopChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) InboxTopChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.InboxTopChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// InboxTopChallengeTemplate is a free data retrieval call binding the contract method 0xf089f1c2.
//
// Solidity: function inboxTopChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) InboxTopChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.InboxTopChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// OneStepProofAddress is a free data retrieval call binding the contract method 0x52ddf4a5.
//
// Solidity: function oneStepProofAddress() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) OneStepProofAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "oneStepProofAddress")
	return *ret0, err
}

// OneStepProofAddress is a free data retrieval call binding the contract method 0x52ddf4a5.
//
// Solidity: function oneStepProofAddress() constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) OneStepProofAddress() (common.Address, error) {
	return _ChallengeFactory.Contract.OneStepProofAddress(&_ChallengeFactory.CallOpts)
}

// OneStepProofAddress is a free data retrieval call binding the contract method 0x52ddf4a5.
//
// Solidity: function oneStepProofAddress() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) OneStepProofAddress() (common.Address, error) {
	return _ChallengeFactory.Contract.OneStepProofAddress(&_ChallengeFactory.CallOpts)
}

// OneStepProofAddress2 is a free data retrieval call binding the contract method 0x4930fdc5.
//
// Solidity: function oneStepProofAddress2() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) OneStepProofAddress2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "oneStepProofAddress2")
	return *ret0, err
}

// OneStepProofAddress2 is a free data retrieval call binding the contract method 0x4930fdc5.
//
// Solidity: function oneStepProofAddress2() constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) OneStepProofAddress2() (common.Address, error) {
	return _ChallengeFactory.Contract.OneStepProofAddress2(&_ChallengeFactory.CallOpts)
}

// OneStepProofAddress2 is a free data retrieval call binding the contract method 0x4930fdc5.
//
// Solidity: function oneStepProofAddress2() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) OneStepProofAddress2() (common.Address, error) {
	return _ChallengeFactory.Contract.OneStepProofAddress2(&_ChallengeFactory.CallOpts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactor) CreateChallenge(opts *bind.TransactOpts, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeFactory.contract.Transact(opts, "createChallenge", _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_ChallengeFactory *ChallengeFactorySession) CreateChallenge(_asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactorSession) CreateChallenge(_asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_ChallengeFactory *ChallengeFactoryFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _ChallengeFactory.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}
