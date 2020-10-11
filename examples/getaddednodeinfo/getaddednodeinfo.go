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

	var gninf kmdgo.GetAddedNodeInfo

	args := make(kmdgo.APIParams, 1)
	args[0] = true
	//args[1] = `207.148.4.183`
	fmt.Println(args)

	gninf, err := appName.GetAddedNodeInfo(args)
	if err != nil {
		fmt.Printf("Code: %v\n", gninf.Error.Code)
		fmt.Printf("Message: %v\n\n", gninf.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("gninf value", gninf)
	fmt.Println("-------")
	fmt.Println(gninf.Result)

	for i, v := range gninf.Result {
		fmt.Println(i)
		fmt.Println(v)
		fmt.Println("Addednode: ", v.Addednode)
		fmt.Println("Connected: ", v.Connected)
		for iv, vv := range v.Addresses {
			fmt.Println("-->", iv)
			fmt.Println("-->Address: ", vv.Address)
			fmt.Println("-->Connected: ", vv.Connected)
		}
	}
}
