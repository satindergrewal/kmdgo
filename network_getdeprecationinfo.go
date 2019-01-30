// Copyright © 2018-2019 Satinderjit Singh.
//
// See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
// the top-level directory of this distribution for the individual copyright
// holder information and the developer policies on copyright and licensing.
//
// Unless otherwise agreed in a custom licensing agreement, no part of the
// kmdgo software, including this file may be copied, modified, propagated.
// or distributed except according to the terms contained in the LICENSE file
//
// Removal or modification of this copyright notice is prohibited.



package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type GetDeprecationInfo struct {
	Result struct {
		Version           int    `json:"version"`
		Subversion        string `json:"subversion"`
		Deprecationheight int    `json:"deprecationheight"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetDeprecationInfo() (GetDeprecationInfo, error) {
	query := APIQuery {
		Method:	`getdeprecationinfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getdeprecationinfo GetDeprecationInfo

	getdeprecationinfoJson := appName.APICall(query)
	//fmt.Println(getdeprecationinfoJson)

	var result APIResult

	json.Unmarshal([]byte(getdeprecationinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getdeprecationinfoJson), &getdeprecationinfo)
		return getdeprecationinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getdeprecationinfoJson), &getdeprecationinfo)
	return getdeprecationinfo, nil
}