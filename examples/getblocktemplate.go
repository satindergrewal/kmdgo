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

	var blksb kmdgo.GetBlockTemplate

	type JsonObjArg struct {
		Mode         string   `json:"mode"`
		Capabilities []string `json:"capabilities"`
	}

	args := make(kmdgo.APIParams, 1)
	args[0] = JsonObjArg{
		"template",
		[]string{"longpoll", "coinbasetxn", "coinbasevalue", "proposal", "serverlist", "workid"},
	}
	fmt.Println(args)

	blksb, err := appName.GetBlockTemplate(args)
	if err != nil {
		fmt.Printf("Code: %v\n", blksb.Error.Code)
		fmt.Printf("Message: %v\n\n", blksb.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("blksb value", blksb)
	fmt.Println("-------")
	fmt.Printf("%0.8f\n", blksb.Result)

	fmt.Println("Capabilities: ", blksb.Result.Capabilities)

	fmt.Println("Version: ", blksb.Result.Version)
	fmt.Println("Previousblockhash: ", blksb.Result.Previousblockhash)

	for i, v := range blksb.Result.Transactions {
		fmt.Println(i)
		fmt.Println(v)
		fmt.Println("Transactions --> Data: ", v.Data)
		fmt.Println("Transactions --> Hash: ", v.Hash)
		fmt.Println("Transactions --> Depends: ", v.Depends)
		fmt.Println("Transactions --> Fee: ", v.Fee)
		fmt.Println("Transactions --> Sigops: ", v.Sigops)
	}

	fmt.Println("Coinbasetxn --> Data: ", blksb.Result.Coinbasetxn.Data)
	fmt.Println("Coinbasetxn --> Hash: ", blksb.Result.Coinbasetxn.Hash)
	fmt.Println("Coinbasetxn --> Depends: ", blksb.Result.Coinbasetxn.Depends)
	fmt.Println("Coinbasetxn --> Fee: ", blksb.Result.Coinbasetxn.Fee)
	fmt.Println("Coinbasetxn --> Sigops: ", blksb.Result.Coinbasetxn.Sigops)
	fmt.Println("Coinbasetxn --> Coinbasevalue: ", blksb.Result.Coinbasetxn.Coinbasevalue)
	fmt.Println("Coinbasetxn --> Required: ", blksb.Result.Coinbasetxn.Required)

	fmt.Println("Longpollid: ", blksb.Result.Longpollid)
	fmt.Println("Target: ", blksb.Result.Target)
	fmt.Println("Mintime: ", blksb.Result.Mintime)

	for mi, mv := range blksb.Result.Mutable {
		fmt.Println("Mutable -->", mi)
		fmt.Println("Mutable -->", mv)
	}

	fmt.Println("Noncerange: ", blksb.Result.Noncerange)
	fmt.Println("Sigoplimit: ", blksb.Result.Sigoplimit)
	fmt.Println("Sizelimit: ", blksb.Result.Sizelimit)
	fmt.Println("Curtime: ", blksb.Result.Curtime)
	fmt.Println("Bits: ", blksb.Result.Bits)
	fmt.Println("Height: ", blksb.Result.Height)
}
