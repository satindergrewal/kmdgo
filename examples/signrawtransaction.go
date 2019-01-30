 //******************************************************************************
 //  Copyright © 2018-2019 Satinderjit Singh.                                   *
 //                                                                             *
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 //  the top-level directory of this distribution for the individual copyright  *
 //  holder information and the developer policies on copyright and licensing.  *
 //                                                                             *
 //  Unless otherwise agreed in a custom licensing agreement, no part of the    *
 //  kmdgo software, including this file may be copied, modified, propagated.   *
 //  or distributed except according to the terms contained in the LICENSE file *
 //                                                                             *
 //  Removal or modification of this copyright notice is prohibited.            *
 //                                                                             *
 //******************************************************************************/
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

	var snrwtx kmdgo.SignRawTransaction

	args := make(kmdgo.APIParams, 1)
	args[0] = `0400008085202f8901fdd6474f4a070e1b7c271f80fe79557f8c84f8eea4271fb66ee1666c2945bad70000000000ffffffff0240420f00000000001976a91478d98c7ad3345e56fbacc49710723b47dc119e8a88ac80841e00000000001976a914b87ba7e8e320fa7eb60f154909e94df1d22af40888ac00000000af5312000000000000000000000000`
	//args[1] = true
	fmt.Println(args)

	snrwtx, err := appName.SignRawTransaction(args)
	if err != nil {
		fmt.Printf("Code: %v\n", snrwtx.Error.Code)
		fmt.Printf("Message: %v\n\n", snrwtx.Error.Message)
		log.Fatalln("Err happened", err)
	}
	fmt.Println("snrwtx value", snrwtx)
	fmt.Println("-------")
	fmt.Println(snrwtx.Result)

	fmt.Println("Hex: ", snrwtx.Result.Hex)
	fmt.Println("Complete: ", snrwtx.Result.Complete)

	if snrwtx.Result.Complete == false {
		fmt.Println("Error --> Txid: ", snrwtx.Result.Errors[0].Txid)
		fmt.Println("Error --> Vout: ", snrwtx.Result.Errors[0].Vout)
		fmt.Println("Error --> ScriptSig: ", snrwtx.Result.Errors[0].ScriptSig)
		fmt.Println("Error --> Sequence: ", snrwtx.Result.Errors[0].Sequence)
		fmt.Println("Error --> Error: ", snrwtx.Result.Errors[0].Error)
	}
}