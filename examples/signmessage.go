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

	var sngmg kmdgo.SignMessage

	taddress := `REAAchKmsc3aUFAwWhMh1eSKTAyGyTCxXb` //This address must be ismine:true in your wallet. Must have it's privkey imported.
	msg := `hello kmd world!`

	sngmg, err := appName.SignMessage(taddress, msg)
	if err != nil {
		fmt.Printf("Code: %v\n", sngmg.Error.Code)
		fmt.Printf("Message: %v\n\n", sngmg.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("sngmg value", sngmg)
	fmt.Println("-------")
	fmt.Println(sngmg.Result)
}

/*
Expected output:
ICDer79Dlio7/F18nkTefxIU7zh9oeplpY/IHnA4TxolcTDrtD4s5VuXDnCrUERk9AMbWCrHwWJzDyVGwNi23AU=
*/