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

	var smny kmdgo.SendMany

	args := make(kmdgo.APIParams, 2)
	args[0] = `` // DO NOT USE account names (accounts are depricated). Can use "*" instead which will just select all accounts from wallet.
	args[1] = map[string]float64{"RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr": 0.01, "RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s": 0.02}
	//args[2] = 1
	//args[3] = `donation`
	//args[4] = []string{"RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr","RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s"}
	fmt.Println(args)

	smny, err := appName.SendMany(args)
	if err != nil {
		fmt.Printf("Code: %v\n", smny.Error.Code)
		fmt.Printf("Message: %v\n\n", smny.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("smny value", smny)
	fmt.Println("-------")
	fmt.Println(smny.Result)
}
