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

	"github.com/satindergrewal/kmdgo/kmdutil/bip39"
	"github.com/satindergrewal/saplinglib"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	iguanaSeed, _ := bip39.NewMnemonic(entropy)

	// // iguanaSeed := "user specified seed phrase"
	// var taddr kmdutil.IguanaPublicAddress
	// taddr = kmdutil.GetTAddress(iguanaSeed)
	// // fmt.Printf("%+v\n", taddr)
	// fmt.Println("Seed: \t\t\t", taddr.Seed)
	// fmt.Println("Transparent Address: \t", taddr.Address)
	// fmt.Println("Pubkey: \t\t", taddr.Pubkey)
	// fmt.Println("WifC: \t\t\t", taddr.WifC)
	// fmt.Println("WifU: \t\t\t", taddr.WifU)

	var zaddr saplinglib.IguanaSaplingAddress
	nohd := false
	zcount := uint(1)
	isIguanaSeed := true
	coinType := uint(141)

	zaddr = saplinglib.GetZAddress(nohd, zcount, iguanaSeed, isIguanaSeed, coinType)
	// fmt.Println(zaddr)

	fmt.Println(zaddr[0].Num)
	fmt.Println(zaddr[0].Address)
	fmt.Println(zaddr[0].PrivateKey)
	fmt.Println(zaddr[0].Seed.HDSeed)
	fmt.Println(zaddr[0].Seed.Path)

	// var wallet kmdutil.IguanaWallet
	// wallet = kmdutil.GetIguanaWallet("user specified seed phrase")

	// fmt.Println("Seed: \t\t\t", wallet.Seed)
	// fmt.Println("Public Address: \t", wallet.Address)
	// fmt.Println("Pubkey: \t\t", wallet.Pubkey)
	// fmt.Println("Wif Compressed: \t", wallet.WifC)
	// fmt.Println("Wif Uncompressed: \t", wallet.WifU)
	// fmt.Println("Shielded Address: \t", wallet.ZAddress)
	// fmt.Println("Shielded PrivateKey: \t", wallet.ZPrivateKey)
}

// OSX:
// env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security" go run iguana_wallet.go

// Windows:
// env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc" CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go run iguana_wallet.go

// Linux:
// env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm" CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go run iguana_wallet.go
