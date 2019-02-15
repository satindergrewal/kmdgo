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

import (
	"fmt"
    "log"
    "github.com/satindergrewal/kmdgo"
)

func main() {
    appName := kmdgo.NewAppType(`ROGUE`)

	var info kmdgo.CCLibInfo

    info, err := appName.CCLibInfo()
    if err != nil {
        fmt.Printf("Code: %v\n", info.Error.Code)
        fmt.Printf("Message: %v\n\n", info.Error.Message)
        log.Fatalln("Err happened", err)
    }

    fmt.Println(info)
    fmt.Println(info.Result)   
}