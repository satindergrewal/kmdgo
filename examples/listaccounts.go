 //******************************************************************************
 //  Copyright © 2018-2019 Satinderjit Singh.                                   *
 //                                                                             *
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 //  the top-level directory of this distribution for the individual copyright  *
 //  holder information and the developer policies on copyright and licensing.  *
 //                                                                             *
 //  Unless otherwise agreed in a custom licensing agreement, no part of the    *
 //  kmdgo software, including this file may be copied, modified, propagated.   *
 //  or distributed except according to the terms contained in the LICENSE file *
 //                                                                             *
 //  Removal or modification of this copyright notice is prohibited.            *
 //                                                                             *
 //******************************************************************************/
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var lstact kmdgo.ListAccounts

	args := make(kmdgo.APIParams, 2)
	//args[0] = 0
	//args[1] = true
	fmt.Println(args)

	lstact, err := appName.ListAccounts(args)
	if err != nil {
		fmt.Printf("Code: %v\n", lstact.Error.Code)
		fmt.Printf("Message: %v\n\n", lstact.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("lstact value", lstact)
	fmt.Println("-------")
	fmt.Println(lstact.Result)

	for i, v := range lstact.Result {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		fmt.Printf("%0.8f\n",v)

	}

}