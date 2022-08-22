package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://kovan.infura.io/v3/c1f511c8b9ed45f095ef00b69e87b758"
)

func main() {
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

	nonce, err := client.PendingNonceAt(context.Background(), address1)
	if err != nil {
		log.Fatal("Error at get PendingNonceAt address 01: ", err)
	}
	// 1 ether = 10^9 Gwei
	// 1 ether = 10^18 wei
	amount := big.NewInt(1000000000)

	// Error at SendTransaction Transaction gas is too low. There is not enough gas to cover minimal cost of the transaction (minimal: 21000,
	// got: 2210). Try increasing supplied gas.
	gasLimit := uint64(22100)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Error at get SuggestGasPrice ", err)
	}

	transaction := types.NewTransaction(nonce, address2, amount, gasLimit, gasPrice, nil)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("Error at get NetworkID ", err)
	}

	b, err := ioutil.ReadFile("wallet/UTC--2022-08-22T02-55-25.435095600Z--9e9c8e9b9fd5a0e3ace06a12b09cccc3372eedcb")
	if err != nil {
		log.Fatal("Error at ReadFile ", err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal("Error at DecryptKey ", err)
	}

	signedTransaction, err := types.SignTx(transaction, types.NewEIP155Signer(chainId), key.PrivateKey)
	if err != nil {
		log.Fatal("Error at SignTx ", err)
	}

	err = client.SendTransaction(context.Background(), signedTransaction)
	if err != nil {
		log.Fatal("Error at SendTransaction ", err)
	}
}
