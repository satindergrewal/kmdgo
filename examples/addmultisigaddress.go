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
/*
$ komodo-cli getnewaddress
RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr
$ komodo-cli getnewaddress
RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s
$ komodo-cli addmultisigaddress 2 '["RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr","RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s"]'
bWDExTNrQZ4kSRRXwUUgHibyYuzPLS6FgP
*/

package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var mtsig kmdgo.AddMultiSigAddress

	args := make(kmdgo.APIParams, 2)
	args[0] = 2 // This example says requires 2 of 2 signatures to create tx from the multisig address
	args[1] = []string{"RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr","RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s"} // Get the address from command `getnewaddress` to test
	fmt.Println(args)

	mtsig, err := appName.AddMultiSigAddress(args)
	if err != nil {
		fmt.Printf("Code: %v\n", mtsig.Error.Code)
		fmt.Printf("Message: %v\n\n", mtsig.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("mtsig value", mtsig)
	fmt.Println("-------")
	fmt.Println(mtsig.Result)
}