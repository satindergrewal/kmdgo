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

type FaucetAddress struct {
	Result struct {
		Result          string `json:"result"`
		FaucetCCaddress string `json:"FaucetCCaddress"`
		Faucetmarker    string `json:"Faucetmarker"`
		GatewaysPubkey  string `json:"GatewaysPubkey"`
		FaucetCCassets  string `json:"FaucetCCassets"`
		MyCCaddress     string `json:"myCCaddress"`
		Myaddress       string `json:"myaddress"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) FaucetAddress(params APIParams) (FaucetAddress, error) {
	var params_json string

	if params[0] == "" || params[0] == nil {
		params_json = `[]`
		//fmt.Println(params_json)
	} else {
		params_bytes, _ := json.Marshal(params)
		params_json = string(params_bytes)
		//fmt.Println(params_json)
	}

	query := APIQuery {
		Method:	`faucetaddress`,
		Params:	params_json,
	}
	//fmt.Println(query)

	var faucetaddress FaucetAddress

	faucetaddressJson := appName.APICall(query)
	//fmt.Println(faucetaddressJson)

	var result APIResult

	json.Unmarshal([]byte(faucetaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetaddressJson), &faucetaddress)
		return faucetaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(faucetaddressJson), &faucetaddress)
	return faucetaddress, nil
}