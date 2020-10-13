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

	var GetCurncy kmdgo.GetCurrencyState

	args := make(kmdgo.APIParams, 1)
	args[0] = 20932
	// fmt.Println(args)

	GetCurncy, err := appName.GetCurrencyState(args)
	if err != nil {
		fmt.Printf("Code: %v\n", GetCurncy.Error.Code)
		fmt.Printf("Message: %v\n\n", GetCurncy.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println(GetCurncy.Result)
	fmt.Println("Height -", GetCurncy.Result.Height)
	fmt.Println("Blocktime -", GetCurncy.Result.Blocktime)
	// fmt.Println("Currencystate -", GetCurncy.Result.Currencystate)
	fmt.Println("Currencystate > Flags -", GetCurncy.Result.Currencystate.Flags)
	fmt.Println("Currencystate > Currencyid -", GetCurncy.Result.Currencystate.Currencyid)
	fmt.Println("Currencystate > Initialsupply -", GetCurncy.Result.Currencystate.Initialsupply)
	fmt.Println("Currencystate > Emitted -", GetCurncy.Result.Currencystate.Emitted)
	fmt.Println("Currencystate > Supply -", GetCurncy.Result.Currencystate.Supply)
	fmt.Println("Currencystate > Currencies -", GetCurncy.Result.Currencystate.Currencies)
	fmt.Println("Currencystate > Nativefees -", GetCurncy.Result.Currencystate.Nativefees)
	fmt.Println("Currencystate > Nativeconversionfees -", GetCurncy.Result.Currencystate.Nativeconversionfees)
}
