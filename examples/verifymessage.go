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

	var vfymg kmdgo.VerifyMessage

	args := make(kmdgo.APIParams, 3)
	args[0] = `REAAchKmsc3aUFAwWhMh1eSKTAyGyTCxXb` //The address from which the signed message is from. Check SignMessage example.
	args[1] = `ICDer79Dlio7/F18nkTefxIU7zh9oeplpY/IHnA4TxolcTDrtD4s5VuXDnCrUERk9AMbWCrHwWJzDyVGwNi23AU=`
	args[2] = `hello kmd world!`
	fmt.Println(args)

	vfymg, err := appName.VerifyMessage(args)
	if err != nil {
		fmt.Printf("Code: %v\n", vfymg.Error.Code)
		fmt.Printf("Message: %v\n\n", vfymg.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("vfymg value", vfymg)
	fmt.Println("-------")
	fmt.Println(vfymg.Result)
}

/*
Expected Output:
true or false
*/