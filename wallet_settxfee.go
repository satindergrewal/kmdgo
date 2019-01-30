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
	"strconv"
)

type SetTxFee struct {
	Result bool `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) SetTxFee(amount float64) (SetTxFee, error) {
	query := APIQuery {
		Method:	`settxfee`,
		Params:	`[`+strconv.FormatFloat(amount, 'f', 8, 64)+`]`,
	}
	//fmt.Println(query)

	var settxfee SetTxFee

	settxfeeJson := appName.APICall(query)
	//fmt.Println(settxfeeJson)

	var result APIResult

	json.Unmarshal([]byte(settxfeeJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(settxfeeJson), &settxfee)
		return settxfee, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(settxfeeJson), &settxfee)
	return settxfee, nil
}