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
	"github.com/satindergrewal/kmdgo"
	"log"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var stbn kmdgo.SetBan

	args := make(kmdgo.APIParams, 4) // Change length size from 4 to 2 if args 2 and 3 is disabled or adjust the length accordingly.
	args[0] = `192.168.0.0/24`       //"ip(/netmask)" (string, required)
	args[1] = `remove`               // "command" (string, required) 'add' to add a IP/Subnet to the list, 'remove' to remove a IP/Subnet from the list
	//args[1] = `add`
	args[2] = 86400 // (numeric, optional) time in seconds
	args[3] = false // (boolean, optional) If set, the bantime must be a absolute timestamp in seconds since epoch (Jan 1 1970 GMT)
	fmt.Println(args)

	stbn, err := appName.SetBan(args)
	if err != nil {
		fmt.Printf("Code: %v\n", stbn.Error.Code)
		fmt.Printf("Message: %v\n\n", stbn.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("stbn value", stbn)
	fmt.Println("-------")
	fmt.Println(stbn.Result)
}
