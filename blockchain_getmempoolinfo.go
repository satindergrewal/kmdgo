package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type GetMempoolInfo struct {
	Result struct {
		Size  int `json:"size"`
		Bytes int `json:"bytes"`
		Usage int `json:"usage"`
	}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) GetMempoolInfo() (GetMempoolInfo, error) {
	query := APIQuery {
		Method:	`getmempoolinfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var getmempoolinfo GetMempoolInfo

	getmempoolinfoJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getmempoolinfoJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getmempoolinfoJson), &getmempoolinfo)
		return getmempoolinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getmempoolinfoJson), &getmempoolinfo)
	return getmempoolinfo, nil
}