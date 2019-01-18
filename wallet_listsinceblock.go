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

type ListSinceBlock struct {
	Result struct {
		Transactions []struct {
			Account         string        `json:"account"`
			Address         string        `json:"address"`
			Category        string        `json:"category"`
			Amount          float64       `json:"amount"`
			Vout            int           `json:"vout"`
			Confirmations   int           `json:"confirmations"`
			Generated       bool          `json:"generated"`
			Blockhash       string        `json:"blockhash"`
			Blockindex      int           `json:"blockindex"`
			Blocktime       int           `json:"blocktime"`
			Expiryheight    int           `json:"expiryheight"`
			Txid            string        `json:"txid"`
			Walletconflicts []interface{} `json:"walletconflicts"`
			Time            int           `json:"time"`
			Timereceived    int           `json:"timereceived"`
			Vjoinsplit      []interface{} `json:"vjoinsplit"`
			Size            int           `json:"size"`
		} `json:"transactions"`
		Lastblock string `json:"lastblock"`
		} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ListSinceBlock(bhash string, tconf int, incwch bool) (ListSinceBlock, error) {
	query := APIQuery {
		Method:	`listsinceblock`,
		Params:	`["`+bhash+`", `+strconv.Itoa(tconf)+`, `+strconv.FormatBool(incwch)+`]`,
	}
	//fmt.Println(query)

	var listsinceblock ListSinceBlock

	listsinceblockJson := appName.APICall(query)
	//fmt.Println(listsinceblockJson)

	var result APIResult

	json.Unmarshal([]byte(listsinceblockJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listsinceblockJson), &listsinceblock)
		return listsinceblock, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listsinceblockJson), &listsinceblock)
	return listsinceblock, nil
}