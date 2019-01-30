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
	var appName kmdgo.AppType
	appName = `komodo`

	var lstgrp kmdgo.ListAddressGroupings

	lstgrp, err := appName.ListAddressGroupings()
	if err != nil {
		fmt.Printf("Code: %v\n", lstgrp.Error.Code)
		fmt.Printf("Message: %v\n\n", lstgrp.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("lstgrp value", lstgrp)
	fmt.Println("-------")
	fmt.Println(lstgrp.Result)

	for i, v := range lstgrp.Result[0] {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		fmt.Printf("Address: %s\n",v[0])
		fmt.Printf("Amount: %0.8f\n",v[1])
		//fmt.Printf("Account Name: %s\n",v[2])

	}

}