// Copyright © 2018-2020 Satinderjit Singh.
//
// See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
// the top-level directory of this distribution for the individual copyright
// holder information and the developer policies on copyright and licensing.
//
// Unless otherwise agreed in a custom licensing agreement, no part of the
// kmdgo software, including this file may be copied, modified, propagated.
// or distributed except according to the terms contained in the LICENSE file
//
// Removal or modification of this copyright notice is prohibited.

package main

import (
	"fmt"
	"log"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `DEX`

	var get kmdgo.DEXGet

	args := make(kmdgo.APIParams, 1)
	args[0] = "38289120"
	// args[1] = "0138d849d6bc81ff1c5389aae9a60ba3ee9cfd7858d93a3864679c25937e70951f"
	// args[2] = "BTC"
	// args[3] = "KMD"
	fmt.Println(args)

	get, err := appName.DEXGet(args)
	if err != nil {
		fmt.Printf("Code: %v\n", get.Error.Code)
		fmt.Printf("Message: %v\n\n", get.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("get value", get)
	fmt.Println("-------")
	// fmt.Println(get.Result)
	// fmt.Println("-------")

	fmt.Println("Timestamp", get.Result.Timestamp)
	fmt.Println("ID", get.Result.ID)
	fmt.Println("Hash", get.Result.Hash)
	fmt.Println("TagA", get.Result.TagA)
	fmt.Println("TagB", get.Result.TagB)
	fmt.Println("Pubkey", get.Result.Pubkey)
	fmt.Println("Payload", get.Result.Payload)
	fmt.Println("Hex", get.Result.Hex)
	fmt.Println("Decrypted", get.Result.Decrypted)
	fmt.Println("Decryptedhex", get.Result.Decryptedhex)
	fmt.Println("Senderpub", get.Result.Senderpub)
	fmt.Println("AmountA", get.Result.AmountA)
	fmt.Println("AmountB", get.Result.AmountB)
	fmt.Println("Priority", get.Result.Priority)
	fmt.Println("Recvtime", get.Result.Recvtime)
	fmt.Println("Cancelled", get.Result.Cancelled)
}
