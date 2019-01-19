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

type ZShieldCoinbase struct {
	Result struct {
		RemainingUTXOs int     `json:"remainingUTXOs"`
		RemainingValue float64 `json:"remainingValue"`
		ShieldingUTXOs int     `json:"shieldingUTXOs"`
		ShieldingValue float64 `json:"shieldingValue"`
		Opid           string  `json:"opid"`
		} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZShieldCoinbase(frmaddr string, toaddr string, fee float64, limit int) (ZShieldCoinbase, error) {
	query := APIQuery {
		Method:	`z_shieldcoinbase`,
		Params:	`["`+frmaddr+`", "`+toaddr+`", `+strconv.FormatFloat(fee, 'f', 8, 64)+`, `+strconv.Itoa(limit)+`]`,
	}
	//fmt.Println(query)

	var z_shieldcoinbase ZShieldCoinbase

	z_shieldcoinbaseJson := appName.APICall(query)
	//fmt.Println(z_shieldcoinbaseJson)

	var result APIResult

	json.Unmarshal([]byte(z_shieldcoinbaseJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_shieldcoinbaseJson), &z_shieldcoinbase)
		return z_shieldcoinbase, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_shieldcoinbaseJson), &z_shieldcoinbase)
	return z_shieldcoinbase, nil
}