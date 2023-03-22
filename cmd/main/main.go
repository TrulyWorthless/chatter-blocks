package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/trulyworthless/chatter-blocks/pkg/crypt"
	"github.com/trulyworthless/chatter-blocks/pkg/env"
	"github.com/trulyworthless/chatter-blocks/pkg/web3"
)

type Profile struct {
	name            string
	client          *ethclient.Client
	privateRSAKey   *rsa.PrivateKey
	privateECDSAKey *ecdsa.PrivateKey
}

func main() {
	fmt.Println("\nWelcome to Chatter Blocks!!!")
	fmt.Println("Are you logging in or creating an identity?")
	fmt.Println("Type 'login' or 'signup'")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()

		switch scanner.Text() {
		// TODO: add inital config option
		case "login", "l":
			login("")
			return
		case "signup", "sign up", "s":
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
			crypt.GenerateRSAPrivateKeyFile(name)
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
	privateRSAKey := crypt.RetrieveRSAPrivateKey(name)

	//load ecdsa key
	ganacheKEY := env.GoDotEnv(name)
	privateECDSAKey, err := crypto.HexToECDSA(ganacheKEY)
	if err != nil {
		log.Fatal(err)
	}

	// construct(client, privateRSAKey, privateECDSAKey)
	logic(Profile{name: name, client: client, privateRSAKey: privateRSAKey, privateECDSAKey: privateECDSAKey})
}

func logic(profile Profile) {
	fmt.Println("Would you like to do?")
	fmt.Println("A. Create new contact")
	fmt.Println("B. Create a new chat for an existing contract")
	fmt.Println("C. Chatter with an existing contact")
	fmt.Println("D. Export an identity")

	for {
		var command string
		fmt.Scanln(&command)

		switch command {
		case "A", "a":
			createContact()
		case "B", "b":
			address := env.GoDotEnv("CHANNEL_ADDRESS")
			message(profile.client, profile.privateRSAKey, profile.privateECDSAKey, web3.GetContract(profile.client, common.HexToAddress(address)))
		case "C", "c":
			_, _, instance := web3.DeployContract(profile.client, profile.privateECDSAKey)
			message(profile.client, profile.privateRSAKey, profile.privateECDSAKey, instance)
		case "D", "d":
			export(profile)
		default:
			fmt.Printf("I'm sorry, but I dont know that response")
		}
	}
}

// func createcontact() {
// 	fmt.Println("Lets create a new contact!")
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Println("What is your contact's name?")
// 	scanner.Scan()
// 	contactName := scanner.Text()

// 	fmt.Println("What is your contact's blockchain address?")
// 	scanner.Scan()
// 	contactBlockchainAddress := scanner.Text()
// 	web3.GenerateBlockchainAddressFile(contactName, common.HexToAddress(contactBlockchainAddress))

// 	//eventually streamline with imports and blockchain downloads
// 	fmt.Println("What is your contact's public key?")
// 	fmt.Println("First lets ask for the modulus:")
// 	scanner.Scan()
// 	modulus := scanner.Bytes()

// 	fmt.Println("Next lets have the public exponent:")
// 	scanner.Scan()
// 	exponent := scanner.Bytes()

// 	contactPublicKey := rsa.PublicKey{
// 		N: big.NewInt(0).SetBytes(modulus),
// 		E: int(big.NewInt(0).SetBytes(exponent).Int64()),
// 	}

// 	crypt.GenerateRSAPublicKeyFile(contactName, &contactPublicKey)

// 	fmt.Printf("Contract: %s", web3.RetrieveBlockchainAddress(contactName))
// }

// func construct(client *ethclient.Client, privateRSAKey *rsa.PrivateKey, privateECDSAKey *ecdsa.PrivateKey) {
// 	fmt.Println("Would you like to create a new thread? (y/n)")

// 	for {
// 		var command string
// 		fmt.Scanln(&command)

// 		switch command {
// 		case "Y", "y":
// 			_, _, instance := web3.DeployContract(client, privateECDSAKey)
// 			message(client, privateRSAKey, privateECDSAKey, instance)
// 		case "N", "n":
// 			address := env.GoDotEnv("CHANNEL_ADDRESS")
// 			message(client, privateRSAKey, privateECDSAKey, web3.GetContract(client, common.HexToAddress(address)))
// 		default:
// 			fmt.Printf("I'm sorry, but I dont know that response")
// 		}
// 	}
// }

// func message(client *ethclient.Client, privateRSAKey *rsa.PrivateKey, privateECDSAKey *ecdsa.PrivateKey, contract *simplechannel.Simplechannel) {
// 	//info
// 	publicKey := privateECDSAKey.Public()
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

// 	index := big.NewInt(1)
// 	fmt.Println("What would you like to say?")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		scanner.Scan()
// 		message := scanner.Text()

// 		encryptedMessage := crypt.EncryptMessage(privateRSAKey, message)
// 		web3.SubmitTransaction(client, privateECDSAKey, contract, string(encryptedMessage))

// 		response, err := contract.ReadResponseAt(&a, index)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		encryptedResponse := []byte(response.Text)

// 		decryptedResponse := crypt.DecryptMessage(privateRSAKey, encryptedResponse)
// 		fmt.Println("decrypted message: ", decryptedResponse)
// 		index = index.Add(index, big.NewInt(1))
// 	}
// }
