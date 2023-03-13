package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/trulyworthless/chatter-blocks/pkg/env"
)

func main() {
	// apiKEY := env.GoDotEnvVariable("API_KEY")
	// client, err := ethclient.Dial("https://mainnet.gateway.tenderly.co/" + apiKEY)
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client

	// key flow
	GenerateKey()
	privateRSAKey := GetKey()
	encryptedBytes := Encrypt(privateRSAKey, "super secret message")
	decryptedMessage := Decrypt(privateRSAKey, encryptedBytes)
	fmt.Println("decrypted message: ", decryptedMessage)

	// //wallet flow
	// privateECDSAKey := GenerateWallet()
	// CreateKeysStore(privateECDSAKey, "password")
	// ImportKeysStore("wallets/please", "password")

	//transaction flow
	//dummy destination
	bKEY := env.GoDotEnvVariable("G_KEY")
	privateKey, err := crypto.HexToECDSA(bKEY)
	if err != nil {
		log.Fatal(err)
	}
	transferETH(client, privateKey, big.NewInt(1), common.HexToAddress("0xe8148308ef1692e0f169F3FB8C0608a6E9625603"))

}

// createKs()
// importKs()
// emitContract()
// emitBlocks()
// emitTransactions()
// emitTransfer()
// emitTokens()
// emitSubscribe()
// contractDeploy()
// loadContract()
// readContract()
