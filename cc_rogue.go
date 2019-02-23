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

const (
	CCLIB_METHOD   = `cclib`
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
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

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
// newgame: Crypto-Conditions (CC) specific method. You can see all available methods via cclib with following command
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
			params[1] = ROGUE_EVALCODE
		}
	}

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
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
	Error Error  `json:"error"`
	ID    string `json:"id"`
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

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
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
	Error Error  `json:"error"`
	ID    string `json:"id"`
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

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
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
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) RGPending() (RGPending, error) {
	params := make(APIParams, 2)
	params[0] = `pending`
	params[1] = ROGUE_EVALCODE

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
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

func (appName AppType) RGBailout(params APIParams) (RGBailout, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `bailout`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = ROGUE_EVALCODE
		}
	}

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var bail RGBailout

	bailJson := appName.APICall(query)
	//fmt.Println(bailJson)

	var result APIResult

	json.Unmarshal([]byte(bailJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(bailJson), &bail)
		return bail, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(bailJson), &bail)
	return bail, nil
}

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

func (appName AppType) RGHighLander(params APIParams) (RGHighLander, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `highlander`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = ROGUE_EVALCODE
		}
	}

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var hland RGHighLander

	hlandJson := appName.APICall(query)
	//fmt.Println(hlandJson)

	var result APIResult

	json.Unmarshal([]byte(hlandJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(hlandJson), &hland)
		return hland, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(hlandJson), &hland)
	return hland, nil
}

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

func (appName AppType) RGPlayers() (RGPlayers, error) {
	params := make(APIParams, 2)
	params[0] = `players`
	params[1] = ROGUE_EVALCODE

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var players RGPlayers

	playersJson := appName.APICall(query)
	//fmt.Println(playersJson)

	var result APIResult

	json.Unmarshal([]byte(playersJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(playersJson), &players)
		return players, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(playersJson), &players)
	return players, nil
}

type RGSetName struct {
	Result struct {
		Name   string `json:"name"`
		Method string `json:"method"`
		Result string `json:"result"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) RGSetName(params APIParams) (RGSetName, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `setname`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = ROGUE_EVALCODE
		}
	}

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var pname RGSetName

	pnameJson := appName.APICall(query)
	//fmt.Println(pnameJson)

	var result APIResult

	json.Unmarshal([]byte(pnameJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(pnameJson), &pname)
		return pname, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(pnameJson), &pname)
	return pname, nil
}

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

func (appName AppType) RGPlayerInfo(params APIParams) (RGPlayerInfo, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `playerinfo`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = ROGUE_EVALCODE
		}
	}

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var pinfo RGPlayerInfo

	pinfoJson := appName.APICall(query)
	//fmt.Println(pinfoJson)

	var result APIResult

	json.Unmarshal([]byte(pinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(pinfoJson), &pinfo)
		return pinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(pinfoJson), &pinfo)
	return pinfo, nil
}

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

func (appName AppType) RGGames() (RGGames, error) {
	params := make(APIParams, 2)

	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = `games`
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = ROGUE_EVALCODE
		}
	}

	params_json, _ := json.Marshal(params)
	fmt.Println(string(params_json))

	query := APIQuery{
		Method: CCLIB_METHOD,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var gms RGGames

	gmsJson := appName.APICall(query)
	//fmt.Println(gmsJson)

	var result APIResult

	json.Unmarshal([]byte(gmsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gmsJson), &gms)
		return gms, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(gmsJson), &gms)
	return gms, nil
}
