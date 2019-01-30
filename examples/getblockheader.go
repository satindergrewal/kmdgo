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
	var appName kmdgo.AppType
	appName = `komodo`

	var bh kmdgo.GetBlockHeader

	args := make(kmdgo.APIParams, 2)
	args[0] = "0a203fce865faa53e1414b26c3020341a17d4e2136565d9b747f83169c938ac5"
	//args[1] = false
	fmt.Println(args)

	bh, err := appName.GetBlockHeader(args)
	if err != nil {
		fmt.Printf("Code: %v\n", bh.Error.Code)
		fmt.Printf("Message: %v\n\n", bh.Error.Message)
		log.Fatalln("Err happened", err)
	}

	//fmt.Println("bh value", bh)
	//fmt.Println(bh.Result)
	fmt.Println(bh.Result.Solution)
}