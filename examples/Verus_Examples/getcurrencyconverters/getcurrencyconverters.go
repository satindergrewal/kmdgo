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

	var GetCurncyCnvrts kmdgo.GetCurrencyConverters

	args := make(kmdgo.APIParams, 2)
	args[0] = "BTC"
	args[1] = "ETH"
	// fmt.Println(args)

	GetCurncyCnvrts, err := appName.GetCurrencyConverters(args)
	if err != nil {
		fmt.Printf("Code: %v\n", GetCurncyCnvrts.Error.Code)
		fmt.Printf("Message: %v\n\n", GetCurncyCnvrts.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println(GetCurncyCnvrts.Result)
	for i, v := range GetCurncyCnvrts.Result {
		fmt.Println("------")
		fmt.Println(i)
		if v.CurrencyInfo.Name != "" {
			fmt.Println("v.CurrencyInfo -- ", v.CurrencyInfo)
		}
		if v.Multifractional.Name != "" {
			fmt.Println("v.Multifractional -- ", v.Multifractional)
		}
		fmt.Println("v.Lastnotarization -- ", v.Lastnotarization)
	}
}
