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
/*
The `minerids` method returns information about the notary nodes and 
external miners at a specific block height. The response will 
calculate results according to the 2000 blocks proceeding the 
indicated "height" block.
*/

package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type MinerIDs struct {
	Result	struct {
		Mined []struct {
			Notaryid   int    `json:"notaryid,omitempty"`
			KMDaddress string `json:"KMDaddress,omitempty"`
			Pubkey     string `json:"pubkey"`
			Blocks     int    `json:"blocks"`
		} `json:"mined"`
		Numnotaries int `json:"numnotaries"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) MinerIDs(ht string) (MinerIDs, error) {
	query := APIQuery {
		Method:	`minerids`,
		Params:	`["`+ht+`"]`,
	}
	//fmt.Println(query)

	var minerids MinerIDs

	mineridsJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(mineridsJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(mineridsJson), &minerids)
		return minerids, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(mineridsJson), &minerids)
	return minerids, nil
}