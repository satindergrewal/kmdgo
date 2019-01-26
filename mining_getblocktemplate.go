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
)

type GetBlockTemplate struct {
	Result struct {
		Capabilities      []string `json:"capabilities"`
		Version           int      `json:"version"`
		Previousblockhash string   `json:"previousblockhash"`
		Transactions      []struct {
			Data    string        `json:"data"`
			Hash    string        `json:"hash"`
			Depends []interface{} `json:"depends"`
			Fee     int           `json:"fee"`
			Sigops  int           `json:"sigops"`
		} `json:"transactions"`
		Coinbasetxn struct {
			Data          string        `json:"data"`
			Hash          string        `json:"hash"`
			Depends       []interface{} `json:"depends"`
			Fee           int           `json:"fee"`
			Sigops        int           `json:"sigops"`
			Coinbasevalue int           `json:"coinbasevalue"`
			Required      bool          `json:"required"`
		} `json:"coinbasetxn"`
		Longpollid string   `json:"longpollid"`
		Target     string   `json:"target"`
		Mintime    int      `json:"mintime"`
		Mutable    []string `json:"mutable"`
		Noncerange string   `json:"noncerange"`
		Sigoplimit int      `json:"sigoplimit"`
		Sizelimit  int      `json:"sizelimit"`
		Curtime    int      `json:"curtime"`
		Bits       string   `json:"bits"`
		Height     int      `json:"height"`
	} `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) GetBlockTemplate(params APIParams) (GetBlockTemplate, error) {
	var params_json string

	if params[0] == "" || params[0] == nil {
		params_json = `[]`
		//fmt.Println(params_json)
	} else {
		params_bytes, _ := json.Marshal(params)
		params_json = string(params_bytes)
		//fmt.Println(params_json)
	}

	query := APIQuery {
		Method:	`getblocktemplate`,
		Params:	params_json,
	}
	//fmt.Println(query)

	var getblocktemplate GetBlockTemplate

	getblocktemplateJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getblocktemplateJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblocktemplateJson), &getblocktemplate)
		return getblocktemplate, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblocktemplateJson), &getblocktemplate)
	return getblocktemplate, nil
}