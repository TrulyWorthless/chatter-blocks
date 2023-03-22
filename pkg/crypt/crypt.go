package crypt

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"

	"github.com/trulyworthless/chatter-blocks/pkg/filesystem"
)

func GenerateRSAPrivateKeyFile(name string) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	pemfile, err := filesystem.CreateFile("identities", name, ".pem")
	if err != nil {
		panic(err)
	}

	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		panic(err)
	}
}

func RetrieveRSAPrivateKey(name string) *rsa.PrivateKey {
	privateKeyFile, err := filesystem.OpenFile("identities", name, ".pem")
	if err != nil {
		panic(err)
	}

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

func GenerateRSAPublicKeyFile(name string, publicKey *rsa.PublicKey) {
	pemfile, err := filesystem.CreateFile("contacts", name, ".pem")
	if err != nil {
		panic(err)
	}

	var pemkey = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey)}

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		panic(err)
	}
}

func RetrieveRSAPublicKey(name string) *rsa.PublicKey {
	publicKeyFile, err := filesystem.OpenFile("contacts", name, ".pem")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		panic(err)
	}

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyImported, err := x509.ParsePKCS1PublicKey(data.Bytes)
	if err != nil {
		panic(err)
	}

	return publicKeyImported
}

func EncryptMessage(privateKey *rsa.PrivateKey, message string) []byte {
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

func DecryptMessage(privateKey *rsa.PrivateKey, encryptedBytes []byte) string {
	decryptBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}

	return string(decryptBytes)
}
