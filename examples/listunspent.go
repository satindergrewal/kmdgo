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

	var lunspt kmdgo.ListUnspent

	args := make(kmdgo.APIParams, 3)
	args[0] = 0
	args[1] = 99999999
	args[2] = []string{"RUdERUdW8aeEtRehZseDco9GTWN494LWW3","RAfyD4vUX4iUVvVEoxQ3hipHFNAvSw8Gnp"}
	fmt.Println(args)

	lunspt, err := appName.ListUnspent(args)
	if err != nil {
		fmt.Printf("Code: %v\n", lunspt.Error.Code)
		fmt.Printf("Message: %v\n\n", lunspt.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("lunspt value", lunspt)
	fmt.Println("-------")
	fmt.Println(lunspt.Result)
	fmt.Println("-------")
	
	for i, v := range lunspt.Result {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		fmt.Println("Txid: ", v.Txid)
		fmt.Println("Vout: ", v.Vout)
		fmt.Println("Generated: ", v.Generated)
		fmt.Println("Address: ", v.Address)
		fmt.Println("ScriptPubKey: ", v.ScriptPubKey)
		fmt.Println("Amount: ", v.Amount)
		fmt.Println("Interest: ", v.Interest)
		fmt.Println("Confirmations: ", v.Confirmations)
		fmt.Println("Spendable: ", v.Spendable)
	}

}