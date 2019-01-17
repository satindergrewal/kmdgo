package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
	"strconv"
)

type GetTxOut struct {
	Result	struct {
		Bestblock        string  `json:"bestblock"`
		Confirmations    int     `json:"confirmations"`
		Rawconfirmations int     `json:"rawconfirmations"`
		Value            float64 `json:"value"`
		ScriptPubKey     struct {
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
			Addresses []string `json:"addresses"`
		} `json:"scriptPubKey"`
		Version  int  `json:"version"`
		Coinbase bool `json:"coinbase"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) GetTxOut(txid string, n int) (GetTxOut, error) {
	query := APIQuery {
		Method:	`gettxout`,
		Params:	`["`+txid+`", `+strconv.Itoa(n)+`]`,
	}
	//fmt.Println(query)

	var gettxout GetTxOut

	gettxoutJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(gettxoutJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettxoutJson), &gettxout)
		return gettxout, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(gettxoutJson), &gettxout)
	return gettxout, nil
}