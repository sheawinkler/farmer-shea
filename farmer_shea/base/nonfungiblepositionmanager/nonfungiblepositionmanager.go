// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nonfungiblepositionmanager

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

// INonfungiblePositionManagerMintParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerMintParams struct {
	Token0         common.Address
	Token1         common.Address
	Fee            *big.Int
	TickLower      *big.Int
	TickUpper      *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Recipient      common.Address
	Deadline       *big.Int
}

// NonfungiblePositionManagerMetaData contains all meta data concerning the NonfungiblePositionManager contract.
var NonfungiblePositionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenDescriptor_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"DecreaseLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"IncreaseLiquidity\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Desired\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Desired\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount0Min\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Min\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NonfungiblePositionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NonfungiblePositionManagerMetaData.ABI instead.
var NonfungiblePositionManagerABI = NonfungiblePositionManagerMetaData.ABI

// NonfungiblePositionManager is an auto generated Go binding around an Ethereum contract.
type NonfungiblePositionManager struct {
	NonfungiblePositionManagerCaller     // Read-only binding to the contract
	NonfungiblePositionManagerTransactor // Write-only binding to the contract
	NonfungiblePositionManagerFilterer   // Log filterer for contract events
}

// NonfungiblePositionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonfungiblePositionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonfungiblePositionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonfungiblePositionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonfungiblePositionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonfungiblePositionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonfungiblePositionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonfungiblePositionManagerSession struct {
	Contract     *NonfungiblePositionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// NonfungiblePositionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonfungiblePositionManagerCallerSession struct {
	Contract *NonfungiblePositionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// NonfungiblePositionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonfungiblePositionManagerTransactorSession struct {
	Contract     *NonfungiblePositionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// NonfungiblePositionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonfungiblePositionManagerRaw struct {
	Contract *NonfungiblePositionManager // Generic contract binding to access the raw methods on
}

// NonfungiblePositionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonfungiblePositionManagerCallerRaw struct {
	Contract *NonfungiblePositionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NonfungiblePositionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonfungiblePositionManagerTransactorRaw struct {
	Contract *NonfungiblePositionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonfungiblePositionManager creates a new instance of NonfungiblePositionManager, bound to a specific deployed contract.
func NewNonfungiblePositionManager(address common.Address, backend bind.ContractBackend) (*NonfungiblePositionManager, error) {
	contract, err := bindNonfungiblePositionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonfungiblePositionManager{NonfungiblePositionManagerCaller: NonfungiblePositionManagerCaller{contract: contract}, NonfungiblePositionManagerTransactor: NonfungiblePositionManagerTransactor{contract: contract}, NonfungiblePositionManagerFilterer: NonfungiblePositionManagerFilterer{contract: contract}}, nil
}

// NewNonfungiblePositionManagerCaller creates a new read-only instance of NonfungiblePositionManager, bound to a specific deployed contract.
func NewNonfungiblePositionManagerCaller(address common.Address, caller bind.ContractCaller) (*NonfungiblePositionManagerCaller, error) {
	contract, err := bindNonfungiblePositionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonfungiblePositionManagerCaller{contract: contract}, nil
}

// NewNonfungiblePositionManagerTransactor creates a new write-only instance of NonfungiblePositionManager, bound to a specific deployed contract.
func NewNonfungiblePositionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NonfungiblePositionManagerTransactor, error) {
	contract, err := bindNonfungiblePositionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonfungiblePositionManagerTransactor{contract: contract}, nil
}

// NewNonfungiblePositionManagerFilterer creates a new log filterer instance of NonfungiblePositionManager, bound to a specific deployed contract.
func NewNonfungiblePositionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NonfungiblePositionManagerFilterer, error) {
	contract, err := bindNonfungiblePositionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonfungiblePositionManagerFilterer{contract: contract}, nil
}

// bindNonfungiblePositionManager binds a generic wrapper to an already deployed contract.
func bindNonfungiblePositionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NonfungiblePositionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonfungiblePositionManager *NonfungiblePositionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonfungiblePositionManager.Contract.NonfungiblePositionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonfungiblePositionManager *NonfungiblePositionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonfungiblePositionManager.Contract.NonfungiblePositionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonfungiblePositionManager *NonfungiblePositionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonfungiblePositionManager.Contract.NonfungiblePositionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonfungiblePositionManager *NonfungiblePositionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonfungiblePositionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonfungiblePositionManager *NonfungiblePositionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonfungiblePositionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonfungiblePositionManager *NonfungiblePositionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonfungiblePositionManager.Contract.contract.Transact(opts, method, params...)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerCaller) Positions(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _NonfungiblePositionManager.contract.Call(opts, &out, "positions", tokenId)

	outstruct := new(struct {
		Nonce                    *big.Int
		Operator                 common.Address
		Token0                   common.Address
		Token1                   common.Address
		Fee                      *big.Int
		TickLower                *big.Int
		TickUpper                *big.Int
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TickLower = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TickUpper = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _NonfungiblePositionManager.Contract.Positions(&_NonfungiblePositionManager.CallOpts, tokenId)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerCallerSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _NonfungiblePositionManager.Contract.Positions(&_NonfungiblePositionManager.CallOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x9bd84318.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint128,uint128,uint128,uint128,address,uint256) params) returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerTransactor) Mint(opts *bind.TransactOpts, params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _NonfungiblePositionManager.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x9bd84318.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint128,uint128,uint128,uint128,address,uint256) params) returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _NonfungiblePositionManager.Contract.Mint(&_NonfungiblePositionManager.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x9bd84318.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint128,uint128,uint128,uint128,address,uint256) params) returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerTransactorSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _NonfungiblePositionManager.Contract.Mint(&_NonfungiblePositionManager.TransactOpts, params)
}

// NonfungiblePositionManagerCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the NonfungiblePositionManager contract.
type NonfungiblePositionManagerCollectIterator struct {
	Event *NonfungiblePositionManagerCollect // Event containing the contract specifics and raw log

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
func (it *NonfungiblePositionManagerCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonfungiblePositionManagerCollect)
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
		it.Event = new(NonfungiblePositionManagerCollect)
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
func (it *NonfungiblePositionManagerCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonfungiblePositionManagerCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonfungiblePositionManagerCollect represents a Collect event raised by the NonfungiblePositionManager contract.
type NonfungiblePositionManagerCollect struct {
	TokenId   *big.Int
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x4d8babf9b22e68d8f3c8653392a91073d3f3d246ad70593d8c8ed3fe381b3c96.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint128 amount0, uint128 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerFilterer) FilterCollect(opts *bind.FilterOpts, tokenId []*big.Int) (*NonfungiblePositionManagerCollectIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NonfungiblePositionManager.contract.FilterLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NonfungiblePositionManagerCollectIterator{contract: _NonfungiblePositionManager.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x4d8babf9b22e68d8f3c8653392a91073d3f3d246ad70593d8c8ed3fe381b3c96.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint128 amount0, uint128 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *NonfungiblePositionManagerCollect, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NonfungiblePositionManager.contract.WatchLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonfungiblePositionManagerCollect)
				if err := _NonfungiblePositionManager.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x4d8babf9b22e68d8f3c8653392a91073d3f3d246ad70593d8c8ed3fe381b3c96.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint128 amount0, uint128 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerFilterer) ParseCollect(log types.Log) (*NonfungiblePositionManagerCollect, error) {
	event := new(NonfungiblePositionManagerCollect)
	if err := _NonfungiblePositionManager.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NonfungiblePositionManagerDecreaseLiquidityIterator is returned from FilterDecreaseLiquidity and is used to iterate over the raw logs and unpacked data for DecreaseLiquidity events raised by the NonfungiblePositionManager contract.
type NonfungiblePositionManagerDecreaseLiquidityIterator struct {
	Event *NonfungiblePositionManagerDecreaseLiquidity // Event containing the contract specifics and raw log

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
func (it *NonfungiblePositionManagerDecreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonfungiblePositionManagerDecreaseLiquidity)
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
		it.Event = new(NonfungiblePositionManagerDecreaseLiquidity)
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
func (it *NonfungiblePositionManagerDecreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonfungiblePositionManagerDecreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonfungiblePositionManagerDecreaseLiquidity represents a DecreaseLiquidity event raised by the NonfungiblePositionManager contract.
type NonfungiblePositionManagerDecreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDecreaseLiquidity is a free log retrieval operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerFilterer) FilterDecreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*NonfungiblePositionManagerDecreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NonfungiblePositionManager.contract.FilterLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NonfungiblePositionManagerDecreaseLiquidityIterator{contract: _NonfungiblePositionManager.contract, event: "DecreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchDecreaseLiquidity is a free log subscription operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerFilterer) WatchDecreaseLiquidity(opts *bind.WatchOpts, sink chan<- *NonfungiblePositionManagerDecreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NonfungiblePositionManager.contract.WatchLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonfungiblePositionManagerDecreaseLiquidity)
				if err := _NonfungiblePositionManager.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
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

// ParseDecreaseLiquidity is a log parse operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerFilterer) ParseDecreaseLiquidity(log types.Log) (*NonfungiblePositionManagerDecreaseLiquidity, error) {
	event := new(NonfungiblePositionManagerDecreaseLiquidity)
	if err := _NonfungiblePositionManager.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NonfungiblePositionManagerIncreaseLiquidityIterator is returned from FilterIncreaseLiquidity and is used to iterate over the raw logs and unpacked data for IncreaseLiquidity events raised by the NonfungiblePositionManager contract.
type NonfungiblePositionManagerIncreaseLiquidityIterator struct {
	Event *NonfungiblePositionManagerIncreaseLiquidity // Event containing the contract specifics and raw log

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
func (it *NonfungiblePositionManagerIncreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonfungiblePositionManagerIncreaseLiquidity)
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
		it.Event = new(NonfungiblePositionManagerIncreaseLiquidity)
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
func (it *NonfungiblePositionManagerIncreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonfungiblePositionManagerIncreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonfungiblePositionManagerIncreaseLiquidity represents a IncreaseLiquidity event raised by the NonfungiblePositionManager contract.
type NonfungiblePositionManagerIncreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreaseLiquidity is a free log retrieval operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerFilterer) FilterIncreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*NonfungiblePositionManagerIncreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NonfungiblePositionManager.contract.FilterLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NonfungiblePositionManagerIncreaseLiquidityIterator{contract: _NonfungiblePositionManager.contract, event: "IncreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchIncreaseLiquidity is a free log subscription operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerFilterer) WatchIncreaseLiquidity(opts *bind.WatchOpts, sink chan<- *NonfungiblePositionManagerIncreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NonfungiblePositionManager.contract.WatchLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonfungiblePositionManagerIncreaseLiquidity)
				if err := _NonfungiblePositionManager.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
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

// ParseIncreaseLiquidity is a log parse operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonfungiblePositionManager *NonfungiblePositionManagerFilterer) ParseIncreaseLiquidity(log types.Log) (*NonfungiblePositionManagerIncreaseLiquidity, error) {
	event := new(NonfungiblePositionManagerIncreaseLiquidity)
	if err := _NonfungiblePositionManager.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
