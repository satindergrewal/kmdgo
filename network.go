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
)

type AddNode struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) AddNode(params APIParams) (AddNode, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `addnode`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var addnode AddNode

	addnodeJSON := appName.APICall(&query)
	if addnodeJSON == "EMPTY RPC INFO" {
		return addnode, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(addnodeJSON)

	var result APIResult

	json.Unmarshal([]byte(addnodeJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(addnodeJSON), &addnode)
		return addnode, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(addnodeJSON), &addnode)
	return addnode, nil
}

type ClearBanned struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ClearBanned() (ClearBanned, error) {
	query := APIQuery{
		Method: `clearbanned`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var clearbanned ClearBanned

	clearbannedJSON := appName.APICall(&query)
	if clearbannedJSON == "EMPTY RPC INFO" {
		return clearbanned, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(clearbannedJSON)

	var result APIResult

	json.Unmarshal([]byte(clearbannedJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(clearbannedJSON), &clearbanned)
		return clearbanned, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(clearbannedJSON), &clearbanned)
	return clearbanned, nil
}

type DisconnectNode struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) DisconnectNode(node string) (DisconnectNode, error) {
	query := APIQuery{
		Method: `disconnectnode`,
		Params: `["` + node + `"]`,
	}
	//fmt.Println(query)

	var disconnectnode DisconnectNode

	disconnectnodeJSON := appName.APICall(&query)
	if disconnectnodeJSON == "EMPTY RPC INFO" {
		return disconnectnode, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(disconnectnodeJSON)

	var result APIResult

	json.Unmarshal([]byte(disconnectnodeJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(disconnectnodeJSON), &disconnectnode)
		return disconnectnode, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(disconnectnodeJSON), &disconnectnode)
	return disconnectnode, nil
}

type GetAddedNodeInfo struct {
	Result []struct {
		Addednode string `json:"addednode"`
		Connected bool   `json:"connected"`
		Addresses []struct {
			Address   string `json:"address"`
			Connected string `json:"connected"`
		} `json:"addresses"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetAddedNodeInfo(params APIParams) (GetAddedNodeInfo, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getaddednodeinfo`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getaddednodeinfo GetAddedNodeInfo

	getaddednodeinfoJSON := appName.APICall(&query)
	if getaddednodeinfoJSON == "EMPTY RPC INFO" {
		return getaddednodeinfo, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getaddednodeinfoJSON)

	var result APIResult

	json.Unmarshal([]byte(getaddednodeinfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getaddednodeinfoJSON), &getaddednodeinfo)
		return getaddednodeinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getaddednodeinfoJSON), &getaddednodeinfo)
	return getaddednodeinfo, nil
}

type GetConnectionCount struct {
	Result int    `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetConnectionCount() (GetConnectionCount, error) {
	query := APIQuery{
		Method: `getconnectioncount`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getconnectioncount GetConnectionCount

	getconnectioncountJSON := appName.APICall(&query)
	if getconnectioncountJSON == "EMPTY RPC INFO" {
		return getconnectioncount, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getconnectioncountJSON)

	var result APIResult

	json.Unmarshal([]byte(getconnectioncountJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getconnectioncountJSON), &getconnectioncount)
		return getconnectioncount, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getconnectioncountJSON), &getconnectioncount)
	return getconnectioncount, nil
}

type GetDeprecationInfo struct {
	Result struct {
		Version           int    `json:"version"`
		Subversion        string `json:"subversion"`
		Deprecationheight int    `json:"deprecationheight"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetDeprecationInfo() (GetDeprecationInfo, error) {
	query := APIQuery{
		Method: `getdeprecationinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getdeprecationinfo GetDeprecationInfo

	getdeprecationinfoJSON := appName.APICall(&query)
	if getdeprecationinfoJSON == "EMPTY RPC INFO" {
		return getdeprecationinfo, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getdeprecationinfoJSON)

	var result APIResult

	json.Unmarshal([]byte(getdeprecationinfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getdeprecationinfoJSON), &getdeprecationinfo)
		return getdeprecationinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getdeprecationinfoJSON), &getdeprecationinfo)
	return getdeprecationinfo, nil
}

type GetNetTotals struct {
	Result struct {
		Totalbytesrecv int   `json:"totalbytesrecv"`
		Totalbytessent int   `json:"totalbytessent"`
		Timemillis     int64 `json:"timemillis"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetNetTotals() (GetNetTotals, error) {
	query := APIQuery{
		Method: `getnettotals`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getnettotals GetNetTotals

	getnettotalsJSON := appName.APICall(&query)
	if getnettotalsJSON == "EMPTY RPC INFO" {
		return getnettotals, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getnettotalsJSON)

	var result APIResult

	json.Unmarshal([]byte(getnettotalsJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnettotalsJSON), &getnettotals)
		return getnettotals, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getnettotalsJSON), &getnettotals)
	return getnettotals, nil
}

type GetNetworkInfo struct {
	Result struct {
		Version         int    `json:"version"`
		Subversion      string `json:"subversion"`
		Protocolversion int    `json:"protocolversion"`
		Localservices   string `json:"localservices"`
		Timeoffset      int    `json:"timeoffset"`
		Connections     int    `json:"connections"`
		Networks        []struct {
			Name                      string `json:"name"`
			Limited                   bool   `json:"limited"`
			Reachable                 bool   `json:"reachable"`
			Proxy                     string `json:"proxy"`
			ProxyRandomizeCredentials bool   `json:"proxy_randomize_credentials"`
		} `json:"networks"`
		Relayfee       float64       `json:"relayfee"`
		Localaddresses []interface{} `json:"localaddresses"`
		Warnings       string        `json:"warnings"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetNetworkInfo() (GetNetworkInfo, error) {
	query := APIQuery{
		Method: `getnetworkinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getnetworkinfo GetNetworkInfo

	getnetworkinfoJSON := appName.APICall(&query)
	if getnetworkinfoJSON == "EMPTY RPC INFO" {
		return getnetworkinfo, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getnetworkinfoJSON)

	var result APIResult

	json.Unmarshal([]byte(getnetworkinfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnetworkinfoJSON), &getnetworkinfo)
		return getnetworkinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getnetworkinfoJSON), &getnetworkinfo)
	return getnetworkinfo, nil
}

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
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetPeerInfo() (GetPeerInfo, error) {
	query := APIQuery{
		Method: `getpeerinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getpeerinfo GetPeerInfo

	getpeerinfoJSON := appName.APICall(&query)
	if getpeerinfoJSON == "EMPTY RPC INFO" {
		return getpeerinfo, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getpeerinfoJSON)

	var result APIResult

	json.Unmarshal([]byte(getpeerinfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getpeerinfoJSON), &getpeerinfo)
		return getpeerinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getpeerinfoJSON), &getpeerinfo)
	return getpeerinfo, nil
}

type ListBanned struct {
	Result []struct {
		Address     string `json:"address"`
		BannedUntil int    `json:"banned_until"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ListBanned() (ListBanned, error) {
	query := APIQuery{
		Method: `listbanned`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var listbanned ListBanned

	listbannedJSON := appName.APICall(&query)
	if listbannedJSON == "EMPTY RPC INFO" {
		return listbanned, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(listbannedJSON)

	var result APIResult

	json.Unmarshal([]byte(listbannedJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listbannedJSON), &listbanned)
		return listbanned, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(listbannedJSON), &listbanned)
	return listbanned, nil
}

type Ping struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) Ping() (Ping, error) {
	query := APIQuery{
		Method: `ping`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var ping Ping

	pingJSON := appName.APICall(&query)
	if pingJSON == "EMPTY RPC INFO" {
		return ping, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(pingJSON)

	var result APIResult

	json.Unmarshal([]byte(pingJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(pingJSON), &ping)
		return ping, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(pingJSON), &ping)
	return ping, nil
}

type SetBan struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SetBan(params APIParams) (SetBan, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `setban`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var setban SetBan

	setbanJSON := appName.APICall(&query)
	if setbanJSON == "EMPTY RPC INFO" {
		return setban, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(setbanJSON)

	var result APIResult

	json.Unmarshal([]byte(setbanJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(setbanJSON), &setban)
		return setban, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(setbanJSON), &setban)
	return setban, nil
}
