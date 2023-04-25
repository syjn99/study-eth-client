package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const PoolCreateABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
		FromBlock: big.NewInt(8760000),
		ToBlock:   big.NewInt(8900000),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(PoolCreateABI)))
	if err != nil {
		log.Fatal(err)
	}

	var firstTopic [32]byte

	for _, vLog := range logs {
		fmt.Println("Block Number:", vLog.BlockNumber)
		fmt.Println("TxHash: ", vLog.TxHash.Hex())

		firstTopic = vLog.Topics[0]

		fmt.Println("token0 address: ", vLog.Topics[1].Hex())
		fmt.Println("token1 address: ", vLog.Topics[2].Hex())
		fmt.Println("fee: ", vLog.Topics[3].Hex())

		data, err := contractAbi.Unpack("PoolCreated", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("tickSpacing: ", data[0])
		fmt.Println("pool address: ", data[1])

	}

	eventSignature := []byte("PoolCreated(address,address,uint24,int24,address)")
	hash := crypto.Keccak256Hash(eventSignature)
	if hash.Hex() == common.Hash(firstTopic).Hex() {
		fmt.Println("First Topic and hash are equal")
		fmt.Println("First Topic: ", common.Hash(firstTopic).Hex())
		fmt.Println("Hash: ", hash.Hex())
	}
}
