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

	var comit kmdgo.RegisterNameCommitment

	args := make(kmdgo.APIParams, 3)
	args[0] = `sattest01` // name
	args[1] = `satinder@` // controladdress
	args[2] = `satinder@` // referralidentity
	fmt.Println(args)

	comit, err := appName.RegisterNameCommitment(args)
	if err != nil {
		fmt.Printf("Code: %v\n", comit.Error.Code)
		fmt.Printf("Message: %v\n\n", comit.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println(comit.Result)
	fmt.Println("Txid -", comit.Result.Txid)
	fmt.Println("Namereservation -", comit.Result.Namereservation)
	fmt.Println("Namereservation > Name -", comit.Result.Namereservation.Name)
	fmt.Println("Namereservation > Salt -", comit.Result.Namereservation.Salt)
	fmt.Println("Namereservation > Referral -", comit.Result.Namereservation.Referral)
	fmt.Println("Namereservation > Parent -", comit.Result.Namereservation.Parent)
	fmt.Println("Namereservation > Nameid -", comit.Result.Namereservation.Nameid)
}

// Example Output:
//
// Txid - 547514e2dd4483c719ef94e2d8171d9a445975a406264f4aea3008f2bb880d33
// Namereservation - {sattest01 b0ef0c7ca7e9aafdd908b13c8e30ac1b38a95ad6f62c2fae9a70530815561d39 i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18  iG9h8GM17z1ZByGvKMF7CmF88NVCdgbpmr}
// Namereservation > Name - sattest01
// Namereservation > Salt - b0ef0c7ca7e9aafdd908b13c8e30ac1b38a95ad6f62c2fae9a70530815561d39
// Namereservation > Referral - i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18
// Namereservation > Parent -
// Namereservation > Nameid - iG9h8GM17z1ZByGvKMF7CmF88NVCdgbpmr
