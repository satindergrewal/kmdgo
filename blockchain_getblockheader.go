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
	"strconv"
	"errors"
)

type GetBlockHeader struct {
	Result struct {
		Hash              string  `json:"hash"`
		Confirmations     int     `json:"confirmations"`
		Rawconfirmations  int     `json:"rawconfirmations"`
		Height            int     `json:"height"`
		Version           int     `json:"version"`
		Merkleroot        string  `json:"merkleroot"`
		Finalsaplingroot  string  `json:"finalsaplingroot"`
		Time              int     `json:"time"`
		Nonce             string  `json:"nonce"`
		Solution          string  `json:"solution"`
		Bits              string  `json:"bits"`
		Difficulty        float64 `json:"difficulty"`
		Chainwork         string  `json:"chainwork"`
		Segid             int     `json:"segid"`
		Previousblockhash string  `json:"previousblockhash"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetBlockHeader(h string, b bool) (GetBlockHeader, error) {
	query := APIQuery {
		Method:	`getblockheader`,
		Params:	`["`+h+`",`+strconv.FormatBool(b)+`]`,
	}
	//fmt.Println(query)

	var getblockheader GetBlockHeader

	getblockheaderJson := appName.APICall(query)


	var result APIResult

	json.Unmarshal([]byte(getblockheaderJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockheaderJson), &getblockheader)
		return getblockheader, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblockheaderJson), &getblockheader)
	return getblockheader, nil
}