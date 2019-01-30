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



package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type SetPubKey struct {
	Result struct {
		Pubkey   string `json:"pubkey"`
		Ismine   bool   `json:"ismine"`
		RAddress string `json:"R-address"`
		} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) SetPubKey(pubkey string) (SetPubKey, error) {
	query := APIQuery {
		Method:	`setpubkey`,
		Params:	`["`+pubkey+`"]`,
	}
	//fmt.Println(query)

	var setpubkey SetPubKey

	setpubkeyJson := appName.APICall(query)
	//fmt.Println(setpubkeyJson)

	var result APIResult

	json.Unmarshal([]byte(setpubkeyJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(setpubkeyJson), &setpubkey)
		return setpubkey, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(setpubkeyJson), &setpubkey)
	return setpubkey, nil
}