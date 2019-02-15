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

	var cclibinfo CCLibAddress

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