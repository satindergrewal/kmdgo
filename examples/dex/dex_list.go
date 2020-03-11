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

	var list kmdgo.DEXList

	args := make(kmdgo.APIParams, 10)
	// stopat
	args[0] = ""
	// minpriority
	args[1] = "0"
	// tagA
	args[2] = "BTC"
	// tagB
	args[3] = ""
	// pubkey33
	args[4] = ""
	// minA
	args[5] = ""
	// maxA
	args[6] = ""
	// minB
	args[7] = ""
	// maxB
	args[8] = ""
	// stophash
	args[9] = ""
	fmt.Println(args)

	list, err := appName.DEXList(args)
	if err != nil {
		fmt.Printf("Code: %v\n", list.Error.Code)
		fmt.Printf("Message: %v\n\n", list.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("list value", list)
	fmt.Println("-------")
	fmt.Println(list.Result.Matches)
	fmt.Println("-------")

	for i, v := range list.Result.Matches {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		fmt.Println("Timestamp", v.Timestamp)
		fmt.Println("ID", v.ID)
		fmt.Println("Hash", v.Hash)
		fmt.Println("TagA", v.TagA)
		fmt.Println("TagB", v.TagB)
		fmt.Println("Pubkey", v.Pubkey)
		fmt.Println("Payload", v.Payload)
		fmt.Println("Hex", v.Hex)
		fmt.Println("Senderpub", v.Senderpub)
		fmt.Println("Error", v.Error)
		fmt.Println("AmountA", v.AmountA)
		fmt.Println("AmountB", v.AmountB)
		fmt.Println("Priority", v.Priority)
		fmt.Println("Recvtime", v.Recvtime)
		fmt.Println("Cancelled", v.Cancelled)
	}

	fmt.Println("-------")
	fmt.Println("TagA", TagA)
	fmt.Println("TagB", TagB)
	fmt.Println("Pubkey", Pubkey)
	fmt.Println("N", N)
}
