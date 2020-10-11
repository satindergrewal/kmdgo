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

	var publ kmdgo.DEXPublish

	args := make(kmdgo.APIParams, 2)
	// filename
	args[0] = "dex_list.go"
	// priority
	args[1] = "0"
	// // sliceid
	// args[2] = ""
	fmt.Println(args)

	publ, err := appName.DEXPublish(args)
	if err != nil {
		fmt.Printf("Code: %v\n", publ.Error.Code)
		fmt.Printf("Message: %v\n\n", publ.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("publ value", publ)
	fmt.Println("-------")
	// fmt.Println(publ.Result)
	// fmt.Println("-------")

	fmt.Println("Fname", publ.Result.Fname)
	fmt.Println("ID", publ.Result.ID)
	fmt.Println("Senderpub", publ.Result.Senderpub)
	fmt.Println("Filesize", publ.Result.Filesize)
	fmt.Println("Fragments", publ.Result.Fragments)
	fmt.Println("Numlocators", publ.Result.Numlocators)
	fmt.Println("Filehash", publ.Result.Filehash)
	fmt.Println("Checkhash", publ.Result.Checkhash)
	fmt.Println("Result", publ.Result.Result)
}
