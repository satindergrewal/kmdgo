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
	appName := kmdgo.NewAppType(`ROGUE`)

	var rgpdg kmdgo.RGPending

	rgpdg, err := appName.RGPending()
	if err != nil {
		fmt.Printf("Code: %v\n", rgpdg.Error.Code)
		fmt.Printf("Message: %v\n\n", rgpdg.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("rgpdg value", rgpdg)
	fmt.Println("-------")
	fmt.Println(rgpdg.Result)

	fmt.Println("\n")

	fmt.Println("Result: ", rgpdg.Result.Result)
	fmt.Println("Name: ", rgpdg.Result.Name)
	fmt.Println("Method: ", rgpdg.Result.Method)
	fmt.Println("Numpending: ", rgpdg.Result.Numpending)

	for i, v := range rgpdg.Result.Pending {
		fmt.Println(i)
		fmt.Println(v)
	}
}