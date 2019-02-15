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

	var rgifo kmdgo.RGGameInfo

	args := make(kmdgo.APIParams, 3)
	args[0] = `gameinfo`
	args[1] = `17`
	args[2] = `["2ad99222dedd2ed4439501df13e2451e1497e148e2116ba2bb9afd7d42797812"]`
	fmt.Println(args)

	rgifo, err := appName.RGGameInfo(args)
	if err != nil {
		fmt.Printf("Code: %v\n", rgifo.Error.Code)
		fmt.Printf("Message: %v\n\n", rgifo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("rgifo value", rgifo)
	fmt.Println("-------")
	fmt.Println(rgifo.Result)

	fmt.Println("\n")

	fmt.Println("Name: ", rgifo.Result.Name)
	fmt.Println("Method: ", rgifo.Result.Method)
	fmt.Println("Gametxid: ", rgifo.Result.Gametxid)
	fmt.Println("Result: ", rgifo.Result.Result)
	fmt.Println("Gameheight: ", rgifo.Result.Gameheight)
	fmt.Println("Height: ", rgifo.Result.Height)
	fmt.Println("Start: ", rgifo.Result.Start)
	fmt.Println("Starthash: ", rgifo.Result.Starthash)
	fmt.Println("Seed: ", rgifo.Result.Seed)
	fmt.Println("Run: ", rgifo.Result.Run)
	fmt.Println("Alive: ", rgifo.Result.Alive)
	fmt.Println("Numplayers: ", rgifo.Result.Numplayers)
	fmt.Println("Maxplayers: ", rgifo.Result.Maxplayers)
	fmt.Println("Buyin: ", rgifo.Result.Buyin)
	fmt.Println("Players: ", rgifo.Result.Players)
}