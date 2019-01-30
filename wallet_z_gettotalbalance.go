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

type ZGetTotalBalance struct {
	Result struct {
		Transparent float64 `json:"transparent"`
		Interest    float64 `json:"interest"`
		Private     float64 `json:"private"`
		Total       float64 `json:"total"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZGetTotalBalance(params APIParams) (ZGetTotalBalance, error) {
	if params[0] == nil {
		params[0] = 1
	}
	if params[1] == nil {
		params[1] = false
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`z_gettotalbalance`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var z_gettotalbalance ZGetTotalBalance

	z_gettotalbalanceJson := appName.APICall(query)
	//fmt.Println(z_gettotalbalanceJson)

	var result APIResult

	json.Unmarshal([]byte(z_gettotalbalanceJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_gettotalbalanceJson), &z_gettotalbalance)
		return z_gettotalbalance, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_gettotalbalanceJson), &z_gettotalbalance)
	return z_gettotalbalance, nil
}