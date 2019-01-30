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

type GetReceivedByAccount struct {
	Result	int64	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) GetReceivedByAccount(params APIParams) (GetReceivedByAccount, error) {
	
	// In call cases set account to blank. ACCOUNTS feature is DEPRICATED. It Should not be used anymore.
	if params[0] == nil {
		params[0] = ``
	} else {
		params[0] = ``
	}

	if params[1] == nil {
		params[1] = 1
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`getreceivedbyaccount`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var getreceivedbyaccount GetReceivedByAccount

	getreceivedbyaccountJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getreceivedbyaccountJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getreceivedbyaccountJson), &getreceivedbyaccount)
		return getreceivedbyaccount, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getreceivedbyaccountJson), &getreceivedbyaccount)
	return getreceivedbyaccount, nil
}