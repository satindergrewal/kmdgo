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
	"strconv"
)

type ZSendMany struct {
	Result string `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZSendMany(frmaddr string, amounts string, minconf int, fee float64) (ZSendMany, error) {
	query := APIQuery {
		Method:	`z_sendmany`,
		Params:	`["`+frmaddr+`", `+amounts+`, `+strconv.Itoa(minconf)+`, `+strconv.FormatFloat(fee, 'f', 8, 64)+`]`,
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