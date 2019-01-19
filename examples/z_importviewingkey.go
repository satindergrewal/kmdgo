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

	var ziky kmdgo.ZImportViewingKey

	zc_vewing_key := `ZiVKmp1MnmE78xJqL5x1MvUpTHYPF1UUetcgAwe5gYzqr98uDiC6qiVCLa6fQEbNMYUDcmoK6geLGQBDhc6FHoRoGGLKYcS9Y`
	rescan := `whenkeyisnew` // Rescan the wallet for transactions - Expected values are "yes", "no" or "whenkeyisnew"
	start_height := 0

	ziky, err := appName.ZImportViewingKey(zc_vewing_key, rescan, start_height)
	if err != nil {
		fmt.Printf("Code: %v\n", ziky.Error.Code)
		fmt.Printf("Message: %v\n\n", ziky.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("ziky value", ziky)
	fmt.Println("-------")
	fmt.Println(ziky.Result)
}