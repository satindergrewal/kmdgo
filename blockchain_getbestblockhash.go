package kmdgo

import (
    "errors"
    "encoding/json"
)

type GetBestBlockhash struct {
    Result interface{} `json:"result"`
    Error  Error       `json:"error"`
    ID     string      `json:"id"`
}

func (appName AppType) GetBestBlockhash() (GetBestBlockhash, error) {
    query := APIQuery {
        Method:     `getbestblockhash`,
        Params:   `[]`,
    }

    var getbestblockhash GetBestBlockhash

    getbestblockhashJson := appName.APICall(query)
    
    
    var result APIResult

    json.Unmarshal([]byte(getbestblockhashJson), &result)
    
    if result.Result == nil {
        answer_error, err := json.Marshal(result.Error)
        if err != nil {
        }
        json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
        return getbestblockhash, errors.New(string(answer_error))
    }

    json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
    return getbestblockhash, nil
}