package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cPath := flag.String("c", "config/config.json", "path to config for micro-service")
	flag.Parse()
	if *cPath == "" {
		log.Fatal("No config provided")
	}
	cfg := LoadConfig(*cPath)

	clientURL := fmt.Sprintf("https://mainnet.infura.io/v3/%s", cfg.ProjectID)
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(cfg.Address)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("B: (%v)\n", balance)
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)
}
