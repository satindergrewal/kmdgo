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

type WalletLock struct {
	Result	string	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) WalletLock() (WalletLock, error) {
	query := APIQuery {
		Method:	`walletlock`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var walletlock WalletLock

	walletlockJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(walletlockJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(walletlockJson), &walletlock)
		return walletlock, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(walletlockJson), &walletlock)
	return walletlock, nil
}