 //******************************************************************************
 //  Copyright © 2018-2019 Satinderjit Singh.                                   *
 //                                                                             *
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 //  the top-level directory of this distribution for the individual copyright  *
 //  holder information and the developer policies on copyright and licensing.  *
 //                                                                             *
 //  Unless otherwise agreed in a custom licensing agreement, no part of the    *
 //  kmdgo software, including this file may be copied, modified, propagated.   *
 //  or distributed except according to the terms contained in the LICENSE file *
 //                                                                             *
 //  Removal or modification of this copyright notice is prohibited.            *
 //                                                                             *
 //******************************************************************************/
package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type FaucetFund struct {
	Result struct {
		Result			string `json:"result"` // This should have been a `int` value or `float64` value. But since the API only accepts string, for now this is input as string. It must be updated/changed as the reported changes are refelected via API.
		Hex				string `json:"hex"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) FaucetFund(params APIParams) (FaucetFund, error) {
	
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`faucetfund`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var faucetfund FaucetFund

	faucetfundJson := appName.APICall(query)
	//fmt.Println(faucetfundJson)

	var result APIResult

	json.Unmarshal([]byte(faucetfundJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetfundJson), &faucetfund)
		return faucetfund, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(faucetfundJson), &faucetfund)
	return faucetfund, nil
}