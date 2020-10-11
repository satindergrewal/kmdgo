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

	var rgng kmdgo.RGNewGame

	args := make(kmdgo.APIParams, 3)
	//args[0] = `newgame`
	//args[1] = `17`
	args[2] = `[1]`
	fmt.Println(args)

	rgng, err := appName.RGNewGame(args)
	if err != nil {
		fmt.Printf("Code: %v\n", rgng.Error.Code)
		fmt.Printf("Message: %v\n\n", rgng.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("rgng value", rgng)
	fmt.Println("-------")
	fmt.Println(rgng.Result)

	fmt.Println("\n")

	fmt.Println("Name: ", rgng.Result.Name)
	fmt.Println("Method: ", rgng.Result.Method)
	fmt.Println("Maxplayers: ", rgng.Result.Maxplayers)
	fmt.Println("Buyin: ", rgng.Result.Buyin)
	fmt.Println("Type: ", rgng.Result.Type)
	fmt.Println("Hex: ", rgng.Result.Hex)
	fmt.Println("Txid: ", rgng.Result.Txid)
	fmt.Println("Result: ", rgng.Result.Result)
}