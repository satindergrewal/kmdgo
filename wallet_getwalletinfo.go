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
)

type GetWalletInfo struct {
	Result	struct {
		Walletversion      int     `json:"walletversion"`
		Balance            float64 `json:"balance"`
		UnconfirmedBalance float64 `json:"unconfirmed_balance"`
		ImmatureBalance    float64 `json:"immature_balance"`
		Txcount            int     `json:"txcount"`
		Keypoololdest      int     `json:"keypoololdest"`
		Keypoolsize        int     `json:"keypoolsize"`
		UnlockedUntil      int     `json:"unlocked_until"`
		Paytxfee           float64 `json:"paytxfee"`
		Seedfp             string  `json:"seedfp"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) GetWalletInfo() (GetWalletInfo, error) {
	query := APIQuery {
		Method:	`getwalletinfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getwalletinfo GetWalletInfo

	getwalletinfoJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getwalletinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getwalletinfoJson), &getwalletinfo)
		return getwalletinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getwalletinfoJson), &getwalletinfo)
	return getwalletinfo, nil
}