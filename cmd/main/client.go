package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	simplechannel "github.com/trulyworthless/chatter-blocks/bindings"
	"github.com/trulyworthless/chatter-blocks/pkg/crypt"
	"github.com/trulyworthless/chatter-blocks/pkg/env"
	"github.com/trulyworthless/chatter-blocks/pkg/web3"
)

func main() {
	fmt.Println("\nWelcome to Chatter Blocks!!!")
	fmt.Println("Are you logging in or creating an identity?")
	fmt.Println("Type 'login' or 'signup'")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()

		switch scanner.Text() {
		// TODO: add inital config option
		case "login":
			login("")
			return
		case "signup", "sign up":
			signUp()
			return
		default:
			fmt.Println("I'm sorry, but I dont know that response")
		}
	}
}

func signUp() {
	fmt.Println("Whats your name?")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		name := scanner.Text()

		if name == "alice" || name == "bob" {
			crypt.GenerateRSAKeyFile(name)
			login(name)
			return
		} else {
			fmt.Println("Sorry, could you repeat that? (alice/bob)")
		}
	}
}

func login(name string) {
	if len(name) == 0 {
		fmt.Println("Whats your name?")
		scanner := bufio.NewScanner(os.Stdin)

		for {
			scanner.Scan()
			newName := scanner.Text()

			if newName == "alice" || newName == "bob" {
				name = newName
				break
			} else {
				fmt.Println("Sorry, could you repeat that? (alice/bob)")
			}
		}
	}

	//connect
	url := "HTTP://127.0.0.1:7545"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	//load rsa key
	privateRSAKey := crypt.RetrieveRSAKey(name)

	//load ecdsa key
	ganacheKEY := env.GoDotEnv(name)
	privateECDSAKey, err := crypto.HexToECDSA(ganacheKEY)
	if err != nil {
		log.Fatal(err)
	}

	construct(client, privateRSAKey, privateECDSAKey)
}

func construct(client *ethclient.Client, privateRSAKey *rsa.PrivateKey, privateECDSAKey *ecdsa.PrivateKey) {
	fmt.Println("Would you like to create a new thread? (y/n)")

	for {
		var command string
		fmt.Scanln(&command)

		switch command {
		case "Y", "y":
			_, _, instance := web3.DeployContract(client, privateECDSAKey)
			fmt.Println("Temporary closed")
			_ = instance
			return
			// message(client, privateRSAKey, privateECDSAKey, instance)
		case "N", "n":
			address := env.GoDotEnv("CHANNEL_ADDRESS")
			message(client, privateRSAKey, privateECDSAKey, web3.GetContract(client, common.HexToAddress(address)))
		default:
			fmt.Printf("I'm sorry, but I dont know that response")
		}
	}
}

func message(client *ethclient.Client, privateRSAKey *rsa.PrivateKey, privateECDSAKey *ecdsa.PrivateKey, contract *simplechannel.Simplechannel) {
	//info
	publicKey := privateECDSAKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	a := bind.CallOpts{
		Pending:     false,
		From:        fromAddress,
		BlockNumber: nil,
		Context:     context.Background(),
	}

	index := big.NewInt(1)
	fmt.Println("What would you like to say?")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		message := scanner.Text()

		encryptedMessage := crypt.EncryptMessage(privateRSAKey, message)
		web3.SubmitTransaction(client, privateECDSAKey, contract, string(encryptedMessage))

		response, err := contract.ReadResponseAt(&a, index)
		if err != nil {
			log.Fatal(err)
		}

		encryptedResponse := []byte(response.Text)

		decryptedResponse := crypt.DecryptMessage(privateRSAKey, encryptedResponse)
		fmt.Println("decrypted message: ", decryptedResponse)
		index = index.Add(index, big.NewInt(1))
	}
}

// func origin() {
// 	// apiKEY := env.GoDotEnv("API_KEY")
// 	// client, err := ethclient.Dial("https://mainnet.gateway.tenderly.co/" + apiKEY)
// 	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("we have a connection")
// 	_ = client

// 	// key flow
// 	GenerateKey("alice")
// 	privateRSAKey := GetKey("alice")
// 	encryptedBytes := Encrypt(privateRSAKey, "super secret message again")
// 	decryptedMessage := Decrypt(privateRSAKey, encryptedBytes)
// 	fmt.Println("decrypted message: ", decryptedMessage)

// 	// //wallet flow
// 	// privateECDSAKey := GenerateWallet()
// 	// CreateKeysStore(privateECDSAKey, "password")
// 	// ImportKeysStore("wallets/please", "password")

// 	//transaction flow
// 	//dummy destination
// 	gKEY := env.GoDotEnv("G_KEY")
// 	privateKey, err := crypto.HexToECDSA(gKEY)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	TransferETH(client, privateKey, big.NewInt(1), common.HexToAddress("0xe8148308ef1692e0f169F3FB8C0608a6E9625603"))

// 	//contract interactions
// 	address, _, _ := DeployContract(client, privateKey, big.NewInt(0))
// 	instance := GetContract(client, address)

// 	tx := SubmitTransaction(client, privateKey, instance, "Hello World")

// 	_ = tx

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("error casting public key to ECDSA")
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	a := bind.CallOpts{
// 		Pending:     false,
// 		From:        fromAddress,
// 		BlockNumber: nil,
// 		Context:     context.Background(),
// 	}

// 	message, err := instance.ReadResponseAt(&a, big.NewInt(1))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("\nmessage: %s", message)

// 	// submit encrypted message
// 	// myString := string(encryptedBytes[:])
// 	// fmt.Printf("\nmessage: %s", myString)
// 	// fmt.Printf("\nmessage: %x", myString)

// 	// fmt.Printf("\nmessage: %s", encryptedBytes)
// 	// fmt.Printf("\nmessage: %x", encryptedBytes)

// 	tx2 := SubmitTransaction(client, privateKey, instance, string(encryptedBytes))

// 	_ = tx2

// 	message2, err := instance.ReadResponseAt(&a, big.NewInt(2))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("\nmessage: %s", message2)

// 	b := []byte(message2.Text)

// 	decryptedMessage2 := Decrypt(privateRSAKey, b)
// 	fmt.Println("\ndecrypted message2: ", decryptedMessage2)
// }
