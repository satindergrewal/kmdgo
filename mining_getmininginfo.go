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

type GetMiningInfo struct {
	Result struct {
		Blocks           int     `json:"blocks"`
		Currentblocksize int     `json:"currentblocksize"`
		Currentblocktx   int     `json:"currentblocktx"`
		Difficulty       float64 `json:"difficulty"`
		Errors           string  `json:"errors"`
		Genproclimit     int     `json:"genproclimit"`
		Localsolps       int     `json:"localsolps"`
		Networksolps     int     `json:"networksolps"`
		Networkhashps    int     `json:"networkhashps"`
		Pooledtx         int     `json:"pooledtx"`
		Testnet          bool    `json:"testnet"`
		Chain            string  `json:"chain"`
		Generate         bool    `json:"generate"`
		Numthreads       int     `json:"numthreads"`
	} `json:"result"`
	Error Error `json:"error"`
	ID    string      `json:"id"`
}

func (appName AppType) GetMiningInfo() (GetMiningInfo, error) {
	query := APIQuery {
		Method:	`getmininginfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getmininginfo GetMiningInfo

	getmininginfoJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getmininginfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getmininginfoJson), &getmininginfo)
		return getmininginfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getmininginfoJson), &getmininginfo)
	return getmininginfo, nil
}