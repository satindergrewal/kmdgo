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

	var ct kmdgo.GetChainTips
	ct, err := appName.GetChainTips()
	if err != nil {
		fmt.Printf("Code: %v\n", ct.Error.Code)
		fmt.Printf("Message: %v\n\n", ct.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("ct value", ct)
	//fmt.Println(ct.Result)
}
