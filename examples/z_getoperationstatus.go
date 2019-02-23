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
	"github.com/satindergrewal/kmdgo"
	"log"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var oprst kmdgo.ZGetOperationStatus

	args := make(kmdgo.APIParams, 1)
	args[0] = []string{"opid-6e581ee5-4e90-4e70-8961-f95d8d28748c", "opid-6e581ee5-3e17-2e84-8961-j9fd8d28748r"}
	fmt.Println(args)

	oprst, err := appName.ZGetOperationStatus(args)
	if err != nil {
		fmt.Printf("Code: %v\n", oprst.Error.Code)
		fmt.Printf("Message: %v\n\n", oprst.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("oprst value", oprst)
	fmt.Println("-------")
	fmt.Println(oprst.Result)
	fmt.Println("-------")

	for i, v := range oprst.Result {
		fmt.Println(i)
		fmt.Println(v)
		fmt.Println("ID: ", v.ID)
		fmt.Println("Status: ", v.Status)
		fmt.Println("CreationTime: ", v.CreationTime)
		fmt.Println("Result Txid: ", v.Result.Txid)
		fmt.Println("ExecutionSecs: ", v.ExecutionSecs)
		fmt.Println("Method: ", v.Method)
		fmt.Println("Params->Fromaddress: ", v.Params.Fromaddress)
		fmt.Println("Minconf: ", v.Params.Minconf)
		fmt.Println("Fee: ", v.Params.Fee)

		for pi, pv := range v.Params.Amounts {
			fmt.Println(pi)
			fmt.Println("->Amounts-->Address: ", pv.Address)
			fmt.Println("->Amounts-->Amount: ", pv.Amount)
		}
	}

}
