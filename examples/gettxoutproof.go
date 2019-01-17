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

	var pf kmdgo.GetTxOutProof

	txids := `["0660e479396b71715ad29d6fdbf24632bb22a3719240230a77b44b83ef6fd8f2", "9d2e64dd827e30e38c06e060b51e76e38c6bc01c612ea82d6a7b239cc713ae68"]`

	pf, err := appName.GetTxOutProof(txids)
	if err != nil {
		fmt.Printf("Code: %v\n", pf.Error.Code)
		fmt.Printf("Message: %v\n\n", pf.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("pf value", pf)
	fmt.Println("-------")
	fmt.Println(pf.Result)
}