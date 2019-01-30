 //******************************************************************************
 //  Copyright © 2018-2019 Satinderjit Singh.                                   *
 //                                                                             *
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 //  the top-level directory of this distribution for the individual copyright  *
 //  holder information and the developer policies on copyright and licensing.  *
 //                                                                             *
 //  Unless otherwise agreed in a custom licensing agreement, no part of the    *
 //  kmdgo software, including this file may be copied, modified, propagated.   *
 //  or distributed except according to the terms contained in the LICENSE file *
 //                                                                             *
 //  Removal or modification of this copyright notice is prohibited.            *
 //                                                                             *
 //******************************************************************************/
package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type SendMany struct {
	Result string `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) SendMany(params APIParams) (SendMany, error) {
	if len(params) > 2 {
		if params[2] == nil {
			params[2] = 1
		}
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`sendmany`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var sendmany SendMany

	sendmanyJson := appName.APICall(query)
	//fmt.Println(sendmanyJson)

	var result APIResult

	json.Unmarshal([]byte(sendmanyJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendmanyJson), &sendmany)
		return sendmany, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(sendmanyJson), &sendmany)
	return sendmany, nil
}