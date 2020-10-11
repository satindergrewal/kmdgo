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

	var rgifo kmdgo.RGGameInfo

	args := make(kmdgo.APIParams, 3)
	args[2] = `["a469b0660b852692cbbe18083f9b9ab0ffada2ff05d5c29bce1b22a0550b715e"]`
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

	for i, v := range rgifo.Result.Players {
		fmt.Println(i)
		fmt.Println(v)

		fmt.Println("\n")

		fmt.Println("Slot: ", v.Slot)
		fmt.Println("Status: ", v.Status)
		fmt.Println("Baton: ", v.Baton)
		fmt.Println("Tokenid: ", v.Tokenid)
		fmt.Println("Batonaddr: ", v.Batonaddr)
		fmt.Println("Ismine: ", v.Ismine)
		fmt.Println("Batonvout: ", v.Batonvout)
		fmt.Println("Batonvalue: ", v.Batonvalue)
		fmt.Println("Batonht: ", v.Batonht)
	}
}