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

	var stbn kmdgo.SetBan

	ip_network := `192.168.0.0/24`
	command := `remove`
	//command := `add`
	ban_time := 86400
	absolute := false

	stbn, err := appName.SetBan(ip_network, command, ban_time, absolute)
	if err != nil {
		fmt.Printf("Code: %v\n", stbn.Error.Code)
		fmt.Printf("Message: %v\n\n", stbn.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("stbn value", stbn)
	fmt.Println("-------")
	fmt.Println(stbn.Result)
}