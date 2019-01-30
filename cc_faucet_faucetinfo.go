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

type FaucetInfo struct {
	Result struct {
		Result  string `json:"result"`
		Name    string `json:"name"`
		Funding string `json:"funding"` // This should have been a `int` value or `float64` value. But since the API only accepts string, for now this is input as string. It must be updated/changed as the reported changes are refelected via API.
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) FaucetInfo() (FaucetInfo, error) {
	query := APIQuery {
		Method:	`faucetinfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var faucetinfo FaucetInfo

	faucetinfoJson := appName.APICall(query)
	//fmt.Println(faucetinfoJson)

	var result APIResult

	json.Unmarshal([]byte(faucetinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetinfoJson), &faucetinfo)
		return faucetinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(faucetinfoJson), &faucetinfo)
	return faucetinfo, nil
}