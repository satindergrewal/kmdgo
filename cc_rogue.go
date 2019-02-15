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

const (
	CCLIB_METHOD = `cclib`
	ROGUE_EVALCODE = `17`
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

// cclib method has these command line options:
// 		$ komodo-cli -ac_name=ROGUE help cclib 
// 		cclib method [evalcode] [JSON params]
// 
// Example command
//		komodo-cli -ac_name=ROGUE cclib newgame 17 "[1]"
//
// Explantion
//		Command line Utility: komodo-cli
//		Assetchain name parameter: -ac_name=ROGUE
//		Command Method: cclib
//		Command Sub-Method: newgame
//		evalcode used by the Crypto-Condition: 17
//		JSON Parameters passed for this crypto-conditions as string value: "[1]"
func (appName AppType) RGNewGame(params APIParams) (RGNewGame, error) {

	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `newgame`
		}
	}
	
	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = ROGUE_EVALCODE
		}
	}
	
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	CCLIB_METHOD,
		Params:	string(params_json),
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
		Name       string  `json:"name"`
		Method     string  `json:"method"`
		Gametxid   string  `json:"gametxid"`
		Result     string  `json:"result"`
		Gameheight int     `json:"gameheight"`
		Height     int     `json:"height"`
		Start      int     `json:"start"`
		Starthash  string  `json:"starthash"`
		Seed       int64   `json:"seed"`
		Run        string  `json:"run"`
		Alive      int     `json:"alive"`
		Numplayers int     `json:"numplayers"`
		Maxplayers int     `json:"maxplayers"`
		Buyin      float64 `json:"buyin"`
		Players    []struct {
			Slot       int     `json:"slot"`
			Status     string  `json:"status"`
			Baton      string  `json:"baton"`
			Tokenid    string  `json:"tokenid"`
			Batonaddr  string  `json:"batonaddr"`
			Ismine     bool    `json:"ismine"`
			Batonvout  int     `json:"batonvout"`
			Batonvalue float64 `json:"batonvalue"`
			Batonht    int     `json:"batonht"`
		} `json:"players"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) RGGameInfo(params APIParams) (RGGameInfo, error) {
		if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `gameinfo`
		}
	}
	
	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = ROGUE_EVALCODE
		}
	}
	
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery {
		Method:	CCLIB_METHOD,
		Params:	string(params_json),
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
		Txid       string  `json:"txid,omitempty"`
		Result     string  `json:"result,omitempty"`
		Status     string  `json:"status,omitempty"`
		Error      string  `json:"error,omitempty"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) RGRegister(params APIParams) (RGRegister, error) {
		if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `register`
		}
	}
	
	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = ROGUE_EVALCODE
		}
	}
	
	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery {
		Method:	CCLIB_METHOD,
		Params:	string(params_json),
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


type RGPending struct {
	Result struct {
		Result     string   `json:"result"`
		Name       string   `json:"name"`
		Method     string   `json:"method"`
		Pending    []string `json:"pending"`
		Numpending int      `json:"numpending"`
	} `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) RGPending() (RGPending, error) {
	var params APIParams
	params[0] = `pending`
	params[1] = ROGUE_EVALCODE

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery {
		Method:	CCLIB_METHOD,
		Params:	string(params_json),
	}
	//fmt.Println(query)


	var pending RGPending

	pendingJson := appName.APICall(query)
	//fmt.Println(pendingJson)

	var result APIResult

	json.Unmarshal([]byte(pendingJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(pendingJson), &pending)
		return pending, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(pendingJson), &pending)
	return pending, nil
}