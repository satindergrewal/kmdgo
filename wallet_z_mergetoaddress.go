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

type ZMergeToAddress struct {
	Result struct {
		RemainingUTXOs            int     `json:"remainingUTXOs"`
		RemainingTransparentValue float64 `json:"remainingTransparentValue"`
		RemainingNotes            int     `json:"remainingNotes"`
		RemainingShieldedValue    float64 `json:"remainingShieldedValue"`
		MergingUTXOs              int     `json:"mergingUTXOs"`
		MergingTransparentValue   float64 `json:"mergingTransparentValue"`
		MergingNotes              int     `json:"mergingNotes"`
		MergingShieldedValue      float64 `json:"mergingShieldedValue"`
		Opid                      string  `json:"opid"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZMergeToAddress(params APIParams) (ZMergeToAddress, error) {

	if len(params) >= 3 {
		if params[2] == nil {
			params[2] = 0.0001
		}
	}

	if len(params) >= 4 {
		if params[3] == nil {
			params[3] = 50
		}
	}

	if len(params) >= 5 {
		if params[4] == nil {
			params[4] = 10
		}
	}

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`z_mergetoaddress`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var z_mergetoaddress ZMergeToAddress

	z_mergetoaddressJson := appName.APICall(query)
	//fmt.Println(z_mergetoaddressJson)

	var result APIResult

	json.Unmarshal([]byte(z_mergetoaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_mergetoaddressJson), &z_mergetoaddress)
		return z_mergetoaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_mergetoaddressJson), &z_mergetoaddress)
	return z_mergetoaddress, nil
}