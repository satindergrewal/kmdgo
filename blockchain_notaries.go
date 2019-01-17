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
/*
The notaries method returns the public key, BTC address, and KMD address for each Komodo notary node.
Either or both of the height and timestamp parameters will suffice.
*/

package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type Notaries struct {
	Result	struct {
		Notaries []struct {
			Pubkey     string `json:"pubkey"`
			BTCaddress string `json:"BTCaddress"`
			KMDaddress string `json:"KMDaddress"`
		} `json:"notaries"`
		Numnotaries int `json:"numnotaries"`
		Height      int `json:"height"`
		Timestamp   int `json:"timestamp"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) Notaries(ht string) (Notaries, error) {
	query := APIQuery {
		Method:	`notaries`,
		Params:	`["`+ht+`"]`,
	}
	//fmt.Println(query)

	var notaries Notaries

	notariesJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(notariesJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(notariesJson), &notaries)
		return notaries, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(notariesJson), &notaries)
	return notaries, nil
}