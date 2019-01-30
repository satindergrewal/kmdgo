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

type ListTransactions struct {
	Result	[]struct {
		InvolvesWatchonly bool          `json:"involvesWatchonly"`
		Account           string        `json:"account"`
		Address           string        `json:"address"`
		Category          string        `json:"category"`
		Amount            float64       `json:"amount"`
		Vout              int           `json:"vout"`
		Fee               float64       `json:"fee,omitempty"`
		Rawconfirmations  int           `json:"rawconfirmations"`
		Confirmations     int           `json:"confirmations"`
		Blockhash         string        `json:"blockhash"`
		Blockindex        int           `json:"blockindex"`
		Blocktime         int           `json:"blocktime"`
		Expiryheight      int           `json:"expiryheight"`
		Txid              string        `json:"txid"`
		Walletconflicts   []interface{} `json:"walletconflicts"`
		Time              int           `json:"time"`
		Timereceived      int           `json:"timereceived"`
		Vjoinsplit        []interface{} `json:"vjoinsplit"`
		Size              int           `json:"size"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) ListTransactions(params APIParams) (ListTransactions, error) {
	if params[0] == nil {
		params[0] = "*"
	}
	if params[1] == nil {
		params[1] = 10
	}

	if params[2] == nil {
		params[2] = 0
	}
	if params[3] == nil {
		params[3] = false
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	`listtransactions`,
		Params:	string(params_json),
	}
	//fmt.Println(query)

	var listtransactions ListTransactions

	listtransactionsJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(listtransactionsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listtransactionsJson), &listtransactions)
		return listtransactions, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listtransactionsJson), &listtransactions)
	return listtransactions, nil
}