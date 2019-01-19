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

type ListBanned struct {
	Result []struct {
		Address     string `json:"address"`
		BannedUntil int    `json:"banned_until"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ListBanned() (ListBanned, error) {
	query := APIQuery {
		Method:	`listbanned`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var listbanned ListBanned

	listbannedJson := appName.APICall(query)
	//fmt.Println(listbannedJson)

	var result APIResult

	json.Unmarshal([]byte(listbannedJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listbannedJson), &listbanned)
		return listbanned, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listbannedJson), &listbanned)
	return listbanned, nil
}