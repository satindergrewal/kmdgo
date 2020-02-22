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
	Error Error  `json:"error"`
	ID    string `json:"id"`
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

	query := APIQuery{
		Method: `faucetaddress`,
		Params: params_json,
	}
	//fmt.Println(query)

	var faucetaddress FaucetAddress

	faucetaddressJson := appName.APICall(query)
	if faucetaddressJson == "EMPTY RPC INFO!" {
		return faucetaddress, errors.New("EMPTY RPC INFO!")
	}
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

// This should have been a `int` value or `float64` value.
// But since the API only accepts string, for now this is input as string.
type FaucetFund struct {
	Result struct {
		Result string `json:"result"`
		Hex    string `json:"hex"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) FaucetFund(params APIParams) (FaucetFund, error) {

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `faucetfund`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var faucetfund FaucetFund

	faucetfundJson := appName.APICall(query)
	if faucetfundJson == "EMPTY RPC INFO!" {
		return faucetfund, errors.New("EMPTY RPC INFO!")
	}
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

type FaucetGet struct {
	Result struct {
		Result string `json:"result"`
		Hex    string `json:"hex"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) FaucetGet() (FaucetGet, error) {
	query := APIQuery{
		Method: `faucetget`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var faucetget FaucetGet

	faucetgetJson := appName.APICall(query)
	if faucetgetJson == "EMPTY RPC INFO!" {
		return faucetget, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(faucetgetJson)

	var result APIResult

	json.Unmarshal([]byte(faucetgetJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetgetJson), &faucetget)
		return faucetget, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(faucetgetJson), &faucetget)
	return faucetget, nil
}

// This should have been a `int` value or `float64` value.
// But since the API only accepts string, for now this is input as string.
type FaucetInfo struct {
	Result struct {
		Result  string `json:"result"`
		Name    string `json:"name"`
		Funding string `json:"funding"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) FaucetInfo() (FaucetInfo, error) {
	query := APIQuery{
		Method: `faucetinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var faucetinfo FaucetInfo

	faucetinfoJson := appName.APICall(query)
	if faucetinfoJson == "EMPTY RPC INFO!" {
		return faucetinfo, errors.New("EMPTY RPC INFO!")
	}
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
