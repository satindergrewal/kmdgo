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
	appName = `pirate`

	var zvldadr kmdgo.ZValidateAddress

	zaddress := `zs12znwna63m9rad32a0wqrycxnzc7v99lp2d8yk9yfqtzyyflz4md4zjk6hlv792vpx48tkgyx9dw`

	zvldadr, err := appName.ZValidateAddress(zaddress)
	if err != nil {
		fmt.Printf("Code: %v\n", zvldadr.Error.Code)
		fmt.Printf("Message: %v\n\n", zvldadr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("zvldadr value", zvldadr)
	fmt.Println("-------")
	fmt.Println(zvldadr.Result)
	fmt.Println("-------")
	fmt.Println("Isvalid: ", zvldadr.Result.Isvalid)
	fmt.Println("Address: ", zvldadr.Result.Address)
	fmt.Println("Type: ", zvldadr.Result.Type)
	fmt.Println("Diversifier: ", zvldadr.Result.Diversifier)
	fmt.Println("Diversifiedtransmissionkey: ", zvldadr.Result.Diversifiedtransmissionkey)
	fmt.Println("Ismine: ", zvldadr.Result.Ismine)
}