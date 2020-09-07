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
	"github.com/satindergrewal/kmdgo"
	"log"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var zlunspt kmdgo.ZListUnspent

	args := make(kmdgo.APIParams, 4)
	//args[0] = 12
	//args[1] = 200
	//args[2] = true
	args[3] = []string{"zs1xhnpvrggh3n5a0pt5hwgdcvmuhn9wqjg6kn25jddduvq5qyy3nhusfjufljp0ul55d66vkxajt3", "zs1tyaqq9nstpk2ezvj5ayxg6nfkhlrc80dcset6v6jmpk9gft384v6rpgmxhu00u3aalygqgk77eg"}
	fmt.Println(args)

	zlunspt, err := appName.ZListUnspent(args)
	if err != nil {
		fmt.Printf("Code: %v\n", zlunspt.Error.Code)
		fmt.Printf("Message: %v\n\n", zlunspt.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("zlunspt value", zlunspt)
	fmt.Println("-------")
	fmt.Println(zlunspt.Result)
	fmt.Println("-------")

	for i, v := range zlunspt.Result {
		fmt.Println(i)
		fmt.Println("Txid: ", v.Txid)
		fmt.Println("Jsindex: ", v.Jsindex)
		fmt.Println("Jsoutindex: ", v.Jsoutindex)
		fmt.Println("Outindex: ", v.Outindex)
		fmt.Println("Confirmations: ", v.Confirmations)
		fmt.Println("Spendable: ", v.Spendable)
		fmt.Println("Address: ", v.Address)
		fmt.Printf("Amount: %0.8f\n", v.Amount)
		fmt.Println("Memo: ", v.Memo)
		fmt.Println("Change: ", v.Change)
	}
}
