 //  Copyright © 2018-2019 Satinderjit Singh.
 //
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
 //  the top-level directory of this distribution for the individual copyright
 //  holder information and the developer policies on copyright and licensing.
 //
 //  Unless otherwise agreed in a custom licensing agreement, no part of the
 //  kmdgo software, including this file may be copied, modified, propagated.
 //  or distributed except according to the terms contained in the LICENSE file
 //
 //  Removal or modification of this copyright notice is prohibited.
package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

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
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) SignRawTransaction(params APIParams) (SignRawTransaction, error) {

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`signrawtransaction`,
		Params:	string(params_json),
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