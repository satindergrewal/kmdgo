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

	var listCur kmdgo.ListCurrencies

	args := make(kmdgo.APIParams, 1)
	args[0] = "VRSC-BTC"
	// fmt.Println(args)

	listCur, err := appName.ListCurrencies(args)
	if err != nil {
		fmt.Printf("Code: %v\n", listCur.Error.Code)
		fmt.Printf("Message: %v\n\n", listCur.Error.Message)
		log.Fatalln("Err happened", err)
	}

	// fmt.Println(listCur.Result)
	for i, v := range listCur.Result {
		fmt.Println("-----")
		fmt.Println(i)
		fmt.Println("Currencydefinition")
		fmt.Println("\tName -", v.Currencydefinition.Name)
		fmt.Println("\tVersion -", v.Currencydefinition.Version)
		fmt.Println("\tOptions -", v.Currencydefinition.Options)
		fmt.Println("\tParent -", v.Currencydefinition.Parent)
		fmt.Println("\tSystemid -", v.Currencydefinition.Systemid)
		fmt.Println("\tCurrencyid -", v.Currencydefinition.Currencyid)
		fmt.Println("\tNotarizationprotocol -", v.Currencydefinition.Notarizationprotocol)
		fmt.Println("\tProofprotocol -", v.Currencydefinition.Proofprotocol)
		fmt.Println("\tIdregistrationprice -", v.Currencydefinition.Idregistrationprice)
		fmt.Println("\tIdreferrallevels -", v.Currencydefinition.Idreferrallevels)
		fmt.Println("\tMinnotariesconfirm -", v.Currencydefinition.Minnotariesconfirm)
		fmt.Println("\tBillingperiod -", v.Currencydefinition.Billingperiod)
		fmt.Println("\tNotarizationreward -", v.Currencydefinition.Notarizationreward)
		fmt.Println("\tStartblock -", v.Currencydefinition.Startblock)
		fmt.Println("\tEndblock -", v.Currencydefinition.Endblock)
		fmt.Println("\tCurrencies -", v.Currencydefinition.Currencies)
		fmt.Println("\tWeights -", v.Currencydefinition.Weights)
		fmt.Println("\tConversions -", v.Currencydefinition.Conversions)
		fmt.Println("\tInitialsupply -", v.Currencydefinition.Initialsupply)
		fmt.Println("\tMinpreconversion -", v.Currencydefinition.Minpreconversion)
		// fmt.Println("\tPreallocation -", v.Currencydefinition.Preallocation)
		for pi, pv := range v.Currencydefinition.Preallocation {
			fmt.Println("\t\t", pi, ":", pv)
		}
		fmt.Println("\tPrelaunchcarveout -", v.Currencydefinition.Prelaunchcarveout)
		fmt.Println("\tInitialcontributions -", v.Currencydefinition.Initialcontributions)
		fmt.Println("\tPreconversions -", v.Currencydefinition.Preconversions)
		fmt.Println("\tEras -", v.Currencydefinition.Eras)
		fmt.Println("Bestheight -", v.Bestheight)
		fmt.Println("Besttxid -", v.Besttxid)
		fmt.Println("Bestcurrencystate")
		fmt.Println("\tFlags -", v.Bestcurrencystate.Flags)
		fmt.Println("\tCurrencyid -", v.Bestcurrencystate.Currencyid)
		// fmt.Println("\tReservecurrencies -", v.Bestcurrencystate.Reservecurrencies)
		for ii, vv := range v.Bestcurrencystate.Reservecurrencies {
			fmt.Println("\t\t-----")
			fmt.Println("\t\t", ii)
			// fmt.Println("\t\t-----", vv)
			fmt.Println("\t\tCurrencyid -", vv.Currencyid)
			fmt.Println("\t\tWeight -", vv.Weight)
			fmt.Println("\t\tReserves -", vv.Reserves)
			fmt.Println("\t\tPriceinreserve -", vv.Priceinreserve)
		}
		fmt.Println("\tInitialsupply -", v.Bestcurrencystate.Initialsupply)
		fmt.Println("\tEmitted -", v.Bestcurrencystate.Emitted)
		fmt.Println("\tSupply -", v.Bestcurrencystate.Supply)
		// fmt.Println("\tCurrencies -", v.Bestcurrencystate.Currencies)
		for ci, cv := range v.Bestcurrencystate.Currencies {
			fmt.Println("\t\tCurrency", ci)
			// fmt.Println(cv)
			fmt.Println("\t\tReservein -", cv.Reservein)
			fmt.Println("\t\tNativein -", cv.Nativein)
			fmt.Println("\t\tReserveout -", cv.Reserveout)
			fmt.Println("\t\tLastconversionprice -", cv.Lastconversionprice)
			fmt.Println("\t\tViaconversionprice -", cv.Viaconversionprice)
			fmt.Println("\t\tFees -", cv.Fees)
			fmt.Println("\t\tConversionfees -", cv.Conversionfees)
		}
		fmt.Println("\tNativefees -", v.Bestcurrencystate.Nativefees)
		fmt.Println("\tNativeconversionfees -", v.Bestcurrencystate.Nativeconversionfees)
	}
}
