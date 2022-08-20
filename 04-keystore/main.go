package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// ----------------------------------------------------------------------------------------------
	// !!! Notes: Whole commented codes are used for generating UTC-... file in wallet folder
	// N – iterations count (affects memory and CPU usage), e.g. 16384 or 2048
	// p – parallelism factor (threads to run in parallel - affects the memory, CPU usage), usually 1
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// password := "password"
	// acc, err := key.NewAccount(password)

	// if err != nil {
	// 	log.Fatal("Error in generating: ", err)
	// }
	// log.Fatal(acc.Address)
	// ----------------------------------------------------------------------------------------------

	// After generating the UTC file
	password := "password"
	b, err := ioutil.ReadFile("./wallet/UTC--2022-08-20T05-29-16.058730900Z--f410208c8b63c02734200fbf6759e60fe261fd91")
	if err != nil {
		log.Fatal("Error in reading file phase: ", err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal("Error in decrypting key: ", err)
	}
	pvk := key.PrivateKey
	// Private key
	prvData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(prvData))

	// Public key
	pubData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(pubData))

	// Public address
	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
