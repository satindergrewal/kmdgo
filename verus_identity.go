// Copyright © 2018-2020 Satinderjit Singh.
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
	"fmt"
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
// "name"                           (string, required)  the unique name to commit to. creating a name commitment is not a registration, and if one is created for a name that exists, it may succeed, but will never be able to be used.
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

// RegisterIdentity type
type RegisterIdentity struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

// RegisterIdentityData holds the JSON data supplied for registering VerusID
type RegisterIdentityData struct {
	Txid            string `json:"txid"`
	Namereservation struct {
		Name     string `json:"name"`
		Salt     string `json:"salt"`
		Referral string `json:"referral"`
		Parent   string `json:"parent"`
		Nameid   string `json:"nameid"`
	} `json:"namereservation"`
	Identity struct {
		Name                string   `json:"name"`
		Primaryaddresses    []string `json:"primaryaddresses"`
		Minimumsignatures   int      `json:"minimumsignatures"`
		Revocationauthority string   `json:"revocationauthority"`
		Recoveryauthority   string   `json:"recoveryauthority"`
		Privateaddress      string   `json:"privateaddress"`
	} `json:"identity"`
}

// RegisterIdentity gets the values from RegisterNameCommitment output and uses those to register VerusID
//
// registeridentity "jsonidregistration" feeoffer
//
// Arguments
// {
// 	"txid": "hexid",			(hex)    the transaction ID of the name committment for this ID name - Taken from RegisterNameCommitment's output - txid
// 	"namereservation": {
// 	  "name": "namestr",		(string) the unique name in this commitment - Taken from RegisterNameCommitment's output - name
// 	  "salt": "hexstr",			(hex)    salt used to hide the commitment - Taken from RegisterNameCommitment's output - salt
// 	  "referral": "identityID",	(name@ or address) must be a valid ID to use as a referrer to receive a discount - Taken from RegisterNameCommitment's output - referral
// 	  "parent": "",				(name@ or address) must be a valid ID. This ID can be used to revoke and recover the nameID we regsiter with this current command - Taken from RegisterNameCommitment's output - parent
// 	  "nameid": "nameID"		(base58) identity address for this identity if it is created - Taken from RegisterNameCommitment's output - nameid
// 	},
// 	"identity": {
// 	  "name": "namestr",		(string) the unique name for this identity - Taken from RegisterNameCommitment's output - name
// 	  "primaryaddresses": [		(array of strings) the trasparent/public address(es)
// 		"hexstr"
// 	  ],
// 	  "minimumsignatures": 1,	(int) MofN signatures required out of the primary addresses list to sign transactions
//    "revocationauthority": "namestr", (name@ or address) The ID entered here will be able to disable your created ID in case of loss or theft. It is some existing ID which either is under your own control or the ID you trust can help you revoke in case of this ID's theft.
//    "recoveryauthority": "namestr",	(name@ or address) The ID entered here will be able to revive your created ID if it is revoked. It is some existing ID which either is under your own control or the ID you trust can help you revive in case of this ID's revoked.
// 	  "privateaddress": "hexstr"	(string) shielded address associated with the VerusID being made
// 	}
// }
// feeoffer                           (amount, optional) amount to offer miner/staker for the registration fee, if missing, uses standard price
func (appName AppType) RegisterIdentity(params APIParams) (RegisterIdentity, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println("paramsJSON", string(paramsJSON))

	query := APIQuery{
		Method: `registeridentity`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var RegID RegisterIdentity

	RegIDJSON := appName.APICall(&query)
	if RegIDJSON == "EMPTY RPC INFO" {
		return RegID, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(RegIDJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(RegIDJSON), &RegID)
		return RegID, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(RegIDJSON), &RegID)
	return RegID, nil
}

// UpdateIdentity type
type UpdateIdentity struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

// UpdateIdentityData type
type UpdateIdentityData struct {
	Parent              string                 `json:"parent"`
	Name                string                 `json:"name"`
	Flags               int                    `json:"flags"`
	Timelock            int                    `json:"timelock,omitempty"`
	Primaryaddresses    []string               `json:"primaryaddresses"`
	Minimumsignatures   int                    `json:"minimumsignatures"`
	Contentmap          map[string]interface{} `json:"contentmap"`
	Revocationauthority string                 `json:"revocationauthority"`
	Recoveryauthority   string                 `json:"recoveryauthority"`
	Privateaddress      string                 `json:"privateaddress"`
}

// UpdateIdentity gets the values from GetIdentity output and uses those to update VerusID
//
// updateidentity "jsonidregistration" (returntx)
//
// Arguments
// {
// 	"parent": "address",		(name@ or address) must be a valid ID. This ID can be used to revoke and recover the nameID we regsiter with this current command - Taken from GetIdentity's output - parent
// 	"name": "namestr",			(string) the unique name in this commitment - Taken from GetIdentity's output - name
//	"flags": number,			(int) default value is 0.
//	"timelock": 1219040,		(int) block height to define.
// 	"primaryaddresses": [		(array of strings) the trasparent/public address(es)
// 		"hexstr"
// 	],
// 	"minimumsignatures": 1,		(int) MofN signatures required out of the primary addresses list to sign transactions
//	"contentmap": {},			(string) it takes a 20 byte key and 32 byte value, expected to be a key hash and whatever other hash for value you want for content addressable storage. It needs to be a 20 byte (40 char)hex value key and a 32 byte (64) hex value value. It is an object of key/values.
// 	"revocationauthority": "namestr",	(name@ or address) The ID entered here will be able to disable your created ID in case of loss or theft. It is some existing ID which either is under your own control or the ID you trust can help you revoke in case of this ID's theft.
// 	"recoveryauthority": "namestr",		(name@ or address) The ID entered here will be able to revive your created ID if it is revoked. It is some existing ID which either is under your own control or the ID you trust can help you revive in case of this ID's revoked.
// 	"privateaddress": "hexstr"			(string) shielded address associated with the VerusID being made
// }
// "returntx"					(bool,   optional) defaults to false and transaction is sent, if true, transaction is signed by this wallet and returned
//
// ## Command Examples
// Time Lock: The timelock parameter defines the unlock height of the identity.
// 	verus -chain=VRSCTEST updateidentity '{"name": "ID@", "flags": 0, "timelock": <Unlock block height>, "minimumsignatures": 1, "primaryaddresses": ["t-address"]}'
//
// Time Delay: The timelock parameter defines how many blocks to delay an ID's unlock when the flags are set back to an unlocked state.
//	verus -chain=VRSCTEST updateidentity '{"name": "ID@", "flags": 2, "timelock": <Unlock block delay>, "minimumsignatures": 1, "primaryaddresses": ["t-address"]}'
func (appName AppType) UpdateIdentity(params APIParams) (UpdateIdentity, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println("paramsJSON", string(paramsJSON))

	query := APIQuery{
		Method: `updateidentity`,
		Params: string(paramsJSON),
	}
	fmt.Println(query)

	var RegID UpdateIdentity

	RegIDJSON := appName.APICall(&query)
	if RegIDJSON == "EMPTY RPC INFO" {
		return RegID, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(RegIDJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(RegIDJSON), &RegID)
		return RegID, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(RegIDJSON), &RegID)
	return RegID, nil
}
