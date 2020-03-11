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

	var broadcast kmdgo.DEXBroadcast

	args := make(kmdgo.APIParams, 7)
	args[0] = "hello"
	args[1] = "5"
	args[2] = "BTC"
	args[3] = "KMD"
	args[4] = "0138d849d6bc81ff1c5389aae9a60ba3ee9cfd7858d93a3864679c25937e70951f"
	args[5] = "0.1"
	args[6] = "100"
	fmt.Println(args)

	broadcast, err := appName.DEXBroadcast(args)
	if err != nil {
		fmt.Printf("Code: %v\n", broadcast.Error.Code)
		fmt.Printf("Message: %v\n\n", broadcast.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("broadcast value", broadcast)
	fmt.Println("-------")
	// fmt.Println(broadcast.Result)
	// fmt.Println("-------")

	fmt.Println("Timestamp", broadcast.Result.Timestamp)
	fmt.Println("ID", broadcast.Result.ID)
	fmt.Println("Hash", broadcast.Result.Hash)
	fmt.Println("TagA", broadcast.Result.TagA)
	fmt.Println("TagB", broadcast.Result.TagB)
	fmt.Println("Pubkey", broadcast.Result.Pubkey)
	fmt.Println("Payload", broadcast.Result.Payload)
	fmt.Println("Hex", broadcast.Result.Hex)
	fmt.Println("Decrypted", broadcast.Result.Decrypted)
	fmt.Println("Decryptedhex", broadcast.Result.Decryptedhex)
	fmt.Println("Senderpub", broadcast.Result.Senderpub)
	fmt.Println("AmountA", broadcast.Result.AmountA)
	fmt.Println("AmountB", broadcast.Result.AmountB)
	fmt.Println("Priority", broadcast.Result.Priority)
	fmt.Println("Recvtime", broadcast.Result.Recvtime)
	fmt.Println("Cancelled", broadcast.Result.Cancelled)
}
