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
	// Define appName type from kmdgo package
	var appName kmdgo.AppType
		
	// Define appname variable. The name value must be the matching value of it's data directory name.
	// Example Komodo's data directory is `komodo`, VerusCoin's data directory is `VRSC` and so on.
	appName = `komodo`
	
	// define the variable with GetBlockchainInfo struct from pacakge kmdgo
	var gb kmdgo.GetBlockchainInfo

	// Get output of the GetBlockchainInfo() method and store it to GetBlockchainInfo struct type variable
	gb, err := appName.GetBlockchainInfo()
	if err != nil {
		fmt.Printf("Code: %v\n", gb.Error.Code)
		fmt.Printf("Message: %v\n\n", gb.Error.Message)
		log.Fatalln("Err happened", err)
	}
	
	// Can print and use the struct variable outputs in further code logic. Check GetBlockchainInfo struct in package file.
	//fmt.Println(gb)
	//fmt.Println(gb.Result)

	fmt.Printf("Chain %v\n", gb.Result.Chain)
	fmt.Printf("Blocks %v\n", gb.Result.Blocks)
	fmt.Printf("Headers %v\n", gb.Result.Headers)
	fmt.Printf("Bestblockhash %v\n", gb.Result.Bestblockhash)
	fmt.Printf("Difficulty %0.8f\n", gb.Result.Difficulty)
	fmt.Printf("Verificationprogress %0.16f\n", gb.Result.Verificationprogress)
	fmt.Printf("Chainwork %v\n", gb.Result.Chainwork)
	fmt.Printf("Pruned %v\n", gb.Result.Pruned)
	fmt.Printf("Commitments %v\n", gb.Result.Commitments)
	fmt.Printf("ValuePools %v\n", gb.Result.ValuePools)


	//fmt.Println(gb.Error)
	//fmt.Println(gb.ID)
}