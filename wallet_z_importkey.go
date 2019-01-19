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

type ZImportKey struct {
	Result interface{} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZImportKey(zkey string, rescan string, startHeight int) (ZImportKey, error) {
	query := APIQuery {
		Method:	`z_importkey`,
		Params:	`["`+zkey+`", "`+rescan+`", `+strconv.Itoa(startHeight)+`]`,
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