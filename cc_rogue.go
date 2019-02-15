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
	"fmt"
	"encoding/json"
	"errors"
)

type RGNewGame struct {
	Result struct {
		Name       string  `json:"name"`
		Method     string  `json:"method"`
		Maxplayers int     `json:"maxplayers"`
		Buyin      float64 `json:"buyin"`
		Type       string  `json:"type"`
		Hex        string  `json:"hex"`
		Txid       string  `json:"txid"`
		Result     string  `json:"result"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) RGNewGame(params APIParams) (RGNewGame, error) {
	var params_json string

	if params[0] == "" || params[0] == nil {
		params_json = `[]`
		//fmt.Println(params_json)
	} else {
		params_bytes, _ := json.Marshal(params)
		params_json = string(params_bytes)
		fmt.Println(params_json)
	}

	query := APIQuery {
		Method:	`cclib`,
		Params:	params_json,
	}
	//fmt.Println(query)

	var faucetaddress RGNewGame

	faucetaddressJson := appName.APICall(query)
	//fmt.Println(faucetaddressJson)

	var result APIResult

	json.Unmarshal([]byte(faucetaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetaddressJson), &faucetaddress)
		return faucetaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(faucetaddressJson), &faucetaddress)
	return faucetaddress, nil
}
