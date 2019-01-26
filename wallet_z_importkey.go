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

type ZImportKey struct {
	Result interface{} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZImportKey(params APIParams) (ZImportKey, error) {
	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = "whenkeyisnew"
		}
	}

	if len(params) == 3 {
		if params[2] == nil {
			params[2] = 0
		}
	}

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`z_importkey`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var z_importkey ZImportKey

	z_importkeyJson := appName.APICall(query)
	//fmt.Println(z_importkeyJson)

	var result APIResult

	json.Unmarshal([]byte(z_importkeyJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_importkeyJson), &z_importkey)
		return z_importkey, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_importkeyJson), &z_importkey)
	return z_importkey, nil
}