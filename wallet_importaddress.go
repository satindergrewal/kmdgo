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
	"strconv"
)

type ImportAddress struct {
	Result	string	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) ImportAddress(taddr string, lbl string, rscn bool) (ImportAddress, error) {
	query := APIQuery {
		Method:	`importaddress`,
		Params:	`["`+taddr+`", "`+lbl+`", `+strconv.FormatBool(rscn)+`]`,
	}
	//fmt.Println(query)

	var importaddress ImportAddress

	importaddressJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(importaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(importaddressJson), &importaddress)
		return importaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(importaddressJson), &importaddress)
	return importaddress, nil
}