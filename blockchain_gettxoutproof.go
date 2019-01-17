package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type GetTxOutProof struct {
	Result	string	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) GetTxOutProof(txids string) (GetTxOutProof, error) {
	query := APIQuery {
		Method:	`gettxoutproof`,
		Params:	`[`+txids+`]`,
	}
	//fmt.Println(query)

	var gettxoutproof GetTxOutProof

	gettxoutproofJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(gettxoutproofJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettxoutproofJson), &gettxoutproof)
		return gettxoutproof, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(gettxoutproofJson), &gettxoutproof)
	return gettxoutproof, nil
}