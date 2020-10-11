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
	"github.com/satindergrewal/kmdgo"
	"log"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var sfrm kmdgo.SendFrom

	args := make(kmdgo.APIParams, 4)
	args[0] = `*` // DO NOT USE account names (accounts are depricated). Can use "*" instead which will just select all accounts from wallet.
	args[1] = `RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr`
	args[2] = 0.01
	//args[3] = 0
	//args[4] = `donation`
	//args[5] = `Non Profit Org`
	fmt.Println(args)

	sfrm, err := appName.SendFrom(args)
	if err != nil {
		fmt.Printf("Code: %v\n", sfrm.Error.Code)
		fmt.Printf("Message: %v\n\n", sfrm.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("sfrm value", sfrm)
	fmt.Println("-------")
	fmt.Println(sfrm.Result)
}
