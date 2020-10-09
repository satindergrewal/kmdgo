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

// Example Output:
// -------
// 0
// Identity > Version - 2
// Identity > Flags - 0
// Identity > Primaryaddresses - [RLG8i7pX62iBcF15CP8mbT8ZmbyHtdqZMc]
// Identity > Minimumsignatures - 1
// Identity > Identityaddress - i5QMirsAsC6oRUCPF8HbF7mU2jUHzf6pdH
// Identity > Parent - iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq
// Identity > Name - grewal
// Identity > Contentmap - {}
// Identity > Revocationauthority - i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18
// Identity > Recoveryauthority - i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18
// Identity > Privateaddress - zs19eqgy366y5lhnvq8v2v7gaqa9uenjx2ryqf3guy34ffy9apdfcazflwyg8ncxk56dzvfz6xxc94
// Identity > Timelock - 0
// Blockheight - 20213
// Txid - cc0f3c8fc6eeb7f097ae293359947a8b798614441090f5642df554cc95554cc3
// Status - active
// Canspendfor - true
// Cansignfor - true

// -------
// 1
// Identity > Version - 2
// Identity > Flags - 0
// Identity > Primaryaddresses - [RPi1Dm5czxzcSb3VZq9mtDD2dUw5xUq3qB]
// Identity > Minimumsignatures - 1
// Identity > Identityaddress - i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18
// Identity > Parent - iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq
// Identity > Name - satinder
// Identity > Contentmap - {}
// Identity > Revocationauthority - i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18
// Identity > Recoveryauthority - i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18
// Identity > Privateaddress - zs19eqgy366y5lhnvq8v2v7gaqa9uenjx2ryqf3guy34ffy9apdfcazflwyg8ncxk56dzvfz6xxc94
// Identity > Timelock - 0
// Blockheight - 20168
// Txid - 8e28e30c03faf8a5b1562d4e320bfde6f222e94929b87263767133ba51b19351
// Status - active
// Canspendfor - true
// Cansignfor - true
