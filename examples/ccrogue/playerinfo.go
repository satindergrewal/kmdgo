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

	var rgpfo kmdgo.RGPlayerInfo

	args := make(kmdgo.APIParams, 3)
	args[2] = `["09035cd939b1c5f38cd08ddf04c908f918eb855875cad975196988c527a439e6"]`
	fmt.Println(args)

	rgpfo, err := appName.RGPlayerInfo(args)
	if err != nil {
		fmt.Printf("Code: %v\n", rgpfo.Error.Code)
		fmt.Printf("Message: %v\n\n", rgpfo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("rgpfo value", rgpfo)
	fmt.Println("-------")
	fmt.Println(rgpfo.Result)

	fmt.Println("\n")
	
	fmt.Println("Name: ", rgpfo.Result.Name)
	fmt.Println("Method: ", rgpfo.Result.Method)
	fmt.Println("Result: ", rgpfo.Result.Result)
	fmt.Println("Player->Playertxid: ", rgpfo.Result.Player.Playertxid)
	fmt.Println("Player->Tokenid: ", rgpfo.Result.Player.Tokenid)
	fmt.Println("Player->Data: ", rgpfo.Result.Player.Data)

	for pi, pv := range rgpfo.Result.Player.Pack {
		fmt.Printf("%v=> Player -> Pack: %s", pi, pv)
	}
	
	fmt.Println("Player -> Packsize: ", rgpfo.Result.Player.Packsize)
	fmt.Println("Player -> Hitpoints: ", rgpfo.Result.Player.Hitpoints)
	fmt.Println("Player -> Strength: ", rgpfo.Result.Player.Strength)
	fmt.Println("Player -> Level: ", rgpfo.Result.Player.Level)
	fmt.Println("Player -> Experience: ", rgpfo.Result.Player.Experience)
	fmt.Println("Player -> Dungeonlevel: ", rgpfo.Result.Player.Dungeonlevel)
	fmt.Println("Player -> Chain: ", rgpfo.Result.Player.Chain)
	fmt.Println("Player -> Pname: ", rgpfo.Result.Player.Pname)

	
}