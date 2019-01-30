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
package kmdgo

import (
	"fmt"
	"github.com/satindergrewal/kmdgo/kmdutil"
)

func main() {
	appName := "komodo"
	dir := kmdutil.AppDataDir(appName, false)
	fmt.Println(dir)

	rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(appName)
	fmt.Printf("RPC User: %s\nRPC Password: %s\nRPC Port: %s\n", rpcuser, rpcpass, rpcport)
}