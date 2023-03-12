package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client
	GenerateKey()
	privateKey := GetKey()
	encryptedBytes := Encrypt(privateKey, "super secret message")
	decryptedMessage := Decrypt(privateKey, encryptedBytes)
	fmt.Println("decrypted message: ", decryptedMessage)
}
