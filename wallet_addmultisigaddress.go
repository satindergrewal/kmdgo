package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
	"strconv"
)

type AddMultiSigAddress struct {
	Result	string	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) AddMultiSigAddress(nr int, addr string) (AddMultiSigAddress, error) {
	query := APIQuery {
		Method:	`addmultisigaddress`,
		Params:	`[`+strconv.Itoa(nr)+`, `+addr+`]`,
	}
	//fmt.Println(query)

	var addmultisigaddress AddMultiSigAddress

	addmultisigaddressJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(addmultisigaddressJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(addmultisigaddressJson), &addmultisigaddress)
		return addmultisigaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(addmultisigaddressJson), &addmultisigaddress)
	return addmultisigaddress, nil
}