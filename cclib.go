// Copyright © 2018-2020 Satinderjit Singh.
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
	"encoding/json"
	"errors"
	"fmt"
)

// CCLibInfo type
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
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// CCLibInfo method displays all the methods of all the modules that are available in the current library.
// The library is loaded at runtime using the -ac_cclib parameter.
func (appName AppType) CCLibInfo() (CCLibInfo, error) {
	query := APIQuery{
		Method: `cclibinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var cclibinfo CCLibInfo

	cclibinfoJSON := appName.APICall(&query)
	if cclibinfoJSON == "EMPTY RPC INFO" {
		return cclibinfo, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(cclibinfoJSON)

	var result APIResult

	json.Unmarshal([]byte(cclibinfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(cclibinfoJSON), &cclibinfo)
		return cclibinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(cclibinfoJSON), &cclibinfo)
	return cclibinfo, nil
}

// CCLibAddress type
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
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// CCLibAddress method returns information about the addresses related to the specified pubkey,
// and according to the Antara Module associated with the specified evalcode. If no pubkey is provided,
// the pubkey used to the launch the daemon is the default.
func (appName AppType) CCLibAddress(params APIParams) (CCLibAddress, error) {

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `cclibaddress`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var cclibaddr CCLibAddress

	cclibaddrJSON := appName.APICall(&query)
	if cclibaddrJSON == "EMPTY RPC INFO" {
		return cclibaddr, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(cclibaddrJSON)

	var result APIResult

	json.Unmarshal([]byte(cclibaddrJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(cclibaddrJSON), &cclibaddr)
		return cclibaddr, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(cclibaddrJSON), &cclibaddr)
	return cclibaddr, nil
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
/*func (appName AppType) CCLib(params APIParams) (interface{}, error) {

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery {
		Method:	`cclib`,
		Params:	string(paramsJSON),
	}
	//fmt.Println(query)

	//var cclb CCLib

	cclbJson := appName.APICall(&query)
	//fmt.Println(cclbJson)

	var result APIResult

	json.Unmarshal([]byte(cclbJson), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(cclbJson), &cclb)
		return cclb, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(cclbJson), &cclb)
	return cclb, nil
}*/
