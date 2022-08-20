package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("Error generating key:", err)
	}

	// Private key
	prvData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(prvData))

	// Public key
	pubData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(pubData))

	// Public address
	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
