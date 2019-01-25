/******************************************************************************
 * Copyright © 2018-2019 Satinderjit Singh.                                   *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * kmdgo software, including this file may be copied, modified, propagated.   *
 * or distributed except according to the terms contained in the LICENSE file *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/
package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type ZSendMany struct {
	Result string `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZSendMany(params APIParams) (ZSendMany, error) {
	if len(params) >= 3 {
		if params[2] == nil {
			params[2] = 1
		}
	}

	if len(params) >= 4 {
		if params[3] == nil {
			params[3] = 0.0001
		}
	}

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`z_sendmany`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var z_sendmany ZSendMany

	z_sendmanyJson := appName.APICall(query)
	//fmt.Println(z_sendmanyJson)

	var result APIResult

	json.Unmarshal([]byte(z_sendmanyJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_sendmanyJson), &z_sendmany)
		return z_sendmany, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_sendmanyJson), &z_sendmany)
	return z_sendmany, nil
}