package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraUrl = "https://mainnet.infura.io/v3/c1f511c8b9ed45f095ef00b69e87b758"
var ganacheUrl = "http://localhost:8545"

func main() {

	// Background returns a non-nil, empty Context.
	// It is never canceled, has no values, and has no deadline.
	// It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
	client, err := ethclient.DialContext(context.Background(), infuraUrl)

	// ganache is local blockchain
	// use ganacheUrl for interaction with local blockchain
	// client, err := ethclient.DialContext(context.Background(), ganacheUrl)

	if err != nil {
		// %v verb means to use the default format which can be overridden
		log.Fatal("Error to create a ether client: %v", err)
	}

	defer client.Close()

	// parameters include id of block
	// return latest block if id is nil
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("Error to get a block: %v")
	}

	// you can check result at https://etherscan.io/blocks
	fmt.Println("Block number: ", block.Number().String())

	addr := "0x3EcEf08D0e2DaD803847E052249bb4F8bFf2D5bB"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal("Error to get a block: %v")
	}
	fmt.Println("Balance: ", balance)

	// 1 ether = 10^18 wei
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println(fBalance)

	balanceEther := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(balanceEther)
}
