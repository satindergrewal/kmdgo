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
	// Define appName type from kmdgo package
	var appName kmdgo.AppType
		
	// Define appname variable. The name value must be the matching value of it's data directory name.
	// Example Komodo's data directory is `komodo`, VerusCoin's data directory is `VRSC` and so on.
	appName = `komodo`
	
	// define the variable with CoinSupply struct from pacakge kmdgo
	var cs kmdgo.CoinSupply

	args := make(kmdgo.APIParams, 1)
	args[0] = "10"
	fmt.Println(args)

	// Get output of the CoinSupply() method and store it to CoinSupply struct type variable
	cs, err := appName.CoinSupply(args)
	if err != nil {
		fmt.Printf("Code: %v\n", cs.Error.Code)
		fmt.Printf("Message: %v\n\n", cs.Error.Message)
		log.Fatalln("Err happened", err)
	}
	
	// Can print and use the struct variable outputs in further code logic. Check CoinSupply struct in package file.
	fmt.Println(cs)
	fmt.Println(cs.Result)
	
	fmt.Println("Result:", cs.Result.Result)
	fmt.Println("Coin:", cs.Result.Coin)
	fmt.Println("Height:", cs.Result.Height)
	fmt.Printf("Supply: %0.8f\n", cs.Result.Supply)
	fmt.Printf("Zfunds: %0.8f\n", cs.Result.Zfunds)
	fmt.Printf("Sprout: %0.8f\n", cs.Result.Sprout)
	fmt.Printf("Total: %0.8f\n", cs.Result.Total)
	fmt.Println(cs.Error)
	fmt.Println(cs.ID)
}