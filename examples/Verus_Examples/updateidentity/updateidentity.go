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
	"encoding/json"
	"fmt"
	"log"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSCTEST`

	var UpdateID kmdgo.UpdateIdentity

	// Define data to submit for updating VerusID. Parts of it taken from getidentity's output
	// Example output from getidentity like this:
	// ➜  ~ verustest getidentity "sattest01@"
	// {
	// "identity": {
	// 	"version": 2,
	// 	"flags": 0,
	// 	"primaryaddresses": [
	// 	"RCGQNeJNBaueVVGorbuAXKVEJxKEwfSdkV"
	// 	],
	// 	"minimumsignatures": 1,
	// 	"identityaddress": "iG9h8GM17z1ZByGvKMF7CmF88NVCdgbpmr",
	// 	"parent": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
	// 	"name": "sattest01",
	// 	"contentmap": {},
	// 	"revocationauthority": "i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18",
	// 	"recoveryauthority": "i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18",
	// 	"privateaddress": "zs19eqgy366y5lhnvq8v2v7gaqa9uenjx2ryqf3guy34ffy9apdfcazflwyg8ncxk56dzvfz6xxc94",
	// 	"timelock": 0
	// },
	// "status": "active",
	// "canspendfor": true,
	// "cansignfor": true,
	// "blockheight": 20499,
	// "txid": "c325219a8965a4e996573ea23df5099bd941927d4e6ff7fda37221aae01635c6",
	// "vout": 0
	// }
	var UpdateIDData kmdgo.UpdateIdentityData
	UpdateIDData.Parent = "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq" // from getidentity output
	UpdateIDData.Name = "sattest01"                            // name ID you are registering
	// UpdateIDData.Flags = 2
	// UpdateIDData.Timelock = 1219040
	UpdateIDData.Primaryaddresses = []string{"RCGQNeJNBaueVVGorbuAXKVEJxKEwfSdkV", "RYbMjn1tpn9LLteodFX6p1XoY3VSHzMnv5"} // List of addresses which you can make a list of multisig address too
	UpdateIDData.Minimumsignatures = 1                                                                                   // MofN signatures required from Primary addresses list
	UpdateIDData.Revocationauthority = "grewal@"                                                                         // Recovaction ID
	UpdateIDData.Recoveryauthority = "satinder@"                                                                         // Recovery ID
	UpdateIDData.Privateaddress = "zs19eqgy366y5lhnvq8v2v7gaqa9uenjx2ryqf3guy34ffy9apdfcazflwyg8ncxk56dzvfz6xxc94"       // Shielded Address

	cmap := make(map[string]interface{})
	cmap["key1"] = "ehlo value"
	// fmt.Println("map:", cmap)
	// UpdateIDData.Contentmap = cmap

	paramsJSON, _ := json.MarshalIndent(UpdateIDData, "", "  ")
	fmt.Println(string(paramsJSON))

	args := make(kmdgo.APIParams, 1)
	args[0] = UpdateIDData
	// fmt.Println(args)

	UpdateID, err := appName.UpdateIdentity(args)
	if err != nil {
		fmt.Printf("Code: %v\n", UpdateID.Error.Code)
		fmt.Printf("Message: %v\n\n", UpdateID.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println(UpdateID.Result)

}

// Example output will be just a txid:
//
// 273427d3af2c3badb9ac25ffd69f09105bb163917ae57d2ac6f2eae16ab140ae
