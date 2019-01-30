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

	var smny kmdgo.SendToAddress

	args := make(kmdgo.APIParams, 2)
	args[0] = `RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s`
	args[1] = 0.01
	//args[2] = `donation`
	//args[3] = `Non Profit Org`
	//args[4] = true
	fmt.Println(args)

	smny, err := appName.SendToAddress(args)
	if err != nil {
		fmt.Printf("Code: %v\n", smny.Error.Code)
		fmt.Printf("Message: %v\n\n", smny.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("smny value", smny)
	fmt.Println("-------")
	fmt.Println(smny.Result)
}