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
	"encoding/json"
	"errors"
)

// ListIdentities struct type to render result for ListIdentities func
type ListIdentities struct {
	Result []struct {
		Identity struct {
			Version           int      `json:"version"`
			Flags             int      `json:"flags"`
			Primaryaddresses  []string `json:"primaryaddresses"`
			Minimumsignatures int      `json:"minimumsignatures"`
			Identityaddress   string   `json:"identityaddress"`
			Parent            string   `json:"parent"`
			Name              string   `json:"name"`
			Contentmap        struct {
			} `json:"contentmap"`
			Revocationauthority string `json:"revocationauthority"`
			Recoveryauthority   string `json:"recoveryauthority"`
			Privateaddress      string `json:"privateaddress"`
			Timelock            int    `json:"timelock"`
		} `json:"identity"`
		Blockheight int    `json:"blockheight"`
		Txid        string `json:"txid"`
		Status      string `json:"status"`
		Canspendfor bool   `json:"canspendfor"`
		Cansignfor  bool   `json:"cansignfor"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// ListIdentities lists identities available in your wallet
//
// listidentities (includecansign) (includewatchonly)
//
// Arguments
// "includecanspend"    (bool, optional, default=true)    Include identities for which we can spend/authorize
// "includecansign"     (bool, optional, default=true)    Include identities that we can only sign for but not spend
// "includewatchonly"   (bool, optional, default=false)   Include identities that we can neither sign nor spend, but are either watched or are co-signers with us
func (appName AppType) ListIdentities(params APIParams) (ListIdentities, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `listidentities`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var ListIDs ListIdentities

	ListIDsJSON := appName.APICall(&query)
	if ListIDsJSON == "EMPTY RPC INFO!" {
		return ListIDs, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(ListIDsJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(ListIDsJSON), &ListIDs)
		return ListIDs, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(ListIDsJSON), &ListIDs)
	return ListIDs, nil
}
