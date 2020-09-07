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

	var dcspt kmdgo.DecodeScript

	args := make(kmdgo.APIParams, 1)
	args[0] = `0400008085202f8901fdd6474f4a070e1b7c271f80fe79557f8c84f8eea4271fb66ee1666c2945bad70000000000ffffffff0240420f00000000001976a91478d98c7ad3345e56fbacc49710723b47dc119e8a88ac80841e00000000001976a914b87ba7e8e320fa7eb60f154909e94df1d22af40888ac000000001e0000000000000000000000000000`
	fmt.Println(args)

	dcspt, err := appName.DecodeScript(args)
	if err != nil {
		fmt.Printf("Code: %v\n", dcspt.Error.Code)
		fmt.Printf("Message: %v\n\n", dcspt.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("dcspt value", dcspt)
	//fmt.Println("-------")
	//fmt.Println(dcspt.Result)
	fmt.Println("-------")

	fmt.Println("Asm: ", dcspt.Result.Asm)
	fmt.Println("Type: ", dcspt.Result.Type)
	fmt.Println("P2Sh: ", dcspt.Result.P2Sh)
}
