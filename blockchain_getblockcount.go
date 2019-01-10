package kmdgo

import (
    "errors"
    "encoding/json"
)

type GetBlockCount struct {
    Result interface{} `json:"result"`
    Error  Error       `json:"error"`
    ID     string      `json:"id"`
}

func (appName AppType) GetBlockCount() (GetBlockCount, error) {
    query := APIQuery {
        Method:     `getblockcount`,
        Params:   `[]`,
    }

    var getblockcount GetBlockCount

    getbestblockhashJson := appName.APICall(query)
    
    
    var result APIResult

    json.Unmarshal([]byte(getbestblockhashJson), &result)
    
    if result.Result == nil {
        answer_error, err := json.Marshal(result.Error)
        if err != nil {
        }
        json.Unmarshal([]byte(getbestblockhashJson), &getblockcount)
        return getblockcount, errors.New(string(answer_error))
    }

    json.Unmarshal([]byte(getbestblockhashJson), &getblockcount)
    return getblockcount, nil
}