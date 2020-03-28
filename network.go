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

type AddNode struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) AddNode(params APIParams) (AddNode, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `addnode`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var addnode AddNode

	addnodeJson := appName.APICall(&query)
	if addnodeJson == "EMPTY RPC INFO!" {
		return addnode, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(addnodeJson)

	var result APIResult

	json.Unmarshal([]byte(addnodeJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(addnodeJson), &addnode)
		return addnode, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(addnodeJson), &addnode)
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

	clearbannedJson := appName.APICall(&query)
	if clearbannedJson == "EMPTY RPC INFO!" {
		return clearbanned, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(clearbannedJson)

	var result APIResult

	json.Unmarshal([]byte(clearbannedJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(clearbannedJson), &clearbanned)
		return clearbanned, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(clearbannedJson), &clearbanned)
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

	disconnectnodeJson := appName.APICall(&query)
	if disconnectnodeJson == "EMPTY RPC INFO!" {
		return disconnectnode, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(disconnectnodeJson)

	var result APIResult

	json.Unmarshal([]byte(disconnectnodeJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(disconnectnodeJson), &disconnectnode)
		return disconnectnode, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(disconnectnodeJson), &disconnectnode)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `getaddednodeinfo`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var getaddednodeinfo GetAddedNodeInfo

	getaddednodeinfoJson := appName.APICall(&query)
	if getaddednodeinfoJson == "EMPTY RPC INFO!" {
		return getaddednodeinfo, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(getaddednodeinfoJson)

	var result APIResult

	json.Unmarshal([]byte(getaddednodeinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getaddednodeinfoJson), &getaddednodeinfo)
		return getaddednodeinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getaddednodeinfoJson), &getaddednodeinfo)
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

	getconnectioncountJson := appName.APICall(&query)
	if getconnectioncountJson == "EMPTY RPC INFO!" {
		return getconnectioncount, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(getconnectioncountJson)

	var result APIResult

	json.Unmarshal([]byte(getconnectioncountJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getconnectioncountJson), &getconnectioncount)
		return getconnectioncount, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getconnectioncountJson), &getconnectioncount)
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

	getdeprecationinfoJson := appName.APICall(&query)
	if getdeprecationinfoJson == "EMPTY RPC INFO!" {
		return getdeprecationinfo, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(getdeprecationinfoJson)

	var result APIResult

	json.Unmarshal([]byte(getdeprecationinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getdeprecationinfoJson), &getdeprecationinfo)
		return getdeprecationinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getdeprecationinfoJson), &getdeprecationinfo)
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

	getnettotalsJson := appName.APICall(&query)
	if getnettotalsJson == "EMPTY RPC INFO!" {
		return getnettotals, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(getnettotalsJson)

	var result APIResult

	json.Unmarshal([]byte(getnettotalsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnettotalsJson), &getnettotals)
		return getnettotals, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getnettotalsJson), &getnettotals)
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

	getnetworkinfoJson := appName.APICall(&query)
	if getnetworkinfoJson == "EMPTY RPC INFO!" {
		return getnetworkinfo, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(getnetworkinfoJson)

	var result APIResult

	json.Unmarshal([]byte(getnetworkinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnetworkinfoJson), &getnetworkinfo)
		return getnetworkinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getnetworkinfoJson), &getnetworkinfo)
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

	getpeerinfoJson := appName.APICall(&query)
	if getpeerinfoJson == "EMPTY RPC INFO!" {
		return getpeerinfo, errors.New("EMPTY RPC INFO!")
	}
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

	listbannedJson := appName.APICall(&query)
	if listbannedJson == "EMPTY RPC INFO!" {
		return listbanned, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(listbannedJson)

	var result APIResult

	json.Unmarshal([]byte(listbannedJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listbannedJson), &listbanned)
		return listbanned, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listbannedJson), &listbanned)
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

	pingJson := appName.APICall(&query)
	if pingJson == "EMPTY RPC INFO!" {
		return ping, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(pingJson)

	var result APIResult

	json.Unmarshal([]byte(pingJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(pingJson), &ping)
		return ping, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(pingJson), &ping)
	return ping, nil
}

type SetBan struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SetBan(params APIParams) (SetBan, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `setban`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var setban SetBan

	setbanJson := appName.APICall(&query)
	if setbanJson == "EMPTY RPC INFO!" {
		return setban, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(setbanJson)

	var result APIResult

	json.Unmarshal([]byte(setbanJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(setbanJson), &setban)
		return setban, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(setbanJson), &setban)
	return setban, nil
}
