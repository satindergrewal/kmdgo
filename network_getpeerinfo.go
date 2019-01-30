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

type GetPeerInfo struct {
	Result []struct {
		ID             int           `json:"id"`
		Addr           string        `json:"addr"`
		Addrlocal      string        `json:"addrlocal"`
		Services       string        `json:"services"`
		Lastsend       int           `json:"lastsend"`
		Lastrecv       int           `json:"lastrecv"`
		Bytessent      int           `json:"bytessent"`
		Bytesrecv      int           `json:"bytesrecv"`
		Conntime       int           `json:"conntime"`
		Timeoffset     int           `json:"timeoffset"`
		Pingtime       float64       `json:"pingtime"`
		Version        int           `json:"version"`
		Subver         string        `json:"subver"`
		Inbound        bool          `json:"inbound"`
		Startingheight int           `json:"startingheight"`
		Banscore       int           `json:"banscore"`
		SyncedHeaders  int           `json:"synced_headers"`
		SyncedBlocks   int           `json:"synced_blocks"`
		Inflight       []interface{} `json:"inflight"`
		Whitelisted    bool          `json:"whitelisted"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetPeerInfo() (GetPeerInfo, error) {
	query := APIQuery {
		Method:	`getpeerinfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getpeerinfo GetPeerInfo

	getpeerinfoJson := appName.APICall(query)
	//fmt.Println(getpeerinfoJson)

	var result APIResult

	json.Unmarshal([]byte(getpeerinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getpeerinfoJson), &getpeerinfo)
		return getpeerinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getpeerinfoJson), &getpeerinfo)
	return getpeerinfo, nil
}