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

type GetNetTotals struct {
	Result struct {
		Totalbytesrecv int   `json:"totalbytesrecv"`
		Totalbytessent int   `json:"totalbytessent"`
		Timemillis     int64 `json:"timemillis"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetNetTotals() (GetNetTotals, error) {
	query := APIQuery {
		Method:	`getnettotals`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getnettotals GetNetTotals

	getnettotalsJson := appName.APICall(query)
	//fmt.Println(getnettotalsJson)

	var result APIResult

	json.Unmarshal([]byte(getnettotalsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnettotalsJson), &getnettotals)
		return getnettotals, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getnettotalsJson), &getnettotals)
	return getnettotals, nil
}