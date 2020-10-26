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

	var GetNotarData kmdgo.GetNotarizationData

	args := make(kmdgo.APIParams, 1)
	args[0] = "VRSC-BTC"
	// fmt.Println(args)

	GetNotarData, err := appName.GetNotarizationData(args)
	if err != nil {
		fmt.Printf("Code: %v\n", GetNotarData.Error.Code)
		fmt.Printf("Message: %v\n\n", GetNotarData.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println(GetNotarData.Result)
	fmt.Println("Version - ", GetNotarData.Result.Version)
	// fmt.Println("Notarizations - ", GetNotarData.Result.Notarizations)
	for i, v := range GetNotarData.Result.Notarizations {
		fmt.Println("\t----")
		fmt.Println("\t", i)
		// fmt.Println(v)
		fmt.Println("\tIndex - ", v.Index)
		fmt.Println("\tTxid - ", v.Txid)
		// fmt.Println("\tNotarization - ", v.Notarization)
		fmt.Println("\t\tVersion - ", v.Notarization.Version)
		fmt.Println("\t\tCurrencyid - ", v.Notarization.Currencyid)
		fmt.Println("\t\tNotaryaddress - ", v.Notarization.Notaryaddress)
		fmt.Println("\t\tNotarizationheight - ", v.Notarization.Notarizationheight)
		fmt.Println("\t\tMmrroot - ", v.Notarization.Mmrroot)
		fmt.Println("\t\tNotarizationprehash - ", v.Notarization.Notarizationprehash)
		fmt.Println("\t\tWork - ", v.Notarization.Work)
		fmt.Println("\t\tStake - ", v.Notarization.Stake)
		// fmt.Println("\t\tCurrencystate - ", v.Notarization.Currencystate)
		fmt.Println("\t\t\tFlags- ", v.Notarization.Currencystate.Flags)
		fmt.Println("\t\t\tCurrencyid- ", v.Notarization.Currencystate.Currencyid)
		// fmt.Println("\t\t\tReservecurrencies- ", v.Notarization.Currencystate.Reservecurrencies)
		for ii, vv := range v.Notarization.Currencystate.Reservecurrencies {
			fmt.Println("\t\t\t----")
			fmt.Println("\t\t\t", ii)
			// fmt.Println("\t\t\t", vv)
			fmt.Println("\t\t\tCurrencyid - ", vv.Currencyid)
			fmt.Println("\t\t\tWeight - ", vv.Weight)
			fmt.Println("\t\t\tReserves - ", vv.Reserves)
			fmt.Println("\t\t\tPriceinreserve - ", vv.Priceinreserve)
		}
		fmt.Println("\t\t\tInitialsupply- ", v.Notarization.Currencystate.Initialsupply)
		fmt.Println("\t\t\tEmitted- ", v.Notarization.Currencystate.Emitted)
		fmt.Println("\t\t\tSupply- ", v.Notarization.Currencystate.Supply)
		// fmt.Println("\t\t\tCurrencies- ", v.Notarization.Currencystate.Currencies)
		for ci, cv := range v.Notarization.Currencystate.Currencies {
			fmt.Println("\t\t\t\tCurrency", ci)
			// fmt.Println(cv)
			fmt.Println("\t\t\t\tReservein -", cv.Reservein)
			fmt.Println("\t\t\t\tNativein -", cv.Nativein)
			fmt.Println("\t\t\t\tReserveout -", cv.Reserveout)
			fmt.Println("\t\t\t\tLastconversionprice -", cv.Lastconversionprice)
			fmt.Println("\t\t\t\tViaconversionprice -", cv.Viaconversionprice)
			fmt.Println("\t\t\t\tFees -", cv.Fees)
			fmt.Println("\t\t\t\tConversionfees -", cv.Conversionfees)
		}
		fmt.Println("\t\t\tNativefees- ", v.Notarization.Currencystate.Nativefees)
		fmt.Println("\t\t\tNativeconversionfees- ", v.Notarization.Currencystate.Nativeconversionfees)
		fmt.Println("\t\tPrevnotarization - ", v.Notarization.Prevnotarization)
		fmt.Println("\t\tPrevheight - ", v.Notarization.Prevheight)
		fmt.Println("\t\tCrossnotarization - ", v.Notarization.Crossnotarization)
		fmt.Println("\t\tCrossheight - ", v.Notarization.Crossheight)
		fmt.Println("\t\tNodes - ", v.Notarization.Nodes)
	}
	fmt.Println("Forks - ", GetNotarData.Result.Forks)
	fmt.Println("Lastconfirmedheight - ", GetNotarData.Result.Lastconfirmedheight)
	fmt.Println("Lastconfirmed - ", GetNotarData.Result.Lastconfirmed)
	fmt.Println("Bestchain - ", GetNotarData.Result.Bestchain)
}
