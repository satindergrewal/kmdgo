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

type SendFrom struct {
	Result string `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) SendFrom(params APIParams) (SendFrom, error) {
	if params[3] == nil {
		params[3] = 1
	}

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`sendfrom`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var sendfrom SendFrom

	sendfromJson := appName.APICall(query)
	//fmt.Println(sendfromJson)

	var result APIResult

	json.Unmarshal([]byte(sendfromJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendfromJson), &sendfrom)
		return sendfrom, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(sendfromJson), &sendfrom)
	return sendfrom, nil
}