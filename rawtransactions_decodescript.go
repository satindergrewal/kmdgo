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
	"fmt"
	"encoding/json"
	"errors"
)

type DecodeScript struct {
	Result struct {
		Asm  string `json:"asm"`
		Type string `json:"type"`
		P2Sh string `json:"p2sh"`
	} `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) DecodeScript(params APIParams) (DecodeScript, error) {
	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`decodescript`,
		Params:	string(params_json),
	}
	fmt.Println(query)

	var decodescript DecodeScript

	decodescriptJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(decodescriptJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(decodescriptJson), &decodescript)
		return decodescript, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(decodescriptJson), &decodescript)
	return decodescript, nil
}