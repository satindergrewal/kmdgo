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

	var lrbaddr kmdgo.ListReceivedByAddress

	min_confs := 0
	include_empty := true
	include_watch_only := true

	lrbaddr, err := appName.ListReceivedByAddress(min_confs, include_empty, include_watch_only)
	if err != nil {
		fmt.Printf("Code: %v\n", lrbaddr.Error.Code)
		fmt.Printf("Message: %v\n\n", lrbaddr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("lrbaddr value", lrbaddr)
	fmt.Println("-------")
	fmt.Println(lrbaddr.Result)

	for i, v := range lrbaddr.Result {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		fmt.Println("Address: ", v.Address)
		fmt.Println("Account: ", v.Account)
		fmt.Println("Amount: ", v.Amount)
		fmt.Println("Rawconfirmations: ", v.Rawconfirmations)
		fmt.Println("Confirmations: ", v.Confirmations)
		fmt.Println("Txids: ", v.Txids)
	}

}