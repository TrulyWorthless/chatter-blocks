package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
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
)

func GenerateWallet() *ecdsa.PrivateKey {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	return privateKey
}

func GetAddress(privateKey *ecdsa.PrivateKey) string {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}

// TODO: file name?
func CreateKeysStore(privateKey *ecdsa.PrivateKey, pass string) accounts.Account {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, pass)
	if err != nil {
		log.Fatal(err)
	}

	return account
}

// TODO: fix how it reads and writes
func ImportKeysStore(file, pass string) {
	ks := keystore.NewKeyStore("/wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	account, err := ks.Import(jsonBytes, pass, pass)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func CheckBalance(client *ethclient.Client, address string) *big.Float {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(21000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	tx := types.NewTransaction(nonce, destination, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	/* //raw transaction data
	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	*/

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("transaction hash: %s", signedTx.Hash().Hex())
}

// todo make not specific to channel
func DeployContract(client *ethclient.Client, privateKey *ecdsa.PrivateKey, value *big.Int) (common.Address, *types.Transaction, *simplechannel.Simplechannel) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
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
	auth.Value = value              // in wei
	auth.GasLimit = uint64(2000000) // in units TODO make dynamic
	auth.GasPrice = gasPrice

	address, tx, instance, err := simplechannel.DeploySimplechannel(auth, client, common.HexToAddress("0xE5b4C8bcA8237D9F2D7201f7bC744b697aCF8A23"), common.HexToAddress("0xe8148308ef1692e0f169F3FB8C0608a6E9625603"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\ncontract address: %s", address)

	_ = instance
	_ = tx
	return address, tx, instance
}

// todo make not specific to channel
func GetContract(client *ethclient.Client, address common.Address) *simplechannel.Simplechannel {
	instance, err := simplechannel.NewSimplechannel(address, client)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}

func SubmitTransaction(client *ethclient.Client, privateKey *ecdsa.PrivateKey, instance *simplechannel.Simplechannel, message string) *types.Transaction {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	return tx
}

//hashing new changes
// hash := sha3.NewLegacyKeccak256()
