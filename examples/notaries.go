 //******************************************************************************
 //  Copyright © 2018-2019 Satinderjit Singh.                                   *
 //                                                                             *
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 //  the top-level directory of this distribution for the individual copyright  *
 //  holder information and the developer policies on copyright and licensing.  *
 //                                                                             *
 //  Unless otherwise agreed in a custom licensing agreement, no part of the    *
 //  kmdgo software, including this file may be copied, modified, propagated.   *
 //  or distributed except according to the terms contained in the LICENSE file *
 //                                                                             *
 //  Removal or modification of this copyright notice is prohibited.            *
 //                                                                             *
 //******************************************************************************/
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var ninfo kmdgo.Notaries

	param := `1187601` // Block Height
	//param := `1536365515` // Block Height Timestamp

	
	// Either use Block Height or Block Height Timestamp as input string value
	ninfo, err := appName.Notaries(param)
	if err != nil {
		fmt.Printf("Code: %v\n", ninfo.Error.Code)
		fmt.Printf("Message: %v\n\n", ninfo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	//fmt.Println("ninfo value", ninfo)
	fmt.Println("-------")
	fmt.Println(ninfo.Result)
	fmt.Println("-------")
	fmt.Println("Numnotaries: ", ninfo.Result.Numnotaries)
	fmt.Println("Height: ", ninfo.Result.Height)
	fmt.Println("Timestamp: ", ninfo.Result.Timestamp)
	fmt.Printf("\n\n\n\n")

	for i, v := range ninfo.Result.Notaries {
		fmt.Println(i)
		fmt.Println("Pubkey: ", v.Pubkey)
		fmt.Println("KMDaddress: ", v.KMDaddress)
		fmt.Println("BTCaddress: ", v.BTCaddress)
	}
}