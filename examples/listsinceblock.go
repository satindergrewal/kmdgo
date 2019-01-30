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

	var lrbaddr kmdgo.ListSinceBlock

	args := make(kmdgo.APIParams, 3)
	args[0] = `02a06c8a9535c3f7fb0db832b177c8fa3e8a4eb5034aeb698c37d2e66a570648`
	args[1] = 6
	//args[2] = true
	fmt.Println(args)

	lrbaddr, err := appName.ListSinceBlock(args)
	if err != nil {
		fmt.Printf("Code: %v\n", lrbaddr.Error.Code)
		fmt.Printf("Message: %v\n\n", lrbaddr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("lrbaddr value", lrbaddr)
	fmt.Println("-------")
	fmt.Println(lrbaddr.Result)
	fmt.Println("-------")
	fmt.Println("Lastblock", lrbaddr.Result.Lastblock)

	for i, v := range lrbaddr.Result.Transactions {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		fmt.Println("Account: ", v.Account)
		fmt.Println("Address: ", v.Address)
		fmt.Println("Category: ", v.Category)
		fmt.Println("Amount: ", v.Amount)
		fmt.Println("Vout: ", v.Vout)
		fmt.Println("Confirmations: ", v.Confirmations)
		fmt.Println("Generated: ", v.Generated)
		fmt.Println("Blockhash: ", v.Blockhash)
		fmt.Println("Blockindex: ", v.Blockindex)
		fmt.Println("Blocktime: ", v.Blocktime)
		fmt.Println("Expiryheight: ", v.Expiryheight)
		fmt.Println("Txid: ", v.Txid)
		fmt.Println("Walletconflicts: ", v.Walletconflicts)
		fmt.Println("Time: ", v.Time)
		fmt.Println("Timereceived: ", v.Timereceived)
		fmt.Println("Vjoinsplit: ", v.Vjoinsplit)
		fmt.Println("Size: ", v.Size)
	}

}