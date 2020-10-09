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

	var lsids kmdgo.ListIdentities

	args := make(kmdgo.APIParams, 0)
	// fmt.Println(args)

	lsids, err := appName.ListIdentities(args)
	if err != nil {
		fmt.Printf("Code: %v\n", lsids.Error.Code)
		fmt.Printf("Message: %v\n\n", lsids.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println("lsids value", lsids)
	// fmt.Println("-------")
	// fmt.Println(lsids.Result)
	// fmt.Println("-------")

	for i, v := range lsids.Result {
		fmt.Printf("\n-------\n")
		fmt.Println(i)
		// fmt.Println("Identity", v.Identity)
		fmt.Println("Identity > Version -", v.Identity.Version)
		fmt.Println("Identity > Flags -", v.Identity.Flags)
		fmt.Println("Identity > Primaryaddresses -", v.Identity.Primaryaddresses)
		fmt.Println("Identity > Minimumsignatures -", v.Identity.Minimumsignatures)
		fmt.Println("Identity > Identityaddress -", v.Identity.Identityaddress)
		fmt.Println("Identity > Parent -", v.Identity.Parent)
		fmt.Println("Identity > Name -", v.Identity.Name)
		fmt.Println("Identity > Contentmap -", v.Identity.Contentmap)
		fmt.Println("Identity > Revocationauthority -", v.Identity.Revocationauthority)
		fmt.Println("Identity > Recoveryauthority -", v.Identity.Recoveryauthority)
		fmt.Println("Identity > Privateaddress -", v.Identity.Privateaddress)
		fmt.Println("Identity > Timelock -", v.Identity.Timelock)
		fmt.Println("Blockheight -", v.Blockheight)
		fmt.Println("Txid -", v.Txid)
		fmt.Println("Status -", v.Status)
		fmt.Println("Canspendfor -", v.Canspendfor)
		fmt.Println("Cansignfor -", v.Cansignfor)

	}
}
