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
	//"strconv"
)

type ValidateAddress struct {
	Result	struct {
		Isvalid      bool   `json:"isvalid"`
		Address      string `json:"address"`
		ScriptPubKey string `json:"scriptPubKey"`
		Segid        int    `json:"segid"`
		Ismine       bool   `json:"ismine"`
		Iswatchonly  bool   `json:"iswatchonly"`
		Isscript     bool   `json:"isscript"`
		Pubkey       string `json:"pubkey"`
		Iscompressed bool   `json:"iscompressed"`
		Account      string `json:"account"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) ValidateAddress(taddr string) (ValidateAddress, error) {
	query := APIQuery {
		Method:	`validateaddress`,
		Params:	`["`+taddr+`"]`,
	}
	//fmt.Println(query)

	var validateaddress ValidateAddress

	validateaddressJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(validateaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(validateaddressJson), &validateaddress)
		return validateaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(validateaddressJson), &validateaddress)
	return validateaddress, nil
}