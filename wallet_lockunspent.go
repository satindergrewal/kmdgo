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

type LockUnspent struct {
	Result bool `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) LockUnspent(unlock bool, transactions string) (LockUnspent, error) {
	query := APIQuery {
		Method:	`lockunspent`,
		Params:	`[`+strconv.FormatBool(unlock)+`, `+transactions+`]`,
	}
	//fmt.Println(query)

	var lockunspent LockUnspent

	lockunspentJson := appName.APICall(query)
	//fmt.Println(lockunspentJson)

	var result APIResult

	json.Unmarshal([]byte(lockunspentJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(lockunspentJson), &lockunspent)
		return lockunspent, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(lockunspentJson), &lockunspent)
	return lockunspent, nil
}