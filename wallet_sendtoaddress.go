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

type SendToAddress struct {
	Result string `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) SendToAddress(taddr string, amount float64, comment string, commentto string, subtractfeefromamount bool) (SendToAddress, error) {
	query := APIQuery {
		Method:	`sendtoaddress`,
		Params:	`["`+taddr+`", `+strconv.FormatFloat(amount, 'f', 8, 64)+`, "`+comment+`", "`+commentto+`", `+strconv.FormatBool(subtractfeefromamount)+`]`,
	}
	//fmt.Println(query)

	var sendtoaddress SendToAddress

	sendtoaddressJson := appName.APICall(query)
	//fmt.Println(sendtoaddressJson)

	var result APIResult

	json.Unmarshal([]byte(sendtoaddressJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendtoaddressJson), &sendtoaddress)
		return sendtoaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(sendtoaddressJson), &sendtoaddress)
	return sendtoaddress, nil
}