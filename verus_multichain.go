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
	}
	return GetCurCovrts, nil
}

// GetExports type
type GetExports struct {
	Result []struct {
		Blockheight int    `json:"blockheight"`
		Exportid    string `json:"exportid"`
		Description struct {
			Version          int                `json:"version"`
			Exportcurrencyid string             `json:"exportcurrencyid"`
			Numinputs        int                `json:"numinputs"`
			Totalamounts     map[string]float64 `json:"totalamounts"`
			Totalfees        map[string]float64 `json:"totalfees"`
		} `json:"description"`
		Transfers []struct {
			Version               int     `json:"version"`
			Currencyid            string  `json:"currencyid"`
			Value                 float64 `json:"value"`
			Flags                 int     `json:"flags"`
			Preconvert            bool    `json:"preconvert,omitempty"`
			Fees                  float64 `json:"fees"`
			Destinationcurrencyid string  `json:"destinationcurrencyid"`
			Destination           string  `json:"destination"`
			Feeoutput             bool    `json:"feeoutput,omitempty"`
		} `json:"transfers"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetExports returns all pending export transfers that are not yet provable with confirmed notarizations.
// These are the transactions that are crossing from one to another currency. In other words: conversions
// It's output behaves like a mempool transactions, and the output of results disappear after a while, and new ones shows up.
//
// getexports "chainname"
//
// Arguments
// 1. "chainname"                     (string, optional) name of the chain to look for. no parameter returns current chain in daemon.
//
// Example Result:
// [
//   {
//     "blockheight": 144,
//     "exportid": "ea087427e81352bd84887ff90d370e8cf5c51b61f694c673b75b64696391d777",
//     "description": {
//       "version": 1,
//       "exportcurrencyid": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
//       "numinputs": 2,
//       "totalamounts": {
//         "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": 94.15731371,
//         "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": 1000250.06291562
//       },
//       "totalfees": {
//         "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": 0.02353932,
//         "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": 250.06291562
//       }
//     },
//     "transfers": [
//       {
//         "version": 1,
//         "currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
//         "value": 1000250.06251562,
//         "flags": 4101,
//         "preconvert": true,
//         "fees": 0.0002,
//         "destinationcurrencyid": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
//         "destination": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt"
//       },
//       {
//         "version": 1,
//         "currencyid": "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
//         "value": 94.15731371,
//         "flags": 4101,
//         "preconvert": true,
//         "fees": 0.0002,
//         "destinationcurrencyid": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
//         "destination": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt"
//       },
//       {
//         "version": 1,
//         "currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
//         "value": 0,
//         "flags": 9,
//         "feeoutput": true,
//         "fees": 0,
//         "destinationcurrencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
//         "destination": "RCdrw4BL7B8rKJ2iQqftvBA4SAtwGA3eBc"
//       }
//     ]
//   },
//   {},
//   ...
// ]
//
// Examples:
// > verus getexports "chainname"
// > curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getexports", "params": ["chainname"] }' -H 'content-type: text/plain;' http://127.0.0.1:27486/
func (appName AppType) GetExports(params APIParams) (GetExports, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getexports`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var GetExp GetExports

	GetExpJSON := appName.APICall(&query)
	if GetExpJSON == "EMPTY RPC INFO" {
		return GetExp, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(GetExpJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(GetExpJSON), &GetExp)
		return GetExp, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(GetExpJSON), &GetExp)
	return GetExp, nil
}

// GetImports type
type GetImports struct {
	Result []struct {
		Blockheight int    `json:"blockheight"`
		Importid    string `json:"importid"`
		Description struct {
			Version          int                `json:"version"`
			Sourcesystemid   string             `json:"sourcesystemid"`
			Importcurrencyid string             `json:"importcurrencyid"`
			Valuein          map[string]float64 `json:"valuein"`
			Tokensout        map[string]float64 `json:"tokensout"`
		} `json:"description"`
		Transfers []struct {
			Version               int     `json:"version"`
			Currencyid            string  `json:"currencyid"`
			Value                 float64 `json:"value"`
			Flags                 int     `json:"flags"`
			Preconvert            bool    `json:"preconvert,omitempty"`
			Fees                  float64 `json:"fees"`
			Destinationcurrencyid string  `json:"destinationcurrencyid"`
			Destination           string  `json:"destination"`
			Feeoutput             bool    `json:"feeoutput,omitempty"`
		} `json:"transfers"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetImports returns all imports from a specific chain.
//
// getimports "chainname"
//
// Arguments
// 1. "chainname"                     (string, optional) name of the chain to look for. no parameter returns current chain in daemon.
//
// Example Result:
// [
//   {
//     "blockheight": 149,
//     "importid": "c5b5aa070b57b6599ea8714692187f06261c215c641d13e06dacd74ed40272a3",
//     "description": {
//       "version": 1,
//       "sourcesystemid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
//       "importcurrencyid": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
//       "valuein": {
//         "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt": 1999933.86859995,
//         "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": 94.15731371,
//         "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": 1000500.06191861
//       },
//       "tokensout": {
//         "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": 94.15731371
//       }
//     },
//     "transfers": [
//       {
//         "version": 1,
//         "currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
//         "value": 1000250.06251562,
//         "flags": 4101,
//         "preconvert": true,
//         "fees": 0.0002,
//         "destinationcurrencyid": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
//         "destination": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt"
//       },
//       {
//         "version": 1,
//         "currencyid": "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
//         "value": 94.15731371,
//         "flags": 4101,
//         "preconvert": true,
//         "fees": 0.0002,
//         "destinationcurrencyid": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
//         "destination": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt"
//       },
//       {
//         "version": 1,
//         "currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
//         "value": 0,
//         "flags": 9,
//         "feeoutput": true,
//         "fees": 0,
//         "destinationcurrencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
//         "destination": "RCdrw4BL7B8rKJ2iQqftvBA4SAtwGA3eBc"
//       }
//     ]
//   },
//   {},
//   ...
// ]
//
// Examples:
// > verus getimports "chainname"
// > curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getimports", "params": ["chainname"] }' -H 'content-type: text/plain;' http://127.0.0.1:27486/
func (appName AppType) GetImports(params APIParams) (GetImports, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getimports`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var GetImp GetImports

	GetImpJSON := appName.APICall(&query)
	if GetImpJSON == "EMPTY RPC INFO" {
		return GetImp, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(GetImpJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(GetImpJSON), &GetImp)
		return GetImp, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(GetImpJSON), &GetImp)
	return GetImp, nil
}

// GetInitialCurrencyState type
type GetInitialCurrencyState struct {
	Result struct {
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
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetInitialCurrencyState returns the total amount of preconversions that have been confirmed on the blockchain for the specified PBaaS chain.
// This should be used to get information about chains that are not this chain, but are being launched by it.
//
// getinitialcurrencystate "name"
//
// Arguments
//    "name"                    (string, required) name or chain ID of the chain to get the export transactions for
//
// Example Result:
// {
// 	"flags": 11,
// 	"currencyid": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
// 	"reservecurrencies": [
// 		{
// 			"currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
// 			"weight": 0.50000000,
// 			"reserves": 999750.00099701,
// 			"priceinreserve": 0.99975000
// 		},
// 		{
// 			"currencyid": "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
// 			"weight": 0.50000000,
// 			"reserves": 94.15731371,
// 			"priceinreserve": 0.00009415
// 		}
// 	],
// 	"initialsupply": 2000000.00000000,
// 	"emitted": 0.00000000,
// 	"supply": 2000000.00000000,
// 	"currencies": {
// 		"iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": {
// 			"reservein": 1000000.00000000,
// 			"nativein": 0.00000000,
// 			"reserveout": 249.99900299,
// 			"lastconversionprice": 1.00000000,
// 			"viaconversionprice": 0.99981249,
// 			"fees": 250.06291562,
// 			"conversionfees": 250.06251562
// 		},
// 		"iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": {
// 			"reservein": 94.15731371,
// 			"nativein": 0.00000000,
// 			"reserveout": 0.00000000,
// 			"lastconversionprice": 0.00009414,
// 			"viaconversionprice": 0.00009414,
// 			"fees": 0.02353932,
// 			"conversionfees": 0.02353932
// 		}
// 	},
// 	"nativefees": 50006191861,
// 	"nativeconversionfees": 50006151861
// }
//
// Examples:
// > verus getinitialcurrencystate name
// > curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getinitialcurrencystate", "params": [name] }' -H 'content-type: text/plain;' http://127.0.0.1:27486/
func (appName AppType) GetInitialCurrencyState(params APIParams) (GetInitialCurrencyState, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getinitialcurrencystate`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var GetInCurSt GetInitialCurrencyState

	GetInCurStJSON := appName.APICall(&query)
	if GetInCurStJSON == "EMPTY RPC INFO" {
		return GetInCurSt, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(GetInCurStJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(GetInCurStJSON), &GetInCurSt)
		return GetInCurSt, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(GetInCurStJSON), &GetInCurSt)
	return GetInCurSt, nil
}

// GetLastImportin type
type GetLastImportin struct {
	Result struct {
		Lastimporttransaction     string             `json:"lastimporttransaction"`
		Lastconfirmednotarization string             `json:"lastconfirmednotarization"`
		Importtxtemplate          string             `json:"importtxtemplate"`
		Nativeimportavailable     int64              `json:"nativeimportavailable"`
		Tokenimportavailable      map[string]float64 `json:"tokenimportavailable"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetLastImportin returns the last import transaction from the chain specified and a blank transaction template to use when making new
// import transactions. Since the typical use for this call is to make new import transactions from the other chain that will be then
// broadcast to this chain, we include the template by default.
//
// getlastimportin "fromname"
//
// Arguments
//    "fromname"                (string, required) name of the chain to get the last import transaction in from
//
// Result:
//    {
//        "lastimporttransaction": "hex"      Hex encoded serialized import transaction
//        "lastconfirmednotarization" : "hex" Hex encoded last confirmed notarization transaction
//        "importtxtemplate": "hex"           Hex encoded import template for new import transactions
//        "nativeimportavailable": "amount"   Total amount of native import currency available to import as native
//        "tokenimportavailable": "array"     ([{"currencyid":amount},..], required) tokens available to import, if controlled by this chain
//    }
//
// Examples:
// > verus getlastimportin jsondefinition
// > curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getlastimportin", "params": [jsondefinition] }' -H 'content-type: text/plain;' http://127.0.0.1:27486/
func (appName AppType) GetLastImportin(params APIParams) (GetLastImportin, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getlastimportin`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var GetLastImpIn GetLastImportin

	GetLastImpInJSON := appName.APICall(&query)
	if GetLastImpInJSON == "EMPTY RPC INFO" {
		return GetLastImpIn, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(GetLastImpInJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(GetLastImpInJSON), &GetLastImpIn)
		return GetLastImpIn, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(GetLastImpInJSON), &GetLastImpIn)
	return GetLastImpIn, nil
}

// GetNotarizationData type
type GetNotarizationData struct {
	Result struct {
		Version       int `json:"version"`
		Notarizations []struct {
			Index        int    `json:"index"`
			Txid         string `json:"txid"`
			Notarization struct {
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
					Nativefees           int                   `json:"nativefees"`
					Nativeconversionfees int                   `json:"nativeconversionfees"`
				} `json:"currencystate"`
				Prevnotarization  string        `json:"prevnotarization"`
				Prevheight        int           `json:"prevheight"`
				Crossnotarization string        `json:"crossnotarization"`
				Crossheight       int           `json:"crossheight"`
				Nodes             []interface{} `json:"nodes"`
			} `json:"notarization"`
		} `json:"notarizations"`
		Forks               [][]int `json:"forks"`
		Lastconfirmedheight int     `json:"lastconfirmedheight"`
		Lastconfirmed       int     `json:"lastconfirmed"`
		Bestchain           int     `json:"bestchain"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetNotarizationData returns the latest PBaaS notarization data for the specifed currencyid.
//
// getnotarizationdata "currencyid" accepted
//
// Arguments
// 1. "currencyid"                  (string, required) the hex-encoded ID or string name  search for notarizations on
// 2. "accepted"                    (bool, optional) accepted, not earned notarizations, default: true if on
//                                                     VRSC or VRSCTEST, false otherwise
//
// Example Result:
// {
// 	"version": 1,
// 	"notarizations": [
// 	  {
// 		"index": 0,
// 		"txid": "936e0f4f318bf21c0cd804037526e336f8162bdf2409f72134868119752efd8f",
// 		"notarization": {
// 		  "version": 1,
// 		  "currencyid": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
// 		  "notaryaddress": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
// 		  "notarizationheight": 31112,
// 		  "mmrroot": "9f95c86ffc6d7f076370dcb27c198c18505281e1abbf4bac742e46a005597548",
// 		  "notarizationprehash": "4cbe080e1d9c188ab1bd56f4915298b6671aa556be1161ed8b0bebb5085bc5d6",
// 		  "work": "00000000000000000000000000000000000000000000000000000000195f61be",
// 		  "stake": "0000000000000000000000000000000000000000000000000000000000000000",
// 		  "currencystate": {
// 			"flags": 3,
// 			"currencyid": "i84mndBk2Znydpgm9T9pTjVvBnHkhErzLt",
// 			"reservecurrencies": [
// 			  {
// 				"currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
// 				"weight": 0.50000000,
// 				"reserves": 999894.45679682,
// 				"priceinreserve": 0.99983971
// 			  },
// 			  {
// 				"currencyid": "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
// 				"weight": 0.50000000,
// 				"reserves": 94.15401986,
// 				"priceinreserve": 0.00009414
// 			  }
// 			],
// 			"initialsupply": 2000000.00000000,
// 			"emitted": 0.00000000,
// 			"supply": 2000109.49758289,
// 			"currencies": {
// 			  "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": {
// 				"reservein": 0.00000000,
// 				"nativein": 0.99923512,
// 				"reserveout": 0.99907519,
// 				"lastconversionprice": 0.99983995,
// 				"viaconversionprice": 0.99982722,
// 				"fees": 0.00000000,
// 				"conversionfees": 0.00000000
// 			  },
// 			  "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": {
// 				"reservein": 0.00000000,
// 				"nativein": 0.00000000,
// 				"reserveout": 0.00000000,
// 				"lastconversionprice": 0.00009414,
// 				"viaconversionprice": 0.00009414,
// 				"fees": 0.00000000,
// 				"conversionfees": 0.00000000
// 			  }
// 			},
// 			"nativefees": 12490,
// 			"nativeconversionfees": 0
// 		  },
// 		  "prevnotarization": "a694a22cc3aee0212105021f5363803f87ff2fc49689749d3feb563fd1c9c2ab",
// 		  "prevheight": 31081,
// 		  "crossnotarization": "0000000000000000000000000000000000000000000000000000000000000000",
// 		  "crossheight": 0,
// 		  "nodes": [
// 		  ]
// 		}
// 	  }
// 	],
// 	"forks": [
// 	  [
// 		0
// 	  ]
// 	],
// 	"lastconfirmedheight": 31112,
// 	"lastconfirmed": 0,
// 	"bestchain": 0
// }
//
// Examples:
// > verus getnotarizationdata "currencyid" true
// > curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getnotarizationdata", "params": ["currencyid"] }' -H 'content-type: text/plain;' http://127.0.0.1:27486/
func (appName AppType) GetNotarizationData(params APIParams) (GetNotarizationData, error) {

	// fmt.Println("params[0]", params[0])

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getnotarizationdata`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var GetNotarData GetNotarizationData

	GetNotarDataJSON := appName.APICall(&query)
	if GetNotarDataJSON == "EMPTY RPC INFO" {
		return GetNotarData, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(GetNotarDataJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(GetNotarDataJSON), &GetNotarData)
		return GetNotarData, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(GetNotarDataJSON), &GetNotarData)
	return GetNotarData, nil
}
