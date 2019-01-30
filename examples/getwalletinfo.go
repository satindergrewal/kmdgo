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

	var wltnf kmdgo.GetWalletInfo

	wltnf, err := appName.GetWalletInfo()
	if err != nil {
		fmt.Printf("Code: %v\n", wltnf.Error.Code)
		fmt.Printf("Message: %v\n\n", wltnf.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("wltnf value", wltnf)
	fmt.Println("-------")
	fmt.Println(wltnf.Result)
	fmt.Println("-------")
	
	fmt.Println("Walletversion: ", wltnf.Result.Walletversion)
	fmt.Println("Balance: ", wltnf.Result.Balance)
	fmt.Println("UnconfirmedBalance: ", wltnf.Result.UnconfirmedBalance)
	fmt.Println("ImmatureBalance: ", wltnf.Result.ImmatureBalance)
	fmt.Println("Txcount: ", wltnf.Result.Txcount)
	fmt.Println("Keypoololdest: ", wltnf.Result.Keypoololdest)
	fmt.Println("Keypoolsize: ", wltnf.Result.Keypoolsize)
	fmt.Println("UnlockedUntil: ", wltnf.Result.UnlockedUntil)
	fmt.Println("Paytxfee: ", wltnf.Result.Paytxfee)
	fmt.Println("Seedfp: ", wltnf.Result.Seedfp)
}