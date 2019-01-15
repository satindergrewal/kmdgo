package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
	"strconv"
)

type GetRawMempool struct {
	Result map[string]RawMempoolTrue	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

type RawMempoolTrue struct {
	Size             int           `json:"size"`
	Fee              float64       `json:"fee"`
	Time             int           `json:"time"`
	Height           int           `json:"height"`
	Startingpriority float64       `json:"startingpriority"`
	Currentpriority  float64       `json:"currentpriority"`
	Depends          []interface{} `json:"depends"`
}

func (appName AppType) GetRawMempool(b bool) (GetRawMempool, error) {
	query := APIQuery {
		Method:	`getrawmempool`,
		Params:	`[`+strconv.FormatBool(b)+`]`,
	}
	//fmt.Println(query)

	var getrawmempool GetRawMempool

	getrawmempoolJson := appName.APICall(query)
	//fmt.Println(getrawmempoolJson)

	var result APIResult

	json.Unmarshal([]byte(getrawmempoolJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
		return getrawmempool, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
	return getrawmempool, nil
}