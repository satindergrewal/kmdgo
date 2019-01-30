 //  Copyright © 2018-2019 Satinderjit Singh.
 //
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
 //  the top-level directory of this distribution for the individual copyright
 //  holder information and the developer policies on copyright and licensing.
 //
 //  Unless otherwise agreed in a custom licensing agreement, no part of the
 //  kmdgo software, including this file may be copied, modified, propagated.
 //  or distributed except according to the terms contained in the LICENSE file
 //
 //  Removal or modification of this copyright notice is prohibited.
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var lkunspt kmdgo.LockUnspent

	type Txes []struct {
		Txid string `json:"txid"`
		Vout int    `json:"vout"`
	}

	args := make(kmdgo.APIParams, 2)
	args[0] = true
	args[1] = Txes{{"d7ba45296c66e16eb61f27a4eef8848c7f5579fe801f277c1b0e074a4f47d6fd", 0}}
	fmt.Println(args)

	lkunspt, err := appName.LockUnspent(args)
	if err != nil {
		fmt.Printf("Code: %v\n", lkunspt.Error.Code)
		fmt.Printf("Message: %v\n\n", lkunspt.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("lkunspt value", lkunspt)
	fmt.Println("-------")
	fmt.Println(lkunspt.Result)
}