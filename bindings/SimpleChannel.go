// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// ISimpleChannelMessage is an auto generated low-level Go binding around an user-defined struct.
type ISimpleChannelMessage struct {
	Sender    common.Address
	Text      string
	Timestamp *big.Int
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumISimpleChannel.Actor\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"MessageSubmit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"Zeek\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readResponse\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structISimpleChannel.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"readResponseAt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structISimpleChannel.Message\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readThread\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structISimpleChannel.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"}],\"name\":\"submitMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006001553480156200005757600080fd5b5060405162001cee38038062001cee83398181016040528101906200007d91906200026d565b81600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460405180606001604052803073ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600981526020017f696e697420636861740000000000000000000000000000000000000000000000815250815260200142815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081620001ee91906200052e565b50604082015181600201555050505062000615565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620002358262000208565b9050919050565b620002478162000228565b81146200025357600080fd5b50565b60008151905062000267816200023c565b92915050565b6000806040838503121562000287576200028662000203565b5b6000620002978582860162000256565b9250506020620002aa8582860162000256565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200033657607f821691505b6020821081036200034c576200034b620002ee565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003b67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000377565b620003c2868362000377565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200040f620004096200040384620003da565b620003e4565b620003da565b9050919050565b6000819050919050565b6200042b83620003ee565b620004436200043a8262000416565b84845462000384565b825550505050565b600090565b6200045a6200044b565b6200046781848462000420565b505050565b5b818110156200048f576200048360008262000450565b6001810190506200046d565b5050565b601f821115620004de57620004a88162000352565b620004b38462000367565b81016020851015620004c3578190505b620004db620004d28562000367565b8301826200046c565b50505b505050565b600082821c905092915050565b60006200050360001984600802620004e3565b1980831691505092915050565b60006200051e8383620004f0565b9150826002028217905092915050565b6200053982620002b4565b67ffffffffffffffff811115620005555762000554620002bf565b5b6200056182546200031d565b6200056e82828562000493565b600060209050601f831160018114620005a6576000841562000591578287015190505b6200059d858262000510565b8655506200060d565b601f198416620005b68662000352565b60005b82811015620005e057848901518255600182019150602085019450602081019050620005b9565b86831015620006005784890151620005fc601f891682620004f0565b8355505b6001600288020188555050505b505050505050565b6116c980620006256000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806327f811d6146100675780633dbcc8d114610097578063708b34fe146100b55780639db91092146100e5578063afebf50214610103578063d46347bb14610121575b600080fd5b610081600480360381019061007c9190610c06565b610151565b60405161008e9190610d63565b60405180910390f35b61009f6103b0565b6040516100ac9190610d94565b60405180910390f35b6100cf60048036038101906100ca9190610ee4565b6103bd565b6040516100dc9190610f48565b60405180910390f35b6100ed6105c8565b6040516100fa9190611075565b60405180910390f35b61010b610801565b6040516101189190610d63565b60405180910390f35b61013b600480360381019061013691906110c3565b610b0d565b604051610148919061113a565b60405180910390f35b610159610b85565b3373ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061020257503373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610241576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610238906111ce565b60405180910390fd5b6102496103b0565b821061028a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102819061123a565b60405180910390fd5b6004828154811061029e5761029d61125a565b5b90600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461031d906112b8565b80601f0160208091040260200160405190810160405280929190818152602001828054610349906112b8565b80156103965780601f1061036b57610100808354040283529160200191610396565b820191906000526020600020905b81548152906001019060200180831161037957829003601f168201915b505050505081526020016002820154815250509050919050565b6000600480549050905090565b60003373ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061046857503373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6104a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049e906111ce565b60405180910390fd5b600460405180606001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184815260200142815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161055f9190611495565b506040820151816002015550507f1d3f6866315a304121810f27b07b22280edd2d79eae157fad4f357ee97b1e44660016105976103b0565b6105a19190611596565b6105a9610b26565b6040516105b7929190611641565b60405180910390a160019050919050565b60603373ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061067357503373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6106b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a9906111ce565b60405180910390fd5b6004805480602002602001604051908101604052809291908181526020016000905b828210156107f857838290600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461075d906112b8565b80601f0160208091040260200160405190810160405280929190818152602001828054610789906112b8565b80156107d65780601f106107ab576101008083540402835291602001916107d6565b820191906000526020600020905b8154815290600101906020018083116107b957829003601f168201915b50505050508152602001600282015481525050815260200190600101906106d4565b50505050905090565b610809610b85565b3373ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614806108b257503373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6108f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e8906111ce565b60405180910390fd5b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160405180602001604052806000815250815260200160008152509050600060016109416103b0565b61094b9190611596565b90505b60008110610b08573373ffffffffffffffffffffffffffffffffffffffff16600482815481106109815761098061125a565b5b906000526020600020906003020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610af557600481815481106109e2576109e161125a565b5b90600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610a61906112b8565b80601f0160208091040260200160405190810160405280929190818152602001828054610a8d906112b8565b8015610ada5780601f10610aaf57610100808354040283529160200191610ada565b820191906000526020600020905b815481529060010190602001808311610abd57829003601f168201915b50505050508152602001600282015481525050915050610b0a565b8080610b009061166a565b91505061094e565b505b90565b6060604051806020016040528060008152509050919050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1603610b8257600190505b90565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600081525090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610be381610bd0565b8114610bee57600080fd5b50565b600081359050610c0081610bda565b92915050565b600060208284031215610c1c57610c1b610bc6565b5b6000610c2a84828501610bf1565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c5e82610c33565b9050919050565b610c6e81610c53565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610cae578082015181840152602081019050610c93565b60008484015250505050565b6000601f19601f8301169050919050565b6000610cd682610c74565b610ce08185610c7f565b9350610cf0818560208601610c90565b610cf981610cba565b840191505092915050565b610d0d81610bd0565b82525050565b6000606083016000830151610d2b6000860182610c65565b5060208301518482036020860152610d438282610ccb565b9150506040830151610d586040860182610d04565b508091505092915050565b60006020820190508181036000830152610d7d8184610d13565b905092915050565b610d8e81610bd0565b82525050565b6000602082019050610da96000830184610d85565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610df182610cba565b810181811067ffffffffffffffff82111715610e1057610e0f610db9565b5b80604052505050565b6000610e23610bbc565b9050610e2f8282610de8565b919050565b600067ffffffffffffffff821115610e4f57610e4e610db9565b5b610e5882610cba565b9050602081019050919050565b82818337600083830152505050565b6000610e87610e8284610e34565b610e19565b905082815260208101848484011115610ea357610ea2610db4565b5b610eae848285610e65565b509392505050565b600082601f830112610ecb57610eca610daf565b5b8135610edb848260208601610e74565b91505092915050565b600060208284031215610efa57610ef9610bc6565b5b600082013567ffffffffffffffff811115610f1857610f17610bcb565b5b610f2484828501610eb6565b91505092915050565b60008115159050919050565b610f4281610f2d565b82525050565b6000602082019050610f5d6000830184610f39565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000606083016000830151610fa76000860182610c65565b5060208301518482036020860152610fbf8282610ccb565b9150506040830151610fd46040860182610d04565b508091505092915050565b6000610feb8383610f8f565b905092915050565b6000602082019050919050565b600061100b82610f63565b6110158185610f6e565b93508360208202850161102785610f7f565b8060005b8581101561106357848403895281516110448582610fdf565b945061104f83610ff3565b925060208a0199505060018101905061102b565b50829750879550505050505092915050565b6000602082019050818103600083015261108f8184611000565b905092915050565b6110a081610f2d565b81146110ab57600080fd5b50565b6000813590506110bd81611097565b92915050565b6000602082840312156110d9576110d8610bc6565b5b60006110e7848285016110ae565b91505092915050565b600082825260208201905092915050565b600061110c82610c74565b61111681856110f0565b9350611126818560208601610c90565b61112f81610cba565b840191505092915050565b600060208201905081810360008301526111548184611101565b905092915050565b7f4368616e6e656c3a2063616c6c6572206973206e6f742061206d657373616e6760008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b60006111b86022836110f0565b91506111c38261115c565b604082019050919050565b600060208201905081810360008301526111e7816111ab565b9050919050565b7f4368616e6e656c3a20696e646578206f7574206f6620626f756e647300000000600082015250565b6000611224601c836110f0565b915061122f826111ee565b602082019050919050565b6000602082019050818103600083015261125381611217565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806112d057607f821691505b6020821081036112e3576112e2611289565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261134b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261130e565b611355868361130e565b95508019841693508086168417925050509392505050565b6000819050919050565b600061139261138d61138884610bd0565b61136d565b610bd0565b9050919050565b6000819050919050565b6113ac83611377565b6113c06113b882611399565b84845461131b565b825550505050565b600090565b6113d56113c8565b6113e08184846113a3565b505050565b5b81811015611404576113f96000826113cd565b6001810190506113e6565b5050565b601f8211156114495761141a816112e9565b611423846112fe565b81016020851015611432578190505b61144661143e856112fe565b8301826113e5565b50505b505050565b600082821c905092915050565b600061146c6000198460080261144e565b1980831691505092915050565b6000611485838361145b565b9150826002028217905092915050565b61149e82610c74565b67ffffffffffffffff8111156114b7576114b6610db9565b5b6114c182546112b8565b6114cc828285611408565b600060209050601f8311600181146114ff57600084156114ed578287015190505b6114f78582611479565b86555061155f565b601f19841661150d866112e9565b60005b8281101561153557848901518255600182019150602085019450602081019050611510565b86831015611552578489015161154e601f89168261145b565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006115a182610bd0565b91506115ac83610bd0565b92508282039050818111156115c4576115c3611567565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811061160a576116096115ca565b5b50565b600081905061161b826115f9565b919050565b600061162b8261160d565b9050919050565b61163b81611620565b82525050565b60006040820190506116566000830185610d85565b6116636020830184611632565b9392505050565b600061167582610bd0565b91506000820361168857611687611567565b5b60018203905091905056fea2646970667358221220dccfd83f36d8780880e5e6ec48969d953f7686e35176399d52aa0b1d907e3b6e64736f6c63430008130033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, a common.Address, b common.Address) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, a, b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Zeek is a free data retrieval call binding the contract method 0xd46347bb.
//
// Solidity: function Zeek(bool ) pure returns(string)
func (_Store *StoreCaller) Zeek(opts *bind.CallOpts, arg0 bool) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "Zeek", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Zeek is a free data retrieval call binding the contract method 0xd46347bb.
//
// Solidity: function Zeek(bool ) pure returns(string)
func (_Store *StoreSession) Zeek(arg0 bool) (string, error) {
	return _Store.Contract.Zeek(&_Store.CallOpts, arg0)
}

// Zeek is a free data retrieval call binding the contract method 0xd46347bb.
//
// Solidity: function Zeek(bool ) pure returns(string)
func (_Store *StoreCallerSession) Zeek(arg0 bool) (string, error) {
	return _Store.Contract.Zeek(&_Store.CallOpts, arg0)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Store *StoreCaller) MessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "messageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Store *StoreSession) MessageCount() (*big.Int, error) {
	return _Store.Contract.MessageCount(&_Store.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Store *StoreCallerSession) MessageCount() (*big.Int, error) {
	return _Store.Contract.MessageCount(&_Store.CallOpts)
}

// ReadResponse is a free data retrieval call binding the contract method 0xafebf502.
//
// Solidity: function readResponse() view returns((address,string,uint256) message)
func (_Store *StoreCaller) ReadResponse(opts *bind.CallOpts) (ISimpleChannelMessage, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "readResponse")

	if err != nil {
		return *new(ISimpleChannelMessage), err
	}

	out0 := *abi.ConvertType(out[0], new(ISimpleChannelMessage)).(*ISimpleChannelMessage)

	return out0, err

}

// ReadResponse is a free data retrieval call binding the contract method 0xafebf502.
//
// Solidity: function readResponse() view returns((address,string,uint256) message)
func (_Store *StoreSession) ReadResponse() (ISimpleChannelMessage, error) {
	return _Store.Contract.ReadResponse(&_Store.CallOpts)
}

// ReadResponse is a free data retrieval call binding the contract method 0xafebf502.
//
// Solidity: function readResponse() view returns((address,string,uint256) message)
func (_Store *StoreCallerSession) ReadResponse() (ISimpleChannelMessage, error) {
	return _Store.Contract.ReadResponse(&_Store.CallOpts)
}

// ReadResponseAt is a free data retrieval call binding the contract method 0x27f811d6.
//
// Solidity: function readResponseAt(uint256 index) view returns((address,string,uint256))
func (_Store *StoreCaller) ReadResponseAt(opts *bind.CallOpts, index *big.Int) (ISimpleChannelMessage, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "readResponseAt", index)

	if err != nil {
		return *new(ISimpleChannelMessage), err
	}

	out0 := *abi.ConvertType(out[0], new(ISimpleChannelMessage)).(*ISimpleChannelMessage)

	return out0, err

}

// ReadResponseAt is a free data retrieval call binding the contract method 0x27f811d6.
//
// Solidity: function readResponseAt(uint256 index) view returns((address,string,uint256))
func (_Store *StoreSession) ReadResponseAt(index *big.Int) (ISimpleChannelMessage, error) {
	return _Store.Contract.ReadResponseAt(&_Store.CallOpts, index)
}

// ReadResponseAt is a free data retrieval call binding the contract method 0x27f811d6.
//
// Solidity: function readResponseAt(uint256 index) view returns((address,string,uint256))
func (_Store *StoreCallerSession) ReadResponseAt(index *big.Int) (ISimpleChannelMessage, error) {
	return _Store.Contract.ReadResponseAt(&_Store.CallOpts, index)
}

// ReadThread is a free data retrieval call binding the contract method 0x9db91092.
//
// Solidity: function readThread() view returns((address,string,uint256)[])
func (_Store *StoreCaller) ReadThread(opts *bind.CallOpts) ([]ISimpleChannelMessage, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "readThread")

	if err != nil {
		return *new([]ISimpleChannelMessage), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISimpleChannelMessage)).(*[]ISimpleChannelMessage)

	return out0, err

}

// ReadThread is a free data retrieval call binding the contract method 0x9db91092.
//
// Solidity: function readThread() view returns((address,string,uint256)[])
func (_Store *StoreSession) ReadThread() ([]ISimpleChannelMessage, error) {
	return _Store.Contract.ReadThread(&_Store.CallOpts)
}

// ReadThread is a free data retrieval call binding the contract method 0x9db91092.
//
// Solidity: function readThread() view returns((address,string,uint256)[])
func (_Store *StoreCallerSession) ReadThread() ([]ISimpleChannelMessage, error) {
	return _Store.Contract.ReadThread(&_Store.CallOpts)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x708b34fe.
//
// Solidity: function submitMessage(string text) returns(bool)
func (_Store *StoreTransactor) SubmitMessage(opts *bind.TransactOpts, text string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "submitMessage", text)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x708b34fe.
//
// Solidity: function submitMessage(string text) returns(bool)
func (_Store *StoreSession) SubmitMessage(text string) (*types.Transaction, error) {
	return _Store.Contract.SubmitMessage(&_Store.TransactOpts, text)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x708b34fe.
//
// Solidity: function submitMessage(string text) returns(bool)
func (_Store *StoreTransactorSession) SubmitMessage(text string) (*types.Transaction, error) {
	return _Store.Contract.SubmitMessage(&_Store.TransactOpts, text)
}

// StoreMessageSubmitIterator is returned from FilterMessageSubmit and is used to iterate over the raw logs and unpacked data for MessageSubmit events raised by the Store contract.
type StoreMessageSubmitIterator struct {
	Event *StoreMessageSubmit // Event containing the contract specifics and raw log

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
func (it *StoreMessageSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreMessageSubmit)
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
		it.Event = new(StoreMessageSubmit)
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
func (it *StoreMessageSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreMessageSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreMessageSubmit represents a MessageSubmit event raised by the Store contract.
type StoreMessageSubmit struct {
	Index *big.Int
	Role  uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMessageSubmit is a free log retrieval operation binding the contract event 0x1d3f6866315a304121810f27b07b22280edd2d79eae157fad4f357ee97b1e446.
//
// Solidity: event MessageSubmit(uint256 index, uint8 role)
func (_Store *StoreFilterer) FilterMessageSubmit(opts *bind.FilterOpts) (*StoreMessageSubmitIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "MessageSubmit")
	if err != nil {
		return nil, err
	}
	return &StoreMessageSubmitIterator{contract: _Store.contract, event: "MessageSubmit", logs: logs, sub: sub}, nil
}

// WatchMessageSubmit is a free log subscription operation binding the contract event 0x1d3f6866315a304121810f27b07b22280edd2d79eae157fad4f357ee97b1e446.
//
// Solidity: event MessageSubmit(uint256 index, uint8 role)
func (_Store *StoreFilterer) WatchMessageSubmit(opts *bind.WatchOpts, sink chan<- *StoreMessageSubmit) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "MessageSubmit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreMessageSubmit)
				if err := _Store.contract.UnpackLog(event, "MessageSubmit", log); err != nil {
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

// ParseMessageSubmit is a log parse operation binding the contract event 0x1d3f6866315a304121810f27b07b22280edd2d79eae157fad4f357ee97b1e446.
//
// Solidity: event MessageSubmit(uint256 index, uint8 role)
func (_Store *StoreFilterer) ParseMessageSubmit(log types.Log) (*StoreMessageSubmit, error) {
	event := new(StoreMessageSubmit)
	if err := _Store.contract.UnpackLog(event, "MessageSubmit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
