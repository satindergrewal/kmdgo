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

	var zlrbadr kmdgo.ZListReceivedByAddress

	args := make(kmdgo.APIParams, 1)
	args[0] = `zs1tyaqq9nstpk2ezvj5ayxg6nfkhlrc80dcset6v6jmpk9gft384v6rpgmxhu00u3aalygqgk77eg`
	//args[1] = 1
	fmt.Println(args)

	zlrbadr, err := appName.ZListReceivedByAddress(args)
	if err != nil {
		fmt.Printf("Code: %v\n", zlrbadr.Error.Code)
		fmt.Printf("Message: %v\n\n", zlrbadr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("zlrbadr value", zlrbadr)
	fmt.Println("-------")
	fmt.Println(zlrbadr.Result)
	fmt.Println("-------")
	
	for i, v := range zlrbadr.Result {
		fmt.Println(i)
		fmt.Println("Txid: ", v.Txid)
		fmt.Printf("Amount: %0.8f\n", v.Amount)
		fmt.Println("Memo: ", v.Memo)
	}
}