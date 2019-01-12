package kmdgo

import (
	"errors"
	"encoding/json"
	"strconv"
)

type GetBlockHash struct {
	Result	string	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}


func (appName AppType) GetBlockHash(h int) (GetBlockHash, error) {
	query := APIQuery {
		Method:	`getblockhash`,
		Params:	`[`+strconv.Itoa(h)+`]`,
	}
	//fmt.Println(query)

	var getblockhash GetBlockHash

	getblockhashJson := appName.APICall(query)


	var result APIResult

	json.Unmarshal([]byte(getblockhashJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockhashJson), &getblockhash)
		return getblockhash, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblockhashJson), &getblockhash)
	return getblockhash, nil
}