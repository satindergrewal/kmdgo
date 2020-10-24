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

type CreateRawTransaction struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) CreateRawTransaction(params APIParams) (CreateRawTransaction, error) {
	if len(params) >= 3 {
		if params[2] == nil {
			params[2] = 0
		}
	}

	if len(params) == 4 {
		if params[3] == nil {
			params[3] = 20
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `createrawtransaction`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var createrawtransaction CreateRawTransaction

	createrawtransactionJSON := appName.APICall(&query)
	if createrawtransactionJSON == "EMPTY RPC INFO" {
		return createrawtransaction, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(createrawtransactionJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(createrawtransactionJSON), &createrawtransaction)
		return createrawtransaction, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(createrawtransactionJSON), &createrawtransaction)
	return createrawtransaction, nil
}

type DecodeRawTransaction struct {
	Result struct {
		Txid     string `json:"txid"`
		Size     int    `json:"size"`
		Version  int    `json:"version"`
		Locktime int    `json:"locktime"`
		Vin      []struct {
			Txid      string `json:"txid"`
			Vout      int    `json:"vout"`
			ScriptSig struct {
				Asm string `json:"asm"`
				Hex string `json:"hex"`
			} `json:"scriptSig"`
			Sequence int64 `json:"sequence"`
		} `json:"vin"`
		Vout []struct {
			Value        float64 `json:"value"`
			ValueSat     int     `json:"valueSat"`
			N            int     `json:"n"`
			ScriptPubKey struct {
				Asm       string   `json:"asm"`
				Hex       string   `json:"hex"`
				ReqSigs   int      `json:"reqSigs"`
				Type      string   `json:"type"`
				Addresses []string `json:"addresses"`
			} `json:"scriptPubKey"`
		} `json:"vout"`
		Vjoinsplit []interface{} `json:"vjoinsplit"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) DecodeRawTransaction(params APIParams) (DecodeRawTransaction, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `decoderawtransaction`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var decoderawtransaction DecodeRawTransaction

	decoderawtransactionJSON := appName.APICall(&query)
	if decoderawtransactionJSON == "EMPTY RPC INFO" {
		return decoderawtransaction, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(decoderawtransactionJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(decoderawtransactionJSON), &decoderawtransaction)
		return decoderawtransaction, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(decoderawtransactionJSON), &decoderawtransaction)
	return decoderawtransaction, nil
}

type DecodeScript struct {
	Result struct {
		Asm  string `json:"asm"`
		Type string `json:"type"`
		P2Sh string `json:"p2sh"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) DecodeScript(params APIParams) (DecodeScript, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `decodescript`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var decodescript DecodeScript

	decodescriptJSON := appName.APICall(&query)
	if decodescriptJSON == "EMPTY RPC INFO" {
		return decodescript, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(decodescriptJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(decodescriptJSON), &decodescript)
		return decodescript, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(decodescriptJSON), &decodescript)
	return decodescript, nil
}

type FundRawTransaction struct {
	Result struct {
		Hex       string  `json:"hex"`
		Changepos int     `json:"changepos"`
		Fee       float64 `json:"fee"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) FundRawTransaction(params APIParams) (FundRawTransaction, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `fundrawtransaction`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var fundrawtransaction FundRawTransaction

	fundrawtransactionJSON := appName.APICall(&query)
	if fundrawtransactionJSON == "EMPTY RPC INFO" {
		return fundrawtransaction, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(fundrawtransactionJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(fundrawtransactionJSON), &fundrawtransaction)
		return fundrawtransaction, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(fundrawtransactionJSON), &fundrawtransaction)
	return fundrawtransaction, nil
}

type GetRawTransaction struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

type GetRawTransactionDetailed struct {
	Result GetRawTransactionFull `json:"result"`
	Error  Error                 `json:"error"`
	ID     string                `json:"id"`
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

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getrawtransaction`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getrawtransaction GetRawTransaction

	getrawtransactionJSON := appName.APICall(&query)
	if getrawtransactionJSON == "EMPTY RPC INFO" {
		return getrawtransaction, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getrawtransactionJSON)

	var result APIResult

	json.Unmarshal([]byte(getrawtransactionJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawtransactionJSON), &getrawtransaction)
		return getrawtransaction, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getrawtransactionJSON), &getrawtransaction)
	return getrawtransaction, nil
}

func (appName AppType) GetRawTransactionDetailed(params APIParams) (GetRawTransactionDetailed, error) {

	if len(params) == 2 {
		if params[1] == nil {
			params[1] = 0
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getrawtransaction`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getrawtransaction GetRawTransactionDetailed

	getrawtransactionJSON := appName.APICall(&query)
	if getrawtransactionJSON == "EMPTY RPC INFO" {
		return getrawtransaction, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getrawtransactionJSON)

	var result APIResult

	json.Unmarshal([]byte(getrawtransactionJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawtransactionJSON), &getrawtransaction)
		return getrawtransaction, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getrawtransactionJSON), &getrawtransaction)
	return getrawtransaction, nil
}

type SendRawTransaction struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SendRawTransaction(params APIParams) (SendRawTransaction, error) {

	if len(params) == 2 {
		if params[1] == nil {
			params[1] = false
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `sendrawtransaction`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var sendrawtransaction SendRawTransaction

	sendrawtransactionJSON := appName.APICall(&query)
	if sendrawtransactionJSON == "EMPTY RPC INFO" {
		return sendrawtransaction, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(sendrawtransactionJSON)

	var result APIResult

	json.Unmarshal([]byte(sendrawtransactionJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendrawtransactionJSON), &sendrawtransaction)
		return sendrawtransaction, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(sendrawtransactionJSON), &sendrawtransaction)
	return sendrawtransaction, nil
}

type SignRawTransaction struct {
	Result struct {
		Hex      string `json:"hex"`
		Complete bool   `json:"complete"`
		Errors   []struct {
			Txid      string `json:"txid"`
			Vout      int    `json:"vout"`
			ScriptSig string `json:"scriptSig"`
			Sequence  int64  `json:"sequence"`
			Error     string `json:"error"`
		} `json:"errors"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) SignRawTransaction(params APIParams) (SignRawTransaction, error) {

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `signrawtransaction`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var signrawtransaction SignRawTransaction

	signrawtransactionJSON := appName.APICall(&query)
	if signrawtransactionJSON == "EMPTY RPC INFO" {
		return signrawtransaction, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(signrawtransactionJSON)

	var result APIResult

	json.Unmarshal([]byte(signrawtransactionJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(signrawtransactionJSON), &signrawtransaction)
		return signrawtransaction, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(signrawtransactionJSON), &signrawtransaction)
	return signrawtransaction, nil
}
