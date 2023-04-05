package web3

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	simplechannel "github.com/trulyworthless/chatter-blocks/bindings"
	"github.com/trulyworthless/chatter-blocks/pkg/filesystem"
)

func GenerateBlockchainAddressFile(name string, address common.Address) {
	//TODO use common.Address
	file, _ := json.MarshalIndent(address, "", " ")
	jsonfile, err := filesystem.CreateFile("exports", name, ".json")
	if err != nil {
		panic(err)
	}

	jsonfile.Write(file)
}

func RetrieveBlockchainAddressFromFile(name string) common.Address {
	var address common.Address
	jsonfile, err := filesystem.OpenFile("contacts", name, ".json")
	if err != nil {
		panic(err)
	}

	jsonfileinfo, _ := jsonfile.Stat()
	var size int64 = jsonfileinfo.Size()
	jsonbytes := make([]byte, size)

	buffer := bufio.NewReader(jsonfile)
	_, err = buffer.Read(jsonbytes)
	if err != nil {
		panic(err)
	}

	//TODO use commo.Address
	json.Unmarshal(jsonbytes, &address)

	return address
}

func GenerateEOA() *ecdsa.PrivateKey {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	return privateKey
}

func GetBalance(client *ethclient.Client, address string) *big.Float {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue
}

func TransferETH(client *ethclient.Client, privateKey *ecdsa.PrivateKey, value *big.Int, destination common.Address) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Errorf("web3"))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	gasLimit := uint64(21000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	var data []byte
	tx := types.NewTransaction(nonce, destination, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		panic(err)
	}

	/* //raw transaction data
	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}
	*/

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}
}

// todo make not specific to channel
func DeployContract(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (common.Address, *types.Transaction, *simplechannel.Simplechannel) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Errorf("web3"))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(2000000) // in units TODO make dynamic
	auth.GasPrice = gasPrice

	address, tx, instance, err := simplechannel.DeploySimplechannel(auth, client, common.HexToAddress("0xE5b4C8bcA8237D9F2D7201f7bC744b697aCF8A23"), common.HexToAddress("0xe8148308ef1692e0f169F3FB8C0608a6E9625603"))
	if err != nil {
		panic(err)
	}

	return address, tx, instance
}

// todo make not specific to channel
func GetContract(client *ethclient.Client, address common.Address) *simplechannel.Simplechannel {
	instance, err := simplechannel.NewSimplechannel(address, client)
	if err != nil {
		panic(err)
	}

	return instance
}

func SubmitTransaction(client *ethclient.Client, privateKey *ecdsa.PrivateKey, instance *simplechannel.Simplechannel, message string) *types.Transaction {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Errorf("web3"))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(2000000) // in units TODO make dynamic
	auth.GasPrice = gasPrice

	tx, err := instance.SubmitMessage(auth, message)
	if err != nil {
		panic(err)
	}

	return tx
}

func EmitSubscribe() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("env loaded")
	}

	apiKEY := os.Getenv("ALCHEMY_API_KEY")
	client, err := ethclient.Dial("wss://eth-mainnet.g.alchemy.com/v2/" + apiKEY)
	if err != nil {
		panic(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				panic(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
