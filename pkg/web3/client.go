package web3

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/trulyworthless/chatter-blocks/pkg/config"
)

var HttpClient *ethclient.Client
var WsClient *ethclient.Client

func Init() {
	//http
	// httpClient, err := ethclient.Dial(config.Config("HTTP_PROVIDER_URL_SEPOLIA") + config.Config("TENDERLY_API_KEY"))
	httpClient, err := ethclient.Dial(config.Config("AVALANCHE_RPC_FUJI"))
	if err != nil {
		panic(err)
	}

	HttpClient = httpClient
}
