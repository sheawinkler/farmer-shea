// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv3factory

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

// Uniswapv3factoryMetaData contains all meta data concerning the Uniswapv3factory contract.
var Uniswapv3factoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"FeeAmountEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Uniswapv3factoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv3factoryMetaData.ABI instead.
var Uniswapv3factoryABI = Uniswapv3factoryMetaData.ABI

// Uniswapv3factory is an auto generated Go binding around an Ethereum contract.
type Uniswapv3factory struct {
	Uniswapv3factoryCaller     // Read-only binding to the contract
	Uniswapv3factoryTransactor // Write-only binding to the contract
	Uniswapv3factoryFilterer   // Log filterer for contract events
}

// Uniswapv3factoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv3factoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3factoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv3factoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3factoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv3factoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3factorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv3factorySession struct {
	Contract     *Uniswapv3factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv3factoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv3factoryCallerSession struct {
	Contract *Uniswapv3factoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Uniswapv3factoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv3factoryTransactorSession struct {
	Contract     *Uniswapv3factoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Uniswapv3factoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv3factoryRaw struct {
	Contract *Uniswapv3factory // Generic contract binding to access the raw methods on
}

// Uniswapv3factoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv3factoryCallerRaw struct {
	Contract *Uniswapv3factoryCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv3factoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv3factoryTransactorRaw struct {
	Contract *Uniswapv3factoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv3factory creates a new instance of Uniswapv3factory, bound to a specific deployed contract.
func NewUniswapv3factory(address common.Address, backend bind.ContractBackend) (*Uniswapv3factory, error) {
	contract, err := bindUniswapv3factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3factory{Uniswapv3factoryCaller: Uniswapv3factoryCaller{contract: contract}, Uniswapv3factoryTransactor: Uniswapv3factoryTransactor{contract: contract}, Uniswapv3factoryFilterer: Uniswapv3factoryFilterer{contract: contract}}, nil
}

// NewUniswapv3factoryCaller creates a new read-only instance of Uniswapv3factory, bound to a specific deployed contract.
func NewUniswapv3factoryCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv3factoryCaller, error) {
	contract, err := bindUniswapv3factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3factoryCaller{contract: contract}, nil
}

// NewUniswapv3factoryTransactor creates a new write-only instance of Uniswapv3factory, bound to a specific deployed contract.
func NewUniswapv3factoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv3factoryTransactor, error) {
	contract, err := bindUniswapv3factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3factoryTransactor{contract: contract}, nil
}

// NewUniswapv3factoryFilterer creates a new log filterer instance of Uniswapv3factory, bound to a specific deployed contract.
func NewUniswapv3factoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv3factoryFilterer, error) {
	contract, err := bindUniswapv3factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3factoryFilterer{contract: contract}, nil
}

// bindUniswapv3factory binds a generic wrapper to an already deployed contract.
func bindUniswapv3factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Uniswapv3factoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3factory *Uniswapv3factoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3factory.Contract.Uniswapv3factoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3factory *Uniswapv3factoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.Uniswapv3factoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3factory *Uniswapv3factoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.Uniswapv3factoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3factory *Uniswapv3factoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3factory *Uniswapv3factoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3factory *Uniswapv3factoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.contract.Transact(opts, method, params...)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address token0, address token1, uint24 fee) view returns(address pool)
func (_Uniswapv3factory *Uniswapv3factoryCaller) GetPool(opts *bind.CallOpts, token0 common.Address, token1 common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3factory.contract.Call(opts, &out, "getPool", token0, token1, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address token0, address token1, uint24 fee) view returns(address pool)
func (_Uniswapv3factory *Uniswapv3factorySession) GetPool(token0 common.Address, token1 common.Address, fee *big.Int) (common.Address, error) {
	return _Uniswapv3factory.Contract.GetPool(&_Uniswapv3factory.CallOpts, token0, token1, fee)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address token0, address token1, uint24 fee) view returns(address pool)
func (_Uniswapv3factory *Uniswapv3factoryCallerSession) GetPool(token0 common.Address, token1 common.Address, fee *big.Int) (common.Address, error) {
	return _Uniswapv3factory.Contract.GetPool(&_Uniswapv3factory.CallOpts, token0, token1, fee)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Uniswapv3factory *Uniswapv3factoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Uniswapv3factory *Uniswapv3factorySession) Owner() (common.Address, error) {
	return _Uniswapv3factory.Contract.Owner(&_Uniswapv3factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Uniswapv3factory *Uniswapv3factoryCallerSession) Owner() (common.Address, error) {
	return _Uniswapv3factory.Contract.Owner(&_Uniswapv3factory.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_Uniswapv3factory *Uniswapv3factoryCaller) Parameters(opts *bind.CallOpts) (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	var out []interface{}
	err := _Uniswapv3factory.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		Factory     common.Address
		Token0      common.Address
		Token1      common.Address
		Fee         *big.Int
		TickSpacing *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickSpacing = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_Uniswapv3factory *Uniswapv3factorySession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _Uniswapv3factory.Contract.Parameters(&_Uniswapv3factory.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_Uniswapv3factory *Uniswapv3factoryCallerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _Uniswapv3factory.Contract.Parameters(&_Uniswapv3factory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_Uniswapv3factory *Uniswapv3factoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Uniswapv3factory.contract.Transact(opts, "createPool", tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_Uniswapv3factory *Uniswapv3factorySession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.CreatePool(&_Uniswapv3factory.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_Uniswapv3factory *Uniswapv3factoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.CreatePool(&_Uniswapv3factory.TransactOpts, tokenA, tokenB, fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0xbcb928ec.
//
// Solidity: function enableFeeAmount(uint24 fee) returns()
func (_Uniswapv3factory *Uniswapv3factoryTransactor) EnableFeeAmount(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Uniswapv3factory.contract.Transact(opts, "enableFeeAmount", fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0xbcb928ec.
//
// Solidity: function enableFeeAmount(uint24 fee) returns()
func (_Uniswapv3factory *Uniswapv3factorySession) EnableFeeAmount(fee *big.Int) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.EnableFeeAmount(&_Uniswapv3factory.TransactOpts, fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0xbcb928ec.
//
// Solidity: function enableFeeAmount(uint24 fee) returns()
func (_Uniswapv3factory *Uniswapv3factoryTransactorSession) EnableFeeAmount(fee *big.Int) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.EnableFeeAmount(&_Uniswapv3factory.TransactOpts, fee)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Uniswapv3factory *Uniswapv3factoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Uniswapv3factory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Uniswapv3factory *Uniswapv3factorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.SetOwner(&_Uniswapv3factory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Uniswapv3factory *Uniswapv3factoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Uniswapv3factory.Contract.SetOwner(&_Uniswapv3factory.TransactOpts, _owner)
}

// Uniswapv3factoryFeeAmountEnabledIterator is returned from FilterFeeAmountEnabled and is used to iterate over the raw logs and unpacked data for FeeAmountEnabled events raised by the Uniswapv3factory contract.
type Uniswapv3factoryFeeAmountEnabledIterator struct {
	Event *Uniswapv3factoryFeeAmountEnabled // Event containing the contract specifics and raw log

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
func (it *Uniswapv3factoryFeeAmountEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Uniswapv3factoryFeeAmountEnabled)
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
		it.Event = new(Uniswapv3factoryFeeAmountEnabled)
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
func (it *Uniswapv3factoryFeeAmountEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Uniswapv3factoryFeeAmountEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Uniswapv3factoryFeeAmountEnabled represents a FeeAmountEnabled event raised by the Uniswapv3factory contract.
type Uniswapv3factoryFeeAmountEnabled struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountEnabled is a free log retrieval operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_Uniswapv3factory *Uniswapv3factoryFilterer) FilterFeeAmountEnabled(opts *bind.FilterOpts, fee []*big.Int, tickSpacing []*big.Int) (*Uniswapv3factoryFeeAmountEnabledIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _Uniswapv3factory.contract.FilterLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3factoryFeeAmountEnabledIterator{contract: _Uniswapv3factory.contract, event: "FeeAmountEnabled", logs: logs, sub: sub}, nil
}

// WatchFeeAmountEnabled is a free log subscription operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_Uniswapv3factory *Uniswapv3factoryFilterer) WatchFeeAmountEnabled(opts *bind.WatchOpts, sink chan<- *Uniswapv3factoryFeeAmountEnabled, fee []*big.Int, tickSpacing []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _Uniswapv3factory.contract.WatchLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Uniswapv3factoryFeeAmountEnabled)
				if err := _Uniswapv3factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
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

// ParseFeeAmountEnabled is a log parse operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_Uniswapv3factory *Uniswapv3factoryFilterer) ParseFeeAmountEnabled(log types.Log) (*Uniswapv3factoryFeeAmountEnabled, error) {
	event := new(Uniswapv3factoryFeeAmountEnabled)
	if err := _Uniswapv3factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Uniswapv3factoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Uniswapv3factory contract.
type Uniswapv3factoryPoolCreatedIterator struct {
	Event *Uniswapv3factoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *Uniswapv3factoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Uniswapv3factoryPoolCreated)
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
		it.Event = new(Uniswapv3factoryPoolCreated)
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
func (it *Uniswapv3factoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Uniswapv3factoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Uniswapv3factoryPoolCreated represents a PoolCreated event raised by the Uniswapv3factory contract.
type Uniswapv3factoryPoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_Uniswapv3factory *Uniswapv3factoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, fee []*big.Int) (*Uniswapv3factoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Uniswapv3factory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3factoryPoolCreatedIterator{contract: _Uniswapv3factory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_Uniswapv3factory *Uniswapv3factoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *Uniswapv3factoryPoolCreated, token0 []common.Address, token1 []common.Address, fee []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Uniswapv3factory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Uniswapv3factoryPoolCreated)
				if err := _Uniswapv3factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_Uniswapv3factory *Uniswapv3factoryFilterer) ParsePoolCreated(log types.Log) (*Uniswapv3factoryPoolCreated, error) {
	event := new(Uniswapv3factoryPoolCreated)
	if err := _Uniswapv3factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
