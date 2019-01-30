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
package kmdgo

import (
	"errors"
	"encoding/json"
)

type GetBlockchainInfo struct {
	Result struct {
		Chain                string  `json:"chain"`
		Blocks               int     `json:"blocks"`
		Headers              int     `json:"headers"`
		Bestblockhash        string  `json:"bestblockhash"`
		Difficulty           float64 `json:"difficulty"`
		Verificationprogress float64 `json:"verificationprogress"`
		Chainwork            string  `json:"chainwork"`
		Pruned               bool    `json:"pruned"`
		Commitments          int     `json:"commitments"`
		ValuePools           []struct {
			ID            string  `json:"id"`
			Monitored     bool    `json:"monitored"`
			ChainValue    float64 `json:"chainValue"`
			ChainValueZat int64   `json:"chainValueZat"`
		} `json:"valuePools"`
		Softforks []struct {
			ID      string `json:"id"`
			Version int    `json:"version"`
			Enforce struct {
				Status   bool `json:"status"`
				Found    int  `json:"found"`
				Required int  `json:"required"`
				Window   int  `json:"window"`
			} `json:"enforce"`
			Reject struct {
				Status   bool `json:"status"`
				Found    int  `json:"found"`
				Required int  `json:"required"`
				Window   int  `json:"window"`
			} `json:"reject"`
		} `json:"softforks"`
		Upgrades struct {
			FiveBa81B19 struct {
				Name             string `json:"name"`
				Activationheight int    `json:"activationheight"`
				Status           string `json:"status"`
				Info             string `json:"info"`
			} `json:"5ba81b19"`
			Seven6B809Bb struct {
				Name             string `json:"name"`
				Activationheight int    `json:"activationheight"`
				Status           string `json:"status"`
				Info             string `json:"info"`
			} `json:"76b809bb"`
		} `json:"upgrades"`
		Consensus struct {
			Chaintip  string `json:"chaintip"`
			Nextblock string `json:"nextblock"`
		} `json:"consensus"`
	}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}


func (appName AppType) GetBlockchainInfo() (GetBlockchainInfo, error) {
	query := APIQuery {
		Method:	`getblockchaininfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getblockchaininfo GetBlockchainInfo

	getblockchaininfoJson := appName.APICall(query)


	var result APIResult

	json.Unmarshal([]byte(getblockchaininfoJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockchaininfoJson), &getblockchaininfo)
		return getblockchaininfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblockchaininfoJson), &getblockchaininfo)
	return getblockchaininfo, nil
}