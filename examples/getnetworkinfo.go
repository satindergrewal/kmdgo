/******************************************************************************
 * Copyright © 2018-2019 Satinderjit Singh.                                   *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * kmdgo software, including this file may be copied, modified, propagated.   *
 * or distributed except according to the terms contained in the LICENSE file *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var gntifo kmdgo.GetNetworkInfo

	gntifo, err := appName.GetNetworkInfo()
	if err != nil {
		fmt.Printf("Code: %v\n", gntifo.Error.Code)
		fmt.Printf("Message: %v\n\n", gntifo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("gntifo value", gntifo)
	fmt.Println("-------")
	fmt.Println(gntifo.Result)
	fmt.Println("-------")

	fmt.Println("Version: ", gntifo.Result.Version)
	fmt.Println("Subversion: ", gntifo.Result.Subversion)
	fmt.Println("Protocolversion: ", gntifo.Result.Protocolversion)
	fmt.Println("Localservices: ", gntifo.Result.Localservices)
	fmt.Println("Timeoffset: ", gntifo.Result.Timeoffset)
	fmt.Println("Connections: ", gntifo.Result.Connections)
	fmt.Println("Networks: ", gntifo.Result.Networks)

	for i, v := range gntifo.Result.Networks {
		fmt.Println(i)
		fmt.Println("Networks->Name: ", v.Name)
		fmt.Println("Networks->Limited: ", v.Limited)
		fmt.Println("Networks->Reachable: ", v.Reachable)
		fmt.Println("Networks->Proxy: ", v.Proxy)
		fmt.Println("Networks->ProxyRandomizeCredentials: ", v.ProxyRandomizeCredentials)
	}

	fmt.Println("Relayfee: ", gntifo.Result.Relayfee)
	fmt.Println("Localaddresses: ", gntifo.Result.Localaddresses)
	fmt.Println("Warnings: ", gntifo.Result.Warnings)
}