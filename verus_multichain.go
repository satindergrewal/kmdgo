// Copyright © 2018-2020 Satinderjit Singh.
//
// See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
// the top-level directory of this distribution for the individual copyright
// holder information and the developer policies on copyright and licensing.
//
// Unless otherwise agreed in a custom licensing agreement, no part of the
// kmdgo software, including this file may be copied, modified, propagated.
// or distributed except according to the terms contained in the LICENSE file
//
// Removal or modification of this copyright notice is prohibited.

package kmdgo

import (
	"encoding/json"
	"errors"
)

// GetCurrency type
type GetCurrency struct {
	Result struct {
		Name                 string        `json:"name"`
		Version              int           `json:"version"`
		Options              int           `json:"options"`
		Parent               string        `json:"parent"`
		Systemid             string        `json:"systemid"`
		Currencyid           string        `json:"currencyid"`
		Notarizationprotocol int           `json:"notarizationprotocol"`
		Proofprotocol        int           `json:"proofprotocol"`
		Idregistrationprice  int           `json:"idregistrationprice"`
		Idreferrallevels     int           `json:"idreferrallevels"`
		Minnotariesconfirm   int           `json:"minnotariesconfirm"`
		Billingperiod        int           `json:"billingperiod"`
		Notarizationreward   int           `json:"notarizationreward"`
		Startblock           int           `json:"startblock"`
		Endblock             int           `json:"endblock"`
		Currencies           []string      `json:"currencies"`
		Weights              []float64     `json:"weights"`
		Conversions          []float64     `json:"conversions"`
		Initialsupply        float64       `json:"initialsupply"`
		Prelaunchcarveout    float64       `json:"prelaunchcarveout"`
		Initialcontributions []float64     `json:"initialcontributions"`
		Preconversions       []float64     `json:"preconversions"`
		Eras                 []interface{} `json:"eras"`
		Definitiontxid       string        `json:"definitiontxid"`
		Bestcurrencystate    struct {
			Flags             int    `json:"flags"`
			Currencyid        string `json:"currencyid"`
			Reservecurrencies []struct {
				Currencyid     string  `json:"currencyid"`
				Weight         float64 `json:"weight"`
				Reserves       float64 `json:"reserves"`
				Priceinreserve float64 `json:"priceinreserve"`
			} `json:"reservecurrencies"`
			Initialsupply        float64               `json:"initialsupply"`
			Emitted              float64               `json:"emitted"`
			Supply               float64               `json:"supply"`
			Currencies           map[string]Currencies `json:"currencies"`
			Nativefees           int                   `json:"nativefees"`
			Nativeconversionfees int                   `json:"nativeconversionfees"`
		} `json:"bestcurrencystate"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// Currencies type
type Currencies struct {
	Reservein           float64 `json:"reservein"`
	Nativein            float64 `json:"nativein"`
	Reserveout          float64 `json:"reserveout"`
	Lastconversionprice float64 `json:"lastconversionprice"`
	Viaconversionprice  float64 `json:"viaconversionprice"`
	Fees                float64 `json:"fees"`
	Conversionfees      float64 `json:"conversionfees"`
}

// GetCurrency returns a complete definition for any given chain if it is registered on the blockchain. If the chain requested is NULL, chain definition of the current chain is returned.
//
// getcurrency "chainname"
//
// Arguments
// 1. "chainname"                     (string, optional) name of the chain to look for. no parameter returns current chain in daemon.
func (appName AppType) GetCurrency(params APIParams) (GetCurrency, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getcurrency`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var GetCrncy GetCurrency

	GetCrncyJSON := appName.APICall(&query)
	if GetCrncyJSON == "EMPTY RPC INFO" {
		return GetCrncy, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(GetCrncyJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(GetCrncyJSON), &GetCrncy)
		return GetCrncy, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(GetCrncyJSON), &GetCrncy)
	return GetCrncy, nil
}

// GetCurrencyState type
type GetCurrencyState struct {
	Result struct {
		Height        int `json:"height"`
		Blocktime     int `json:"blocktime"`
		Currencystate struct {
			Flags         int     `json:"flags"`
			Currencyid    string  `json:"currencyid"`
			Initialsupply float64 `json:"initialsupply"`
			Emitted       float64 `json:"emitted"`
			Supply        float64 `json:"supply"`
			Currencies    struct {
			} `json:"currencies"`
			Nativefees           int `json:"nativefees"`
			Nativeconversionfees int `json:"nativeconversionfees"`
		} `json:"currencystate"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetCurrencyState returns the total amount of preconversions that have been confirmed on the blockchain for the specified chain.
//
// getcurrencystate "n"
//
// Arguments
//    "n" or "m,n" or "m,n,o"         (int or string, optional) height or inclusive range with optional step at which to get the currency state. If not specified, the latest currency state and height is returned
//
func (appName AppType) GetCurrencyState(params APIParams) (GetCurrencyState, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getcurrencystate`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var GetCurSt GetCurrencyState

	GetCurStJSON := appName.APICall(&query)
	if GetCurStJSON == "EMPTY RPC INFO" {
		return GetCurSt, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(GetCurStJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(GetCurStJSON), &GetCurSt)
		return GetCurSt, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(GetCurStJSON), &GetCurSt)
	return GetCurSt, nil
}
