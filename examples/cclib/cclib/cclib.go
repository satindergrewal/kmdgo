// Copyright © 2018-2019 Satinderjit Singh.
//
// See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
// the top-level directory of this distribution for the individual copyright
// holder information and the developer policies on copyright and licensing.
//
// Unless otherwise agreed in a custom licensing agreement, no part of the
// kmdgo software, including this file may be copied, modified, propagated.
// or distributed except according to the terms contained in the LICENSE file
//
// Removal or modification of this copyright notice is prohibited.

package main

// import (
// 	"fmt"
// 	"log"
// 	"github.com/satindergrewal/kmdgo"
// )

// func main() {
// 	appName := kmdgo.NewAppType(`ROGUE`)

// 	var clb kmdgo.RGGameInfo

// 	args := make(kmdgo.APIParams, 3)
// 	args[0] = `gameinfo`
// 	args[1] = `17`
// 	args[2] = `["2ad99222dedd2ed4439501df13e2451e1497e148e2116ba2bb9afd7d42797812"]`
// 	fmt.Println(args)

// 	clb, err := appName.CCLib(args, )
// 	if err != nil {
// 		fmt.Printf("Code: %v\n", clb.Error.Code)
// 		fmt.Printf("Message: %v\n\n", clb.Error.Message)
// 		log.Fatalln("Err happened", err)
// 	}

// 	fmt.Println("clb value", clb)
// 	fmt.Println("-------")
// 	fmt.Println(clb.Result)
// }
