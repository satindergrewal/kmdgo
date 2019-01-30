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
	//"encoding/json"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	args := make(kmdgo.APIParams, 2)
	args[0] = `8d0790c3dcce42187a0e0ab710c9fcf376566e0e64fa800980b2b553a72da0c4`
	//args[1] = 1
	fmt.Println(args)

	if args[1] != 1 {
		var grwtx kmdgo.GetRawTransaction
		grwtx, err := appName.GetRawTransaction(args)
		if err != nil {
			fmt.Printf("Code: %v\n", grwtx.Error.Code)
			fmt.Printf("Message: %v\n\n", grwtx.Error.Message)
			log.Fatalln("Err happened", err)
		}
		//fmt.Println("grwtx value", grwtx)
		//fmt.Println("-------")
		fmt.Println(grwtx.Result)
	} else {
		var grwtx kmdgo.GetRawTransactionDetailed
		grwtx, err := appName.GetRawTransactionDetailed(args)
		if err != nil {
			fmt.Printf("Code: %v\n", grwtx.Error.Code)
			fmt.Printf("Message: %v\n\n", grwtx.Error.Message)
			log.Fatalln("Err happened", err)
		}
		//fmt.Println("grwtx value", grwtx)
		//fmt.Println("-------")
		//fmt.Println(grwtx.Result)
		fmt.Println("Hex: ", grwtx.Result.Hex)
		fmt.Println("Txid: ", grwtx.Result.Txid)
		fmt.Println("Overwintered: ", grwtx.Result.Overwintered)
		fmt.Println("Version: ", grwtx.Result.Version)
		fmt.Println("Locktime: ", grwtx.Result.Locktime)

		fmt.Println("-------")

		for in, vn := range grwtx.Result.Vin {
			fmt.Println(in)
			fmt.Println(vn)

			fmt.Println("Vin --> Txid: ", vn.Txid)
			fmt.Println("Vin --> Vout: ", vn.Vout)
			fmt.Println("Vin --> Address: ", vn.Address)
				fmt.Println("Vin --> ScriptSig --> Asm: ", vn.ScriptSig.Asm)
				fmt.Println("Vin --> ScriptSig --> Hex: ", vn.ScriptSig.Hex)
			fmt.Println("Vin --> Sequence: ", vn.Sequence)
		}

		fmt.Println("-------")

		for it, vt := range grwtx.Result.Vout {
			fmt.Println(it)
			fmt.Println(vt)
			fmt.Println("Vout --> Value: ", vt.Value)
			fmt.Println("Vout --> Interest: ", vt.Interest)
			fmt.Println("Vout --> ValueSat: ", vt.ValueSat)
			fmt.Println("Vout --> N: ", vt.N)
				fmt.Println("Vout --> ScriptPubKey --> Asm: ", vt.ScriptPubKey.Asm)
				fmt.Println("Vout --> ScriptPubKey --> Hex: ", vt.ScriptPubKey.Hex)
				fmt.Println("Vout --> ScriptPubKey --> ReqSigs: ", vt.ScriptPubKey.ReqSigs)
				fmt.Println("Vout --> ScriptPubKey --> Type: ", vt.ScriptPubKey.Type)
				fmt.Println("Vout --> ScriptPubKey --> Addresses: ", vt.ScriptPubKey.Addresses)
		}

		fmt.Println("Vjoinsplit: ", grwtx.Result.Vjoinsplit)
		fmt.Println("Blockhash: ", grwtx.Result.Blockhash)
		fmt.Println("Height: ", grwtx.Result.Height)
		fmt.Println("Confirmations: ", grwtx.Result.Confirmations)
		fmt.Println("Rawconfirmations: ", grwtx.Result.Rawconfirmations)
		fmt.Println("Time: ", grwtx.Result.Time)
		fmt.Println("Blocktime: ", grwtx.Result.Blocktime)
	}
}