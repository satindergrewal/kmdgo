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

	var GetImp kmdgo.GetImports

	args := make(kmdgo.APIParams, 1)
	args[0] = "VRSC-BTC"
	// fmt.Println(args)

	GetImp, err := appName.GetImports(args)
	if err != nil {
		fmt.Printf("Code: %v\n", GetImp.Error.Code)
		fmt.Printf("Message: %v\n\n", GetImp.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println(GetImp.Result)

	for i, v := range GetImp.Result {
		fmt.Println("------")
		fmt.Println(i)
		// fmt.Println(v)
		fmt.Println("Blockheight - ", v.Blockheight)
		fmt.Println("Importid - ", v.Importid)
		// fmt.Println("Description - ", v.Description)
		fmt.Println("\tDescription.Version -- ", v.Description.Version)
		fmt.Println("\tDescription.Sourcesystemid -- ", v.Description.Sourcesystemid)
		fmt.Println("\tDescription.Importcurrencyid -- ", v.Description.Importcurrencyid)
		fmt.Println("\tDescription.Valuein -- ", v.Description.Valuein)
		fmt.Println("\tDescription.Tokensout -- ", v.Description.Tokensout)

		for ii, vv := range v.Transfers {
			fmt.Println("\t   ------")
			fmt.Println("\t   ", ii)
			// fmt.Println(vv)
			fmt.Println("\t   Transfers - Version -- ", vv.Version)
			fmt.Println("\t   Transfers - Currencyid -- ", vv.Currencyid)
			fmt.Println("\t   Transfers - Value -- ", vv.Value)
			fmt.Println("\t   Transfers - Flags -- ", vv.Flags)
			fmt.Println("\t   Transfers - Preconvert -- ", vv.Preconvert)
			fmt.Println("\t   Transfers - Fees -- ", vv.Fees)
			fmt.Println("\t   Transfers - Destinationcurrencyid -- ", vv.Destinationcurrencyid)
			fmt.Println("\t   Transfers - Destination -- ", vv.Destination)
			fmt.Println("\t   Transfers - Feeoutput -- ", vv.Feeoutput)
		}
	}
}
