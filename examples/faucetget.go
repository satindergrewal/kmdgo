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

/*
The faucetget method requests the faucet contract to send coins.

The method returns a hex value which must then be broadcast using the sendrawtransaction method.

The faucetget command yields 0.1 coins and requires about 30 seconds of CPU time to execute.
*/


package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `TXSCLCC`

	var fctfnd kmdgo.FaucetGet

	fctfnd, err := appName.FaucetGet()
	if err != nil {
		fmt.Printf("Code: %v\n", fctfnd.Error.Code)
		fmt.Printf("Message: %v\n\n", fctfnd.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("fctfnd value", fctfnd)
	fmt.Println("-------")
	fmt.Println(fctfnd.Result)

	fmt.Println("Result: ", fctfnd.Result.Result)
	fmt.Println("Hex: ", fctfnd.Result.Hex)
}
