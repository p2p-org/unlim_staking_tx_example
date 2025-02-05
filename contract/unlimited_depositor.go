// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package directstakingv3

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FeeRecipient is an auto generated low-level Go binding around an user-defined struct.
type FeeRecipient struct {
	BasisPoints *big.Int
	Recipient   common.Address
}

// DirectStakingDepositorV3MetaData contains all meta data concerning the DirectStakingDepositorV3 contract.
var DirectStakingDepositorV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeDistributorFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"P2pOrgUnlimitedEthDepositor__AmountOfParametersError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_client\",\"type\":\"address\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__CallerNotClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eip7251Enabler\",\"type\":\"address\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__CallerNotEip7251Enabler\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"P2pOrgUnlimitedEthDepositor__DoNotSendEthDirectlyHere\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"P2pOrgUnlimitedEthDepositor__Eip7251NotEnabledYet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmountPerValidatorInWei\",\"type\":\"uint256\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__EthAmountPerValidatorInWeiOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"P2pOrgUnlimitedEthDepositor__EtherValueError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__FailedToSendEth\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_passedPrefix\",\"type\":\"bytes1\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__IncorrectWithdrawalCredentialsPrefix\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__NoDepositToReject\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"P2pOrgUnlimitedEthDepositor__NoSmallDeposits\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_passedAddress\",\"type\":\"address\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__NotFactory\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_passedAddress\",\"type\":\"address\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__NotFeeDistributor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__ShouldNotBeRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"P2pOrgUnlimitedEthDepositor__ValidatorCountError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_expiration\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_now\",\"type\":\"uint40\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__WaitForExpiration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_eth2WithdrawalCredentials\",\"type\":\"bytes32\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__WithdrawalCredentialsBytesNotZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_feeDistributorInstance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_eth2WithdrawalCredentials\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"_expiration\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_ethAmountPerValidatorInWei\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__ClientEthAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"P2pOrgUnlimitedEthDepositor__Eip7251Enabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__Eth2Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__Eth2DepositCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__Eth2DepositInProgress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_feeDistributorInstance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"P2pOrgUnlimitedEthDepositor__ServiceRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_eth2WithdrawalCredentials\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"_ethAmountPerValidatorInWei\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"_referenceFeeDistributor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint96\",\"name\":\"basisPoints\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFeeRecipient\",\"name\":\"_clientConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint96\",\"name\":\"basisPoints\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFeeRecipient\",\"name\":\"_referrerConfig\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"addEth\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"depositId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feeDistributorInstance\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"}],\"name\":\"depositAmount\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"}],\"name\":\"depositExpiration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"}],\"name\":\"depositStatus\",\"outputs\":[{\"internalType\":\"enumClientDepositStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip7251Enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableEip7251\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_eth2WithdrawalCredentials\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"_ethAmountPerValidatorInWei\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"_feeDistributorInstance\",\"type\":\"address\"}],\"name\":\"getDepositId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_eth2WithdrawalCredentials\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"_ethAmountPerValidatorInWei\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"_referenceFeeDistributor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint96\",\"name\":\"basisPoints\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFeeRecipient\",\"name\":\"_clientConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint96\",\"name\":\"basisPoints\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFeeRecipient\",\"name\":\"_referrerConfig\",\"type\":\"tuple\"}],\"name\":\"getDepositId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_depositContract\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_feeDistributorFactory\",\"outputs\":[{\"internalType\":\"contractIFeeDistributorFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_eth2WithdrawalCredentials\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"_ethAmountPerValidatorInWei\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"_feeDistributorInstance\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"makeBeaconDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_eth2WithdrawalCredentials\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"_ethAmountPerValidatorInWei\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"_feeDistributorInstance\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"rejectService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DirectStakingDepositorV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use DirectStakingDepositorV3MetaData.ABI instead.
var DirectStakingDepositorV3ABI = DirectStakingDepositorV3MetaData.ABI

// DirectStakingDepositorV3 is an auto generated Go binding around an Ethereum contract.
type DirectStakingDepositorV3 struct {
	DirectStakingDepositorV3Caller     // Read-only binding to the contract
	DirectStakingDepositorV3Transactor // Write-only binding to the contract
	DirectStakingDepositorV3Filterer   // Log filterer for contract events
}

// DirectStakingDepositorV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type DirectStakingDepositorV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectStakingDepositorV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DirectStakingDepositorV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectStakingDepositorV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DirectStakingDepositorV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectStakingDepositorV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DirectStakingDepositorV3Session struct {
	Contract     *DirectStakingDepositorV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DirectStakingDepositorV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DirectStakingDepositorV3CallerSession struct {
	Contract *DirectStakingDepositorV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// DirectStakingDepositorV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DirectStakingDepositorV3TransactorSession struct {
	Contract     *DirectStakingDepositorV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// DirectStakingDepositorV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type DirectStakingDepositorV3Raw struct {
	Contract *DirectStakingDepositorV3 // Generic contract binding to access the raw methods on
}

// DirectStakingDepositorV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DirectStakingDepositorV3CallerRaw struct {
	Contract *DirectStakingDepositorV3Caller // Generic read-only contract binding to access the raw methods on
}

// DirectStakingDepositorV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DirectStakingDepositorV3TransactorRaw struct {
	Contract *DirectStakingDepositorV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDirectStakingDepositorV3 creates a new instance of DirectStakingDepositorV3, bound to a specific deployed contract.
func NewDirectStakingDepositorV3(address common.Address, backend bind.ContractBackend) (*DirectStakingDepositorV3, error) {
	contract, err := bindDirectStakingDepositorV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3{DirectStakingDepositorV3Caller: DirectStakingDepositorV3Caller{contract: contract}, DirectStakingDepositorV3Transactor: DirectStakingDepositorV3Transactor{contract: contract}, DirectStakingDepositorV3Filterer: DirectStakingDepositorV3Filterer{contract: contract}}, nil
}

// NewDirectStakingDepositorV3Caller creates a new read-only instance of DirectStakingDepositorV3, bound to a specific deployed contract.
func NewDirectStakingDepositorV3Caller(address common.Address, caller bind.ContractCaller) (*DirectStakingDepositorV3Caller, error) {
	contract, err := bindDirectStakingDepositorV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3Caller{contract: contract}, nil
}

// NewDirectStakingDepositorV3Transactor creates a new write-only instance of DirectStakingDepositorV3, bound to a specific deployed contract.
func NewDirectStakingDepositorV3Transactor(address common.Address, transactor bind.ContractTransactor) (*DirectStakingDepositorV3Transactor, error) {
	contract, err := bindDirectStakingDepositorV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3Transactor{contract: contract}, nil
}

// NewDirectStakingDepositorV3Filterer creates a new log filterer instance of DirectStakingDepositorV3, bound to a specific deployed contract.
func NewDirectStakingDepositorV3Filterer(address common.Address, filterer bind.ContractFilterer) (*DirectStakingDepositorV3Filterer, error) {
	contract, err := bindDirectStakingDepositorV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3Filterer{contract: contract}, nil
}

// bindDirectStakingDepositorV3 binds a generic wrapper to an already deployed contract.
func bindDirectStakingDepositorV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DirectStakingDepositorV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DirectStakingDepositorV3.Contract.DirectStakingDepositorV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.DirectStakingDepositorV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.DirectStakingDepositorV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DirectStakingDepositorV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.contract.Transact(opts, method, params...)
}

// DepositAmount is a free data retrieval call binding the contract method 0x01fecab6.
//
// Solidity: function depositAmount(bytes32 _depositId) view returns(uint112)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) DepositAmount(opts *bind.CallOpts, _depositId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "depositAmount", _depositId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositAmount is a free data retrieval call binding the contract method 0x01fecab6.
//
// Solidity: function depositAmount(bytes32 _depositId) view returns(uint112)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) DepositAmount(_depositId [32]byte) (*big.Int, error) {
	return _DirectStakingDepositorV3.Contract.DepositAmount(&_DirectStakingDepositorV3.CallOpts, _depositId)
}

// DepositAmount is a free data retrieval call binding the contract method 0x01fecab6.
//
// Solidity: function depositAmount(bytes32 _depositId) view returns(uint112)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) DepositAmount(_depositId [32]byte) (*big.Int, error) {
	return _DirectStakingDepositorV3.Contract.DepositAmount(&_DirectStakingDepositorV3.CallOpts, _depositId)
}

// DepositExpiration is a free data retrieval call binding the contract method 0x8d1be0cb.
//
// Solidity: function depositExpiration(bytes32 _depositId) view returns(uint40)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) DepositExpiration(opts *bind.CallOpts, _depositId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "depositExpiration", _depositId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositExpiration is a free data retrieval call binding the contract method 0x8d1be0cb.
//
// Solidity: function depositExpiration(bytes32 _depositId) view returns(uint40)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) DepositExpiration(_depositId [32]byte) (*big.Int, error) {
	return _DirectStakingDepositorV3.Contract.DepositExpiration(&_DirectStakingDepositorV3.CallOpts, _depositId)
}

// DepositExpiration is a free data retrieval call binding the contract method 0x8d1be0cb.
//
// Solidity: function depositExpiration(bytes32 _depositId) view returns(uint40)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) DepositExpiration(_depositId [32]byte) (*big.Int, error) {
	return _DirectStakingDepositorV3.Contract.DepositExpiration(&_DirectStakingDepositorV3.CallOpts, _depositId)
}

// DepositStatus is a free data retrieval call binding the contract method 0x7da9874f.
//
// Solidity: function depositStatus(bytes32 _depositId) view returns(uint8)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) DepositStatus(opts *bind.CallOpts, _depositId [32]byte) (uint8, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "depositStatus", _depositId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DepositStatus is a free data retrieval call binding the contract method 0x7da9874f.
//
// Solidity: function depositStatus(bytes32 _depositId) view returns(uint8)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) DepositStatus(_depositId [32]byte) (uint8, error) {
	return _DirectStakingDepositorV3.Contract.DepositStatus(&_DirectStakingDepositorV3.CallOpts, _depositId)
}

// DepositStatus is a free data retrieval call binding the contract method 0x7da9874f.
//
// Solidity: function depositStatus(bytes32 _depositId) view returns(uint8)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) DepositStatus(_depositId [32]byte) (uint8, error) {
	return _DirectStakingDepositorV3.Contract.DepositStatus(&_DirectStakingDepositorV3.CallOpts, _depositId)
}

// Eip7251Enabled is a free data retrieval call binding the contract method 0xc4ad9195.
//
// Solidity: function eip7251Enabled() view returns(bool)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) Eip7251Enabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "eip7251Enabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Eip7251Enabled is a free data retrieval call binding the contract method 0xc4ad9195.
//
// Solidity: function eip7251Enabled() view returns(bool)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) Eip7251Enabled() (bool, error) {
	return _DirectStakingDepositorV3.Contract.Eip7251Enabled(&_DirectStakingDepositorV3.CallOpts)
}

// Eip7251Enabled is a free data retrieval call binding the contract method 0xc4ad9195.
//
// Solidity: function eip7251Enabled() view returns(bool)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) Eip7251Enabled() (bool, error) {
	return _DirectStakingDepositorV3.Contract.Eip7251Enabled(&_DirectStakingDepositorV3.CallOpts)
}

// GetDepositId is a free data retrieval call binding the contract method 0x17c5d106.
//
// Solidity: function getDepositId(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _feeDistributorInstance) pure returns(bytes32)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) GetDepositId(opts *bind.CallOpts, _eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _feeDistributorInstance common.Address) ([32]byte, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "getDepositId", _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _feeDistributorInstance)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositId is a free data retrieval call binding the contract method 0x17c5d106.
//
// Solidity: function getDepositId(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _feeDistributorInstance) pure returns(bytes32)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) GetDepositId(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _feeDistributorInstance common.Address) ([32]byte, error) {
	return _DirectStakingDepositorV3.Contract.GetDepositId(&_DirectStakingDepositorV3.CallOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _feeDistributorInstance)
}

// GetDepositId is a free data retrieval call binding the contract method 0x17c5d106.
//
// Solidity: function getDepositId(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _feeDistributorInstance) pure returns(bytes32)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) GetDepositId(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _feeDistributorInstance common.Address) ([32]byte, error) {
	return _DirectStakingDepositorV3.Contract.GetDepositId(&_DirectStakingDepositorV3.CallOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _feeDistributorInstance)
}

// GetDepositId0 is a free data retrieval call binding the contract method 0x31745d91.
//
// Solidity: function getDepositId(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _referenceFeeDistributor, (uint96,address) _clientConfig, (uint96,address) _referrerConfig) view returns(bytes32)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) GetDepositId0(opts *bind.CallOpts, _eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _referenceFeeDistributor common.Address, _clientConfig FeeRecipient, _referrerConfig FeeRecipient) ([32]byte, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "getDepositId0", _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _referenceFeeDistributor, _clientConfig, _referrerConfig)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositId0 is a free data retrieval call binding the contract method 0x31745d91.
//
// Solidity: function getDepositId(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _referenceFeeDistributor, (uint96,address) _clientConfig, (uint96,address) _referrerConfig) view returns(bytes32)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) GetDepositId0(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _referenceFeeDistributor common.Address, _clientConfig FeeRecipient, _referrerConfig FeeRecipient) ([32]byte, error) {
	return _DirectStakingDepositorV3.Contract.GetDepositId0(&_DirectStakingDepositorV3.CallOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _referenceFeeDistributor, _clientConfig, _referrerConfig)
}

// GetDepositId0 is a free data retrieval call binding the contract method 0x31745d91.
//
// Solidity: function getDepositId(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _referenceFeeDistributor, (uint96,address) _clientConfig, (uint96,address) _referrerConfig) view returns(bytes32)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) GetDepositId0(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _referenceFeeDistributor common.Address, _clientConfig FeeRecipient, _referrerConfig FeeRecipient) ([32]byte, error) {
	return _DirectStakingDepositorV3.Contract.GetDepositId0(&_DirectStakingDepositorV3.CallOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _referenceFeeDistributor, _clientConfig, _referrerConfig)
}

// IDepositContract is a free data retrieval call binding the contract method 0x8a46d351.
//
// Solidity: function i_depositContract() view returns(address)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) IDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "i_depositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IDepositContract is a free data retrieval call binding the contract method 0x8a46d351.
//
// Solidity: function i_depositContract() view returns(address)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) IDepositContract() (common.Address, error) {
	return _DirectStakingDepositorV3.Contract.IDepositContract(&_DirectStakingDepositorV3.CallOpts)
}

// IDepositContract is a free data retrieval call binding the contract method 0x8a46d351.
//
// Solidity: function i_depositContract() view returns(address)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) IDepositContract() (common.Address, error) {
	return _DirectStakingDepositorV3.Contract.IDepositContract(&_DirectStakingDepositorV3.CallOpts)
}

// IFeeDistributorFactory is a free data retrieval call binding the contract method 0x7a49f5e6.
//
// Solidity: function i_feeDistributorFactory() view returns(address)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) IFeeDistributorFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "i_feeDistributorFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IFeeDistributorFactory is a free data retrieval call binding the contract method 0x7a49f5e6.
//
// Solidity: function i_feeDistributorFactory() view returns(address)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) IFeeDistributorFactory() (common.Address, error) {
	return _DirectStakingDepositorV3.Contract.IFeeDistributorFactory(&_DirectStakingDepositorV3.CallOpts)
}

// IFeeDistributorFactory is a free data retrieval call binding the contract method 0x7a49f5e6.
//
// Solidity: function i_feeDistributorFactory() view returns(address)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) IFeeDistributorFactory() (common.Address, error) {
	return _DirectStakingDepositorV3.Contract.IFeeDistributorFactory(&_DirectStakingDepositorV3.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DirectStakingDepositorV3.Contract.SupportsInterface(&_DirectStakingDepositorV3.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DirectStakingDepositorV3.Contract.SupportsInterface(&_DirectStakingDepositorV3.CallOpts, interfaceId)
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() view returns(uint256)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Caller) TotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DirectStakingDepositorV3.contract.Call(opts, &out, "totalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() view returns(uint256)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) TotalBalance() (*big.Int, error) {
	return _DirectStakingDepositorV3.Contract.TotalBalance(&_DirectStakingDepositorV3.CallOpts)
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() view returns(uint256)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3CallerSession) TotalBalance() (*big.Int, error) {
	return _DirectStakingDepositorV3.Contract.TotalBalance(&_DirectStakingDepositorV3.CallOpts)
}

// AddEth is a paid mutator transaction binding the contract method 0xa49b131b.
//
// Solidity: function addEth(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _referenceFeeDistributor, (uint96,address) _clientConfig, (uint96,address) _referrerConfig, bytes _extraData) payable returns(bytes32 depositId, address feeDistributorInstance)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Transactor) AddEth(opts *bind.TransactOpts, _eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _referenceFeeDistributor common.Address, _clientConfig FeeRecipient, _referrerConfig FeeRecipient, _extraData []byte) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.contract.Transact(opts, "addEth", _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _referenceFeeDistributor, _clientConfig, _referrerConfig, _extraData)
}

// AddEth is a paid mutator transaction binding the contract method 0xa49b131b.
//
// Solidity: function addEth(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _referenceFeeDistributor, (uint96,address) _clientConfig, (uint96,address) _referrerConfig, bytes _extraData) payable returns(bytes32 depositId, address feeDistributorInstance)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) AddEth(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _referenceFeeDistributor common.Address, _clientConfig FeeRecipient, _referrerConfig FeeRecipient, _extraData []byte) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.AddEth(&_DirectStakingDepositorV3.TransactOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _referenceFeeDistributor, _clientConfig, _referrerConfig, _extraData)
}

// AddEth is a paid mutator transaction binding the contract method 0xa49b131b.
//
// Solidity: function addEth(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _referenceFeeDistributor, (uint96,address) _clientConfig, (uint96,address) _referrerConfig, bytes _extraData) payable returns(bytes32 depositId, address feeDistributorInstance)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3TransactorSession) AddEth(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _referenceFeeDistributor common.Address, _clientConfig FeeRecipient, _referrerConfig FeeRecipient, _extraData []byte) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.AddEth(&_DirectStakingDepositorV3.TransactOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _referenceFeeDistributor, _clientConfig, _referrerConfig, _extraData)
}

// EnableEip7251 is a paid mutator transaction binding the contract method 0x28c2f044.
//
// Solidity: function enableEip7251() returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Transactor) EnableEip7251(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.contract.Transact(opts, "enableEip7251")
}

// EnableEip7251 is a paid mutator transaction binding the contract method 0x28c2f044.
//
// Solidity: function enableEip7251() returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) EnableEip7251() (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.EnableEip7251(&_DirectStakingDepositorV3.TransactOpts)
}

// EnableEip7251 is a paid mutator transaction binding the contract method 0x28c2f044.
//
// Solidity: function enableEip7251() returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3TransactorSession) EnableEip7251() (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.EnableEip7251(&_DirectStakingDepositorV3.TransactOpts)
}

// MakeBeaconDeposit is a paid mutator transaction binding the contract method 0x926b0e53.
//
// Solidity: function makeBeaconDeposit(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _feeDistributorInstance, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Transactor) MakeBeaconDeposit(opts *bind.TransactOpts, _eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _feeDistributorInstance common.Address, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.contract.Transact(opts, "makeBeaconDeposit", _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _feeDistributorInstance, _pubkeys, _signatures, _depositDataRoots)
}

// MakeBeaconDeposit is a paid mutator transaction binding the contract method 0x926b0e53.
//
// Solidity: function makeBeaconDeposit(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _feeDistributorInstance, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) MakeBeaconDeposit(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _feeDistributorInstance common.Address, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.MakeBeaconDeposit(&_DirectStakingDepositorV3.TransactOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _feeDistributorInstance, _pubkeys, _signatures, _depositDataRoots)
}

// MakeBeaconDeposit is a paid mutator transaction binding the contract method 0x926b0e53.
//
// Solidity: function makeBeaconDeposit(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _feeDistributorInstance, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3TransactorSession) MakeBeaconDeposit(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _feeDistributorInstance common.Address, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.MakeBeaconDeposit(&_DirectStakingDepositorV3.TransactOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _feeDistributorInstance, _pubkeys, _signatures, _depositDataRoots)
}

// Refund is a paid mutator transaction binding the contract method 0x08bedf95.
//
// Solidity: function refund(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _feeDistributorInstance) returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Transactor) Refund(opts *bind.TransactOpts, _eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _feeDistributorInstance common.Address) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.contract.Transact(opts, "refund", _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _feeDistributorInstance)
}

// Refund is a paid mutator transaction binding the contract method 0x08bedf95.
//
// Solidity: function refund(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _feeDistributorInstance) returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) Refund(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _feeDistributorInstance common.Address) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.Refund(&_DirectStakingDepositorV3.TransactOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _feeDistributorInstance)
}

// Refund is a paid mutator transaction binding the contract method 0x08bedf95.
//
// Solidity: function refund(bytes32 _eth2WithdrawalCredentials, uint96 _ethAmountPerValidatorInWei, address _feeDistributorInstance) returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3TransactorSession) Refund(_eth2WithdrawalCredentials [32]byte, _ethAmountPerValidatorInWei *big.Int, _feeDistributorInstance common.Address) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.Refund(&_DirectStakingDepositorV3.TransactOpts, _eth2WithdrawalCredentials, _ethAmountPerValidatorInWei, _feeDistributorInstance)
}

// RejectService is a paid mutator transaction binding the contract method 0xd77cc535.
//
// Solidity: function rejectService(bytes32 _depositId, string _reason) returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Transactor) RejectService(opts *bind.TransactOpts, _depositId [32]byte, _reason string) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.contract.Transact(opts, "rejectService", _depositId, _reason)
}

// RejectService is a paid mutator transaction binding the contract method 0xd77cc535.
//
// Solidity: function rejectService(bytes32 _depositId, string _reason) returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) RejectService(_depositId [32]byte, _reason string) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.RejectService(&_DirectStakingDepositorV3.TransactOpts, _depositId, _reason)
}

// RejectService is a paid mutator transaction binding the contract method 0xd77cc535.
//
// Solidity: function rejectService(bytes32 _depositId, string _reason) returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3TransactorSession) RejectService(_depositId [32]byte, _reason string) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.RejectService(&_DirectStakingDepositorV3.TransactOpts, _depositId, _reason)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DirectStakingDepositorV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Session) Receive() (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.Receive(&_DirectStakingDepositorV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3TransactorSession) Receive() (*types.Transaction, error) {
	return _DirectStakingDepositorV3.Contract.Receive(&_DirectStakingDepositorV3.TransactOpts)
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAddedIterator is returned from FilterP2pOrgUnlimitedEthDepositorClientEthAdded and is used to iterate over the raw logs and unpacked data for P2pOrgUnlimitedEthDepositorClientEthAdded events raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAddedIterator struct {
	Event *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAdded // Event containing the contract specifics and raw log

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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAdded)
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
		it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAdded)
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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAdded represents a P2pOrgUnlimitedEthDepositorClientEthAdded event raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAdded struct {
	DepositId                  [32]byte
	Sender                     common.Address
	FeeDistributorInstance     common.Address
	Eth2WithdrawalCredentials  [32]byte
	Amount                     *big.Int
	Expiration                 *big.Int
	EthAmountPerValidatorInWei *big.Int
	ExtraData                  []byte
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterP2pOrgUnlimitedEthDepositorClientEthAdded is a free log retrieval operation binding the contract event 0xb68a516afc822ac0bf66ed8ded762ab226ec9f1a6f957997ae2053eebc6b388a.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__ClientEthAdded(bytes32 indexed _depositId, address indexed _sender, address indexed _feeDistributorInstance, bytes32 _eth2WithdrawalCredentials, uint256 _amount, uint40 _expiration, uint96 _ethAmountPerValidatorInWei, bytes _extraData)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) FilterP2pOrgUnlimitedEthDepositorClientEthAdded(opts *bind.FilterOpts, _depositId [][32]byte, _sender []common.Address, _feeDistributorInstance []common.Address) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAddedIterator, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _feeDistributorInstanceRule []interface{}
	for _, _feeDistributorInstanceItem := range _feeDistributorInstance {
		_feeDistributorInstanceRule = append(_feeDistributorInstanceRule, _feeDistributorInstanceItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.FilterLogs(opts, "P2pOrgUnlimitedEthDepositor__ClientEthAdded", _depositIdRule, _senderRule, _feeDistributorInstanceRule)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAddedIterator{contract: _DirectStakingDepositorV3.contract, event: "P2pOrgUnlimitedEthDepositor__ClientEthAdded", logs: logs, sub: sub}, nil
}

// WatchP2pOrgUnlimitedEthDepositorClientEthAdded is a free log subscription operation binding the contract event 0xb68a516afc822ac0bf66ed8ded762ab226ec9f1a6f957997ae2053eebc6b388a.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__ClientEthAdded(bytes32 indexed _depositId, address indexed _sender, address indexed _feeDistributorInstance, bytes32 _eth2WithdrawalCredentials, uint256 _amount, uint40 _expiration, uint96 _ethAmountPerValidatorInWei, bytes _extraData)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) WatchP2pOrgUnlimitedEthDepositorClientEthAdded(opts *bind.WatchOpts, sink chan<- *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAdded, _depositId [][32]byte, _sender []common.Address, _feeDistributorInstance []common.Address) (event.Subscription, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _feeDistributorInstanceRule []interface{}
	for _, _feeDistributorInstanceItem := range _feeDistributorInstance {
		_feeDistributorInstanceRule = append(_feeDistributorInstanceRule, _feeDistributorInstanceItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.WatchLogs(opts, "P2pOrgUnlimitedEthDepositor__ClientEthAdded", _depositIdRule, _senderRule, _feeDistributorInstanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAdded)
				if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__ClientEthAdded", log); err != nil {
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

// ParseP2pOrgUnlimitedEthDepositorClientEthAdded is a log parse operation binding the contract event 0xb68a516afc822ac0bf66ed8ded762ab226ec9f1a6f957997ae2053eebc6b388a.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__ClientEthAdded(bytes32 indexed _depositId, address indexed _sender, address indexed _feeDistributorInstance, bytes32 _eth2WithdrawalCredentials, uint256 _amount, uint40 _expiration, uint96 _ethAmountPerValidatorInWei, bytes _extraData)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) ParseP2pOrgUnlimitedEthDepositorClientEthAdded(log types.Log) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAdded, error) {
	event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorClientEthAdded)
	if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__ClientEthAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251EnabledIterator is returned from FilterP2pOrgUnlimitedEthDepositorEip7251Enabled and is used to iterate over the raw logs and unpacked data for P2pOrgUnlimitedEthDepositorEip7251Enabled events raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251EnabledIterator struct {
	Event *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251Enabled // Event containing the contract specifics and raw log

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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251EnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251Enabled)
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
		it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251Enabled)
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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251EnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251EnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251Enabled represents a P2pOrgUnlimitedEthDepositorEip7251Enabled event raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251Enabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterP2pOrgUnlimitedEthDepositorEip7251Enabled is a free log retrieval operation binding the contract event 0x3a8316ecfbf6f3065062d901f7624822ee8487d310ab29ad723e8a90276009b0.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eip7251Enabled()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) FilterP2pOrgUnlimitedEthDepositorEip7251Enabled(opts *bind.FilterOpts) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251EnabledIterator, error) {

	logs, sub, err := _DirectStakingDepositorV3.contract.FilterLogs(opts, "P2pOrgUnlimitedEthDepositor__Eip7251Enabled")
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251EnabledIterator{contract: _DirectStakingDepositorV3.contract, event: "P2pOrgUnlimitedEthDepositor__Eip7251Enabled", logs: logs, sub: sub}, nil
}

// WatchP2pOrgUnlimitedEthDepositorEip7251Enabled is a free log subscription operation binding the contract event 0x3a8316ecfbf6f3065062d901f7624822ee8487d310ab29ad723e8a90276009b0.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eip7251Enabled()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) WatchP2pOrgUnlimitedEthDepositorEip7251Enabled(opts *bind.WatchOpts, sink chan<- *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251Enabled) (event.Subscription, error) {

	logs, sub, err := _DirectStakingDepositorV3.contract.WatchLogs(opts, "P2pOrgUnlimitedEthDepositor__Eip7251Enabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251Enabled)
				if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Eip7251Enabled", log); err != nil {
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

// ParseP2pOrgUnlimitedEthDepositorEip7251Enabled is a log parse operation binding the contract event 0x3a8316ecfbf6f3065062d901f7624822ee8487d310ab29ad723e8a90276009b0.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eip7251Enabled()
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) ParseP2pOrgUnlimitedEthDepositorEip7251Enabled(log types.Log) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251Enabled, error) {
	event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEip7251Enabled)
	if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Eip7251Enabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositIterator is returned from FilterP2pOrgUnlimitedEthDepositorEth2Deposit and is used to iterate over the raw logs and unpacked data for P2pOrgUnlimitedEthDepositorEth2Deposit events raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositIterator struct {
	Event *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2Deposit // Event containing the contract specifics and raw log

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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2Deposit)
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
		it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2Deposit)
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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2Deposit represents a P2pOrgUnlimitedEthDepositorEth2Deposit event raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2Deposit struct {
	DepositId      [32]byte
	ValidatorCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterP2pOrgUnlimitedEthDepositorEth2Deposit is a free log retrieval operation binding the contract event 0x0a564087509236efe6e8513e8e2ea4ef2012d4986653771a5e8fe8dafa6741cb.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eth2Deposit(bytes32 indexed _depositId, uint256 _validatorCount)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) FilterP2pOrgUnlimitedEthDepositorEth2Deposit(opts *bind.FilterOpts, _depositId [][32]byte) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositIterator, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.FilterLogs(opts, "P2pOrgUnlimitedEthDepositor__Eth2Deposit", _depositIdRule)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositIterator{contract: _DirectStakingDepositorV3.contract, event: "P2pOrgUnlimitedEthDepositor__Eth2Deposit", logs: logs, sub: sub}, nil
}

// WatchP2pOrgUnlimitedEthDepositorEth2Deposit is a free log subscription operation binding the contract event 0x0a564087509236efe6e8513e8e2ea4ef2012d4986653771a5e8fe8dafa6741cb.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eth2Deposit(bytes32 indexed _depositId, uint256 _validatorCount)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) WatchP2pOrgUnlimitedEthDepositorEth2Deposit(opts *bind.WatchOpts, sink chan<- *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2Deposit, _depositId [][32]byte) (event.Subscription, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.WatchLogs(opts, "P2pOrgUnlimitedEthDepositor__Eth2Deposit", _depositIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2Deposit)
				if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Eth2Deposit", log); err != nil {
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

// ParseP2pOrgUnlimitedEthDepositorEth2Deposit is a log parse operation binding the contract event 0x0a564087509236efe6e8513e8e2ea4ef2012d4986653771a5e8fe8dafa6741cb.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eth2Deposit(bytes32 indexed _depositId, uint256 _validatorCount)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) ParseP2pOrgUnlimitedEthDepositorEth2Deposit(log types.Log) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2Deposit, error) {
	event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2Deposit)
	if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Eth2Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompletedIterator is returned from FilterP2pOrgUnlimitedEthDepositorEth2DepositCompleted and is used to iterate over the raw logs and unpacked data for P2pOrgUnlimitedEthDepositorEth2DepositCompleted events raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompletedIterator struct {
	Event *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompleted // Event containing the contract specifics and raw log

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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompleted)
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
		it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompleted)
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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompleted represents a P2pOrgUnlimitedEthDepositorEth2DepositCompleted event raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompleted struct {
	DepositId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterP2pOrgUnlimitedEthDepositorEth2DepositCompleted is a free log retrieval operation binding the contract event 0x21b9552c832765250435d11e81889d766400ef1520579ae27e6e8df3fcc240ef.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eth2DepositCompleted(bytes32 indexed _depositId)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) FilterP2pOrgUnlimitedEthDepositorEth2DepositCompleted(opts *bind.FilterOpts, _depositId [][32]byte) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompletedIterator, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.FilterLogs(opts, "P2pOrgUnlimitedEthDepositor__Eth2DepositCompleted", _depositIdRule)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompletedIterator{contract: _DirectStakingDepositorV3.contract, event: "P2pOrgUnlimitedEthDepositor__Eth2DepositCompleted", logs: logs, sub: sub}, nil
}

// WatchP2pOrgUnlimitedEthDepositorEth2DepositCompleted is a free log subscription operation binding the contract event 0x21b9552c832765250435d11e81889d766400ef1520579ae27e6e8df3fcc240ef.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eth2DepositCompleted(bytes32 indexed _depositId)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) WatchP2pOrgUnlimitedEthDepositorEth2DepositCompleted(opts *bind.WatchOpts, sink chan<- *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompleted, _depositId [][32]byte) (event.Subscription, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.WatchLogs(opts, "P2pOrgUnlimitedEthDepositor__Eth2DepositCompleted", _depositIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompleted)
				if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Eth2DepositCompleted", log); err != nil {
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

// ParseP2pOrgUnlimitedEthDepositorEth2DepositCompleted is a log parse operation binding the contract event 0x21b9552c832765250435d11e81889d766400ef1520579ae27e6e8df3fcc240ef.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eth2DepositCompleted(bytes32 indexed _depositId)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) ParseP2pOrgUnlimitedEthDepositorEth2DepositCompleted(log types.Log) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompleted, error) {
	event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositCompleted)
	if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Eth2DepositCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgressIterator is returned from FilterP2pOrgUnlimitedEthDepositorEth2DepositInProgress and is used to iterate over the raw logs and unpacked data for P2pOrgUnlimitedEthDepositorEth2DepositInProgress events raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgressIterator struct {
	Event *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgress // Event containing the contract specifics and raw log

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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgress)
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
		it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgress)
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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgress represents a P2pOrgUnlimitedEthDepositorEth2DepositInProgress event raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgress struct {
	DepositId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterP2pOrgUnlimitedEthDepositorEth2DepositInProgress is a free log retrieval operation binding the contract event 0xd8f8be349b5615f4a7af6a678a23efb0ffa3fc3d6715344cf07f6b8347fec98d.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eth2DepositInProgress(bytes32 indexed _depositId)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) FilterP2pOrgUnlimitedEthDepositorEth2DepositInProgress(opts *bind.FilterOpts, _depositId [][32]byte) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgressIterator, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.FilterLogs(opts, "P2pOrgUnlimitedEthDepositor__Eth2DepositInProgress", _depositIdRule)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgressIterator{contract: _DirectStakingDepositorV3.contract, event: "P2pOrgUnlimitedEthDepositor__Eth2DepositInProgress", logs: logs, sub: sub}, nil
}

// WatchP2pOrgUnlimitedEthDepositorEth2DepositInProgress is a free log subscription operation binding the contract event 0xd8f8be349b5615f4a7af6a678a23efb0ffa3fc3d6715344cf07f6b8347fec98d.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eth2DepositInProgress(bytes32 indexed _depositId)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) WatchP2pOrgUnlimitedEthDepositorEth2DepositInProgress(opts *bind.WatchOpts, sink chan<- *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgress, _depositId [][32]byte) (event.Subscription, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.WatchLogs(opts, "P2pOrgUnlimitedEthDepositor__Eth2DepositInProgress", _depositIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgress)
				if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Eth2DepositInProgress", log); err != nil {
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

// ParseP2pOrgUnlimitedEthDepositorEth2DepositInProgress is a log parse operation binding the contract event 0xd8f8be349b5615f4a7af6a678a23efb0ffa3fc3d6715344cf07f6b8347fec98d.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Eth2DepositInProgress(bytes32 indexed _depositId)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) ParseP2pOrgUnlimitedEthDepositorEth2DepositInProgress(log types.Log) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgress, error) {
	event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorEth2DepositInProgress)
	if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Eth2DepositInProgress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefundIterator is returned from FilterP2pOrgUnlimitedEthDepositorRefund and is used to iterate over the raw logs and unpacked data for P2pOrgUnlimitedEthDepositorRefund events raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefundIterator struct {
	Event *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefund // Event containing the contract specifics and raw log

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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefund)
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
		it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefund)
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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefund represents a P2pOrgUnlimitedEthDepositorRefund event raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefund struct {
	DepositId              [32]byte
	FeeDistributorInstance common.Address
	Client                 common.Address
	Amount                 *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterP2pOrgUnlimitedEthDepositorRefund is a free log retrieval operation binding the contract event 0x1443e4defcd744ac29619ec2a0bdd3ac7adb4040976d5f49e38b7bc4e5accb57.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Refund(bytes32 indexed _depositId, address indexed _feeDistributorInstance, address indexed _client, uint256 _amount)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) FilterP2pOrgUnlimitedEthDepositorRefund(opts *bind.FilterOpts, _depositId [][32]byte, _feeDistributorInstance []common.Address, _client []common.Address) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefundIterator, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}
	var _feeDistributorInstanceRule []interface{}
	for _, _feeDistributorInstanceItem := range _feeDistributorInstance {
		_feeDistributorInstanceRule = append(_feeDistributorInstanceRule, _feeDistributorInstanceItem)
	}
	var _clientRule []interface{}
	for _, _clientItem := range _client {
		_clientRule = append(_clientRule, _clientItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.FilterLogs(opts, "P2pOrgUnlimitedEthDepositor__Refund", _depositIdRule, _feeDistributorInstanceRule, _clientRule)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefundIterator{contract: _DirectStakingDepositorV3.contract, event: "P2pOrgUnlimitedEthDepositor__Refund", logs: logs, sub: sub}, nil
}

// WatchP2pOrgUnlimitedEthDepositorRefund is a free log subscription operation binding the contract event 0x1443e4defcd744ac29619ec2a0bdd3ac7adb4040976d5f49e38b7bc4e5accb57.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Refund(bytes32 indexed _depositId, address indexed _feeDistributorInstance, address indexed _client, uint256 _amount)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) WatchP2pOrgUnlimitedEthDepositorRefund(opts *bind.WatchOpts, sink chan<- *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefund, _depositId [][32]byte, _feeDistributorInstance []common.Address, _client []common.Address) (event.Subscription, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}
	var _feeDistributorInstanceRule []interface{}
	for _, _feeDistributorInstanceItem := range _feeDistributorInstance {
		_feeDistributorInstanceRule = append(_feeDistributorInstanceRule, _feeDistributorInstanceItem)
	}
	var _clientRule []interface{}
	for _, _clientItem := range _client {
		_clientRule = append(_clientRule, _clientItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.WatchLogs(opts, "P2pOrgUnlimitedEthDepositor__Refund", _depositIdRule, _feeDistributorInstanceRule, _clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefund)
				if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Refund", log); err != nil {
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

// ParseP2pOrgUnlimitedEthDepositorRefund is a log parse operation binding the contract event 0x1443e4defcd744ac29619ec2a0bdd3ac7adb4040976d5f49e38b7bc4e5accb57.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__Refund(bytes32 indexed _depositId, address indexed _feeDistributorInstance, address indexed _client, uint256 _amount)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) ParseP2pOrgUnlimitedEthDepositorRefund(log types.Log) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefund, error) {
	event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorRefund)
	if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejectedIterator is returned from FilterP2pOrgUnlimitedEthDepositorServiceRejected and is used to iterate over the raw logs and unpacked data for P2pOrgUnlimitedEthDepositorServiceRejected events raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejectedIterator struct {
	Event *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejected // Event containing the contract specifics and raw log

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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejected)
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
		it.Event = new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejected)
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
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejected represents a P2pOrgUnlimitedEthDepositorServiceRejected event raised by the DirectStakingDepositorV3 contract.
type DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejected struct {
	DepositId [32]byte
	Reason    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterP2pOrgUnlimitedEthDepositorServiceRejected is a free log retrieval operation binding the contract event 0x320fa1e69da1125204e7c088a134705e00c10e1d0ad6082a0361da2532f6e945.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__ServiceRejected(bytes32 indexed _depositId, string _reason)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) FilterP2pOrgUnlimitedEthDepositorServiceRejected(opts *bind.FilterOpts, _depositId [][32]byte) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejectedIterator, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.FilterLogs(opts, "P2pOrgUnlimitedEthDepositor__ServiceRejected", _depositIdRule)
	if err != nil {
		return nil, err
	}
	return &DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejectedIterator{contract: _DirectStakingDepositorV3.contract, event: "P2pOrgUnlimitedEthDepositor__ServiceRejected", logs: logs, sub: sub}, nil
}

// WatchP2pOrgUnlimitedEthDepositorServiceRejected is a free log subscription operation binding the contract event 0x320fa1e69da1125204e7c088a134705e00c10e1d0ad6082a0361da2532f6e945.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__ServiceRejected(bytes32 indexed _depositId, string _reason)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) WatchP2pOrgUnlimitedEthDepositorServiceRejected(opts *bind.WatchOpts, sink chan<- *DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejected, _depositId [][32]byte) (event.Subscription, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _DirectStakingDepositorV3.contract.WatchLogs(opts, "P2pOrgUnlimitedEthDepositor__ServiceRejected", _depositIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejected)
				if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__ServiceRejected", log); err != nil {
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

// ParseP2pOrgUnlimitedEthDepositorServiceRejected is a log parse operation binding the contract event 0x320fa1e69da1125204e7c088a134705e00c10e1d0ad6082a0361da2532f6e945.
//
// Solidity: event P2pOrgUnlimitedEthDepositor__ServiceRejected(bytes32 indexed _depositId, string _reason)
func (_DirectStakingDepositorV3 *DirectStakingDepositorV3Filterer) ParseP2pOrgUnlimitedEthDepositorServiceRejected(log types.Log) (*DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejected, error) {
	event := new(DirectStakingDepositorV3P2pOrgUnlimitedEthDepositorServiceRejected)
	if err := _DirectStakingDepositorV3.contract.UnpackLog(event, "P2pOrgUnlimitedEthDepositor__ServiceRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
