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

	var anonsend kmdgo.DEXAnonsend

	args := make(kmdgo.APIParams, 3)
	args[0] = "hello"
	args[1] = "6"
	args[2] = "0138d849d6bc81ff1c5389aae9a60ba3ee9cfd7858d93a3864679c25937e70951f"
	fmt.Println(args)

	anonsend, err := appName.DEXAnonsend(args)
	if err != nil {
		fmt.Printf("Code: %v\n", anonsend.Error.Code)
		fmt.Printf("Message: %v\n\n", anonsend.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("anonsend value", anonsend)
	fmt.Println("-------")
	// fmt.Println(anonsend.Result)
	// fmt.Println("-------")

	fmt.Println("Timestamp", anonsend.Result.Timestamp)
	fmt.Println("ID", anonsend.Result.ID)
	fmt.Println("Hash", anonsend.Result.Hash)
	fmt.Println("TagA", anonsend.Result.TagA)
	fmt.Println("TagB", anonsend.Result.TagB)
	fmt.Println("Pubkey", anonsend.Result.Pubkey)
	fmt.Println("Payload", anonsend.Result.Payload)
	fmt.Println("Hex", anonsend.Result.Hex)
	fmt.Println("Decrypted", anonsend.Result.Decrypted)
	fmt.Println("Decryptedhex", anonsend.Result.Decryptedhex)
	fmt.Println("Senderpub", anonsend.Result.Senderpub)
	fmt.Println("AmountA", anonsend.Result.AmountA)
	fmt.Println("AmountB", anonsend.Result.AmountB)
	fmt.Println("Priority", anonsend.Result.Priority)
	fmt.Println("Recvtime", anonsend.Result.Recvtime)
	fmt.Println("Cancelled", anonsend.Result.Cancelled)
}
