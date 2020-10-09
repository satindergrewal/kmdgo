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
	if ListIDsJSON == "EMPTY RPC INFO" {
		return ListIDs, errors.New("EMPTY RPC INFO")
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

// GetIdentity type
type GetIdentity struct {
	Result struct {
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
		Status      string `json:"status"`
		Canspendfor bool   `json:"canspendfor"`
		Cansignfor  bool   `json:"cansignfor"`
		Blockheight int    `json:"blockheight"`
		Txid        string `json:"txid"`
		Vout        int    `json:"vout"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetIdentity searches blockchain and returns if there's matching VerusID found.
// getidentity "name"
func (appName AppType) GetIdentity(params APIParams) (GetIdentity, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getidentity`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var GetID GetIdentity

	GetIDJSON := appName.APICall(&query)
	if GetIDJSON == "EMPTY RPC INFO" {
		return GetID, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(GetIDJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(GetIDJSON), &GetID)
		return GetID, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(GetIDJSON), &GetID)
	return GetID, nil
}

// RegisterNameCommitment type
type RegisterNameCommitment struct {
	Result struct {
		Txid            string `json:"txid"`
		Namereservation struct {
			Name     string `json:"name"`
			Salt     string `json:"salt"`
			Referral string `json:"referral"`
			Parent   string `json:"parent"`
			Nameid   string `json:"nameid"`
		} `json:"namereservation"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RegisterNameCommitment - Registers a name commitment, which is required as a source for the name to be used when registering an identity. The name commitment hides the name itself
// while ensuring that the miner who mines in the registration cannot front-run the name unless they have also registered a name commitment for the same name or
// are willing to forfeit the offer of payment for the chance that a commitment made now will allow them to register the name in the future.
//
// NOTE: Invalid name for commitment. Names must not have leading or trailing spaces and must not include any of the following characters between parentheses (\/:*?"<>|@)
//
// registernamecommitment "name" "controladdress" ("referralidentity")
//
// Arguments
// "name"                           (string, required)  the unique name to commit to. creating a name commitment is not a registration, and if one is
//                                                        created for a name that exists, it may succeed, but will never be able to be used.
// "controladdress"                 (address, required) address that will control this commitment
// "referralidentity"               (identity, optional)friendly name or identity address that is provided as a referral mechanism and to lower network cost of the ID
func (appName AppType) RegisterNameCommitment(params APIParams) (RegisterNameCommitment, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `registernamecommitment`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var RegNameComit RegisterNameCommitment

	RegNameComitJSON := appName.APICall(&query)
	if RegNameComitJSON == "EMPTY RPC INFO" {
		return RegNameComit, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(RegNameComitJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(RegNameComitJSON), &RegNameComit)
		return RegNameComit, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(RegNameComitJSON), &RegNameComit)
	return RegNameComit, nil
}
