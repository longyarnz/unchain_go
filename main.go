package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	unchain "github.com/longyarnz/unchain_go/contracts"
)

// Event data from log
type Event struct {
	From            common.Hash
	To              common.Hash
	Token           string
	TokenID         *big.Int
	Amount          *big.Int
	SenderBalance   *big.Int
	ReceiverBalance *big.Int
	BlockNumber     uint64
	Contract        string
}

func main() {
	wg := &sync.WaitGroup{}

	smartContract(wg)
	createServer()
}

func createServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHomeRoute).Methods("GET")
	router.HandleFunc("/logs", handleLogsRoute).Methods("GET")

	fmt.Printf("Listening on PORT: %d \n", 4000)
	log.Fatal(http.ListenAndServe(":4000", router))
}

func handleHomeRoute(writer http.ResponseWriter, reader *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "text/html")
	writer.Write([]byte("<code>Unchain NFT</code>"))
}

func handleLogsRoute(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Println("Welcome to the logs route")
}

func smartContract(wg *sync.WaitGroup) {
	rpcURL := getEnvVariable("RPC_ENDPOINT")
	contractAddress := getEnvVariable("CONTRACT_ADDRESS")
	address := common.HexToAddress(contractAddress)
	client, err := ethclient.Dial(rpcURL)

	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Listening to contract: %v \n", contractAddress)
	}

	unchainAbi, err := abi.JSON(strings.NewReader(string(unchain.UnchainMetaData.ABI)))
	if err != nil {
		log.Fatal("Invalid abi:", err)
	}

	wg.Add(1)
	go func(ch chan types.Log, wg *sync.WaitGroup) {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)

			case vLog := <-logs:
				data, err := unchainAbi.Unpack("TransferLog", vLog.Data)
				if err != nil {
					fmt.Println(err)
				} else if len(data) > 0 {
					log := Event{
						From:            vLog.Topics[1],
						To:              vLog.Topics[2],
						Token:           data[0].(string),
						TokenID:         data[1].(*big.Int),
						Amount:          data[2].(*big.Int),
						SenderBalance:   data[3].(*big.Int),
						ReceiverBalance: data[4].(*big.Int),
						BlockNumber:     vLog.BlockNumber,
						Contract:        contractAddress,
					}

					fmt.Println(log)
				}
			}
		}
	}(logs, wg)
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
