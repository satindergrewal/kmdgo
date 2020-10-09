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
	appName = `VRSCTEST`

	var getid kmdgo.GetIdentity

	args := make(kmdgo.APIParams, 1)
	args[0] = `mike@`
	fmt.Println(args)

	getid, err := appName.GetIdentity(args)
	if err != nil {
		fmt.Printf("Code: %v\n", getid.Error.Code)
		fmt.Printf("Message: %v\n\n", getid.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println(getid.Result)
	fmt.Println("Identity > Version -", getid.Result.Identity.Version)
	fmt.Println("Identity > Flags -", getid.Result.Identity.Flags)
	fmt.Println("Identity > Primaryaddresses -", getid.Result.Identity.Primaryaddresses)
	fmt.Println("Identity > Minimumsignatures -", getid.Result.Identity.Minimumsignatures)
	fmt.Println("Identity > Identityaddress -", getid.Result.Identity.Identityaddress)
	fmt.Println("Identity > Parent -", getid.Result.Identity.Parent)
	fmt.Println("Identity > Name -", getid.Result.Identity.Name)
	fmt.Println("Identity > Contentmap -", getid.Result.Identity.Contentmap)
	fmt.Println("Identity > Revocationauthority -", getid.Result.Identity.Revocationauthority)
	fmt.Println("Identity > Recoveryauthority -", getid.Result.Identity.Recoveryauthority)
	fmt.Println("Identity > Privateaddress -", getid.Result.Identity.Privateaddress)
	fmt.Println("Identity > Timelock -", getid.Result.Identity.Timelock)
	fmt.Println("Blockheight -", getid.Result.Blockheight)
	fmt.Println("Txid -", getid.Result.Txid)
	fmt.Println("Status -", getid.Result.Status)
	fmt.Println("Canspendfor -", getid.Result.Canspendfor)
	fmt.Println("Cansignfor -", getid.Result.Cansignfor)
}
