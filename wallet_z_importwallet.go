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
)

type ZImportWallet struct {
	Result string `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZImportWallet(wltfile string) (ZImportWallet, error) {
	query := APIQuery {
		Method:	`z_importwallet`,
		Params:	`["`+wltfile+`"]`,
	}
	//fmt.Println(query)

	var z_importwallet ZImportWallet

	z_importwalletJson := appName.APICall(query)
	//fmt.Println(z_importwalletJson)

	var result APIResult

	json.Unmarshal([]byte(z_importwalletJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_importwalletJson), &z_importwallet)
		return z_importwallet, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_importwalletJson), &z_importwallet)
	return z_importwallet, nil
}