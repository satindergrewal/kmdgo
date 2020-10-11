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

	var rggms kmdgo.RGGames

	rggms, err := appName.RGGames()
	if err != nil {
		fmt.Printf("Code: %v\n", rggms.Error.Code)
		fmt.Printf("Message: %v\n\n", rggms.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("rggms value", rggms)
	fmt.Println("-------")
	fmt.Println(rggms.Result)

	fmt.Println("\n")

	fmt.Println("Name: ", rggms.Result.Name)
	fmt.Println("Method: ", rggms.Result.Method)
	fmt.Println("Numgames: ", rggms.Result.Numgames)

	for pi, pv := range rggms.Result.Pastgames {
		fmt.Printf("Pastgame %v: %s\n", pi+1, pv)
	}

	for gi, gv := range rggms.Result.Games {
		fmt.Printf("Game %v: %s\n", gi+1, gv)
	}
}