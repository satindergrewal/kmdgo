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

	var clrbnd kmdgo.ClearBanned

	clrbnd, err := appName.ClearBanned()
	if err != nil {
		fmt.Printf("Code: %v\n", clrbnd.Error.Code)
		fmt.Printf("Message: %v\n\n", clrbnd.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("clrbnd value", clrbnd)
	fmt.Println("-------")
	fmt.Println(clrbnd.Result)
}
