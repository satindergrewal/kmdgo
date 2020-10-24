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

const (
	ccLibMethod   = `cclib`
	rogueEvalCode = `17`
)

// RGNewGame type
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
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGNewGame - Crypto-Conditions (CC) specific method. You can see all available methods via cclib with following command
// Method: newgame
// Params: maxplayers buyin
// Params Required: 0
// Max Parameters 2
//
// Example command
// 		 ./komodo-cli -ac_name=ROGUE cclib newgame 17 \"[3,10]\"
//
// Explanation
//
// cclib: API call method.
//
//		./komodo-cli -ac_name=ROGUE cclibinfo
//
// 17: Evalcode for this CC method
//
// Array value 3: value of `maxplayers`. The maximum players that can participate in this game.
//
// Array value 10: value of `buyin`. Each new player participating in this game need to put 10 ROGUE coins to participate.
//
// The winner will be the player who will be last standing, i.e. not killed, while others are killed.
//
// And this winner takes all the buyin total of 3 players, which is equals to 30 ROGUE, also the GOLD converted to ROGUE from all these players.
func (appName AppType) RGNewGame(params APIParams) (RGNewGame, error) {

	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `newgame`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = rogueEvalCode
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var newgame RGNewGame

	newgameJSON := appName.APICall(&query)
	if newgameJSON == "EMPTY RPC INFO" {
		return newgame, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(newgameJSON)

	var result APIResult

	json.Unmarshal([]byte(newgameJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(newgameJSON), &newgame)
		return newgame, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(newgameJSON), &newgame)
	return newgame, nil
}

// RGGameInfo type
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
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGGameInfo ...
func (appName AppType) RGGameInfo(params APIParams) (RGGameInfo, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `gameinfo`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = rogueEvalCode
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var gameinfo RGGameInfo

	gameinfoJSON := appName.APICall(&query)
	if gameinfoJSON == "EMPTY RPC INFO" {
		return gameinfo, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(gameinfoJSON)

	var result APIResult

	json.Unmarshal([]byte(gameinfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gameinfoJSON), &gameinfo)
		return gameinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(gameinfoJSON), &gameinfo)
	return gameinfo, nil
}

// RGRegister type
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
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGRegister ...
func (appName AppType) RGRegister(params APIParams) (RGRegister, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `register`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = rogueEvalCode
		}
	}

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var reg RGRegister

	regJSON := appName.APICall(&query)
	if regJSON == "EMPTY RPC INFO" {
		return reg, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(regJSON)

	var result APIResult

	json.Unmarshal([]byte(regJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(regJSON), &reg)
		return reg, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(regJSON), &reg)
	return reg, nil
}

// RGPending type
type RGPending struct {
	Result struct {
		Result     string   `json:"result"`
		Name       string   `json:"name"`
		Method     string   `json:"method"`
		Pending    []string `json:"pending"`
		Numpending int      `json:"numpending"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGPending ...
func (appName AppType) RGPending() (RGPending, error) {
	params := make(APIParams, 2)
	params[0] = `pending`
	params[1] = rogueEvalCode

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var pending RGPending

	pendingJSON := appName.APICall(&query)
	if pendingJSON == "EMPTY RPC INFO" {
		return pending, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(pendingJSON)

	var result APIResult

	json.Unmarshal([]byte(pendingJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(pendingJSON), &pending)
		return pending, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(pendingJSON), &pending)
	return pending, nil
}

// RGBailout type
type RGBailout struct {
	Result struct {
		Name        string `json:"name"`
		Method      string `json:"method"`
		Myrogueaddr string `json:"myrogueaddr"`
		Gametxid    string `json:"gametxid"`
		Hex         string `json:"hex"`
		Txid        string `json:"txid"`
		Result      string `json:"result"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGBailout ...
func (appName AppType) RGBailout(params APIParams) (RGBailout, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `bailout`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = rogueEvalCode
		}
	}

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var bail RGBailout

	bailJSON := appName.APICall(&query)
	if bailJSON == "EMPTY RPC INFO" {
		return bail, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(bailJSON)

	var result APIResult

	json.Unmarshal([]byte(bailJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(bailJSON), &bail)
		return bail, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(bailJSON), &bail)
	return bail, nil
}

// RGHighLander type
type RGHighLander struct {
	Result struct {
		Name        string `json:"name"`
		Method      string `json:"method"`
		Myrogueaddr string `json:"myrogueaddr"`
		Gametxid    string `json:"gametxid"`
		Hex         string `json:"hex"`
		Txid        string `json:"txid"`
		Result      string `json:"result"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGHighLander ...
func (appName AppType) RGHighLander(params APIParams) (RGHighLander, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `highlander`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = rogueEvalCode
		}
	}

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var hland RGHighLander

	hlandJSON := appName.APICall(&query)
	if hlandJSON == "EMPTY RPC INFO" {
		return hland, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(hlandJSON)

	var result APIResult

	json.Unmarshal([]byte(hlandJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(hlandJSON), &hland)
		return hland, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(hlandJSON), &hland)
	return hland, nil
}

// RGPlayers type
type RGPlayers struct {
	Result struct {
		Name          string   `json:"name"`
		Method        string   `json:"method"`
		Playerdata    []string `json:"playerdata"`
		Numplayerdata int      `json:"numplayerdata"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGPlayers ...
func (appName AppType) RGPlayers() (RGPlayers, error) {
	params := make(APIParams, 2)
	params[0] = `players`
	params[1] = rogueEvalCode

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var players RGPlayers

	playersJSON := appName.APICall(&query)
	if playersJSON == "EMPTY RPC INFO" {
		return players, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(playersJSON)

	var result APIResult

	json.Unmarshal([]byte(playersJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(playersJSON), &players)
		return players, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(playersJSON), &players)
	return players, nil
}

// RGSetName type
type RGSetName struct {
	Result struct {
		Name   string `json:"name"`
		Method string `json:"method"`
		Result string `json:"result"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGSetName ...
func (appName AppType) RGSetName(params APIParams) (RGSetName, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `setname`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = rogueEvalCode
		}
	}

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var pname RGSetName

	pnameJSON := appName.APICall(&query)
	if pnameJSON == "EMPTY RPC INFO" {
		return pname, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(pnameJSON)

	var result APIResult

	json.Unmarshal([]byte(pnameJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(pnameJSON), &pname)
		return pname, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(pnameJSON), &pname)
	return pname, nil
}

// RGPlayerInfo type
type RGPlayerInfo struct {
	Result struct {
		Result string `json:"result"`
		Name   string `json:"name"`
		Method string `json:"method"`
		Player struct {
			Playertxid   string   `json:"playertxid"`
			Tokenid      string   `json:"tokenid"`
			Data         string   `json:"data"`
			Pack         []string `json:"pack"`
			Packsize     int      `json:"packsize"`
			Hitpoints    int      `json:"hitpoints"`
			Strength     int      `json:"strength"`
			Level        int      `json:"level"`
			Experience   int      `json:"experience"`
			Dungeonlevel int      `json:"dungeonlevel"`
			Chain        string   `json:"chain"`
			Pname        string   `json:"pname"`
		} `json:"player"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGPlayerInfo ...
func (appName AppType) RGPlayerInfo(params APIParams) (RGPlayerInfo, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `playerinfo`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = rogueEvalCode
		}
	}

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var pinfo RGPlayerInfo

	pinfoJSON := appName.APICall(&query)
	if pinfoJSON == "EMPTY RPC INFO" {
		return pinfo, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(pinfoJSON)

	var result APIResult

	json.Unmarshal([]byte(pinfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(pinfoJSON), &pinfo)
		return pinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(pinfoJSON), &pinfo)
	return pinfo, nil
}

// RGGames type
type RGGames struct {
	Result struct {
		Name      string   `json:"name"`
		Method    string   `json:"method"`
		Pastgames []string `json:"pastgames"`
		Games     []string `json:"games"`
		Numgames  int      `json:"numgames"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// RGGames ...
func (appName AppType) RGGames() (RGGames, error) {
	params := make(APIParams, 2)

	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `games`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = rogueEvalCode
		}
	}

	paramsJSON, _ := json.Marshal(params)
	fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: ccLibMethod,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var gms RGGames

	gmsJSON := appName.APICall(&query)
	if gmsJSON == "EMPTY RPC INFO" {
		return gms, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(gmsJSON)

	var result APIResult

	json.Unmarshal([]byte(gmsJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gmsJSON), &gms)
		return gms, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(gmsJSON), &gms)
	return gms, nil
}
