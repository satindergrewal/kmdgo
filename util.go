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
	//"fmt"
	"encoding/json"
	"errors"
	//"strconv"
)

// ValidateAddress struct type to store validate address API output
type ValidateAddress struct {
	Result struct {
		Isvalid      bool   `json:"isvalid"`
		Address      string `json:"address"`
		ScriptPubKey string `json:"scriptPubKey"`
		Segid        int    `json:"segid"`
		Ismine       bool   `json:"ismine"`
		Iswatchonly  bool   `json:"iswatchonly"`
		Isscript     bool   `json:"isscript"`
		Pubkey       string `json:"pubkey"`
		Iscompressed bool   `json:"iscompressed"`
		Account      string `json:"account"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// ValidateAddress method allows to check the information about a single wallet address's status
func (appName AppType) ValidateAddress(taddr string) (ValidateAddress, error) {
	query := APIQuery{
		Method: `validateaddress`,
		Params: `["` + taddr + `"]`,
	}
	//fmt.Println(query)

	var validateaddress ValidateAddress

	validateaddressJSON := appName.APICall(&query)
	if validateaddressJSON == "EMPTY RPC INFO" {
		return validateaddress, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(validateaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(validateaddressJSON), &validateaddress)
		return validateaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(validateaddressJSON), &validateaddress)
	return validateaddress, nil
}

// VerifyMessage type
type VerifyMessage struct {
	Result bool   `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

// VerifyMessage - Verify a signed message
//
// verifymessage "address or identity" "signature" "message" "checklatest"
//
// Arguments:
// 1. "t-addr or identity" (string, required) The transparent address or identity that signed the message.
// 2. "signature"       (string, required) The signature provided by the signer in base 64 encoding (see signmessage).
// 3. "message"         (string, required) The message that was signed.
// 3. "checklatest"     (bool, optional)   If true, checks signature validity based on latest identity. defaults to false,
//                                           which determines validity of signing height stored in signature.
//
// Result:
// true|false   (boolean) If the signature is verified or not.
func (appName AppType) VerifyMessage(params APIParams) (VerifyMessage, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `verifymessage`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var verifymessage VerifyMessage

	verifymessageJSON := appName.APICall(&query)
	if verifymessageJSON == "EMPTY RPC INFO" {
		return verifymessage, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(verifymessageJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(verifymessageJSON), &verifymessage)
		return verifymessage, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(verifymessageJSON), &verifymessage)
	return verifymessage, nil
}

// ZValidateAddress type
type ZValidateAddress struct {
	Result struct {
		Isvalid                    bool   `json:"isvalid"`
		Address                    string `json:"address"`
		Type                       string `json:"type"`
		Diversifier                string `json:"diversifier"`
		Diversifiedtransmissionkey string `json:"diversifiedtransmissionkey"`
		Ismine                     bool   `json:"ismine"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// ZValidateAddress returns information about the given z address.
//
// z_validateaddress "zaddr"
//
// Arguments:
// 1. "zaddr"     (string, required) The z address to validate
//
// Result:
// {
//   "isvalid" : true|false,      (boolean) If the address is valid or not. If not, this is the only property returned.
//   "address" : "zaddr",         (string) The z address validated
//   "type" : "xxxx",             (string) "sprout" or "sapling"
//   "ismine" : true|false,       (boolean) If the address is yours or not
//   "payingkey" : "hex",         (string) [sprout] The hex value of the paying key, a_pk
//   "transmissionkey" : "hex",   (string) [sprout] The hex value of the transmission key, pk_enc
//   "diversifier" : "hex",       (string) [sapling] The hex value of the diversifier, d
//   "diversifiedtransmissionkey" : "hex", (string) [sapling] The hex value of pk_d
// }
func (appName AppType) ZValidateAddress(zaddr string) (ZValidateAddress, error) {
	query := APIQuery{
		Method: `z_validateaddress`,
		Params: `["` + zaddr + `"]`,
	}
	//fmt.Println(query)

	var zvalidateaddress ZValidateAddress

	zvalidateaddressJSON := appName.APICall(&query)
	if zvalidateaddressJSON == "EMPTY RPC INFO" {
		return zvalidateaddress, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(zvalidateaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zvalidateaddressJSON), &zvalidateaddress)
		return zvalidateaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zvalidateaddressJSON), &zvalidateaddress)
	return zvalidateaddress, nil
}
