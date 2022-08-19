package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraUrl = "https://mainnet.infura.io/v3/c1f511c8b9ed45f095ef00b69e87b758"

func main() {

	// Background returns a non-nil, empty Context.
	// It is never canceled, has no values, and has no deadline.
	// It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
	client, err := ethclient.DialContext(context.Background(), infuraUrl)

	if err != nil {
		// %v verb means to use the default format which can be overridden
		log.Fatal("Error to create a ether client: %v, err")
	}

	defer client.Close()

	// parameters include id of block
	// return latest block if id is nil
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("Error to get a block: %v, err")
	}

	// you can check result at https://etherscan.io/blocks
	fmt.Printf(block.Number().String())
}
