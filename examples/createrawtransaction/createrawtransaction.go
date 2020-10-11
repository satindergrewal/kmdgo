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

	var crwtx kmdgo.CreateRawTransaction

	type txes []struct {
		Txid string `json:"txid"`
		Vout int    `json:"vout"`
		//Sequence int    `json:"sequence,omitempty"`
	}

	args := make(kmdgo.APIParams, 2)
	args[0] = txes{{"d7ba45296c66e16eb61f27a4eef8848c7f5579fe801f277c1b0e074a4f47d6fd", 0}}
	args[1] = map[string]float64{"RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr": 0.01, "RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s": 0.02}
	//args[2] = 0
	//args[3] = 30
	fmt.Println(args)

	crwtx, err := appName.CreateRawTransaction(args)
	if err != nil {
		fmt.Printf("Code: %v\n", crwtx.Error.Code)
		fmt.Printf("Message: %v\n\n", crwtx.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("crwtx value", crwtx)
	fmt.Println("-------")
	fmt.Println(crwtx.Result)
}
