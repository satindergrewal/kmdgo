/******************************************************************************
 * Copyright © 2018-2019 Satinderjit Singh.                                   *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * kmdgo software, including this file may be copied, modified, propagated.   *
 * or distributed except according to the terms contained in the LICENSE file *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var lkunspt kmdgo.SendFrom

	from_account := `*` // DO NOT USE account names (accounts are depricated). Can use "*" instead which will just select all accounts from wallet.
	to_address := `RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr`
	amount := 0.01
	minconf := 0
	comment := `donation`
	commentto := `Non Profit Org`

	lkunspt, err := appName.SendFrom(from_account, to_address, amount, minconf, comment, commentto)
	if err != nil {
		fmt.Printf("Code: %v\n", lkunspt.Error.Code)
		fmt.Printf("Message: %v\n\n", lkunspt.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("lkunspt value", lkunspt)
	fmt.Println("-------")
	fmt.Println(lkunspt.Result)
}