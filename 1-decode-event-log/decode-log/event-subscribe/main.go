package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	infuraWebsocketApiKey := os.Getenv("INFURA_WEBSOCKET_API_KEY")
	if infuraWebsocketApiKey == "" {
		log.Fatal("INFURA_WEBSOCKET_API_KEY is not set")
	}

	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/" + infuraWebsocketApiKey)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xD40F0EeeAfb58C8d18E6b0469EFF02f0B737583d")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
