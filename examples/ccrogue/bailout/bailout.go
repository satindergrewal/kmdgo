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

	var rgblt kmdgo.RGBailout

	args := make(kmdgo.APIParams, 3)
	args[2] = `["2ad99222dedd2ed4439501df13e2451e1497e148e2116ba2bb9afd7d42797812"]`
	fmt.Println(args)

	rgblt, err := appName.RGBailout(args)
	if err != nil {
		fmt.Printf("Code: %v\n", rgblt.Error.Code)
		fmt.Printf("Message: %v\n\n", rgblt.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("rgblt value", rgblt)
	fmt.Println("-------")
	fmt.Println(rgblt.Result)

	fmt.Println("\n")
	
	fmt.Println("Name: ", rgblt.Result.Name)
	fmt.Println("Method: ", rgblt.Result.Method)
	fmt.Println("Myrogueaddr: ", rgblt.Result.Myrogueaddr)
	fmt.Println("Gametxid: ", rgblt.Result.Gametxid)
	fmt.Println("Hex: ", rgblt.Result.Hex)
	fmt.Println("Txid: ", rgblt.Result.Txid)
	fmt.Println("Result: ", rgblt.Result.Result)
}