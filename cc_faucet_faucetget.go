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

type FaucetGet struct {
	Result struct {
		Result			string `json:"result"`
		Hex				string `json:"hex"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) FaucetGet() (FaucetGet, error) {
	query := APIQuery {
		Method:	`faucetget`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var faucetget FaucetGet

	faucetgetJson := appName.APICall(query)
	//fmt.Println(faucetgetJson)

	var result APIResult

	json.Unmarshal([]byte(faucetgetJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetgetJson), &faucetget)
		return faucetget, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(faucetgetJson), &faucetget)
	return faucetget, nil
}