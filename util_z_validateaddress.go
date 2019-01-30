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
	//"strconv"
)

type ZValidateAddress struct {
	Result	struct {
		Isvalid                    bool   `json:"isvalid"`
		Address                    string `json:"address"`
		Type                       string `json:"type"`
		Diversifier                string `json:"diversifier"`
		Diversifiedtransmissionkey string `json:"diversifiedtransmissionkey"`
		Ismine                     bool   `json:"ismine"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) ZValidateAddress(zaddr string) (ZValidateAddress, error) {
	query := APIQuery {
		Method:	`z_validateaddress`,
		Params:	`["`+zaddr+`"]`,
	}
	//fmt.Println(query)

	var zvalidateaddress ZValidateAddress

	zvalidateaddressJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(zvalidateaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zvalidateaddressJson), &zvalidateaddress)
		return zvalidateaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(zvalidateaddressJson), &zvalidateaddress)
	return zvalidateaddress, nil
}