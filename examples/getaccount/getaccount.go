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

	var grbact kmdgo.GetAccount

	taddress := `RA1vrNf9Aej4Rt6gu4xNG66Zcy2KgjnPia`

	grbact, err := appName.GetAccount(taddress)
	if err != nil {
		fmt.Printf("Code: %v\n", grbact.Error.Code)
		fmt.Printf("Message: %v\n\n", grbact.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("grbact value", grbact)
	fmt.Println("-------")
	fmt.Println(grbact.Result)
}
