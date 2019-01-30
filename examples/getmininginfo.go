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
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var mngfo kmdgo.GetMiningInfo

	mngfo, err := appName.GetMiningInfo()
	if err != nil {
		fmt.Printf("Code: %v\n", mngfo.Error.Code)
		fmt.Printf("Message: %v\n\n", mngfo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("mngfo value", mngfo)
	fmt.Println("-------")
	fmt.Println(mngfo.Result)

	fmt.Println("Blocks: ", mngfo.Result.Blocks)
	fmt.Println("Currentblocksize: ", mngfo.Result.Currentblocksize)
	fmt.Println("Currentblocktx: ", mngfo.Result.Currentblocktx)
	fmt.Println("Difficulty: ", mngfo.Result.Difficulty)
	fmt.Println("Errors: ", mngfo.Result.Errors)
	fmt.Println("Genproclimit: ", mngfo.Result.Genproclimit)
	fmt.Println("Localsolps: ", mngfo.Result.Localsolps)
	fmt.Println("Networksolps: ", mngfo.Result.Networksolps)
	fmt.Println("Networkhashps: ", mngfo.Result.Networkhashps)
	fmt.Println("Pooledtx: ", mngfo.Result.Pooledtx)
	fmt.Println("Testnet: ", mngfo.Result.Testnet)
	fmt.Println("Chain: ", mngfo.Result.Chain)
	fmt.Println("Generate: ", mngfo.Result.Generate)
	fmt.Println("Numthreads: ", mngfo.Result.Numthreads)
}