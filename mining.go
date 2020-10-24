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

type GetBlockSubsidy struct {
	Result struct {
		Miner float64 `json:"miner"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetBlockSubsidy(params APIParams) (GetBlockSubsidy, error) {
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
		Method: `getblocksubsidy`,
		Params: paramsJSON,
	}
	//fmt.Println(query)

	var getblocksubsidy GetBlockSubsidy

	getblocksubsidyJSON := appName.APICall(&query)
	if getblocksubsidyJSON == "EMPTY RPC INFO" {
		return getblocksubsidy, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getblocksubsidyJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblocksubsidyJSON), &getblocksubsidy)
		return getblocksubsidy, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getblocksubsidyJSON), &getblocksubsidy)
	return getblocksubsidy, nil
}

type GetBlockTemplate struct {
	Result struct {
		Capabilities      []string `json:"capabilities"`
		Version           int      `json:"version"`
		Previousblockhash string   `json:"previousblockhash"`
		Transactions      []struct {
			Data    string        `json:"data"`
			Hash    string        `json:"hash"`
			Depends []interface{} `json:"depends"`
			Fee     int           `json:"fee"`
			Sigops  int           `json:"sigops"`
		} `json:"transactions"`
		Coinbasetxn struct {
			Data          string        `json:"data"`
			Hash          string        `json:"hash"`
			Depends       []interface{} `json:"depends"`
			Fee           int           `json:"fee"`
			Sigops        int           `json:"sigops"`
			Coinbasevalue int           `json:"coinbasevalue"`
			Required      bool          `json:"required"`
		} `json:"coinbasetxn"`
		Longpollid string   `json:"longpollid"`
		Target     string   `json:"target"`
		Mintime    int      `json:"mintime"`
		Mutable    []string `json:"mutable"`
		Noncerange string   `json:"noncerange"`
		Sigoplimit int      `json:"sigoplimit"`
		Sizelimit  int      `json:"sizelimit"`
		Curtime    int      `json:"curtime"`
		Bits       string   `json:"bits"`
		Height     int      `json:"height"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetBlockTemplate(params APIParams) (GetBlockTemplate, error) {
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
		Method: `getblocktemplate`,
		Params: paramsJSON,
	}
	//fmt.Println(query)

	var getblocktemplate GetBlockTemplate

	getblocktemplateJSON := appName.APICall(&query)
	if getblocktemplateJSON == "EMPTY RPC INFO" {
		return getblocktemplate, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getblocktemplateJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblocktemplateJSON), &getblocktemplate)
		return getblocktemplate, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getblocktemplateJSON), &getblocktemplate)
	return getblocktemplate, nil
}

type GetLocalSolps struct {
	Result int    `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetLocalSolps() (GetLocalSolps, error) {
	query := APIQuery{
		Method: `getlocalsolps`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getlocalsolps GetLocalSolps

	getlocalsolpsJSON := appName.APICall(&query)
	if getlocalsolpsJSON == "EMPTY RPC INFO" {
		return getlocalsolps, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getlocalsolpsJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getlocalsolpsJSON), &getlocalsolps)
		return getlocalsolps, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getlocalsolpsJSON), &getlocalsolps)
	return getlocalsolps, nil
}

type GetMiningInfo struct {
	Result struct {
		Blocks           int     `json:"blocks"`
		Currentblocksize int     `json:"currentblocksize"`
		Currentblocktx   int     `json:"currentblocktx"`
		Difficulty       float64 `json:"difficulty"`
		Errors           string  `json:"errors"`
		Genproclimit     int     `json:"genproclimit"`
		Localsolps       int     `json:"localsolps"`
		Networksolps     int     `json:"networksolps"`
		Networkhashps    int     `json:"networkhashps"`
		Pooledtx         int     `json:"pooledtx"`
		Testnet          bool    `json:"testnet"`
		Chain            string  `json:"chain"`
		Generate         bool    `json:"generate"`
		Numthreads       int     `json:"numthreads"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetMiningInfo() (GetMiningInfo, error) {
	query := APIQuery{
		Method: `getmininginfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getmininginfo GetMiningInfo

	getmininginfoJSON := appName.APICall(&query)
	if getmininginfoJSON == "EMPTY RPC INFO" {
		return getmininginfo, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getmininginfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getmininginfoJSON), &getmininginfo)
		return getmininginfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getmininginfoJSON), &getmininginfo)
	return getmininginfo, nil
}

type GetNetworkHashps struct {
	Result int    `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetNetworkHashps(params APIParams) (GetNetworkHashps, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = 120
		}
	}

	if len(params) == 2 {
		if params[1] == nil {
			params[1] = -1
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getnetworkhashps`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getnetworkhashps GetNetworkHashps

	getnetworkhashpsJSON := appName.APICall(&query)
	if getnetworkhashpsJSON == "EMPTY RPC INFO" {
		return getnetworkhashps, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getnetworkhashpsJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnetworkhashpsJSON), &getnetworkhashps)
		return getnetworkhashps, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getnetworkhashpsJSON), &getnetworkhashps)
	return getnetworkhashps, nil
}

type GetNetworkSolps struct {
	Result int    `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetNetworkSolps(params APIParams) (GetNetworkSolps, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = 120
		}
	}

	if len(params) == 2 {
		if params[1] == nil {
			params[1] = -1
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getnetworksolps`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getnetworksolps GetNetworkSolps

	getnetworksolpsJSON := appName.APICall(&query)
	if getnetworksolpsJSON == "EMPTY RPC INFO" {
		return getnetworksolps, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getnetworksolpsJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnetworksolpsJSON), &getnetworksolps)
		return getnetworksolps, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getnetworksolpsJSON), &getnetworksolps)
	return getnetworksolps, nil
}

type PrioritiseTransaction struct {
	Result bool   `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) PrioritiseTransaction(params APIParams) (PrioritiseTransaction, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `prioritisetransaction`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var prioritisetransaction PrioritiseTransaction

	prioritisetransactionJSON := appName.APICall(&query)
	if prioritisetransactionJSON == "EMPTY RPC INFO" {
		return prioritisetransaction, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(prioritisetransactionJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(prioritisetransactionJSON), &prioritisetransaction)
		return prioritisetransaction, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(prioritisetransactionJSON), &prioritisetransaction)
	return prioritisetransaction, nil
}

type SubmitBlock struct {
	Result bool   `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SubmitBlock(params APIParams) (SubmitBlock, error) {

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `submitblock`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var submitblock SubmitBlock

	submitblockJSON := appName.APICall(&query)
	if submitblockJSON == "EMPTY RPC INFO" {
		return submitblock, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(submitblockJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(submitblockJSON), &submitblock)
		return submitblock, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(submitblockJSON), &submitblock)
	return submitblock, nil
}
