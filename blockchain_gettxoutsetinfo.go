package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type GetTxOutSetInfo struct {
	Result	struct {
		Height          int     `json:"height"`
		Bestblock       string  `json:"bestblock"`
		Transactions    int     `json:"transactions"`
		Txouts          int     `json:"txouts"`
		BytesSerialized int     `json:"bytes_serialized"`
		HashSerialized  string  `json:"hash_serialized"`
		TotalAmount     float64 `json:"total_amount"`
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) GetTxOutSetInfo() (GetTxOutSetInfo, error) {
	query := APIQuery {
		Method:	`gettxoutsetinfo`,
		Params:	`[]`,
	}
	//fmt.Println(query)

	var gettxoutsetinfo GetTxOutSetInfo

	gettxoutsetinfoJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(gettxoutsetinfoJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettxoutsetinfoJson), &gettxoutsetinfo)
		return gettxoutsetinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(gettxoutsetinfoJson), &gettxoutsetinfo)
	return gettxoutsetinfo, nil
}