package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	todo "test-deploy/gen"
)

var (
	url = "https://kovan.infura.io/v3/c1f511c8b9ed45f095ef00b69e87b758"
)

func main() {
	b, err := ioutil.ReadFile("../wallet/UTC--2022-08-22T02-55-25.435095600Z--9e9c8e9b9fd5a0e3ace06a12b09cccc3372eedcb")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Error at get SuggestGasPrice ", err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("Error at get NetworkID ", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		log.Fatal("Error at get NewKeyStoreTransactorWithChainID ", err)
	}

	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(111352)
	auth.Nonce = big.NewInt(int64(nonce))

	a, tx, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatal("Error at get NewKeyStoreTransactorWithChainID ", err)
	}

	fmt.Println("--------------------------")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("--------------------------")
}

// go run main.go
// --------------------------
// 0x8F9CE0AA5554EcAC4B8e9Af67f81A93DB5d9Cd80
// 0x2b6e3e81e9550f75816b9244b4decc8deb70400a4e7a814f73ebe20a14fc3403
// 0x5b962B2BA2c41902951d70d5579cfE1a3Ba8C8F3
// 0xcefab2391d2276b4d6207f535e3171eae7b68201cf18b4611ebceae1d56eb88d
// --------------------------
