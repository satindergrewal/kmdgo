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
	"fmt"
	"encoding/json"
	"errors"
)


type CCLibInfo struct {
	Result struct {
		Result  string `json:"result"`
		CClib   string `json:"CClib"`
		Methods []struct {
			Evalcode       int    `json:"evalcode"`
			Funcid         string `json:"funcid"`
			Name           string `json:"name"`
			Method         string `json:"method"`
			Help           string `json:"help"`
			ParamsRequired int    `json:"params_required"`
			ParamsMax      int    `json:"params_max"`
		} `json:"methods"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) CCLibInfo() (CCLibInfo, error) {
	query := APIQuery {
		Method:	`cclibinfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var cclibinfo CCLibInfo

	cclibinfoJson := appName.APICall(query)
	//fmt.Println(cclibinfoJson)

	var result APIResult

	json.Unmarshal([]byte(cclibinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(cclibinfoJson), &cclibinfo)
		return cclibinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(cclibinfoJson), &cclibinfo)
	return cclibinfo, nil
}



type CCLibAddress struct {
	Result struct {
		Result               string  `json:"result"`
		CClibCCAddress       string  `json:"CClibCCAddress"`
		CCbalance            float64 `json:"CCbalance"`
		CClibNormalAddress   string  `json:"CClibNormalAddress"`
		CClibCCTokensAddress string  `json:"CClibCCTokensAddress"`
		MyAddress            string  `json:"myAddress"`
		MyCCAddressCClib     string  `json:"myCCAddress(CClib)"`
		MyCCaddress          string  `json:"myCCaddress"`
		MyCCbalance          float64 `json:"myCCbalance"`
		Myaddress            string  `json:"myaddress"`
		Mybalance            float64 `json:"mybalance"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) CCLibAddress(params APIParams) (CCLibAddress, error) {

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`cclibaddress`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var cclibaddr CCLibAddress

	cclibaddrJson := appName.APICall(query)
	//fmt.Println(cclibaddrJson)

	var result APIResult

	json.Unmarshal([]byte(cclibaddrJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(cclibaddrJson), &cclibaddr)
		return cclibaddr, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(cclibaddrJson), &cclibaddr)
	return cclibinfo, nil
}


// cclib method has these command line options:
// 		$ komodo-cli -ac_name=ROGUE help cclib 
// 		cclib method [evalcode] [JSON params]
// 
// Example command
//		komodo-cli -ac_name=ROGUE cclib newgame 17 "[1]"
//
// Explantion
//		Command line Utility: komodo-cli
//		Assetchain name parameter: -ac_name=ROGUE
//		Command Method: cclib
//		Command Sub-Method: newgame
//		evalcode used by the Crypto-Condition: 17
//		JSON Parameters passed for this crypto-conditions as string value: "[1]"
func (appName AppType) CCLib(params APIParams) (interface{}, error) {

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`cclib`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var cclb CCLib

	cclbJson := appName.APICall(query)
	//fmt.Println(cclbJson)

	var result APIResult

	json.Unmarshal([]byte(cclbJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(cclbJson), &cclb)
		return cclb, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(cclbJson), &cclb)
	return cclibinfo, nil
}