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

package saplinggo

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
