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

	var txdtl kmdgo.GetTransaction

	args := make(kmdgo.APIParams, 2)
	args[0] = `8d0790c3dcce42187a0e0ab710c9fcf376566e0e64fa800980b2b553a72da0c4`
	args[1] = true
	fmt.Println(args)

	txdtl, err := appName.GetTransaction(args)
	if err != nil {
		fmt.Printf("Code: %v\n", txdtl.Error.Code)
		fmt.Printf("Message: %v\n\n", txdtl.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("txdtl value", txdtl)
	fmt.Println("-------")
	fmt.Println(txdtl.Result)
	fmt.Println("-------")

	fmt.Println("Amount: ", txdtl.Result.Amount)
	fmt.Println("Fee: ", txdtl.Result.Fee)
	fmt.Println("Rawconfirmations: ", txdtl.Result.Rawconfirmations)
	fmt.Println("Confirmations: ", txdtl.Result.Confirmations)
	fmt.Println("Blockhash: ", txdtl.Result.Blockhash)
	fmt.Println("Blockindex: ", txdtl.Result.Blockindex)
	fmt.Println("Blocktime: ", txdtl.Result.Blocktime)
	fmt.Println("Expiryheight: ", txdtl.Result.Expiryheight)
	fmt.Println("Txid: ", txdtl.Result.Txid)
	fmt.Println("Walletconflicts: ", txdtl.Result.Walletconflicts)
	fmt.Println("Time: ", txdtl.Result.Time)
	fmt.Println("Timereceived: ", txdtl.Result.Timereceived)
	fmt.Println("Vjoinsplit: ", txdtl.Result.Vjoinsplit)

	fmt.Printf("\n\n\n-------")

	for i, v := range txdtl.Result.Details {
		fmt.Println(i)
		//fmt.Println(v)
		fmt.Println("InvolvesWatchonly: ", v.InvolvesWatchonly)
		fmt.Println("Account: ", v.Account)
		fmt.Println("Address: ", v.Address)
		fmt.Println("Category: ", v.Category)
		fmt.Println("Amount: ", v.Amount)
		fmt.Println("Vout: ", v.Vout)
		fmt.Println("Fee: ", v.Fee)
		fmt.Println("Size: ", v.Size)
	}
}