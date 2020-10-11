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

	var ziky kmdgo.ZImportKey

	args := make(kmdgo.APIParams, 1)
	args[0] = `SKxuTXwjHtZyFEFoc6oPtTzQFrqrwnNU4MpLh15vs1VToGmTPa2G`
	//args[1] = `whenkeyisnew` // Rescan the wallet for transactions - Expected values are "yes", "no" or "whenkeyisnew"
	//args[2] = 0
	fmt.Println(args)

	ziky, err := appName.ZImportKey(args)
	if err != nil {
		fmt.Printf("Code: %v\n", ziky.Error.Code)
		fmt.Printf("Message: %v\n\n", ziky.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("ziky value", ziky)
	fmt.Println("-------")
	fmt.Println(ziky.Result)
}
