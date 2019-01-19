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

	var zmrgtadr kmdgo.ZMergeToAddress

	from_address := `["RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr", "zs1tyaqq9nstpk2ezvj5ayxg6nfkhlrc80dcset6v6jmpk9gft384v6rpgmxhu00u3aalygqgk77eg"]`
	to_address := `zs1krtqdjxt6f40t2d8f6jgeg4l702dg2v73s9e9znez856a7k7y6snu5qh4vp6mhtx0n9x5nx6wjk`
	//to_address := `RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s`
	fee := 0.0001
	transparent_limit := 50
	shielded_limit := 10
	memo := `f600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000`

	zmrgtadr, err := appName.ZMergeToAddress(from_address, to_address, fee, transparent_limit, shielded_limit, memo)
	if err != nil {
		fmt.Printf("Code: %v\n", zmrgtadr.Error.Code)
		fmt.Printf("Message: %v\n\n", zmrgtadr.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("zmrgtadr value", zmrgtadr)
	fmt.Println("-------")
	fmt.Println(zmrgtadr.Result)

	fmt.Println("RemainingUTXOs: ", zmrgtadr.Result.RemainingUTXOs)
	fmt.Println("RemainingTransparentValue: ", zmrgtadr.Result.RemainingTransparentValue)
	fmt.Println("RemainingNotes: ", zmrgtadr.Result.RemainingNotes)
	fmt.Println("RemainingShieldedValue: ", zmrgtadr.Result.RemainingShieldedValue)
	fmt.Println("MergingUTXOs: ", zmrgtadr.Result.MergingUTXOs)
	fmt.Println("MergingTransparentValue: ", zmrgtadr.Result.MergingTransparentValue)
	fmt.Println("MergingNotes: ", zmrgtadr.Result.MergingNotes)
	fmt.Println("MergingShieldedValue: ", zmrgtadr.Result.MergingShieldedValue)
	fmt.Println("Opid: ", zmrgtadr.Result.Opid)
}