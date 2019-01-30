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
	
	// define the variable with GetBlock struct from pacakge kmdgo
	var gb kmdgo.GetBlock

	args := make(kmdgo.APIParams, 2)
	args[0] = "08cddc39c6a04e0a7c9bb57477b1f4fa2ad218aa8c67f7ed406bfb470c4e0acd"
	//args[1] = 2
	fmt.Println(args)

	// Get output of the GetBlock() method and store it to GetBlock struct type variable
	gb, err := appName.GetBlock(args)
	if err != nil {
		fmt.Printf("Code: %v\n", gb.Error.Code)
		fmt.Printf("Message: %v\n\n", gb.Error.Message)
		log.Fatalln("Err happened", err)
	}
	
	// Can print and use the struct variable outputs in further code logic. Check GetBlock struct in package file.
	//fmt.Println(gb)
	//fmt.Println(gb.Result)
	
	fmt.Println("Hash", gb.Result.Hash)
	fmt.Println("Confirmations", gb.Result.Confirmations)
	fmt.Println("Rawconfirmations", gb.Result.Rawconfirmations)
	fmt.Println("Size", gb.Result.Size)
	fmt.Println("Height", gb.Result.Height)
	fmt.Println("Version", gb.Result.Version)
	fmt.Println("Merkleroot", gb.Result.Merkleroot)
	fmt.Println("Segid", gb.Result.Segid)
	fmt.Println("Finalsaplingroot", gb.Result.Finalsaplingroot)
	fmt.Println("Tx", gb.Result.Tx)

	for txi, txv := range gb.Result.Tx {
		fmt.Printf("-->%v\n", txi)
		fmt.Println("Txid", txv.Txid)
		fmt.Println("Overwintered", txv.Overwintered)
		fmt.Println("Version", txv.Version)
		fmt.Println("Locktime", txv.Locktime)
	}
}