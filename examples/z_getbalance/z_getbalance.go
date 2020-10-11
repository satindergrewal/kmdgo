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

	var zblc kmdgo.ZGetBalance

	args := make(kmdgo.APIParams, 2)
	args[0] = `zs1tyaqq9nstpk2ezvj5ayxg6nfkhlrc80dcset6v6jmpk9gft384v6rpgmxhu00u3aalygqgk77eg`
	//args[1] = 1
	fmt.Println(args)

	zblc, err := appName.ZGetBalance(args)
	if err != nil {
		fmt.Printf("Code: %v\n", zblc.Error.Code)
		fmt.Printf("Message: %v\n\n", zblc.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("zblc value", zblc)
	fmt.Println("-------")
	fmt.Printf("\n%0.8f\n", zblc.Result)
}
