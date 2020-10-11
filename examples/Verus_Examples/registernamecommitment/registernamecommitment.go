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
	"encoding/json"
	"fmt"
	"log"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSCTEST`

	var comit kmdgo.RegisterNameCommitment

	args := make(kmdgo.APIParams, 4)
	args[0] = `sattest01`                          // name
	args[1] = `RG7NakxGExpWphZRZtFyHSiBUx7itw8b4s` // controladdress
	args[2] = `satinder@`                          // referralidentity
	args[2] = `satinder@`                          // parent
	// fmt.Println(args)

	comit, err := appName.RegisterNameCommitment(args)
	if err != nil {
		fmt.Printf("Code: %v\n", comit.Error.Code)
		fmt.Printf("Message: %v\n\n", comit.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println(comit.Result)
	comitJSON, _ := json.MarshalIndent(comit.Result, "", "  ")
	fmt.Println(string(comitJSON))
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
// {
// 	"txid": "37f3a18769d72fa1c2c461853ab66d46bae809a5e9ac0178bd395d2333234852",
// 	"namereservation": {
// 	  "name": "sattest01",
// 	  "salt": "9c98592e44017fb5ffcd3fcbb6e1e63fd18b25fcb34cffb4b5b03efdb1c17b61",
// 	  "referral": "i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18",
// 	  "parent": "",
// 	  "nameid": "iG9h8GM17z1ZByGvKMF7CmF88NVCdgbpmr"
// 	}
//   }
//   Txid - 37f3a18769d72fa1c2c461853ab66d46bae809a5e9ac0178bd395d2333234852
//   Namereservation - {sattest01 9c98592e44017fb5ffcd3fcbb6e1e63fd18b25fcb34cffb4b5b03efdb1c17b61 i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18  iG9h8GM17z1ZByGvKMF7CmF88NVCdgbpmr}
//   Namereservation > Name - sattest01
//   Namereservation > Salt - 9c98592e44017fb5ffcd3fcbb6e1e63fd18b25fcb34cffb4b5b03efdb1c17b61
//   Namereservation > Referral - i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18
//   Namereservation > Parent -
//   Namereservation > Nameid - iG9h8GM17z1ZByGvKMF7CmF88NVCdgbpmr
