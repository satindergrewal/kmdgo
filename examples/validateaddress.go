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

	var vldadr kmdgo.ValidateAddress

	taddress := `REAAchKmsc3aUFAwWhMh1eSKTAyGyTCxXb`

	vldadr, err := appName.ValidateAddress(taddress)
	if err != nil {
		fmt.Printf("Code: %v\n", vldadr.Error.Code)
		fmt.Printf("Message: %v\n\n", vldadr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("vldadr value", vldadr)
	fmt.Println("-------")
	fmt.Println(vldadr.Result)
	fmt.Println("-------")
	fmt.Println("Isvalid: ", vldadr.Result.Isvalid)
	fmt.Println("Address: ", vldadr.Result.Address)
	fmt.Println("ScriptPubKey: ", vldadr.Result.ScriptPubKey)
	fmt.Println("Segid: ", vldadr.Result.Segid)
	fmt.Println("Ismine: ", vldadr.Result.Ismine)
	fmt.Println("Iswatchonly: ", vldadr.Result.Iswatchonly)
	fmt.Println("Isscript: ", vldadr.Result.Isscript)
	fmt.Println("Pubkey: ", vldadr.Result.Pubkey)
	fmt.Println("Iscompressed: ", vldadr.Result.Iscompressed)
	fmt.Println("Account: ", vldadr.Result.Account)
}