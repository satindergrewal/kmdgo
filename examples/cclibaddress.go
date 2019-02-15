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
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	appName := kmdgo.NewAppType(`ROGUE`)

	var rgng kmdgo.CCLibAddress

	args := make(kmdgo.APIParams, 2)
	args[0] = `17`
	args[1] = `03778d6e5d4d20482ef8c72f5ee1f458ac0ebf1333a03e57187c436568eaf7ac31`
	fmt.Println(args)

	rgng, err := appName.CCLibAddress(args)
	if err != nil {
		fmt.Printf("Code: %v\n", rgng.Error.Code)
		fmt.Printf("Message: %v\n\n", rgng.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("rgng value", rgng)
	fmt.Println("-------")
	fmt.Println(rgng.Result)

	fmt.Println("\n")

}