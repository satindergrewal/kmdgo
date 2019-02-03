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

type GetBlockSubsidy struct {
	Result	struct {
		Miner float64 `json:"miner"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetBlockSubsidy(params APIParams) (GetBlockSubsidy, error) {
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
		Method:	`getblocksubsidy`,
		Params:	params_json,
	}
	//fmt.Println(query)

	var getblocksubsidy GetBlockSubsidy

	getblocksubsidyJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getblocksubsidyJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblocksubsidyJson), &getblocksubsidy)
		return getblocksubsidy, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblocksubsidyJson), &getblocksubsidy)
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
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) GetBlockTemplate(params APIParams) (GetBlockTemplate, error) {
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
		Method:	`getblocktemplate`,
		Params:	params_json,
	}
	//fmt.Println(query)

	var getblocktemplate GetBlockTemplate

	getblocktemplateJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getblocktemplateJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblocktemplateJson), &getblocktemplate)
		return getblocktemplate, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblocktemplateJson), &getblocktemplate)
	return getblocktemplate, nil
}




type GetLocalSolps struct {
	Result int `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) GetLocalSolps() (GetLocalSolps, error) {
	query := APIQuery {
		Method:	`getlocalsolps`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getlocalsolps GetLocalSolps

	getlocalsolpsJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getlocalsolpsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getlocalsolpsJson), &getlocalsolps)
		return getlocalsolps, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getlocalsolpsJson), &getlocalsolps)
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
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) GetMiningInfo() (GetMiningInfo, error) {
	query := APIQuery {
		Method:	`getmininginfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getmininginfo GetMiningInfo

	getmininginfoJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getmininginfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getmininginfoJson), &getmininginfo)
		return getmininginfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getmininginfoJson), &getmininginfo)
	return getmininginfo, nil
}



type GetNetworkHashps struct {
	Result int `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
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
	
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`getnetworkhashps`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var getnetworkhashps GetNetworkHashps

	getnetworkhashpsJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getnetworkhashpsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnetworkhashpsJson), &getnetworkhashps)
		return getnetworkhashps, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getnetworkhashpsJson), &getnetworkhashps)
	return getnetworkhashps, nil
}



type GetNetworkSolps struct {
	Result int `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
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
	
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`getnetworksolps`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var getnetworksolps GetNetworkSolps

	getnetworksolpsJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getnetworksolpsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnetworksolpsJson), &getnetworksolps)
		return getnetworksolps, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getnetworksolpsJson), &getnetworksolps)
	return getnetworksolps, nil
}



type PrioritiseTransaction struct {
	Result bool `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) PrioritiseTransaction(params APIParams) (PrioritiseTransaction, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`prioritisetransaction`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var prioritisetransaction PrioritiseTransaction

	prioritisetransactionJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(prioritisetransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(prioritisetransactionJson), &prioritisetransaction)
		return prioritisetransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(prioritisetransactionJson), &prioritisetransaction)
	return prioritisetransaction, nil
}



type SubmitBlock struct {
	Result bool `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) SubmitBlock(params APIParams) (SubmitBlock, error) {
	
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`submitblock`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var submitblock SubmitBlock

	submitblockJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(submitblockJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(submitblockJson), &submitblock)
		return submitblock, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(submitblockJson), &submitblock)
	return submitblock, nil
}


