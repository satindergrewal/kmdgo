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

type ZGetOperationResult struct {
	Result []struct {
		ID           string `json:"id"`
		Status       string `json:"status"`
		CreationTime int    `json:"creation_time"`
		Result       struct {
			Txid string `json:"txid"`
		} `json:"result"`
		ExecutionSecs float64 `json:"execution_secs"`
		Method        string  `json:"method"`
		Params        struct {
			Fromaddress string `json:"fromaddress"`
			Amounts     []struct {
				Address string  `json:"address"`
				Amount  float64 `json:"amount"`
			} `json:"amounts"`
			Minconf int     `json:"minconf"`
			Fee     float64 `json:"fee"`
		} `json:"params"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZGetOperationResult(params APIParams) (ZGetOperationResult, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`z_getoperationresult`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var z_getoperationresult ZGetOperationResult

	z_getoperationresultJson := appName.APICall(query)
	//fmt.Println(z_getoperationresultJson)

	var result APIResult

	json.Unmarshal([]byte(z_getoperationresultJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_getoperationresultJson), &z_getoperationresult)
		return z_getoperationresult, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_getoperationresultJson), &z_getoperationresult)
	return z_getoperationresult, nil
}