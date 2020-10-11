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

	var obook kmdgo.DEXOrderbook

	args := make(kmdgo.APIParams, 4)
	// maxentries
	args[0] = "10"
	// minpriority
	args[1] = "0"
	// tagA
	args[2] = "KMD"
	// tagB
	args[3] = "BTC"
	// // pubkey33
	// args[4] = ""
	// // minA - Optional
	// args[5] = ""
	// // maxA - Optional
	// args[6] = ""
	// // minB - Optional
	// args[7] = ""
	// // maxB - Optional
	// args[8] = ""
	fmt.Println(args)

	obook, err := appName.DEXOrderbook(args)
	if err != nil {
		fmt.Printf("Code: %v\n", obook.Error.Code)
		fmt.Printf("Message: %v\n\n", obook.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("obook value", obook)
	fmt.Println("-------")
	// fmt.Println(obook.Result)
	// fmt.Println("-------")

	fmt.Println("--- Asks ---")
	for i, v := range obook.Result.Asks {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		fmt.Println("Price", v.Price)
		fmt.Println("Baseamount", v.Baseamount)
		fmt.Println("Relamount", v.Relamount)
		fmt.Println("Priority", v.Priority)
		fmt.Println("Pubkey", v.Pubkey)
		fmt.Println("Timestamp", v.Timestamp)
		fmt.Println("Hash", v.Hash)
		fmt.Println("ID", v.ID)
	}

	fmt.Println("--- Bids ---")
	for i, v := range obook.Result.Bids {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		fmt.Println("Price", v.Price)
		fmt.Println("Baseamount", v.Baseamount)
		fmt.Println("Relamount", v.Relamount)
		fmt.Println("Priority", v.Priority)
		fmt.Println("Pubkey", v.Pubkey)
		fmt.Println("Timestamp", v.Timestamp)
		fmt.Println("Hash", v.Hash)
		fmt.Println("ID", v.ID)
	}

	fmt.Println("-------")
	fmt.Println("Base", obook.Result.Base)
	fmt.Println("Rel", obook.Result.Rel)
}
