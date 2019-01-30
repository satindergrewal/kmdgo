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

/*
This command gives the "hex" value as ouptut which then should be sent with commands like `sendrawtransaction`. The `txid` provided by this command can be decoded with `decoderawtransaction`.

$ txsclcc sendrawtransaction 0400008085202f8901167b86ccf6a887a52a66aad356f6e340ced215dec21cf174e59bfc7f2dc9e0e40000000049483045022100c4d0082c1cea98ce1a2e6e94de7dc1b41eaab1364c935bbbcd8cce8b20bf90ed02205c9dd41e3953ceb0ec8c1022c0c6d84d16b445a97d0ad2036819cb817b315a2d01ffffffff0200c2eb0b00000000302ea22c8020e029c511da55523565835887e412e5a0c9b920801b007000df45e545f25028248103120c008203000401ccf00f178900000000232102fbccd779e492bcf3b12707c1461ea857b4e2ee0ad990f13059c02189bd5d8edaac00000000576401000000000000000000000000
58a15cab5b14b6eab6daa9fd79f5bf69b5072593ebcfe06fe058ff2aebcbf0b2
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

	var fctfnd kmdgo.FaucetFund

	args := make(kmdgo.APIParams, 1)
	args[0] = `2` // This should have been a `int` value or `float64` value. But since the API only accepts string, for now this is input as string. It must be updated/changed as the reported changes are refelected via API.
	fmt.Println(args)

	fctfnd, err := appName.FaucetFund(args)
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
