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
	"log"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSCTEST`

	var GetLastImpIn kmdgo.GetLastImportin

	args := make(kmdgo.APIParams, 1)
	args[0] = "VRSC-BTC"
	// fmt.Println(args)

	GetLastImpIn, err := appName.GetLastImportin(args)
	if err != nil {
		fmt.Printf("Code: %v\n", GetLastImpIn.Error.Code)
		fmt.Printf("Message: %v\n\n", GetLastImpIn.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println(GetLastImpIn.Result)
	fmt.Println("Lastimporttransaction - ", GetLastImpIn.Result.Lastimporttransaction)
	fmt.Println("Lastconfirmednotarization - ", GetLastImpIn.Result.Lastconfirmednotarization)
	fmt.Println("Importtxtemplate - ", GetLastImpIn.Result.Importtxtemplate)
	fmt.Println("Nativeimportavailable - ", GetLastImpIn.Result.Nativeimportavailable)
	// fmt.Println("Tokenimportavailable - ", GetLastImpIn.Result.Tokenimportavailable)
	for i, v := range GetLastImpIn.Result.Tokenimportavailable {
		fmt.Println("  Tokenimportavailable -", i, ":", v)
	}
}
