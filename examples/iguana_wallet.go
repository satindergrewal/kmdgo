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

	"github.com/satindergrewal/kmdgo/kmdutil"
	"github.com/satindergrewal/kmdgo/kmdutil/bip39"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	iguanaSeed, _ := bip39.NewMnemonic(entropy)

	// iguanaSeed := "user specified seed phrase"
	var taddr kmdutil.IguanaTAddress
	taddr = kmdutil.GetTAddress(iguanaSeed)
	fmt.Println(taddr)
}
