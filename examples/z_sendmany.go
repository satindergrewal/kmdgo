 //  Copyright © 2018-2019 Satinderjit Singh.
 //
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
 //  the top-level directory of this distribution for the individual copyright
 //  holder information and the developer policies on copyright and licensing.
 //
 //  Unless otherwise agreed in a custom licensing agreement, no part of the
 //  kmdgo software, including this file may be copied, modified, propagated.
 //  or distributed except according to the terms contained in the LICENSE file
 //
 //  Removal or modification of this copyright notice is prohibited.
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var zmrgtadr kmdgo.ZSendMany

	type amounts []struct {
		Address string  `json:"address"`
		Amount  float64 `json:"amount"`
		Memo    string  `json:"memo,omitempty"`
	}

	args := make(kmdgo.APIParams, 2)
	args[0] = `RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr`
	args[1] = amounts{
		{"zs1krtqdjxt6f40t2d8f6jgeg4l702dg2v73s9e9znez856a7k7y6snu5qh4vp6mhtx0n9x5nx6wjk", 5.0, "f600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"zs1tyaqq9nstpk2ezvj5ayxg6nfkhlrc80dcset6v6jmpk9gft384v6rpgmxhu00u3aalygqgk77eg", 0.03, ""},
	}
	//args[2] = 2
	//args[3] = 0.0003
	fmt.Println(args)

	zmrgtadr, err := appName.ZSendMany(args)
	if err != nil {
		fmt.Printf("Code: %v\n", zmrgtadr.Error.Code)
		fmt.Printf("Message: %v\n\n", zmrgtadr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("zmrgtadr value", zmrgtadr)
	fmt.Println("-------")
	fmt.Println(zmrgtadr.Result)
}