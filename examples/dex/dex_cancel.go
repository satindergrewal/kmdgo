// Copyright © 2018-2019 Satinderjit Singh.
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

	var cancel kmdgo.DEXCancel

	args := make(kmdgo.APIParams, 4)
	args[0] = "38289120"
	args[1] = "0138d849d6bc81ff1c5389aae9a60ba3ee9cfd7858d93a3864679c25937e70951f"
	args[2] = "BTC"
	args[3] = "KMD"
	fmt.Println(args)

	cancel, err := appName.DEXCancel(args)
	if err != nil {
		fmt.Printf("Code: %v\n", cancel.Error.Code)
		fmt.Printf("Message: %v\n\n", cancel.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("cancel value", cancel)
	fmt.Println("-------")
	// fmt.Println(cancel.Result)
	// fmt.Println("-------")

	fmt.Println("Timestamp", cancel.Result.Timestamp)
	fmt.Println("ID", cancel.Result.ID)
	fmt.Println("Hash", cancel.Result.Hash)
	fmt.Println("TagA", cancel.Result.TagA)
	fmt.Println("TagB", cancel.Result.TagB)
	fmt.Println("Pubkey", cancel.Result.Pubkey)
	fmt.Println("Payload", cancel.Result.Payload)
	fmt.Println("Hex", cancel.Result.Hex)
	fmt.Println("Decrypted", cancel.Result.Decrypted)
	fmt.Println("Decryptedhex", cancel.Result.Decryptedhex)
	fmt.Println("AmountA", cancel.Result.AmountA)
	fmt.Println("AmountB", cancel.Result.AmountB)
	fmt.Println("Priority", cancel.Result.Priority)
	fmt.Println("Recvtime", cancel.Result.Recvtime)
	fmt.Println("Cancelled", cancel.Result.Cancelled)
}
