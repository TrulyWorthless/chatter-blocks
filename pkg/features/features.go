package features

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	simplechannel "github.com/trulyworthless/chatter-blocks/bindings"
	"github.com/trulyworthless/chatter-blocks/pkg/config"
	"github.com/trulyworthless/chatter-blocks/pkg/crypt"
	"github.com/trulyworthless/chatter-blocks/pkg/filesystem"
	"github.com/trulyworthless/chatter-blocks/pkg/web3"
)

func export(profile Profile) {
	//export public key
	pemfile, err := filesystem.CreateFile("exports", profile.name, ".pem")
	if err != nil {
		panic(err)
	}

	var pemkey = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&profile.privateRSAKey.PublicKey)}

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		panic(err)
	}

	//export blockchain address
	//TODO make with goeth version
	file, _ := json.MarshalIndent(profile.privateECDSAKey.PublicKey, "", " ")
	jsonfile, err := filesystem.CreateFile("exports", profile.name, ".json")
	if err != nil {
		panic(err)
	}

	jsonfile.Write(file)
}

func createContact() {
	fmt.Println("Lets create a new contact!")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("What is your contact's name?")
	scanner.Scan()
	contactName := scanner.Text()

	fmt.Println("What is your contact's blockchain address?")
	scanner.Scan()
	contactBlockchainAddress := scanner.Text()
	web3.GenerateBlockchainAddressFile(contactName, common.HexToAddress(contactBlockchainAddress))

	//eventually streamline with imports and blockchain downloads
	fmt.Println("What is your contact's public key?")
	fmt.Println("First lets ask for the modulus:")
	//TODO does not return when given value, gets stuck
	scanner.Scan()
	modulus := scanner.Bytes()

	fmt.Println("Next lets have the public exponent:")
	scanner.Scan()
	exponent := scanner.Bytes()

	contactPublicKey := rsa.PublicKey{
		N: big.NewInt(0).SetBytes(modulus),
		E: int(big.NewInt(0).SetBytes(exponent).Int64()),
	}

	crypt.GenerateRSAPublicKeyFile(contactName, &contactPublicKey)

	fmt.Println("Contract Created!")
}

func construct(client *ethclient.Client, privateRSAKey *rsa.PrivateKey, privateECDSAKey *ecdsa.PrivateKey) {
	fmt.Println("Would you like to create a new thread? (y/n)")

	for {
		var command string
		fmt.Scanln(&command)

		switch command {
		case "Y", "y":
			_, _, instance := web3.DeployContract(client, privateECDSAKey)
			message(client, privateRSAKey, privateECDSAKey, instance)
		case "N", "n":
			address := config.Config("CHANNEL_ADDRESS")
			message(client, privateRSAKey, privateECDSAKey, web3.GetContract(client, common.HexToAddress(address)))
		default:
			fmt.Printf("I'm sorry, but I dont know that response")
		}
	}
}

func listen(client *ethclient.Client, privateRSAKey *rsa.PrivateKey, privateECDSAKey *ecdsa.PrivateKey, contract common.Address) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contract},
	}

	fmt.Println("Shhh... We're listening in...")

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		panic(err)
	}

	s := web3.GetContract(client, contract)

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case vLog := <-logs:
			c, err := s.ParseMessageSubmit(vLog)
			if err != nil {
				panic(err)
			}

			// fmt.Printf("%s", c.Message)
			encryptedResponse := []byte(c.Message)

			decryptedResponse := crypt.DecryptMessage(privateRSAKey, encryptedResponse)
			fmt.Println("decrypted message: ", decryptedResponse)
		}
	}
}

func message(client *ethclient.Client, privateRSAKey *rsa.PrivateKey, privateECDSAKey *ecdsa.PrivateKey, contract *simplechannel.Simplechannel) {
	//info
	//topic 0xefd2c33b8cf50ad0141e32984c1a365510fcbb88c3bd84e7c158628aa4177765
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
			panic(err)
		}

		encryptedResponse := []byte(response.Text)

		decryptedResponse := crypt.DecryptMessage(privateRSAKey, encryptedResponse)
		fmt.Println("decrypted message: ", decryptedResponse)
		index = index.Add(index, big.NewInt(1))
	}
}

// func message(client *ethclient.Client, privateRSAKey *rsa.PrivateKey, privateECDSAKey *ecdsa.PrivateKey, contract *simplechannel.Simplechannel) {
// 	//info
// 	//topic 0xefd2c33b8cf50ad0141e32984c1a365510fcbb88c3bd84e7c158628aa4177765
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
