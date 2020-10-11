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

	var gdprifo kmdgo.GetDeprecationInfo

	gdprifo, err := appName.GetDeprecationInfo()
	if err != nil {
		fmt.Printf("Code: %v\n", gdprifo.Error.Code)
		fmt.Printf("Message: %v\n\n", gdprifo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("gdprifo value", gdprifo)
	fmt.Println("-------")
	fmt.Println(gdprifo.Result)
	fmt.Println("-------")

	fmt.Println("Version: ", gdprifo.Result.Version)
	fmt.Println("Subversion: ", gdprifo.Result.Subversion)
	fmt.Println("Deprecationheight: ", gdprifo.Result.Deprecationheight)
}
