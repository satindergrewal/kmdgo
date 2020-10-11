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
	appName = `TXSCLCC`

	var fctadr kmdgo.FaucetAddress

	args := make(kmdgo.APIParams, 1)
	//args[0] = `02fbccd779e492bcf3b12707c1461ea857b4e2ee0ad990f13059c02189bd5d8eda`
	fmt.Println(args)

	fctadr, err := appName.FaucetAddress(args)
	if err != nil {
		fmt.Printf("Code: %v\n", fctadr.Error.Code)
		fmt.Printf("Message: %v\n\n", fctadr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("fctadr value", fctadr)
	fmt.Println("-------")
	fmt.Println(fctadr.Result)

	fmt.Println("Result: ", fctadr.Result.Result)
	fmt.Println("FaucetCCaddress: ", fctadr.Result.FaucetCCaddress)
	fmt.Println("Faucetmarker: ", fctadr.Result.Faucetmarker)
	fmt.Println("GatewaysPubkey: ", fctadr.Result.GatewaysPubkey)
	fmt.Println("FaucetCCassets: ", fctadr.Result.FaucetCCassets)
	fmt.Println("MyCCaddress: ", fctadr.Result.MyCCaddress)
	fmt.Println("Myaddress: ", fctadr.Result.Myaddress)
}