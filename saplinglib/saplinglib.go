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

package saplinglib

import (
	"runtime"
)

// IguanaSaplingAddress data type to parse saping address output from saplinglib
type IguanaSaplingAddress []struct {
	Num        int    `json:"num"`
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
	Seed       struct {
		HDSeed string `json:"HDSeed"`
		Path   string `json:"path"`
	} `json:"seed"`
}

// GetZAddress detects the OS and triggers the appropriate function (based on the OS it is running on)
// to generates a shielded sapling address using a seed phrase
func GetZAddress(nohd bool, zcount uint, iguanaSeed string, isIguanaSeed bool) IguanaSaplingAddress {
	var zaddr IguanaSaplingAddress
	switch os := runtime.GOOS; os {
	case "darwin":
		zaddr = GetZAddressOsx(nohd, zcount, iguanaSeed, isIguanaSeed)
	case "linux":
		zaddr = GetZAddressLinux(nohd, zcount, iguanaSeed, isIguanaSeed)
	case "windows":
		zaddr = GetZAddressWin(nohd, zcount, iguanaSeed, isIguanaSeed)
	default:
		return zaddr
	}
	return zaddr
}
