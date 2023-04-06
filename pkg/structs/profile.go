package structs

import (
	"crypto/ecdsa"
	"crypto/rsa"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Profile struct {
	httpClient      *ethclient.Client
	wsClient        *ethclient.Client
	privateRSAKey   *rsa.PrivateKey
	privateECDSAKey *ecdsa.PrivateKey
}
