package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type GetDifficulty struct {
	Result float64	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetDifficulty() (GetDifficulty, error) {
	query := APIQuery {
		Method:	`getdifficulty`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getdifficulty GetDifficulty

	getdifficultyJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getdifficultyJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getdifficultyJson), &getdifficulty)
		return getdifficulty, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getdifficultyJson), &getdifficulty)
	return getdifficulty, nil
}