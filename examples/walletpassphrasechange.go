// Copyright © 2018-2019 Satinderjit Singh.
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

	var wltopn kmdgo.WalletPassPhrasechangeChange

	args := make(kmdgo.APIParams, 2)
	args[0] = `test123`
	args[1] = `test456`
	fmt.Println(args)

	wltopn, err := appName.WalletPassPhrasechangeChange(args)
	if err != nil {
		fmt.Printf("Code: %v\n", wltopn.Error.Code)
		fmt.Printf("Message: %v\n\n", wltopn.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("wltopn value", wltopn)
	fmt.Println("-------")
	fmt.Println(wltopn.Result)
}
