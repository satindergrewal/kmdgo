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
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	appName := kmdgo.NewAppType(`ROGUE`)

	var ccladdr kmdgo.CCLibAddress

	args := make(kmdgo.APIParams, 2)
	args[0] = `17`
	args[1] = `03778d6e5d4d20482ef8c72f5ee1f458ac0ebf1333a03e57187c436568eaf7ac31`
	fmt.Println(args)

	ccladdr, err := appName.CCLibAddress(args)
	if err != nil {
		fmt.Printf("Code: %v\n", ccladdr.Error.Code)
		fmt.Printf("Message: %v\n\n", ccladdr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("ccladdr value", ccladdr)
	fmt.Println("-------")
	fmt.Println(ccladdr.Result)

	fmt.Println("\n")

	fmt.Println("Result: ", ccladdr.Result.Result)
	fmt.Println("CClibCCAddress: ", ccladdr.Result.CClibCCAddress)
	fmt.Printf("CCbalance: %0.8f\n", ccladdr.Result.CCbalance)
	fmt.Println("CClibNormalAddress: ", ccladdr.Result.CClibNormalAddress)
	fmt.Println("CClibCCTokensAddress: ", ccladdr.Result.CClibCCTokensAddress)
	fmt.Println("MyAddress: ", ccladdr.Result.MyAddress)
	fmt.Println("MyCCAddressCClib: ", ccladdr.Result.MyCCAddressCClib)
	fmt.Println("MyCCaddress: ", ccladdr.Result.MyCCaddress)
	fmt.Printf("MyCCbalance: %0.8f\n", ccladdr.Result.MyCCbalance)
	fmt.Println("Myaddress: ", ccladdr.Result.Myaddress)
	fmt.Printf("Mybalance: %0.8f\n", ccladdr.Result.Mybalance)
}