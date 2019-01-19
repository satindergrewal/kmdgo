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

type GetAddedNodeInfo struct {
	Result []struct {
		Addednode string `json:"addednode"`
		Connected bool   `json:"connected"`
		Addresses []struct {
			Address   string `json:"address"`
			Connected string `json:"connected"`
		} `json:"addresses"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetAddedNodeInfo(dns bool, node string) (GetAddedNodeInfo, error) {
	query := APIQuery {
		Method:	`getaddednodeinfo`,
		Params:	`[`+strconv.FormatBool(dns)+`, "`+node+`"]`,
	}
	//fmt.Println(query)

	var getaddednodeinfo GetAddedNodeInfo

	getaddednodeinfoJson := appName.APICall(query)
	//fmt.Println(getaddednodeinfoJson)

	var result APIResult

	json.Unmarshal([]byte(getaddednodeinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getaddednodeinfoJson), &getaddednodeinfo)
		return getaddednodeinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getaddednodeinfoJson), &getaddednodeinfo)
	return getaddednodeinfo, nil
}