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
)

// GetInfo struct type to store getinfo API output
type GetInfo struct {
	Result struct {
		Version             int     `json:"version"`
		Protocolversion     int     `json:"protocolversion"`
		KMDversion          string  `json:"KMDversion"`
		Notarized           int     `json:"notarized"`
		PrevMoMheight       int     `json:"prevMoMheight"`
		Notarizedhash       string  `json:"notarizedhash"`
		Notarizedtxid       string  `json:"notarizedtxid"`
		NotarizedtxidHeight string  `json:"notarizedtxid_height"`
		KMDnotarizedHeight  int     `json:"KMDnotarized_height"`
		NotarizedConfirms   int     `json:"notarized_confirms"`
		Walletversion       int     `json:"walletversion"`
		Balance             float64 `json:"balance"`
		Blocks              int     `json:"blocks"`
		Longestchain        int     `json:"longestchain"`
		Timeoffset          int     `json:"timeoffset"`
		Tiptime             int     `json:"tiptime"`
		Connections         int     `json:"connections"`
		Proxy               string  `json:"proxy"`
		Difficulty          float64 `json:"difficulty"`
		Testnet             bool    `json:"testnet"`
		Keypoololdest       int     `json:"keypoololdest"`
		Keypoolsize         int     `json:"keypoolsize"`
		Paytxfee            float64 `json:"paytxfee"`
		Relayfee            float64 `json:"relayfee"`
		Errors              string  `json:"errors"`
		CCid                int     `json:"CCid"`
		Name                string  `json:"name"`
		P2Pport             int     `json:"p2pport"`
		Rpcport             int     `json:"rpcport"`
		Magic               int     `json:"magic"`
		Premine             int     `json:"premine"`
		Reward              int64   `json:"reward"`
		Halving             int     `json:"halving"`
		Commission          int     `json:"commission"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetInfo gets some basic and essential information about local node's blockchain, wallet, and network specific info
func (appName AppType) GetInfo() (GetInfo, error) {

	query := new(APIQuery)
	query = &APIQuery{
		Method: "getinfo",
		Params: "[]",
	}

	var getinfo GetInfo

	getinfoJSON := appName.APICall(query)
	if getinfoJSON == "EMPTY RPC INFO" {
		return getinfo, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	// fmt.Println("returned getinfoJson: ", string(getinfoJSON))

	json.Unmarshal([]byte(getinfoJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getinfoJSON), &getinfo)
		return getinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getinfoJSON), &getinfo)
	return getinfo, nil
}

// Stop struct type to store output of chain stop command
type Stop struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

// Stop method allows to issue a blockchain daemon service shutdown command
func (appName AppType) Stop() (Stop, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: "stop",
		Params: "[]",
	}

	var stop Stop

	stopJSON := appName.APICall(query)
	if stopJSON == "EMPTY RPC INFO" {
		return stop, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(stopJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(stopJSON), &stop)
		return stop, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(stopJSON), &stop)
	return stop, nil
}
