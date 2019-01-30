 //  Copyright © 2018-2019 Satinderjit Singh.
 //
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
 //  the top-level directory of this distribution for the individual copyright
 //  holder information and the developer policies on copyright and licensing.
 //
 //  Unless otherwise agreed in a custom licensing agreement, no part of the
 //  kmdgo software, including this file may be copied, modified, propagated.
 //  or distributed except according to the terms contained in the LICENSE file
 //
 //  Removal or modification of this copyright notice is prohibited.
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var opst kmdgo.ZListOperationIDs

	args := make(kmdgo.APIParams, 1)
	//args[0] = ``
	//args[0] = `success`
	fmt.Println(args)

	opst, err := appName.ZListOperationIDs(args)
	if err != nil {
		fmt.Printf("Code: %v\n", opst.Error.Code)
		fmt.Printf("Message: %v\n\n", opst.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("opst value", opst)
	fmt.Println("-------")
	fmt.Println(opst.Result)

	for i, v := range opst.Result {
		fmt.Println(i)
		fmt.Println(v)
	}
}