// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balance_op

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

// BalanceOpABI is the input ABI used to generate the binding from.
const BalanceOpABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

var BalanceOpParsedABI, _ = abi.JSON(strings.NewReader(BalanceOpABI))

// BalanceOpBin is the compiled bytecode used for deploying new contracts.
var BalanceOpBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506109e6806100606000396000f3fe60806040526004361061004e5760003560e01c806327e235e31461005a5780632e1a7d4d146100975780638da5cb5b146100c0578063a9059cbb146100eb578063b6b55f251461011457610055565b3661005557005b600080fd5b34801561006657600080fd5b50610081600480360381019061007c919061069d565b610130565b60405161008e91906106e3565b60405180910390f35b3480156100a357600080fd5b506100be60048036038101906100b9919061072a565b610148565b005b3480156100cc57600080fd5b506100d561032a565b6040516100e29190610778565b60405180910390f35b3480156100f757600080fd5b50610112600480360381019061010d9190610793565b61034e565b005b61012e6004803603810190610129919061072a565b610589565b005b60016020528060005260406000206000915090505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101cd90610830565b60405180910390fd5b80600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610279576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102709061089c565b60405180910390fd5b7f4e70a604b23a8edee2b1d0a656e9b9c00b73ad8bb1afc2c59381ee9f69197de7816040516102a891906106e3565b60405180910390a180600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461032091906108eb565b9250508190555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d390610830565b60405180910390fd5b80600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561047f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104769061089c565b60405180910390fd5b7f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de282826040516104b092919061092e565b60405180910390a180600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461052891906108eb565b9250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461057e9190610957565b925050819055505050565b7f4d6ce1e535dbade1c23defba91e23b8f791ce5edc0cc320257a2b364e4e38426816040516105b891906106e3565b60405180910390a180600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106309190610957565b9250508190555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061066a8261063f565b9050919050565b61067a8161065f565b811461068557600080fd5b50565b60008135905061069781610671565b92915050565b6000602082840312156106b3576106b261063a565b5b60006106c184828501610688565b91505092915050565b6000819050919050565b6106dd816106ca565b82525050565b60006020820190506106f860008301846106d4565b92915050565b610707816106ca565b811461071257600080fd5b50565b600081359050610724816106fe565b92915050565b6000602082840312156107405761073f61063a565b5b600061074e84828501610715565b91505092915050565b60006107628261063f565b9050919050565b61077281610757565b82525050565b600060208201905061078d6000830184610769565b92915050565b600080604083850312156107aa576107a961063a565b5b60006107b885828601610688565b92505060206107c985828601610715565b9150509250929050565b600082825260208201905092915050565b7f63616c6c6572206973206e6f74206f776e657200000000000000000000000000600082015250565b600061081a6013836107d3565b9150610825826107e4565b602082019050919050565b600060208201905081810360008301526108498161080d565b9050919050565b7f496e73756666696369656e742066756e64730000000000000000000000000000600082015250565b60006108866012836107d3565b915061089182610850565b602082019050919050565b600060208201905081810360008301526108b581610879565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006108f6826106ca565b9150610901836106ca565b9250828203905081811115610919576109186108bc565b5b92915050565b6109288161065f565b82525050565b6000604082019050610943600083018561091f565b61095060208301846106d4565b9392505050565b6000610962826106ca565b915061096d836106ca565b9250828201905080821115610985576109846108bc565b5b9291505056fea26469706673582212204b3600075f051f4f210f1bad32c250dbf4ec21a281785f61a55aa2c2e87d570864736f6c637827302e382e31362d646576656c6f702e323032322e372e352b636f6d6d69742e61353366313566340058"

// DeployBalanceOp deploys a new Ethereum contract, binding an instance of BalanceOp to it.
func DeployBalanceOp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BalanceOp, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceOpABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BalanceOpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BalanceOp{BalanceOpCaller: BalanceOpCaller{contract: contract}, BalanceOpTransactor: BalanceOpTransactor{contract: contract}, BalanceOpFilterer: BalanceOpFilterer{contract: contract}}, nil
}

// BalanceOp is an auto generated Go binding around an Ethereum contract.
type BalanceOp struct {
	BalanceOpCaller     // Read-only binding to the contract
	BalanceOpTransactor // Write-only binding to the contract
	BalanceOpFilterer   // Log filterer for contract events
}

// BalanceOpCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceOpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceOpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceOpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceOpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceOpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceOpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceOpSession struct {
	Contract     *BalanceOp        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceOpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceOpCallerSession struct {
	Contract *BalanceOpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BalanceOpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceOpTransactorSession struct {
	Contract     *BalanceOpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BalanceOpRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceOpRaw struct {
	Contract *BalanceOp // Generic contract binding to access the raw methods on
}

// BalanceOpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceOpCallerRaw struct {
	Contract *BalanceOpCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceOpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceOpTransactorRaw struct {
	Contract *BalanceOpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceOp creates a new instance of BalanceOp, bound to a specific deployed contract.
func NewBalanceOp(address common.Address, backend bind.ContractBackend) (*BalanceOp, error) {
	contract, err := bindBalanceOp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceOp{BalanceOpCaller: BalanceOpCaller{contract: contract}, BalanceOpTransactor: BalanceOpTransactor{contract: contract}, BalanceOpFilterer: BalanceOpFilterer{contract: contract}}, nil
}

// NewBalanceOpCaller creates a new read-only instance of BalanceOp, bound to a specific deployed contract.
func NewBalanceOpCaller(address common.Address, caller bind.ContractCaller) (*BalanceOpCaller, error) {
	contract, err := bindBalanceOp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceOpCaller{contract: contract}, nil
}

// NewBalanceOpTransactor creates a new write-only instance of BalanceOp, bound to a specific deployed contract.
func NewBalanceOpTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceOpTransactor, error) {
	contract, err := bindBalanceOp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceOpTransactor{contract: contract}, nil
}

// NewBalanceOpFilterer creates a new log filterer instance of BalanceOp, bound to a specific deployed contract.
func NewBalanceOpFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceOpFilterer, error) {
	contract, err := bindBalanceOp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceOpFilterer{contract: contract}, nil
}

// bindBalanceOp binds a generic wrapper to an already deployed contract.
func bindBalanceOp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceOpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceOp *BalanceOpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceOp.Contract.BalanceOpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceOp *BalanceOpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceOp.Contract.BalanceOpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceOp *BalanceOpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceOp.Contract.BalanceOpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceOp *BalanceOpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceOp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceOp *BalanceOpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceOp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceOp *BalanceOpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceOp.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_BalanceOp *BalanceOpCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalanceOp.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_BalanceOp *BalanceOpSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _BalanceOp.Contract.Balances(&_BalanceOp.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_BalanceOp *BalanceOpCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _BalanceOp.Contract.Balances(&_BalanceOp.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BalanceOp *BalanceOpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalanceOp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BalanceOp *BalanceOpSession) Owner() (common.Address, error) {
	return _BalanceOp.Contract.Owner(&_BalanceOp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BalanceOp *BalanceOpCallerSession) Owner() (common.Address, error) {
	return _BalanceOp.Contract.Owner(&_BalanceOp.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_BalanceOp *BalanceOpTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BalanceOp.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_BalanceOp *BalanceOpSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _BalanceOp.Contract.Deposit(&_BalanceOp.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_BalanceOp *BalanceOpTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _BalanceOp.Contract.Deposit(&_BalanceOp.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns()
func (_BalanceOp *BalanceOpTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalanceOp.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns()
func (_BalanceOp *BalanceOpSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalanceOp.Contract.Transfer(&_BalanceOp.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns()
func (_BalanceOp *BalanceOpTransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalanceOp.Contract.Transfer(&_BalanceOp.TransactOpts, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_BalanceOp *BalanceOpTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BalanceOp.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_BalanceOp *BalanceOpSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _BalanceOp.Contract.Withdraw(&_BalanceOp.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_BalanceOp *BalanceOpTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _BalanceOp.Contract.Withdraw(&_BalanceOp.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BalanceOp *BalanceOpTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceOp.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BalanceOp *BalanceOpSession) Receive() (*types.Transaction, error) {
	return _BalanceOp.Contract.Receive(&_BalanceOp.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BalanceOp *BalanceOpTransactorSession) Receive() (*types.Transaction, error) {
	return _BalanceOp.Contract.Receive(&_BalanceOp.TransactOpts)
}

// BalanceOpDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BalanceOp contract.
type BalanceOpDepositIterator struct {
	Event *BalanceOpDeposit // Event containing the contract specifics and raw log

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
func (it *BalanceOpDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceOpDeposit)
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
		it.Event = new(BalanceOpDeposit)
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
func (it *BalanceOpDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceOpDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceOpDeposit represents a Deposit event raised by the BalanceOp contract.
type BalanceOpDeposit struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4d6ce1e535dbade1c23defba91e23b8f791ce5edc0cc320257a2b364e4e38426.
//
// Solidity: event Deposit(uint256 amount)
func (_BalanceOp *BalanceOpFilterer) FilterDeposit(opts *bind.FilterOpts) (*BalanceOpDepositIterator, error) {

	logs, sub, err := _BalanceOp.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BalanceOpDepositIterator{contract: _BalanceOp.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

var DepositTopicHash = "0x4d6ce1e535dbade1c23defba91e23b8f791ce5edc0cc320257a2b364e4e38426"

// WatchDeposit is a free log subscription operation binding the contract event 0x4d6ce1e535dbade1c23defba91e23b8f791ce5edc0cc320257a2b364e4e38426.
//
// Solidity: event Deposit(uint256 amount)
func (_BalanceOp *BalanceOpFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BalanceOpDeposit) (event.Subscription, error) {

	logs, sub, err := _BalanceOp.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceOpDeposit)
				if err := _BalanceOp.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4d6ce1e535dbade1c23defba91e23b8f791ce5edc0cc320257a2b364e4e38426.
//
// Solidity: event Deposit(uint256 amount)
func (_BalanceOp *BalanceOpFilterer) ParseDeposit(log types.Log) (*BalanceOpDeposit, error) {
	event := new(BalanceOpDeposit)
	if err := _BalanceOp.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalanceOpTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BalanceOp contract.
type BalanceOpTransferIterator struct {
	Event *BalanceOpTransfer // Event containing the contract specifics and raw log

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
func (it *BalanceOpTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceOpTransfer)
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
		it.Event = new(BalanceOpTransfer)
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
func (it *BalanceOpTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceOpTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceOpTransfer represents a Transfer event raised by the BalanceOp contract.
type BalanceOpTransfer struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address receiver, uint256 amount)
func (_BalanceOp *BalanceOpFilterer) FilterTransfer(opts *bind.FilterOpts) (*BalanceOpTransferIterator, error) {

	logs, sub, err := _BalanceOp.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &BalanceOpTransferIterator{contract: _BalanceOp.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

var TransferTopicHash = "0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2"

// WatchTransfer is a free log subscription operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address receiver, uint256 amount)
func (_BalanceOp *BalanceOpFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalanceOpTransfer) (event.Subscription, error) {

	logs, sub, err := _BalanceOp.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceOpTransfer)
				if err := _BalanceOp.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address receiver, uint256 amount)
func (_BalanceOp *BalanceOpFilterer) ParseTransfer(log types.Log) (*BalanceOpTransfer, error) {
	event := new(BalanceOpTransfer)
	if err := _BalanceOp.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalanceOpWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the BalanceOp contract.
type BalanceOpWithdrawalIterator struct {
	Event *BalanceOpWithdrawal // Event containing the contract specifics and raw log

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
func (it *BalanceOpWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceOpWithdrawal)
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
		it.Event = new(BalanceOpWithdrawal)
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
func (it *BalanceOpWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceOpWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceOpWithdrawal represents a Withdrawal event raised by the BalanceOp contract.
type BalanceOpWithdrawal struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x4e70a604b23a8edee2b1d0a656e9b9c00b73ad8bb1afc2c59381ee9f69197de7.
//
// Solidity: event Withdrawal(uint256 amount)
func (_BalanceOp *BalanceOpFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*BalanceOpWithdrawalIterator, error) {

	logs, sub, err := _BalanceOp.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &BalanceOpWithdrawalIterator{contract: _BalanceOp.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

var WithdrawalTopicHash = "0x4e70a604b23a8edee2b1d0a656e9b9c00b73ad8bb1afc2c59381ee9f69197de7"

// WatchWithdrawal is a free log subscription operation binding the contract event 0x4e70a604b23a8edee2b1d0a656e9b9c00b73ad8bb1afc2c59381ee9f69197de7.
//
// Solidity: event Withdrawal(uint256 amount)
func (_BalanceOp *BalanceOpFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *BalanceOpWithdrawal) (event.Subscription, error) {

	logs, sub, err := _BalanceOp.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceOpWithdrawal)
				if err := _BalanceOp.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x4e70a604b23a8edee2b1d0a656e9b9c00b73ad8bb1afc2c59381ee9f69197de7.
//
// Solidity: event Withdrawal(uint256 amount)
func (_BalanceOp *BalanceOpFilterer) ParseWithdrawal(log types.Log) (*BalanceOpWithdrawal, error) {
	event := new(BalanceOpWithdrawal)
	if err := _BalanceOp.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
