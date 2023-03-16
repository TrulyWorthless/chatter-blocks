package web3

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	simplechannel "github.com/trulyworthless/chatter-blocks/bindings"
	"github.com/trulyworthless/chatter-blocks/pkg/env"
)

func GenerateEOA() *ecdsa.PrivateKey {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	return privateKey
}

// TODO: file name?
func CreateKeysStore(privateKey *ecdsa.PrivateKey, pass string) accounts.Account {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, pass)
	if err != nil {
		panic(err)
	}

	return account
}

// TODO: fix how it reads and writes
func ImportKeysStore(file, pass string) accounts.Account {
	ks := keystore.NewKeyStore("/wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	// jsonBytes, err := ioutil.ReadFile(file)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	account, err := ks.Import(jsonBytes, pass, pass)
	if err != nil {
		panic(err)
	}

	if err := os.Remove(file); err != nil {
		panic(err)
	}
	return account
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

func CreateKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func ImportKs() {
	file := "tmp/UTC--2023-02-19T04-51-14.589991000Z--db56cf470ee09e536016b87428ed52f7d8b60c7f"
	ks := keystore.NewKeyStore("/tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	// jsonBytes, err := ioutil.ReadFile(file)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		panic(err)
	}
}

func EmitSubscribe() {
	apiKEY := env.GoDotEnv("ALCHEMY_API_KEY")
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
