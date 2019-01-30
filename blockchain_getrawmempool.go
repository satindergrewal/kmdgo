 //******************************************************************************
 //  Copyright © 2018-2019 Satinderjit Singh.                                   *
 //                                                                             *
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 //  the top-level directory of this distribution for the individual copyright  *
 //  holder information and the developer policies on copyright and licensing.  *
 //                                                                             *
 //  Unless otherwise agreed in a custom licensing agreement, no part of the    *
 //  kmdgo software, including this file may be copied, modified, propagated.   *
 //  or distributed except according to the terms contained in the LICENSE file *
 //                                                                             *
 //  Removal or modification of this copyright notice is prohibited.            *
 //                                                                             *
 //******************************************************************************/
package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
	"strconv"
)


type GetRawMempoolTrue struct {
	Result map[string]RawMempoolTrue	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

type RawMempoolTrue struct {
	Size             int           `json:"size"`
	Fee              float64       `json:"fee"`
	Time             int           `json:"time"`
	Height           int           `json:"height"`
	Startingpriority float64       `json:"startingpriority"`
	Currentpriority  float64       `json:"currentpriority"`
	Depends          []interface{} `json:"depends"`
}

type GetRawMempoolFalse struct {
	Result []string	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetRawMempoolTrue(b bool) (GetRawMempoolTrue, error) {
	query := APIQuery {
		Method:	`getrawmempool`,
		Params:	`[`+strconv.FormatBool(b)+`]`,
	}
	//fmt.Println(query)

	var getrawmempool GetRawMempoolTrue

	getrawmempoolJson := appName.APICall(query)
	//fmt.Println(getrawmempoolJson)

	var result APIResult

	json.Unmarshal([]byte(getrawmempoolJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
		return getrawmempool, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
	return getrawmempool, nil
}


func (appName AppType) GetRawMempoolFalse(b bool) (GetRawMempoolFalse, error) {
	query := APIQuery {
		Method:	`getrawmempool`,
		Params:	`[`+strconv.FormatBool(b)+`]`,
	}
	//fmt.Println(query)

	var getrawmempool GetRawMempoolFalse

	getrawmempoolJson := appName.APICall(query)
	//fmt.Println(getrawmempoolJson)

	var result APIResult

	json.Unmarshal([]byte(getrawmempoolJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
		return getrawmempool, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
	return getrawmempool, nil
}