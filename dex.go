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
	"fmt"
)

type DEXAnonsend struct {
	Result struct {
		Timestamp    int    `json:"timestamp"`
		ID           int    `json:"id"`
		Hash         string `json:"hash"`
		TagA         string `json:"tagA"`
		TagB         string `json:"tagB"`
		Pubkey       string `json:"pubkey"`
		Payload      string `json:"payload"`
		Hex          int    `json:"hex"`
		Decrypted    string `json:"decrypted"`
		Decryptedhex int    `json:"decryptedhex"`
		Senderpub    string `json:"senderpub"`
		AmountA      string `json:"amountA"`
		AmountB      string `json:"amountB"`
		Priority     int    `json:"priority"`
		Recvtime     int    `json:"recvtime"`
		Cancelled    int    `json:"cancelled"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) DEXAnonsend(params APIParams) (DEXAnonsend, error) {

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: `DEX_anonsend`,
		Params: string(params_json),
	}
	fmt.Println(query)

	var dexAnonsend DEXAnonsend

	dexAnonsendJson := appName.APICall(query)
	if dexAnonsendJson == "EMPTY RPC INFO!" {
		return dexAnonsend, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(dexAnonsendJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexAnonsendJson), &dexAnonsend)
		return dexAnonsend, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(dexAnonsendJson), &dexAnonsend)
	return dexAnonsend, nil
}

type DEXBroadcast struct {
	Result struct {
		Timestamp    int    `json:"timestamp"`
		ID           int    `json:"id"`
		Hash         string `json:"hash"`
		TagA         string `json:"tagA"`
		TagB         string `json:"tagB"`
		Pubkey       string `json:"pubkey"`
		Payload      string `json:"payload"`
		Hex          int    `json:"hex"`
		Decrypted    string `json:"decrypted"`
		Decryptedhex int    `json:"decryptedhex"`
		Senderpub    string `json:"senderpub"`
		AmountA      string `json:"amountA"`
		AmountB      string `json:"amountB"`
		Priority     int    `json:"priority"`
		Recvtime     int    `json:"recvtime"`
		Cancelled    int    `json:"cancelled"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) DEXBroadcast(params APIParams) (DEXBroadcast, error) {

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: `DEX_broadcast`,
		Params: string(params_json),
	}
	fmt.Println(query)

	var dexBroadcast DEXBroadcast

	dexBroadcastJson := appName.APICall(query)
	if dexBroadcastJson == "EMPTY RPC INFO!" {
		return dexBroadcast, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(dexBroadcastJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexBroadcastJson), &dexBroadcast)
		return dexBroadcast, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(dexBroadcastJson), &dexBroadcast)
	return dexBroadcast, nil
}

type DEXCancel struct {
	Result struct {
		Timestamp    int    `json:"timestamp"`
		ID           int    `json:"id"`
		Hash         string `json:"hash"`
		TagA         string `json:"tagA"`
		TagB         string `json:"tagB"`
		Pubkey       string `json:"pubkey"`
		Payload      string `json:"payload"`
		Hex          int    `json:"hex"`
		Decrypted    string `json:"decrypted"`
		Decryptedhex int    `json:"decryptedhex"`
		AmountA      string `json:"amountA"`
		AmountB      string `json:"amountB"`
		Priority     int    `json:"priority"`
		Recvtime     int    `json:"recvtime"`
		Cancelled    int    `json:"cancelled"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) DEXCancel(params APIParams) (DEXCancel, error) {

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: `DEX_cancel`,
		Params: string(params_json),
	}
	fmt.Println(query)

	var dexCancel DEXCancel

	dexCancelJson := appName.APICall(query)
	if dexCancelJson == "EMPTY RPC INFO!" {
		return dexCancel, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(dexCancelJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexCancelJson), &dexCancel)
		return dexCancel, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(dexCancelJson), &dexCancel)
	return dexCancel, nil
}

type DEXGet struct {
	Result struct {
		Timestamp    int    `json:"timestamp"`
		ID           int    `json:"id"`
		Hash         string `json:"hash"`
		TagA         string `json:"tagA"`
		TagB         string `json:"tagB"`
		Pubkey       string `json:"pubkey"`
		Payload      string `json:"payload"`
		Hex          int    `json:"hex"`
		Decrypted    string `json:"decrypted"`
		Decryptedhex int    `json:"decryptedhex"`
		Senderpub    string `json:"senderpub"`
		AmountA      string `json:"amountA"`
		AmountB      string `json:"amountB"`
		Priority     int    `json:"priority"`
		Recvtime     int    `json:"recvtime"`
		Cancelled    int    `json:"cancelled"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) DEXGet(params APIParams) (DEXGet, error) {

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: `DEX_get`,
		Params: string(params_json),
	}
	fmt.Println(query)

	var dexGet DEXGet

	dexGetJson := appName.APICall(query)
	if dexGetJson == "EMPTY RPC INFO!" {
		return dexGet, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(dexGetJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexGetJson), &dexGet)
		return dexGet, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(dexGetJson), &dexGet)
	return dexGet, nil
}

type DEXList struct {
	Result struct {
		Result  string `json:"result"`
		Matches []struct {
			Timestamp int    `json:"timestamp"`
			ID        int64  `json:"id"`
			Hash      string `json:"hash"`
			TagA      string `json:"tagA"`
			TagB      string `json:"tagB"`
			Pubkey    string `json:"pubkey"`
			Payload   string `json:"payload"`
			Hex       int    `json:"hex"`
			Senderpub string `json:"senderpub"`
			Error     string `json:"error"`
			AmountA   string `json:"amountA"`
			AmountB   string `json:"amountB"`
			Priority  int    `json:"priority"`
			Recvtime  int    `json:"recvtime"`
			Cancelled int    `json:"cancelled"`
		} `json:"matches"`
		TagA   string `json:"tagA"`
		TagB   string `json:"tagB"`
		Pubkey string `json:"pubkey"`
		N      int    `json:"n"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) DEXList(params APIParams) (DEXList, error) {

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: `DEX_list`,
		Params: string(params_json),
	}
	fmt.Println(query)

	var dexList DEXList

	dexListJson := appName.APICall(query)
	if dexListJson == "EMPTY RPC INFO!" {
		return dexList, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(dexListJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexListJson), &dexList)
		return dexList, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(dexListJson), &dexList)
	return dexList, nil
}

type DEXOrderbook struct {
	Result struct {
		Asks []struct {
			Price      string `json:"price"`
			Baseamount string `json:"baseamount"`
			Relamount  string `json:"relamount"`
			Priority   int    `json:"priority"`
			Pubkey     string `json:"pubkey"`
			Timestamp  int    `json:"timestamp"`
			Hash       string `json:"hash"`
			ID         int64  `json:"id"`
		} `json:"asks"`
		Bids []struct {
			Price      string `json:"price"`
			Baseamount string `json:"baseamount"`
			Relamount  string `json:"relamount"`
			Priority   int    `json:"priority"`
			Pubkey     string `json:"pubkey"`
			Timestamp  int    `json:"timestamp"`
			Hash       string `json:"hash"`
			ID         int64  `json:"id"`
		} `json:"bids"`
		Base string `json:"base"`
		Rel  string `json:"rel"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) DEXOrderbook(params APIParams) (DEXOrderbook, error) {

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: `DEX_orderbook`,
		Params: string(params_json),
	}
	fmt.Println(query)

	var dexOrderbook DEXOrderbook

	dexOrderbookJson := appName.APICall(query)
	if dexOrderbookJson == "EMPTY RPC INFO!" {
		return dexOrderbook, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(dexOrderbookJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexOrderbookJson), &dexOrderbook)
		return dexOrderbook, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(dexOrderbookJson), &dexOrderbook)
	return dexOrderbook, nil
}
