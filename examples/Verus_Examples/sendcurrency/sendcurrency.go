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

	var sndCur kmdgo.SendCurrency

	var sndCurInput kmdgo.SendCurrencyInput

	sndCurInput.Address = "satinder@"
	sndCurInput.Amount = 100.0
	sndCurInput.Convertto = "vDEX"
	sndCurInput.Via = "vDEX-VRSC"

	var sndCurInputArray []kmdgo.SendCurrencyInput
	sndCurInputArray = append(sndCurInputArray, sndCurInput)

	args := make(kmdgo.APIParams, 2)
	args[0] = `*`              // "*" to select all transparent addresses balance from wallet.
	args[1] = sndCurInputArray // JSON data input
	// fmt.Println(args)

	sndCur, err := appName.SendCurrency(args)
	if err != nil {
		fmt.Printf("Code: %v\n", sndCur.Error.Code)
		fmt.Printf("Message: %v\n\n", sndCur.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Printf("sndCur value %+v\n", sndCur)
	fmt.Println("-------")
	fmt.Println(sndCur.Result)
}
