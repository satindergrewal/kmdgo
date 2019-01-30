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

	var adnd kmdgo.AddNode

	args := make(kmdgo.APIParams, 2)
	args[0] = `192.169.7.252:7770`
	args[1] = `onetry` // possible values are: "add" or "remove" or "onetry"
	fmt.Println(args)

	adnd, err := appName.AddNode(args)
	if err != nil {
		fmt.Printf("Code: %v\n", adnd.Error.Code)
		fmt.Printf("Message: %v\n\n", adnd.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("adnd value", adnd)
	fmt.Println("-------")
	fmt.Println(adnd.Result)
}