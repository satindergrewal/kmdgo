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

	var GetInCurSt kmdgo.GetInitialCurrencyState

	args := make(kmdgo.APIParams, 1)
	args[0] = "VRSC-BTC"
	// fmt.Println(args)

	GetInCurSt, err := appName.GetInitialCurrencyState(args)
	if err != nil {
		fmt.Printf("Code: %v\n", GetInCurSt.Error.Code)
		fmt.Printf("Message: %v\n\n", GetInCurSt.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println(GetInCurSt.Result)
	fmt.Println("Flags - ", GetInCurSt.Result.Flags)
	fmt.Println("Currencyid - ", GetInCurSt.Result.Currencyid)
	fmt.Println("Reservecurrencies - ", GetInCurSt.Result.Reservecurrencies)

	for i, v := range GetInCurSt.Result.Reservecurrencies {
		fmt.Println("\t------")
		fmt.Println("\t", i)
		// fmt.Println(v)
		fmt.Println("\tCurrencyid - ", v.Currencyid)
		fmt.Println("\tWeight - ", v.Weight)
		fmt.Println("\tReserves - ", v.Reserves)
		fmt.Println("\tPriceinreserve - ", v.Priceinreserve)
	}
	fmt.Println("Initialsupply - ", GetInCurSt.Result.Initialsupply)
	fmt.Println("Emitted - ", GetInCurSt.Result.Emitted)
	fmt.Println("Supply - ", GetInCurSt.Result.Supply)
	// fmt.Println("Currencies - ", GetInCurSt.Result.Currencies)
	for ci, cv := range GetInCurSt.Result.Currencies {
		fmt.Println("Currency", ci)
		// fmt.Println(cv)
		fmt.Println("\tReservein -", cv.Reservein)
		fmt.Println("\tNativein -", cv.Nativein)
		fmt.Println("\tReserveout -", cv.Reserveout)
		fmt.Println("\tLastconversionprice -", cv.Lastconversionprice)
		fmt.Println("\tViaconversionprice -", cv.Viaconversionprice)
		fmt.Println("\tFees -", cv.Fees)
		fmt.Println("\tConversionfees -", cv.Conversionfees)
	}
	fmt.Println("Nativefees - ", GetInCurSt.Result.Nativefees)
	fmt.Println("Nativeconversionfees - ", GetInCurSt.Result.Nativeconversionfees)
}
