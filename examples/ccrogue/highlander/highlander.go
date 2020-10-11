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
	appName := kmdgo.NewAppType(`ROGUE`)

	var rghldr kmdgo.RGHighLander

	args := make(kmdgo.APIParams, 3)
	args[2] = `["2ad99222dedd2ed4439501df13e2451e1497e148e2116ba2bb9afd7d42797812"]`
	fmt.Println(args)

	rghldr, err := appName.RGHighLander(args)
	if err != nil {
		fmt.Printf("Code: %v\n", rghldr.Error.Code)
		fmt.Printf("Message: %v\n\n", rghldr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("rghldr value", rghldr)
	fmt.Println("-------")
	fmt.Println(rghldr.Result)

	fmt.Println("\n")
	
	fmt.Println("Name: ", rghldr.Result.Name)
	fmt.Println("Method: ", rghldr.Result.Method)
	fmt.Println("Myrogueaddr: ", rghldr.Result.Myrogueaddr)
	fmt.Println("Gametxid: ", rghldr.Result.Gametxid)
	fmt.Println("Hex: ", rghldr.Result.Hex)
	fmt.Println("Txid: ", rghldr.Result.Txid)
	fmt.Println("Result: ", rghldr.Result.Result)
}