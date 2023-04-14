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

func GenerateRSAPrivateKeyToBytes() []byte {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	return x509.MarshalPKCS1PrivateKey(privatekey)
}

func GetRSAPrivateKeyFromBytes(keyBytes []byte) *rsa.PrivateKey {
	privatekey, err := x509.ParsePKCS1PrivateKey(keyBytes)
	if err != nil {
		panic(err)
	}

	return privatekey
}

func GenerateRSAPublicKeyFile(fileName string, publicKey *rsa.PublicKey) error {
	pemfile, err := filesystem.CreateFile("contacts", fileName, ".pem")
	if err != nil {
		return err
	}

	defer pemfile.Close()

	var pemkey = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey)}

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveRSAPublicKeyFromFile(fileName string) []byte {
	publicKeyFile, err := filesystem.OpenFile("contacts", fileName, ".pem")
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

	//TODO at risk of bad state
	publicKeyFile.Close()
	err = filesystem.DeleteFile("contacts/", fileName, ".pem")
	if err != nil {
		panic(err)
	}

	return data.Bytes
}

func RetrieveRSAPublicKeyFromBytes(pembytes []byte) *rsa.PublicKey {
	publicKey, err := x509.ParsePKCS1PublicKey(pembytes)
	if err != nil {
		panic(err)
	}

	return publicKey
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
