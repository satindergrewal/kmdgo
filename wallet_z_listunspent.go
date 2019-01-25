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

type ZListUnspent struct {
	Result []struct {
		Txid          string  `json:"txid"`
		Jsindex       int     `json:"jsindex"`
		Jsoutindex    int     `json:"jsoutindex"`
		Outindex      int     `json:"outindex"`
		Confirmations int     `json:"confirmations"`
		Spendable     bool    `json:"spendable"`
		Address       string  `json:"address"`
		Amount        float64 `json:"amount"`
		Memo          string  `json:"memo"`
		Change        bool    `json:"change"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZListUnspent(params APIParams) (ZListUnspent, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = 1
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = 9999999
		}
	}

	if len(params) >= 3 {
		if params[2] == nil {
			params[2] = false
		}
	}
	
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`z_listunspent`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var z_listunspent ZListUnspent

	z_listunspentJson := appName.APICall(query)
	//fmt.Println(z_listunspentJson)

	var result APIResult

	json.Unmarshal([]byte(z_listunspentJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_listunspentJson), &z_listunspent)
		return z_listunspent, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_listunspentJson), &z_listunspent)
	return z_listunspent, nil
}