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
	//"encoding/json"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var sdrwtx kmdgo.SendRawTransaction

	args := make(kmdgo.APIParams, 2)
	args[0] = `0100000001958cb041d8369bbf6c2493accc4d949909a2c669cad883e232038d782eeb4fa4000000006a4730440220242c38740261799f9b6ccbde8f941e2567e86c84108c508d108d062ab9677b6e02206fea089b28c6d66d1c8f2343e1de7960dadafa3cf268c00f7dbe391cd8b9365f01210384c0db4f1eaa142a2745742b942f989375dbec32c55310a793225bb5c43cdc98ffffffff0140420f00000000001976a91456def632e67aa11c25ac16a0ee52893c2e5a2b6a88ac00000000`
	//args[1] = true
	fmt.Println(args)

	sdrwtx, err := appName.SendRawTransaction(args)
	if err != nil {
		fmt.Printf("Code: %v\n", sdrwtx.Error.Code)
		fmt.Printf("Message: %v\n\n", sdrwtx.Error.Message)
		log.Fatalln("Err happened", err)
	}
	fmt.Println("sdrwtx value", sdrwtx)
	fmt.Println("-------")
	fmt.Println(sdrwtx.Result)
}
