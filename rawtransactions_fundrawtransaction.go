 //  Copyright © 2018-2019 Satinderjit Singh.
 //
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
 //  the top-level directory of this distribution for the individual copyright
 //  holder information and the developer policies on copyright and licensing.
 //
 //  Unless otherwise agreed in a custom licensing agreement, no part of the
 //  kmdgo software, including this file may be copied, modified, propagated.
 //  or distributed except according to the terms contained in the LICENSE file
 //
 //  Removal or modification of this copyright notice is prohibited.
package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type FundRawTransaction struct {
	Result struct {
		Hex       string  `json:"hex"`
		Changepos int     `json:"changepos"`
		Fee       float64 `json:"fee"`
	} `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) FundRawTransaction(params APIParams) (FundRawTransaction, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`fundrawtransaction`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var fundrawtransaction FundRawTransaction

	fundrawtransactionJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(fundrawtransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(fundrawtransactionJson), &fundrawtransaction)
		return fundrawtransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(fundrawtransactionJson), &fundrawtransaction)
	return fundrawtransaction, nil
}