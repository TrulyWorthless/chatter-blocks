package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")

	privateKeyAlice, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKeyAlice := privateKeyAlice.PublicKey
	fmt.Println(publicKeyAlice, privateKeyAlice)

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKeyAlice,
		[]byte("super secret message"),
		nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("encrypted bytes: ", encryptedBytes)

	decryptBytes, err := privateKeyAlice.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})

	if err != nil {
		panic(err)
	}

	fmt.Println("decrypted message: ", string(decryptBytes))

}
