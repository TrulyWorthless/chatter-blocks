// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simplechannel

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

// SimplechannelMetaData contains all meta data concerning the Simplechannel contract.
var SimplechannelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumISimpleChannel.Actor\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"MessageSubmit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"Zeek\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readResponse\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structISimpleChannel.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"readResponseAt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structISimpleChannel.Message\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readThread\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structISimpleChannel.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"}],\"name\":\"submitMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006001553480156200005757600080fd5b5060405162001c8f38038062001c8f83398181016040528101906200007d91906200026d565b81600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460405180606001604052803073ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600981526020017f696e697420636861740000000000000000000000000000000000000000000000815250815260200142815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081620001ee91906200052e565b50604082015181600201555050505062000615565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620002358262000208565b9050919050565b620002478162000228565b81146200025357600080fd5b50565b60008151905062000267816200023c565b92915050565b6000806040838503121562000287576200028662000203565b5b6000620002978582860162000256565b9250506020620002aa8582860162000256565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200033657607f821691505b6020821081036200034c576200034b620002ee565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003b67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000377565b620003c2868362000377565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200040f620004096200040384620003da565b620003e4565b620003da565b9050919050565b6000819050919050565b6200042b83620003ee565b620004436200043a8262000416565b84845462000384565b825550505050565b600090565b6200045a6200044b565b6200046781848462000420565b505050565b5b818110156200048f576200048360008262000450565b6001810190506200046d565b5050565b601f821115620004de57620004a88162000352565b620004b38462000367565b81016020851015620004c3578190505b620004db620004d28562000367565b8301826200046c565b50505b505050565b600082821c905092915050565b60006200050360001984600802620004e3565b1980831691505092915050565b60006200051e8383620004f0565b9150826002028217905092915050565b6200053982620002b4565b67ffffffffffffffff811115620005555762000554620002bf565b5b6200056182546200031d565b6200056e82828562000493565b600060209050601f831160018114620005a6576000841562000591578287015190505b6200059d858262000510565b8655506200060d565b601f198416620005b68662000352565b60005b82811015620005e057848901518255600182019150602085019450602081019050620005b9565b86831015620006005784890151620005fc601f891682620004f0565b8355505b6001600288020188555050505b505050505050565b61166a80620006256000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806327f811d6146100675780633dbcc8d114610097578063708b34fe146100b55780639db91092146100e5578063afebf50214610103578063d46347bb14610121575b600080fd5b610081600480360381019061007c9190610c18565b610151565b60405161008e9190610d75565b60405180910390f35b61009f6103b0565b6040516100ac9190610da6565b60405180910390f35b6100cf60048036038101906100ca9190610ef6565b6103bd565b6040516100dc9190610f5a565b60405180910390f35b6100ed6105da565b6040516100fa9190611087565b60405180910390f35b61010b610813565b6040516101189190610d75565b60405180910390f35b61013b600480360381019061013691906110d5565b610b1f565b604051610148919061114c565b60405180910390f35b610159610b97565b3373ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061020257503373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610241576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610238906111e0565b60405180910390fd5b6102496103b0565b821061028a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102819061124c565b60405180910390fd5b6004828154811061029e5761029d61126c565b5b90600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461031d906112ca565b80601f0160208091040260200160405190810160405280929190818152602001828054610349906112ca565b80156103965780601f1061036b57610100808354040283529160200191610396565b820191906000526020600020905b81548152906001019060200180831161037957829003601f168201915b505050505081526020016002820154815250509050919050565b6000600480549050905090565b60003373ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061046857503373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6104a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049e906111e0565b60405180910390fd5b600460405180606001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184815260200142815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161055f91906114a7565b50604082015181600201555050610574610b38565b600181111561058657610585611579565b5b60016105906103b0565b61059a91906115d7565b7fefd2c33b8cf50ad0141e32984c1a365510fcbb88c3bd84e7c158628aa4177765846040516105c9919061114c565b60405180910390a360019050919050565b60603373ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061068557503373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6106c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bb906111e0565b60405180910390fd5b6004805480602002602001604051908101604052809291908181526020016000905b8282101561080a57838290600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461076f906112ca565b80601f016020809104026020016040519081016040528092919081815260200182805461079b906112ca565b80156107e85780601f106107bd576101008083540402835291602001916107e8565b820191906000526020600020905b8154815290600101906020018083116107cb57829003601f168201915b50505050508152602001600282015481525050815260200190600101906106e6565b50505050905090565b61081b610b97565b3373ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614806108c457503373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610903576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108fa906111e0565b60405180910390fd5b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160405180602001604052806000815250815260200160008152509050600060016109536103b0565b61095d91906115d7565b90505b60008110610b1a573373ffffffffffffffffffffffffffffffffffffffff16600482815481106109935761099261126c565b5b906000526020600020906003020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b0757600481815481106109f4576109f361126c565b5b90600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610a73906112ca565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9f906112ca565b8015610aec5780601f10610ac157610100808354040283529160200191610aec565b820191906000526020600020905b815481529060010190602001808311610acf57829003601f168201915b50505050508152602001600282015481525050915050610b1c565b8080610b129061160b565b915050610960565b505b90565b6060604051806020016040528060008152509050919050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1603610b9457600190505b90565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600081525090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610bf581610be2565b8114610c0057600080fd5b50565b600081359050610c1281610bec565b92915050565b600060208284031215610c2e57610c2d610bd8565b5b6000610c3c84828501610c03565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c7082610c45565b9050919050565b610c8081610c65565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610cc0578082015181840152602081019050610ca5565b60008484015250505050565b6000601f19601f8301169050919050565b6000610ce882610c86565b610cf28185610c91565b9350610d02818560208601610ca2565b610d0b81610ccc565b840191505092915050565b610d1f81610be2565b82525050565b6000606083016000830151610d3d6000860182610c77565b5060208301518482036020860152610d558282610cdd565b9150506040830151610d6a6040860182610d16565b508091505092915050565b60006020820190508181036000830152610d8f8184610d25565b905092915050565b610da081610be2565b82525050565b6000602082019050610dbb6000830184610d97565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610e0382610ccc565b810181811067ffffffffffffffff82111715610e2257610e21610dcb565b5b80604052505050565b6000610e35610bce565b9050610e418282610dfa565b919050565b600067ffffffffffffffff821115610e6157610e60610dcb565b5b610e6a82610ccc565b9050602081019050919050565b82818337600083830152505050565b6000610e99610e9484610e46565b610e2b565b905082815260208101848484011115610eb557610eb4610dc6565b5b610ec0848285610e77565b509392505050565b600082601f830112610edd57610edc610dc1565b5b8135610eed848260208601610e86565b91505092915050565b600060208284031215610f0c57610f0b610bd8565b5b600082013567ffffffffffffffff811115610f2a57610f29610bdd565b5b610f3684828501610ec8565b91505092915050565b60008115159050919050565b610f5481610f3f565b82525050565b6000602082019050610f6f6000830184610f4b565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000606083016000830151610fb96000860182610c77565b5060208301518482036020860152610fd18282610cdd565b9150506040830151610fe66040860182610d16565b508091505092915050565b6000610ffd8383610fa1565b905092915050565b6000602082019050919050565b600061101d82610f75565b6110278185610f80565b93508360208202850161103985610f91565b8060005b8581101561107557848403895281516110568582610ff1565b945061106183611005565b925060208a0199505060018101905061103d565b50829750879550505050505092915050565b600060208201905081810360008301526110a18184611012565b905092915050565b6110b281610f3f565b81146110bd57600080fd5b50565b6000813590506110cf816110a9565b92915050565b6000602082840312156110eb576110ea610bd8565b5b60006110f9848285016110c0565b91505092915050565b600082825260208201905092915050565b600061111e82610c86565b6111288185611102565b9350611138818560208601610ca2565b61114181610ccc565b840191505092915050565b600060208201905081810360008301526111668184611113565b905092915050565b7f4368616e6e656c3a2063616c6c6572206973206e6f742061206d657373616e6760008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b60006111ca602283611102565b91506111d58261116e565b604082019050919050565b600060208201905081810360008301526111f9816111bd565b9050919050565b7f4368616e6e656c3a20696e646578206f7574206f6620626f756e647300000000600082015250565b6000611236601c83611102565b915061124182611200565b602082019050919050565b6000602082019050818103600083015261126581611229565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806112e257607f821691505b6020821081036112f5576112f461129b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261135d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611320565b6113678683611320565b95508019841693508086168417925050509392505050565b6000819050919050565b60006113a461139f61139a84610be2565b61137f565b610be2565b9050919050565b6000819050919050565b6113be83611389565b6113d26113ca826113ab565b84845461132d565b825550505050565b600090565b6113e76113da565b6113f28184846113b5565b505050565b5b818110156114165761140b6000826113df565b6001810190506113f8565b5050565b601f82111561145b5761142c816112fb565b61143584611310565b81016020851015611444578190505b61145861145085611310565b8301826113f7565b50505b505050565b600082821c905092915050565b600061147e60001984600802611460565b1980831691505092915050565b6000611497838361146d565b9150826002028217905092915050565b6114b082610c86565b67ffffffffffffffff8111156114c9576114c8610dcb565b5b6114d382546112ca565b6114de82828561141a565b600060209050601f83116001811461151157600084156114ff578287015190505b611509858261148b565b865550611571565b601f19841661151f866112fb565b60005b8281101561154757848901518255600182019150602085019450602081019050611522565b868310156115645784890151611560601f89168261146d565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006115e282610be2565b91506115ed83610be2565b9250828203905081811115611605576116046115a8565b5b92915050565b600061161682610be2565b915060008203611629576116286115a8565b5b60018203905091905056fea26469706673582212205bcce4145a0c8643e0303ace2fb7cf48e3f1c834689336384d93044776a61a7864736f6c63430008130033",
}

// SimplechannelABI is the input ABI used to generate the binding from.
// Deprecated: Use SimplechannelMetaData.ABI instead.
var SimplechannelABI = SimplechannelMetaData.ABI

// SimplechannelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimplechannelMetaData.Bin instead.
var SimplechannelBin = SimplechannelMetaData.Bin

// DeploySimplechannel deploys a new Ethereum contract, binding an instance of Simplechannel to it.
func DeploySimplechannel(auth *bind.TransactOpts, backend bind.ContractBackend, a common.Address, b common.Address) (common.Address, *types.Transaction, *Simplechannel, error) {
	parsed, err := SimplechannelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimplechannelBin), backend, a, b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simplechannel{SimplechannelCaller: SimplechannelCaller{contract: contract}, SimplechannelTransactor: SimplechannelTransactor{contract: contract}, SimplechannelFilterer: SimplechannelFilterer{contract: contract}}, nil
}

// Simplechannel is an auto generated Go binding around an Ethereum contract.
type Simplechannel struct {
	SimplechannelCaller     // Read-only binding to the contract
	SimplechannelTransactor // Write-only binding to the contract
	SimplechannelFilterer   // Log filterer for contract events
}

// SimplechannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplechannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplechannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplechannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplechannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplechannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplechannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplechannelSession struct {
	Contract     *Simplechannel    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimplechannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplechannelCallerSession struct {
	Contract *SimplechannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimplechannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplechannelTransactorSession struct {
	Contract     *SimplechannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimplechannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplechannelRaw struct {
	Contract *Simplechannel // Generic contract binding to access the raw methods on
}

// SimplechannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplechannelCallerRaw struct {
	Contract *SimplechannelCaller // Generic read-only contract binding to access the raw methods on
}

// SimplechannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplechannelTransactorRaw struct {
	Contract *SimplechannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplechannel creates a new instance of Simplechannel, bound to a specific deployed contract.
func NewSimplechannel(address common.Address, backend bind.ContractBackend) (*Simplechannel, error) {
	contract, err := bindSimplechannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simplechannel{SimplechannelCaller: SimplechannelCaller{contract: contract}, SimplechannelTransactor: SimplechannelTransactor{contract: contract}, SimplechannelFilterer: SimplechannelFilterer{contract: contract}}, nil
}

// NewSimplechannelCaller creates a new read-only instance of Simplechannel, bound to a specific deployed contract.
func NewSimplechannelCaller(address common.Address, caller bind.ContractCaller) (*SimplechannelCaller, error) {
	contract, err := bindSimplechannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplechannelCaller{contract: contract}, nil
}

// NewSimplechannelTransactor creates a new write-only instance of Simplechannel, bound to a specific deployed contract.
func NewSimplechannelTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplechannelTransactor, error) {
	contract, err := bindSimplechannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplechannelTransactor{contract: contract}, nil
}

// NewSimplechannelFilterer creates a new log filterer instance of Simplechannel, bound to a specific deployed contract.
func NewSimplechannelFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplechannelFilterer, error) {
	contract, err := bindSimplechannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplechannelFilterer{contract: contract}, nil
}

// bindSimplechannel binds a generic wrapper to an already deployed contract.
func bindSimplechannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimplechannelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplechannel *SimplechannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplechannel.Contract.SimplechannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplechannel *SimplechannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplechannel.Contract.SimplechannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplechannel *SimplechannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplechannel.Contract.SimplechannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplechannel *SimplechannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplechannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplechannel *SimplechannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplechannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplechannel *SimplechannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplechannel.Contract.contract.Transact(opts, method, params...)
}

// Zeek is a free data retrieval call binding the contract method 0xd46347bb.
//
// Solidity: function Zeek(bool ) pure returns(string)
func (_Simplechannel *SimplechannelCaller) Zeek(opts *bind.CallOpts, arg0 bool) (string, error) {
	var out []interface{}
	err := _Simplechannel.contract.Call(opts, &out, "Zeek", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Zeek is a free data retrieval call binding the contract method 0xd46347bb.
//
// Solidity: function Zeek(bool ) pure returns(string)
func (_Simplechannel *SimplechannelSession) Zeek(arg0 bool) (string, error) {
	return _Simplechannel.Contract.Zeek(&_Simplechannel.CallOpts, arg0)
}

// Zeek is a free data retrieval call binding the contract method 0xd46347bb.
//
// Solidity: function Zeek(bool ) pure returns(string)
func (_Simplechannel *SimplechannelCallerSession) Zeek(arg0 bool) (string, error) {
	return _Simplechannel.Contract.Zeek(&_Simplechannel.CallOpts, arg0)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Simplechannel *SimplechannelCaller) MessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simplechannel.contract.Call(opts, &out, "messageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Simplechannel *SimplechannelSession) MessageCount() (*big.Int, error) {
	return _Simplechannel.Contract.MessageCount(&_Simplechannel.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Simplechannel *SimplechannelCallerSession) MessageCount() (*big.Int, error) {
	return _Simplechannel.Contract.MessageCount(&_Simplechannel.CallOpts)
}

// ReadResponse is a free data retrieval call binding the contract method 0xafebf502.
//
// Solidity: function readResponse() view returns((address,string,uint256) message)
func (_Simplechannel *SimplechannelCaller) ReadResponse(opts *bind.CallOpts) (ISimpleChannelMessage, error) {
	var out []interface{}
	err := _Simplechannel.contract.Call(opts, &out, "readResponse")

	if err != nil {
		return *new(ISimpleChannelMessage), err
	}

	out0 := *abi.ConvertType(out[0], new(ISimpleChannelMessage)).(*ISimpleChannelMessage)

	return out0, err

}

// ReadResponse is a free data retrieval call binding the contract method 0xafebf502.
//
// Solidity: function readResponse() view returns((address,string,uint256) message)
func (_Simplechannel *SimplechannelSession) ReadResponse() (ISimpleChannelMessage, error) {
	return _Simplechannel.Contract.ReadResponse(&_Simplechannel.CallOpts)
}

// ReadResponse is a free data retrieval call binding the contract method 0xafebf502.
//
// Solidity: function readResponse() view returns((address,string,uint256) message)
func (_Simplechannel *SimplechannelCallerSession) ReadResponse() (ISimpleChannelMessage, error) {
	return _Simplechannel.Contract.ReadResponse(&_Simplechannel.CallOpts)
}

// ReadResponseAt is a free data retrieval call binding the contract method 0x27f811d6.
//
// Solidity: function readResponseAt(uint256 index) view returns((address,string,uint256))
func (_Simplechannel *SimplechannelCaller) ReadResponseAt(opts *bind.CallOpts, index *big.Int) (ISimpleChannelMessage, error) {
	var out []interface{}
	err := _Simplechannel.contract.Call(opts, &out, "readResponseAt", index)

	if err != nil {
		return *new(ISimpleChannelMessage), err
	}

	out0 := *abi.ConvertType(out[0], new(ISimpleChannelMessage)).(*ISimpleChannelMessage)

	return out0, err

}

// ReadResponseAt is a free data retrieval call binding the contract method 0x27f811d6.
//
// Solidity: function readResponseAt(uint256 index) view returns((address,string,uint256))
func (_Simplechannel *SimplechannelSession) ReadResponseAt(index *big.Int) (ISimpleChannelMessage, error) {
	return _Simplechannel.Contract.ReadResponseAt(&_Simplechannel.CallOpts, index)
}

// ReadResponseAt is a free data retrieval call binding the contract method 0x27f811d6.
//
// Solidity: function readResponseAt(uint256 index) view returns((address,string,uint256))
func (_Simplechannel *SimplechannelCallerSession) ReadResponseAt(index *big.Int) (ISimpleChannelMessage, error) {
	return _Simplechannel.Contract.ReadResponseAt(&_Simplechannel.CallOpts, index)
}

// ReadThread is a free data retrieval call binding the contract method 0x9db91092.
//
// Solidity: function readThread() view returns((address,string,uint256)[])
func (_Simplechannel *SimplechannelCaller) ReadThread(opts *bind.CallOpts) ([]ISimpleChannelMessage, error) {
	var out []interface{}
	err := _Simplechannel.contract.Call(opts, &out, "readThread")

	if err != nil {
		return *new([]ISimpleChannelMessage), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISimpleChannelMessage)).(*[]ISimpleChannelMessage)

	return out0, err

}

// ReadThread is a free data retrieval call binding the contract method 0x9db91092.
//
// Solidity: function readThread() view returns((address,string,uint256)[])
func (_Simplechannel *SimplechannelSession) ReadThread() ([]ISimpleChannelMessage, error) {
	return _Simplechannel.Contract.ReadThread(&_Simplechannel.CallOpts)
}

// ReadThread is a free data retrieval call binding the contract method 0x9db91092.
//
// Solidity: function readThread() view returns((address,string,uint256)[])
func (_Simplechannel *SimplechannelCallerSession) ReadThread() ([]ISimpleChannelMessage, error) {
	return _Simplechannel.Contract.ReadThread(&_Simplechannel.CallOpts)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x708b34fe.
//
// Solidity: function submitMessage(string text) returns(bool)
func (_Simplechannel *SimplechannelTransactor) SubmitMessage(opts *bind.TransactOpts, text string) (*types.Transaction, error) {
	return _Simplechannel.contract.Transact(opts, "submitMessage", text)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x708b34fe.
//
// Solidity: function submitMessage(string text) returns(bool)
func (_Simplechannel *SimplechannelSession) SubmitMessage(text string) (*types.Transaction, error) {
	return _Simplechannel.Contract.SubmitMessage(&_Simplechannel.TransactOpts, text)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x708b34fe.
//
// Solidity: function submitMessage(string text) returns(bool)
func (_Simplechannel *SimplechannelTransactorSession) SubmitMessage(text string) (*types.Transaction, error) {
	return _Simplechannel.Contract.SubmitMessage(&_Simplechannel.TransactOpts, text)
}

// SimplechannelMessageSubmitIterator is returned from FilterMessageSubmit and is used to iterate over the raw logs and unpacked data for MessageSubmit events raised by the Simplechannel contract.
type SimplechannelMessageSubmitIterator struct {
	Event *SimplechannelMessageSubmit // Event containing the contract specifics and raw log

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
func (it *SimplechannelMessageSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplechannelMessageSubmit)
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
		it.Event = new(SimplechannelMessageSubmit)
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
func (it *SimplechannelMessageSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplechannelMessageSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplechannelMessageSubmit represents a MessageSubmit event raised by the Simplechannel contract.
type SimplechannelMessageSubmit struct {
	Index   *big.Int
	Role    uint8
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMessageSubmit is a free log retrieval operation binding the contract event 0xefd2c33b8cf50ad0141e32984c1a365510fcbb88c3bd84e7c158628aa4177765.
//
// Solidity: event MessageSubmit(uint256 indexed index, uint8 indexed role, string message)
func (_Simplechannel *SimplechannelFilterer) FilterMessageSubmit(opts *bind.FilterOpts, index []*big.Int, role []uint8) (*SimplechannelMessageSubmitIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Simplechannel.contract.FilterLogs(opts, "MessageSubmit", indexRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &SimplechannelMessageSubmitIterator{contract: _Simplechannel.contract, event: "MessageSubmit", logs: logs, sub: sub}, nil
}

// WatchMessageSubmit is a free log subscription operation binding the contract event 0xefd2c33b8cf50ad0141e32984c1a365510fcbb88c3bd84e7c158628aa4177765.
//
// Solidity: event MessageSubmit(uint256 indexed index, uint8 indexed role, string message)
func (_Simplechannel *SimplechannelFilterer) WatchMessageSubmit(opts *bind.WatchOpts, sink chan<- *SimplechannelMessageSubmit, index []*big.Int, role []uint8) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Simplechannel.contract.WatchLogs(opts, "MessageSubmit", indexRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplechannelMessageSubmit)
				if err := _Simplechannel.contract.UnpackLog(event, "MessageSubmit", log); err != nil {
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

// ParseMessageSubmit is a log parse operation binding the contract event 0xefd2c33b8cf50ad0141e32984c1a365510fcbb88c3bd84e7c158628aa4177765.
//
// Solidity: event MessageSubmit(uint256 indexed index, uint8 indexed role, string message)
func (_Simplechannel *SimplechannelFilterer) ParseMessageSubmit(log types.Log) (*SimplechannelMessageSubmit, error) {
	event := new(SimplechannelMessageSubmit)
	if err := _Simplechannel.contract.UnpackLog(event, "MessageSubmit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
