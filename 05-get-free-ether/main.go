package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// import (
// 	"log"

// 	"github.com/ethereum/go-ethereum/accounts/keystore"
// )

var (
	url  = "https://kovan.infura.io/v3/c1f511c8b9ed45f095ef00b69e87b758"
	mUrl = "https://mainnet.infura.io/v3/c1f511c8b9ed45f095ef00b69e87b758"
)

func main() {
	// ----------------------------------------------------------------------------------------------
	// !!! Notes: Whole commented codes are used for generating UTC-... file in wallet folder
	// N – iterations count (affects memory and CPU usage), e.g. 16384 or 2048
	// p – parallelism factor (threads to run in parallel - affects the memory, CPU usage), usually 1
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// password := "password"
	// _, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal("Error in generating: ", err)
	// }
	// _, err1 := key.NewAccount(password)
	// if err1 != nil {
	// 	log.Fatal("Error in generating: ", err)
	// }
	// run 02 times go run main.go for generating 02 files UTC--
	// b052ea54b934ea996de291694e3ec337ce1288de
	// 9e9c8e9b9fd5a0e3ace06a12b09cccc3372eedcb
	// !!! Notes: add 0x prefix for searching and sending eth for testing
	// Like: 0x9e9c8e9b9fd5a0e3ace06a12b09cccc3372eedcb
	// Get Kovan Network eth for testing at: https://ethdrop.dev/
	// ----------------------------------------------------------------------------------------------

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	address1 := common.HexToAddress("0x9e9c8e9b9fd5a0e3ace06a12b09cccc3372eedcb")
	address2 := common.HexToAddress("0xb052ea54b934ea996de291694e3ec337ce1288de")

	balance1, err := client.BalanceAt(context.Background(), address1, nil)
	if err != nil {
		log.Fatal("Get balance at address 01: ", err)
	}
	fmt.Println("Balance 01: ", balance1)

	balance2, err := client.BalanceAt(context.Background(), address2, nil)
	if err != nil {
		log.Fatal("Get balance at address 02: ", err)
	}
	fmt.Println("Balance 02: ", balance2)
}
