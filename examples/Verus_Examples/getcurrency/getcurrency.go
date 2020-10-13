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

package main

import (
	"fmt"
	"log"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSCTEST`

	var GetCurncy kmdgo.GetCurrency

	args := make(kmdgo.APIParams, 1)
	args[0] = "VRSC-BTC-ETH-KMD"
	// fmt.Println(args)

	GetCurncy, err := appName.GetCurrency(args)
	if err != nil {
		fmt.Printf("Code: %v\n", GetCurncy.Error.Code)
		fmt.Printf("Message: %v\n\n", GetCurncy.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("Name - ", GetCurncy.Result.Name)
	fmt.Println("Version - ", GetCurncy.Result.Version)
	fmt.Println("Options - ", GetCurncy.Result.Options)
	fmt.Println("Parent - ", GetCurncy.Result.Parent)
	fmt.Println("Systemid - ", GetCurncy.Result.Systemid)
	fmt.Println("Currencyid - ", GetCurncy.Result.Currencyid)
	fmt.Println("Notarizationprotocol - ", GetCurncy.Result.Notarizationprotocol)
	fmt.Println("Proofprotocol - ", GetCurncy.Result.Proofprotocol)
	fmt.Println("Idregistrationprice - ", GetCurncy.Result.Idregistrationprice)
	fmt.Println("Idreferrallevels - ", GetCurncy.Result.Idreferrallevels)
	fmt.Println("Minnotariesconfirm - ", GetCurncy.Result.Minnotariesconfirm)
	fmt.Println("Billingperiod - ", GetCurncy.Result.Billingperiod)
	fmt.Println("Notarizationreward - ", GetCurncy.Result.Notarizationreward)
	fmt.Println("Startblock - ", GetCurncy.Result.Startblock)
	fmt.Println("Endblock - ", GetCurncy.Result.Endblock)
	// fmt.Println("Currencies - ", GetCurncy.Result.Currencies)
	for ci, cv := range GetCurncy.Result.Currencies {
		fmt.Println("Currency", ci, "-", cv)
	}
	// fmt.Println("Weights - ", GetCurncy.Result.Weights)
	for wi, wv := range GetCurncy.Result.Weights {
		fmt.Println("Weights", wi, "-", wv)
	}
	// fmt.Println("Conversions - ", GetCurncy.Result.Conversions)
	for cnvi, cnvv := range GetCurncy.Result.Conversions {
		fmt.Printf("Conversion %v - %.8f\n", cnvi, cnvv)
	}
	fmt.Printf("Initialsupply - %.8f\n", GetCurncy.Result.Initialsupply)
	fmt.Printf("Prelaunchcarveout - %.8f\n", GetCurncy.Result.Prelaunchcarveout)
	// fmt.Println("Initialcontributions - ", GetCurncy.Result.Initialcontributions)
	for initConti, initContv := range GetCurncy.Result.Initialcontributions {
		fmt.Printf("Initialcontribution %v - %.8f\n", initConti, initContv)
	}
	// fmt.Println("Preconversions - ", GetCurncy.Result.Preconversions)
	for preConvi, preConvv := range GetCurncy.Result.Preconversions {
		fmt.Printf("Preconversion %v - %.8f\n", preConvi, preConvv)
	}
	fmt.Println("Eras - ", GetCurncy.Result.Eras)
	fmt.Println("Definitiontxid - ", GetCurncy.Result.Definitiontxid)
	// fmt.Println("Bestcurrencystate - ", GetCurncy.Result.Bestcurrencystate)
	fmt.Println("Bestcurrencystate > Flags - ", GetCurncy.Result.Bestcurrencystate.Flags)
	fmt.Println("Bestcurrencystate > Currencyid - ", GetCurncy.Result.Bestcurrencystate.Currencyid)
	// fmt.Println("Bestcurrencystate > Reservecurrencies - ", GetCurncy.Result.Bestcurrencystate.Reservecurrencies)
	for rsvCrncyi, rsvCrncyv := range GetCurncy.Result.Bestcurrencystate.Reservecurrencies {
		fmt.Println("Reservecurrency", rsvCrncyi, "-")
		fmt.Println("Bestcurrencystate > Reservecurrency > Currencyid -", rsvCrncyv.Currencyid)
		fmt.Printf("Bestcurrencystate > Reservecurrency > Weight - %.8f\n", rsvCrncyv.Weight)
		fmt.Printf("Bestcurrencystate > Reservecurrency > Reserves - %.8f\n", rsvCrncyv.Reserves)
		fmt.Printf("Bestcurrencystate > Reservecurrency > Priceinreserve - %.8f\n", rsvCrncyv.Priceinreserve)
	}
	fmt.Printf("Bestcurrencystate > Reservecurrencies > Initialsupply - %.8f\n", GetCurncy.Result.Bestcurrencystate.Initialsupply)
	fmt.Printf("Bestcurrencystate > Reservecurrencies > Emitted - %.8f\n", GetCurncy.Result.Bestcurrencystate.Emitted)
	fmt.Printf("Bestcurrencystate > Reservecurrencies > Supply - %.8f\n", GetCurncy.Result.Bestcurrencystate.Supply)
	// fmt.Println("Bestcurrencystate > Reservecurrencies > Currencies - ", GetCurncy.Result.Bestcurrencystate.Currencies)
	for i, v := range GetCurncy.Result.Bestcurrencystate.Currencies {
		fmt.Println("------")
		fmt.Println(i)
		// fmt.Println(v)
		fmt.Printf("Reservein - %.8f\n", v.Reservein)
		fmt.Printf("Nativein - %.8f\n", v.Nativein)
		fmt.Printf("Reserveout - %.8f\n", v.Reserveout)
		fmt.Printf("Lastconversionprice - %.8f\n", v.Lastconversionprice)
		fmt.Printf("Viaconversionprice - %.8f\n", v.Viaconversionprice)
		fmt.Printf("Fees - %.8f\n", v.Fees)
		fmt.Printf("Conversionfees - %.8f\n", v.Conversionfees)
	}
	fmt.Println("Bestcurrencystate > Reservecurrencies > Nativefees - ", GetCurncy.Result.Bestcurrencystate.Nativefees)
	fmt.Println("Bestcurrencystate > Reservecurrencies > Nativeconversionfees - ", GetCurncy.Result.Bestcurrencystate.Nativeconversionfees)

}
