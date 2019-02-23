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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `createrawtransaction`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var createrawtransaction CreateRawTransaction

	createrawtransactionJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(createrawtransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(createrawtransactionJson), &createrawtransaction)
		return createrawtransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(createrawtransactionJson), &createrawtransaction)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `decoderawtransaction`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var decoderawtransaction DecodeRawTransaction

	decoderawtransactionJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(decoderawtransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(decoderawtransactionJson), &decoderawtransaction)
		return decoderawtransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(decoderawtransactionJson), &decoderawtransaction)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `decodescript`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var decodescript DecodeScript

	decodescriptJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(decodescriptJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(decodescriptJson), &decodescript)
		return decodescript, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(decodescriptJson), &decodescript)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `fundrawtransaction`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var fundrawtransaction FundRawTransaction

	fundrawtransactionJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(fundrawtransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(fundrawtransactionJson), &fundrawtransaction)
		return fundrawtransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(fundrawtransactionJson), &fundrawtransaction)
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `getrawtransaction`,
		Params: string(params_json),
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

	query := APIQuery{
		Method: `getrawtransaction`,
		Params: string(params_json),
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `sendrawtransaction`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var sendrawtransaction SendRawTransaction

	sendrawtransactionJson := appName.APICall(query)
	//fmt.Println(sendrawtransactionJson)

	var result APIResult

	json.Unmarshal([]byte(sendrawtransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendrawtransactionJson), &sendrawtransaction)
		return sendrawtransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(sendrawtransactionJson), &sendrawtransaction)
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `signrawtransaction`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var signrawtransaction SignRawTransaction

	signrawtransactionJson := appName.APICall(query)
	//fmt.Println(signrawtransactionJson)

	var result APIResult

	json.Unmarshal([]byte(signrawtransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(signrawtransactionJson), &signrawtransaction)
		return signrawtransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(signrawtransactionJson), &signrawtransaction)
	return signrawtransaction, nil
}
