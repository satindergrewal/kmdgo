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
	"fmt"
	"log"
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

// CurrencyInfo type
type CurrencyInfo struct {
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
}

// GetCurrencyConverter type
type GetCurrencyConverter struct {
	CurrencyInfo     CurrencyInfo `json:"-"`
	Lastnotarization struct {
		Version             int    `json:"version"`
		Currencyid          string `json:"currencyid"`
		Notaryaddress       string `json:"notaryaddress"`
		Notarizationheight  int    `json:"notarizationheight"`
		Mmrroot             string `json:"mmrroot"`
		Notarizationprehash string `json:"notarizationprehash"`
		Work                string `json:"work"`
		Stake               string `json:"stake"`
		Currencystate       struct {
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
			Nativefees           int64                 `json:"nativefees"`
			Nativeconversionfees int64                 `json:"nativeconversionfees"`
		} `json:"currencystate"`
		Prevnotarization  string        `json:"prevnotarization"`
		Prevheight        int           `json:"prevheight"`
		Crossnotarization string        `json:"crossnotarization"`
		Crossheight       int           `json:"crossheight"`
		Nodes             []interface{} `json:"nodes"`
	} `json:"lastnotarization"`
	Multifractional struct {
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
	} `json:"multifractional,omitempty"`
}

// GetCurrencyConverters array type
type GetCurrencyConverters struct {
	Result []GetCurrencyConverter `json:"result"`
	Error  Error                  `json:"error"`
	ID     string                 `json:"id"`
}

// GetCurrencyConverters Retrieves all currencies that have at least 1000 VRSC in reserve, are >10% VRSC reserve ratio, and have all listed currencies as reserves
//
// getcurrencyconverters currency1 currency2
// Arguments
//        ["currencyname"    : "string", ...]  (string list, one or more) all selected currencies are returned with their current state
// Result:
//        "[{currency1}, {currency2}]" : "array of objects" (string) All currencies and the last notarization, which are valid converters.
//
// Examples:
// > verus getcurrencyconverters currency1 currency2 ...
func (appName AppType) GetCurrencyConverters(params APIParams) (GetCurrencyConverters, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getcurrencyconverters`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var GetCurCovrts GetCurrencyConverters

	GetCurCovrtsJSON := appName.APICall(&query)
	if GetCurCovrtsJSON == "EMPTY RPC INFO" {
		return GetCurCovrts, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(GetCurCovrtsJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(GetCurCovrtsJSON), &GetCurCovrts)
		return GetCurCovrts, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(GetCurCovrtsJSON), &GetCurCovrts)
	// fmt.Println(GetCurCovrtsJSON)

	var f interface{}
	// m := f.(map[string]interface{})
	err := json.Unmarshal([]byte(GetCurCovrtsJSON), &f)
	if err != nil {
		log.Printf("%v", err)
	}
	// fmt.Printf("%+v\n", f)

	// fmt.Printf("%T\n", f)

	m := f.(map[string]interface{})

	var n interface{}
	for k, v := range m {
		// fmt.Printf("k -- %+v\n", k)
		// fmt.Println("v --- ", v)
		if k == "result" {
			n = v
		}
	}
	// fmt.Println("n -- ", n)

	o := n.([]interface{})

	for k, v := range o {
		// fmt.Printf("k -- %+v\n", k)
		// fmt.Println("v --- ", v)
		if _, ok := v.(map[string]interface{})["result"]; ok {

		}
		p := v.(map[string]interface{})
		if _, ok := v.(map[string]interface{})["lastnotarization"]; ok {
			// fmt.Println("lastnotarization ---", val)
			delete(v.(map[string]interface{}), "lastnotarization")
		}
		if _, ok := v.(map[string]interface{})["multifractional"]; ok {
			// fmt.Println("multifractional ---", val)
			delete(v.(map[string]interface{}), "multifractional")
		}
		for pk, pv := range p {
			// fmt.Printf("pk - %+v\n", pk)
			// fmt.Printf("pv - %T\n", pv)
			switch vv := pv.(type) {
			case string:
				fmt.Println(pk, "is string", vv)
			case float64:
				fmt.Println(pk, "is float64", vv)
			case []interface{}:
				fmt.Println(pk, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				// fmt.Println(pk, "is of a type I don't know how to handle")
				// fmt.Printf("%T\n", vv)
				// fmt.Printf("vv -- %+v\n", vv)
				if val, ok := vv.(map[string]interface{})["name"]; ok {
					// fmt.Println("name ---", val)
					GetCurCovrts.Result[0].CurrencyInfo.Name = val.(string)
				}
				if val, ok := vv.(map[string]interface{})["version"]; ok {
					// fmt.Println("version ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Version = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["options"]; ok {
					// fmt.Println("options ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Options = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["parent"]; ok {
					// fmt.Println("parent ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Parent = val.(string)
				}
				if val, ok := vv.(map[string]interface{})["systemid"]; ok {
					// fmt.Println("systemid ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Systemid = val.(string)
				}
				if val, ok := vv.(map[string]interface{})["currencyid"]; ok {
					// fmt.Println("currencyid ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Currencyid = val.(string)
				}
				if val, ok := vv.(map[string]interface{})["notarizationprotocol"]; ok {
					// fmt.Println("notarizationprotocol ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Notarizationprotocol = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["proofprotocol"]; ok {
					// fmt.Println("proofprotocol ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Proofprotocol = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["idregistrationprice"]; ok {
					// fmt.Println("idregistrationprice ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Idregistrationprice = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["idreferrallevels"]; ok {
					// fmt.Println("idreferrallevels ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Idreferrallevels = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["minnotariesconfirm"]; ok {
					// fmt.Println("minnotariesconfirm ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Minnotariesconfirm = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["billingperiod"]; ok {
					// fmt.Println("billingperiod ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Billingperiod = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["notarizationreward"]; ok {
					// fmt.Println("notarizationreward ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Notarizationreward = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["startblock"]; ok {
					// fmt.Println("startblock ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Startblock = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["endblock"]; ok {
					// fmt.Println("endblock ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Endblock = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["currencies"]; ok {
					// fmt.Printf("currencies type --- %T\n", val)
					// fmt.Println("currencies ---", val)
					var _tmpCurrencies []string
					for _, curv := range val.([]interface{}) {
						// fmt.Println(curv)
						// fmt.Printf("curv -- %T\n", curv)
						_tmpCurrencies = append(_tmpCurrencies, fmt.Sprintf("%v", curv))
					}
					// fmt.Printf("%T\n", _tmpCurrencies)
					// fmt.Printf("%+v\n", _tmpCurrencies)
					GetCurCovrts.Result[k].CurrencyInfo.Currencies = _tmpCurrencies
				}
				if val, ok := vv.(map[string]interface{})["weights"]; ok {
					// fmt.Println("weights ---", val)
					var _tmpWeights []float64
					for _, wghtv := range val.([]interface{}) {
						// fmt.Println(wghtv)
						// fmt.Printf("wghtv -- %T\n", wghtv)
						_tmpWeights = append(_tmpWeights, wghtv.(float64))
					}
					// fmt.Printf("%T\n", _tmpWeights)
					// fmt.Printf("%+v\n", _tmpWeights)
					GetCurCovrts.Result[k].CurrencyInfo.Weights = _tmpWeights
				}
				if val, ok := vv.(map[string]interface{})["conversions"]; ok {
					// fmt.Println("conversions ---", val)
					var _tmpConversions []float64
					for _, cnvrsv := range val.([]interface{}) {
						// fmt.Println(cnvrsv)
						// fmt.Printf("cnvrsv -- %T\n", cnvrsv)
						_tmpConversions = append(_tmpConversions, cnvrsv.(float64))
					}
					// fmt.Printf("%T\n", _tmpConversions)
					// fmt.Printf("%+v\n", _tmpConversions)
					GetCurCovrts.Result[k].CurrencyInfo.Conversions = _tmpConversions
				}
				if val, ok := vv.(map[string]interface{})["initialsupply"]; ok {
					// fmt.Println("initialsupply ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Initialsupply = val.(float64)
				}
				if val, ok := vv.(map[string]interface{})["prelaunchcarveout"]; ok {
					// fmt.Println("prelaunchcarveout ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Prelaunchcarveout = val.(float64)
				}
				if val, ok := vv.(map[string]interface{})["initialcontributions"]; ok {
					// fmt.Println("initialcontributions ---", val)
					var _tmpInitContri []float64
					for _, initcontv := range val.([]interface{}) {
						// fmt.Println(initcontv)
						// fmt.Printf("initcontv -- %T\n", initcontv)
						_tmpInitContri = append(_tmpInitContri, initcontv.(float64))
					}
					// fmt.Printf("%T\n", _tmpInitContri)
					// fmt.Printf("%+v\n", _tmpInitContri)
					GetCurCovrts.Result[k].CurrencyInfo.Initialcontributions = _tmpInitContri
				}
				if val, ok := vv.(map[string]interface{})["preconversions"]; ok {
					// fmt.Println("preconversions ---", val)
					var _tmpPreConv []float64
					for _, preconv := range val.([]interface{}) {
						// fmt.Println(preconv)
						// fmt.Printf("preconv -- %T\n", preconv)
						_tmpPreConv = append(_tmpPreConv, preconv.(float64))
					}
					// fmt.Printf("%T\n", _tmpPreConv)
					// fmt.Printf("%+v\n", _tmpPreConv)
					GetCurCovrts.Result[k].CurrencyInfo.Preconversions = _tmpPreConv
				}
				if val, ok := vv.(map[string]interface{})["eras"]; ok {
					// fmt.Println("eras ---", val)
					GetCurCovrts.Result[k].CurrencyInfo.Eras = val.([]interface{})
				}
			}
		}
		// break
	}
	return GetCurCovrts, nil
}
