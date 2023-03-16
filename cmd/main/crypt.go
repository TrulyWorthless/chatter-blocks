package main

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateKey(name string) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	pemfile, err := os.Create(name + ".pem")
	if err != nil {
		panic(err)
	}
	defer pemfile.Close()

	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		panic(err)
	}
}

func GetKey(name string) *rsa.PrivateKey {
	fileName := name + ".pem"
	fmt.Println(fileName)
	privateKeyFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer privateKeyFile.Close()

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		panic(err)
	}

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func Encrypt(privateKey *rsa.PrivateKey, message string) []byte {
	publicKey := privateKey.PublicKey

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte(message),
		nil)
	if err != nil {
		panic(err)
	}

	return encryptedBytes
}

func Decrypt(privateKey *rsa.PrivateKey, encryptedBytes []byte) string {
	decryptBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}

	return string(decryptBytes)
}
