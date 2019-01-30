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

type GetRawTransaction struct {
	Result string `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

type GetRawTransactionDetailed struct {
	Result GetRawTransactionFull `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

type GetRawTransactionFull struct {
	Hex          string `json:"hex"`
	Txid         string `json:"txid"`
	Overwintered bool   `json:"overwintered"`
	Version      int    `json:"version"`
	Locktime     int    `json:"locktime"`
	Vin          []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		Address   string `json:"address"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence int64 `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		Interest     float64 `json:"interest"`
		ValueSat     int64   `json:"valueSat"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
			Addresses []string `json:"addresses"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
	Vjoinsplit       []interface{} `json:"vjoinsplit"`
	Blockhash        string        `json:"blockhash"`
	Height           int           `json:"height"`
	Confirmations    int           `json:"confirmations"`
	Rawconfirmations int           `json:"rawconfirmations"`
	Time             int           `json:"time"`
	Blocktime        int           `json:"blocktime"`
}

func (appName AppType) GetRawTransaction(params APIParams) (GetRawTransaction, error) {
	
	if len(params) == 2 {
		if params[1] == nil {
			params[1] = 0
		}
	}

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`getrawtransaction`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var getrawtransaction GetRawTransaction

	getrawtransactionJson := appName.APICall(query)
	//fmt.Println(getrawtransactionJson)

	var result APIResult

	json.Unmarshal([]byte(getrawtransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawtransactionJson), &getrawtransaction)
		return getrawtransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getrawtransactionJson), &getrawtransaction)
	return getrawtransaction, nil
}

func (appName AppType) GetRawTransactionDetailed(params APIParams) (GetRawTransactionDetailed, error) {
	
	if len(params) == 2 {
		if params[1] == nil {
			params[1] = 0
		}
	}

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`getrawtransaction`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var getrawtransaction GetRawTransactionDetailed

	getrawtransactionJson := appName.APICall(query)
	//fmt.Println(getrawtransactionJson)

	var result APIResult

	json.Unmarshal([]byte(getrawtransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawtransactionJson), &getrawtransaction)
		return getrawtransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getrawtransactionJson), &getrawtransaction)
	return getrawtransaction, nil
}