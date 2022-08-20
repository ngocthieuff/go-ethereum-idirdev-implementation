package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// import (
// 	"log"

// 	// "github.com/ethereum/go-ethereum/accounts/keystore"
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
	// run 02 times go run main.go for generating 02 files UTC--
	// b3ab37b783d5e18fae4915104826823d5eece8eb
	// 41bcefc86c5756a49261909d8a91733fe9c1c1b3
	// ----------------------------------------------------------------------------------------------

	client, err := ethclient.Dial(mUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	address1 := common.HexToAddress("b3ab37b783d5e18fae4915104826823d5eece8eb")
	address2 := common.HexToAddress("41bcefc86c5756a49261909d8a91733fe9c1c1b3")

	balance1, err := client.BalanceAt(context.Background(), address1, nil)
	if err != nil {
		log.Fatal("Get balance at address 01: ", err)
	}
	println("Balance 01: ", balance1)

	balance2, err := client.BalanceAt(context.Background(), address2, nil)
	if err != nil {
		log.Fatal("Get balance at address 02: ", err)
	}
	println("Balance 02: ", balance2)
}
