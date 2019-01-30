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

	var lstx kmdgo.ListTransactions

	args := make(kmdgo.APIParams, 4)
	args[1] = 10
	args[2] = 100
	args[3] = true
	fmt.Println(args)

	lstx, err := appName.ListTransactions(args)
	if err != nil {
		fmt.Printf("Code: %v\n", lstx.Error.Code)
		fmt.Printf("Message: %v\n\n", lstx.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("lstx value", lstx)
	//fmt.Println("-------")
	//fmt.Println(lstx.Result)
	fmt.Println("-------")
	
	//fmt.Printf("\n\n\n-------")

	for i, v := range lstx.Result {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		//fmt.Println(v)
		fmt.Println("InvolvesWatchonly: ", v.InvolvesWatchonly)
		fmt.Println("Account: ", v.Account)
		fmt.Println("Address: ", v.Address)
		fmt.Println("Category: ", v.Category)
		fmt.Println("Amount: ", v.Amount)
		fmt.Println("Vout: ", v.Vout)
		fmt.Println("Fee: ", v.Fee)
		fmt.Println("Rawconfirmations: ", v.Rawconfirmations)
		fmt.Println("Confirmations: ", v.Confirmations)
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