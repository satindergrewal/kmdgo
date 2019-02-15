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
	"fmt"
	"encoding/json"
	"errors"
)

type RGNewGame struct {
	Result struct {
		Name       string  `json:"name"`
		Method     string  `json:"method"`
		Maxplayers int     `json:"maxplayers"`
		Buyin      float64 `json:"buyin"`
		Type       string  `json:"type"`
		Hex        string  `json:"hex"`
		Txid       string  `json:"txid"`
		Result     string  `json:"result"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) RGNewGame(params APIParams) (RGNewGame, error) {
	var params_json string

	if params[0] == "" || params[0] == nil {
		params_json = `[]`
		//fmt.Println(params_json)
	} else {
		params_bytes, _ := json.Marshal(params)
		params_json = string(params_bytes)
		fmt.Println(params_json)
	}

	query := APIQuery {
		Method:	`cclib`,
		Params:	params_json,
	}
	//fmt.Println(query)

	var newgame RGNewGame

	newgameJson := appName.APICall(query)
	//fmt.Println(newgameJson)

	var result APIResult

	json.Unmarshal([]byte(newgameJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(newgameJson), &newgame)
		return newgame, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(newgameJson), &newgame)
	return newgame, nil
}


type RGGameInfo struct {
	Result struct {
		Name       string        `json:"name"`
		Method     string        `json:"method"`
		Gametxid   string        `json:"gametxid"`
		Result     string        `json:"result"`
		Gameheight int           `json:"gameheight"`
		Height     int           `json:"height"`
		Start      int           `json:"start"`
		Starthash  string        `json:"starthash"`
		Seed       int64         `json:"seed"`
		Run        string        `json:"run"`
		Alive      int           `json:"alive"`
		Numplayers int           `json:"numplayers"`
		Maxplayers int           `json:"maxplayers"`
		Buyin      float64       `json:"buyin"`
		Players    []interface{} `json:"players"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) RGGameInfo(params APIParams) (RGGameInfo, error) {
	var params_json string

	if params[0] == "" || params[0] == nil {
		params_json = `[]`
		//fmt.Println(params_json)
	} else {
		params_bytes, _ := json.Marshal(params)
		params_json = string(params_bytes)
		fmt.Println(params_json)
	}

	query := APIQuery {
		Method:	`cclib`,
		Params:	params_json,
	}
	//fmt.Println(query)

	var gameinfo RGGameInfo

	gameinfoJson := appName.APICall(query)
	//fmt.Println(gameinfoJson)

	var result APIResult

	json.Unmarshal([]byte(gameinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gameinfoJson), &gameinfo)
		return gameinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(gameinfoJson), &gameinfo)
	return gameinfo, nil
}



type RGRegister struct {
	Result struct {
		Name       string  `json:"name"`
		Method     string  `json:"method"`
		Maxplayers int     `json:"maxplayers"`
		Buyin      float64 `json:"buyin"`
		Type       string  `json:"type"`
		Hex        string  `json:"hex"`
		Txid       string  `json:"txid"`
		Result     string  `json:"result"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) RGRegister(params APIParams) (RGRegister, error) {
	var params_json string

	if params[0] == "" || params[0] == nil {
		params_json = `[]`
		//fmt.Println(params_json)
	} else {
		params_bytes, _ := json.Marshal(params)
		params_json = string(params_bytes)
		fmt.Println(params_json)
	}

	query := APIQuery {
		Method:	`cclib`,
		Params:	params_json,
	}
	//fmt.Println(query)

	var reg RGRegister

	regJson := appName.APICall(query)
	//fmt.Println(regJson)

	var result APIResult

	json.Unmarshal([]byte(regJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(regJson), &reg)
		return reg, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(regJson), &reg)
	return reg, nil
}