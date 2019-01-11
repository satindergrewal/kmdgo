package kmdgo

import (
	//"fmt"
	"strconv"
	"errors"
	"encoding/json"
)

type CoinSupply struct {
	Result struct {
		Result	string	`json:"result"`
		Coin	string 	`json:"coin"`
		Height	int		`json:"height"`
		Supply	float64	`json:"supply"`
		Zfunds	float64	`json:"zfunds"`
		Sprout	float64	`json:"sprout"`
		Total	float64	`json:"total"`
	}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) CoinSupply(n int) (CoinSupply, error) {
    query := APIQuery {
        Method:     `coinsupply`,
        Params:   `["`+strconv.Itoa(n)+`"]`,
    }
    //fmt.Println(query)

    var coinsupply CoinSupply

    coinsupplyJson := appName.APICall(query)
    
    
    var result APIResult

    json.Unmarshal([]byte(coinsupplyJson), &result)
    
    if result.Result == nil {
        answer_error, err := json.Marshal(result.Error)
        if err != nil {
        }
        json.Unmarshal([]byte(coinsupplyJson), &coinsupply)
        return coinsupply, errors.New(string(answer_error))
    }

    json.Unmarshal([]byte(coinsupplyJson), &coinsupply)
    return coinsupply, nil
}