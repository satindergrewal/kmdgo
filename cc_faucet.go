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
	//"fmt"
	"encoding/json"
	"errors"
)

// FaucetAddress type
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

// FaucetAddress method returns the Antara address information for the specified pubkey.
// If no pubkey is provided, the method returns information for the pubkey used to launch the daemon
func (appName AppType) FaucetAddress(params APIParams) (FaucetAddress, error) {
	var paramsJSON string

	if params[0] == "" || params[0] == nil {
		paramsJSON = `[]`
		//fmt.Println(paramsJSON)
	} else {
		paramsBytes, _ := json.Marshal(params)
		paramsJSON = string(paramsBytes)
		//fmt.Println(paramsJSON)
	}

	query := APIQuery{
		Method: `faucetaddress`,
		Params: paramsJSON,
	}
	//fmt.Println(query)

	var faucetaddress FaucetAddress

	faucetaddressJSON := appName.APICall(&query)
	if faucetaddressJSON == "EMPTY RPC INFO" {
		return faucetaddress, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(faucetaddressJSON)

	var result APIResult

	json.Unmarshal([]byte(faucetaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetaddressJSON), &faucetaddress)
		return faucetaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(faucetaddressJSON), &faucetaddress)
	return faucetaddress, nil
}

// FaucetFund type
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

// FaucetFund method funds the on-chain faucet.
// The method returns a hex value which must then be broadcast using the sendrawtransaction method
func (appName AppType) FaucetFund(params APIParams) (FaucetFund, error) {

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `faucetfund`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var faucetfund FaucetFund

	faucetfundJSON := appName.APICall(&query)
	if faucetfundJSON == "EMPTY RPC INFO" {
		return faucetfund, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(faucetfundJSON)

	var result APIResult

	json.Unmarshal([]byte(faucetfundJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetfundJSON), &faucetfund)
		return faucetfund, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(faucetfundJSON), &faucetfund)
	return faucetfund, nil
}

// FaucetGet type
type FaucetGet struct {
	Result struct {
		Result string `json:"result"`
		Hex    string `json:"hex"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// FaucetGet method requests the faucet module to send coins.
// The method returns a hex value which must then be broadcast using the sendrawtransaction method.
// The faucetget command yields 0.1 coins and requires about 30 seconds of CPU time to execute
func (appName AppType) FaucetGet() (FaucetGet, error) {
	query := APIQuery{
		Method: `faucetget`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var faucetget FaucetGet

	faucetgetJSON := appName.APICall(&query)
	if faucetgetJSON == "EMPTY RPC INFO" {
		return faucetget, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(faucetgetJSON)

	var result APIResult

	json.Unmarshal([]byte(faucetgetJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetgetJSON), &faucetget)
		return faucetget, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(faucetgetJSON), &faucetget)
	return faucetget, nil
}

// FaucetInfo type
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

// FaucetInfo method displays the balance of funds in the chain's faucet.
func (appName AppType) FaucetInfo() (FaucetInfo, error) {
	query := APIQuery{
		Method: `faucetinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var faucetinfo FaucetInfo

	faucetinfoJSON := appName.APICall(&query)
	if faucetinfoJSON == "EMPTY RPC INFO" {
		return faucetinfo, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(faucetinfoJSON)

	var result APIResult

	json.Unmarshal([]byte(faucetinfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(faucetinfoJSON), &faucetinfo)
		return faucetinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(faucetinfoJSON), &faucetinfo)
	return faucetinfo, nil
}
