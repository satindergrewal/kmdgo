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

	var stats kmdgo.DEXStats

	stats, err := appName.DEXStats()
	if err != nil {
		fmt.Printf("Code: %v\n", stats.Error.Code)
		fmt.Printf("Message: %v\n\n", stats.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("stats value", stats)
	fmt.Println("-------")
	// fmt.Println(stats.Result)
	// fmt.Println("-------")

	fmt.Println("Result", stats.Result.Result)
	fmt.Println("PublishablePubkey", stats.Result.PublishablePubkey)
	fmt.Println("Secpkey", stats.Result.Secpkey)
	fmt.Println("Recvaddr", stats.Result.Recvaddr)
	fmt.Println("RecvZaddr", stats.Result.RecvZaddr)
	fmt.Println("Handle", stats.Result.Handle)
	fmt.Println("Txpowbits", stats.Result.Txpowbits)
	fmt.Println("Vip", stats.Result.Vip)
	fmt.Println("Cmdpriority", stats.Result.Cmdpriority)
	fmt.Println("Progress", stats.Result.Progress)
	fmt.Println("Perfstats", stats.Result.Perfstats)
}
