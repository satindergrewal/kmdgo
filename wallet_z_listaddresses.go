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

type ZListAddresses struct {
	Result []string `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZListAddresses(params APIParams) (ZListAddresses, error) {
	if params[0] == nil {
		params[0] = false
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`z_listaddresses`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var z_listaddresses ZListAddresses

	z_listaddressesJson := appName.APICall(query)
	//fmt.Println(z_listaddressesJson)

	var result APIResult

	json.Unmarshal([]byte(z_listaddressesJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_listaddressesJson), &z_listaddresses)
		return z_listaddresses, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_listaddressesJson), &z_listaddresses)
	return z_listaddresses, nil
}