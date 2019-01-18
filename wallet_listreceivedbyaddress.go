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

type ListReceivedByAddress struct {
	Result []struct {
		Address          string        `json:"address"`
		Account          string        `json:"account"`
		Amount           float64       `json:"amount"`
		Rawconfirmations int           `json:"rawconfirmations"`
		Confirmations    int           `json:"confirmations"`
		Txids            []interface{} `json:"txids"`
		} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ListReceivedByAddress(minconf int, incempty bool, incwch bool) (ListReceivedByAddress, error) {
	query := APIQuery {
		Method:	`listreceivedbyaddress`,
		Params:	`[`+strconv.Itoa(minconf)+`, `+strconv.FormatBool(incempty)+`, `+strconv.FormatBool(incwch)+`]`,
	}
	//fmt.Println(query)

	var listreceivedbyaddress ListReceivedByAddress

	listreceivedbyaddressJson := appName.APICall(query)
	//fmt.Println(listreceivedbyaddressJson)

	var result APIResult

	json.Unmarshal([]byte(listreceivedbyaddressJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listreceivedbyaddressJson), &listreceivedbyaddress)
		return listreceivedbyaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listreceivedbyaddressJson), &listreceivedbyaddress)
	return listreceivedbyaddress, nil
}