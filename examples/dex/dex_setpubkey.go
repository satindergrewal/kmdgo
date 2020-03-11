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
	appName = `DEX`

	var dspkey kmdgo.DEXSetPubKey

	args := make(kmdgo.APIParams, 2)
	// pubkey33
	args[0] = "03732f8ef851ff234c74d0df575c2c5b159e2bab3faca4ec52b3f217d5cda5361d"
	fmt.Println(args)

	dspkey, err := appName.DEXSetPubKey(args)
	if err != nil {
		fmt.Printf("Code: %v\n", dspkey.Error.Code)
		fmt.Printf("Message: %v\n\n", dspkey.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("dspkey value", dspkey)
	fmt.Println("-------")
	// fmt.Println(dspkey.Result)
	// fmt.Println("-------")

	fmt.Println("Result", dspkey.Result.Result)
	fmt.Println("PublishablePubkey", dspkey.Result.PublishablePubkey)
	fmt.Println("Secpkey", dspkey.Result.Secpkey)
	fmt.Println("Recvaddr", dspkey.Result.Recvaddr)
	fmt.Println("RecvZaddr", dspkey.Result.RecvZaddr)
	fmt.Println("Handle", dspkey.Result.Handle)
	fmt.Println("Txpowbits", dspkey.Result.Txpowbits)
	fmt.Println("Vip", dspkey.Result.Vip)
	fmt.Println("Cmdpriority", dspkey.Result.Cmdpriority)
	fmt.Println("Perfstats", dspkey.Result.Perfstats)
}
