 //  Copyright © 2018-2019 Satinderjit Singh.
 //
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
 //  the top-level directory of this distribution for the individual copyright
 //  holder information and the developer policies on copyright and licensing.
 //
 //  Unless otherwise agreed in a custom licensing agreement, no part of the
 //  kmdgo software, including this file may be copied, modified, propagated.
 //  or distributed except according to the terms contained in the LICENSE file
 //
 //  Removal or modification of this copyright notice is prohibited.
package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type GetChainTips struct {
	Result []struct {
		Height    int    `json:"height"`
		Hash      string `json:"hash"`
		Branchlen int    `json:"branchlen"`
		Status    string `json:"status"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetChainTips() (GetChainTips, error) {
	query := APIQuery {
		Method:	`getchaintips`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getchaintips GetChainTips

	getchaintipsJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getchaintipsJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getchaintipsJson), &getchaintips)
		return getchaintips, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getchaintipsJson), &getchaintips)
	return getchaintips, nil
}