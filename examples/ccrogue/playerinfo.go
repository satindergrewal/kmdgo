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

	var rgblt kmdgo.RGPlayerInfo

	args := make(kmdgo.APIParams, 3)
	args[2] = `["09035cd939b1c5f38cd08ddf04c908f918eb855875cad975196988c527a439e6"]`
	fmt.Println(args)

	rgblt, err := appName.RGPlayerInfo(args)
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
	fmt.Println("Result: ", rgblt.Result.Result)

	for i, v := range rgblt.Result.Player {
		fmt.Println(i)
		fmt.Println(v)

		fmt.Println("\n")

		fmt.Println("Playertxid: ", v.Playertxid)
		fmt.Println("Tokenid: ", v.Tokenid)
		fmt.Println("Data: ", v.Data)
		fmt.Println("Packsize: ", v.Packsize)
		fmt.Println("Hitpoints: ", v.Hitpoints)
		fmt.Println("Strength: ", v.Strength)
		fmt.Println("Level: ", v.Level)
		fmt.Println("Experience: ", v.Experience)
		fmt.Println("Dungeonlevel: ", v.Dungeonlevel)
		fmt.Println("Chain: ", v.Chain)
		fmt.Println("Pname: ", v.Pname)

		for pi, pv := range v.Pack {
			fmt.Println("Player -> Pack: ", pv)
		}
	}
}