/******************************************************************************
 * Copyright © 2018-2019 Satinderjit Singh.                                   *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * kmdgo software, including this file may be copied, modified, propagated.   *
 * or distributed except according to the terms contained in the LICENSE file *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/
package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

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
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) DecodeRawTransaction(params APIParams) (DecodeRawTransaction, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`decoderawtransaction`,
		Params:	string(params_json),
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