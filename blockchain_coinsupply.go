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
	"errors"
	"encoding/json"
)

type CoinSupply struct {
	Result struct {
		Result	string	`json:"result"`
		Coin	string 	`json:"coin"`
		Height	int		`json:"height"`
		Supply	float64	`json:"supply"`
		Zfunds	float64	`json:"zfunds"`
		Sprout	float64	`json:"sprout"`
		Total	float64	`json:"total"`
	}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) CoinSupply(params APIParams) (CoinSupply, error) {
    if params[0] == nil {
        params[0] = 100
    }
    params_json, _ := json.Marshal(params)
    //fmt.Println(string(params_json))

    query := APIQuery {
        Method:     `coinsupply`,
        Params:     string(params_json),
    }
    //fmt.Println(query)

    var coinsupply CoinSupply

    coinsupplyJson := appName.APICall(query)
    
    
    var result APIResult

    json.Unmarshal([]byte(coinsupplyJson), &result)
    
    if result.Result == nil {
        answer_error, err := json.Marshal(result.Error)
        if err != nil {
        }
        json.Unmarshal([]byte(coinsupplyJson), &coinsupply)
        return coinsupply, errors.New(string(answer_error))
    }

    json.Unmarshal([]byte(coinsupplyJson), &coinsupply)
    return coinsupply, nil
}