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
	appName = `komodo`

	var gprifo kmdgo.GetPeerInfo

	gprifo, err := appName.GetPeerInfo()
	if err != nil {
		fmt.Printf("Code: %v\n", gprifo.Error.Code)
		fmt.Printf("Message: %v\n\n", gprifo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("gprifo value", gprifo)
	fmt.Println("-------")
	fmt.Println(gprifo.Result)

	fmt.Println("-------")

	for i, v := range gprifo.Result {
		fmt.Println(i)
		fmt.Println("ID: ", v.ID)
		fmt.Println("Addr: ", v.Addr)
		fmt.Println("Addrlocal: ", v.Addrlocal)
		fmt.Println("Services: ", v.Services)
		fmt.Println("Lastsend: ", v.Lastsend)
		fmt.Println("Lastrecv: ", v.Lastrecv)
		fmt.Println("Bytessent: ", v.Bytessent)
		fmt.Println("Bytesrecv: ", v.Bytesrecv)
		fmt.Println("Conntime: ", v.Conntime)
		fmt.Println("Timeoffset: ", v.Timeoffset)
		fmt.Println("Pingtime: ", v.Pingtime)
		fmt.Println("Version: ", v.Version)
		fmt.Println("Subver: ", v.Subver)
		fmt.Println("Inbound: ", v.Inbound)
		fmt.Println("Startingheight: ", v.Startingheight)
		fmt.Println("Banscore: ", v.Banscore)
		fmt.Println("SyncedHeaders: ", v.SyncedHeaders)
		fmt.Println("SyncedBlocks: ", v.SyncedBlocks)
		fmt.Println("Inflight: ", v.Inflight)
		fmt.Println("Whitelisted: ", v.Whitelisted)
	}
}