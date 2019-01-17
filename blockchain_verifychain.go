/*
The `verifychain` method verifies the coin daemon's blockchain database.
*/

package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
	"strconv"
)

type VerifyChain struct {
	Result	bool `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) VerifyChain(cl int, num int) (VerifyChain, error) {
	query := APIQuery {
		Method:	`verifychain`,
		Params:	`[`+strconv.Itoa(cl)+`, `+strconv.Itoa(num)+`]`,
	}
	//fmt.Println(query)

	var verifychain VerifyChain

	verifychainJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(verifychainJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(verifychainJson), &verifychain)
		return verifychain, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(verifychainJson), &verifychain)
	return verifychain, nil
}