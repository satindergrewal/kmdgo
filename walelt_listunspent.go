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

type ListUnspent struct {
	Result []struct {
		Txid          string  `json:"txid"`
		Vout          int     `json:"vout"`
		Generated     bool    `json:"generated"`
		Address       string  `json:"address"`
		ScriptPubKey  string  `json:"scriptPubKey"`
		Amount        float64 `json:"amount"`
		Interest      float64 `json:"interest"`
		Confirmations int     `json:"confirmations"`
		Spendable     bool    `json:"spendable"`
		} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ListUnspent(params APIParams) (ListUnspent, error) {
	if params[0] == nil {
		params[0] = 1
	}
	if params[1] == nil {
		params[1] = 9999999
	}

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`listunspent`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var listunspent ListUnspent

	listunspentJson := appName.APICall(query)
	//fmt.Println(listunspentJson)

	var result APIResult

	json.Unmarshal([]byte(listunspentJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listunspentJson), &listunspent)
		return listunspent, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listunspentJson), &listunspent)
	return listunspent, nil
}