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
	"github.com/satindergrewal/kmdgo"
	"log"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var ztlblc kmdgo.ZGetTotalBalance

	args := make(kmdgo.APIParams, 2)
	//args[0] = 1
	//args[1] = false
	fmt.Println(args)

	ztlblc, err := appName.ZGetTotalBalance(args)
	if err != nil {
		fmt.Printf("Code: %v\n", ztlblc.Error.Code)
		fmt.Printf("Message: %v\n\n", ztlblc.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("ztlblc value", ztlblc)
	fmt.Println("-------")
	fmt.Println(ztlblc.Result)
	fmt.Println("-------")
	fmt.Printf("Transparent: %0.8f\n", ztlblc.Result.Transparent)
	fmt.Printf("Interest: %0.8f\n", ztlblc.Result.Interest)
	fmt.Printf("Private: %0.8f\n", ztlblc.Result.Private)
	fmt.Printf("Total: %0.8f\n", ztlblc.Result.Total)

}
