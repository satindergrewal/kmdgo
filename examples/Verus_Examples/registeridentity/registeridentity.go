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

	var RegID kmdgo.RegisterIdentity

	// Define data to submit for registering VerusID. Parts of it taken from RegisterNameCommitment's output
	// Example output from RegisterNameCommitment like this:
	// {
	// 	"txid": "0e6239539b041920e62573c3dd04ae22a6e0832ff87af8266139d030a5d7d16d",
	// 	"namereservation": {
	// 	  "name": "sattest01",
	// 	  "salt": "a3a084c680199ac0567dac697e1b77f6f1178a6e6dfa22dfac6a64ee568670fa",
	// 	  "referral": "i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18",
	// 	  "parent": "",
	// 	  "nameid": "iG9h8GM17z1ZByGvKMF7CmF88NVCdgbpmr"
	// 	}
	// }
	var RegIDData kmdgo.RegisterIdentityData
	RegIDData.Txid = "0e6239539b041920e62573c3dd04ae22a6e0832ff87af8266139d030a5d7d16d"                                  // from RegisterNameCommitment output
	RegIDData.Namereservation.Name = "sattest01"                                                                         // from RegisterNameCommitment output
	RegIDData.Namereservation.Salt = "a3a084c680199ac0567dac697e1b77f6f1178a6e6dfa22dfac6a64ee568670fa"                  // from RegisterNameCommitment output
	RegIDData.Namereservation.Referral = "i92G15J9LuimZ3ZtqQFnajDUsaaaC8qj18"                                            // from RegisterNameCommitment output
	RegIDData.Namereservation.Parent = ""                                                                                // from RegisterNameCommitment output
	RegIDData.Namereservation.Nameid = "iG9h8GM17z1ZByGvKMF7CmF88NVCdgbpmr"                                              // from RegisterNameCommitment output
	RegIDData.Identity.Name = "sattest01"                                                                                // name ID you are registering
	RegIDData.Identity.Primaryaddresses = []string{"RCGQNeJNBaueVVGorbuAXKVEJxKEwfSdkV"}                                 // List of addresses which you can make a list of multisig address too
	RegIDData.Identity.Minimumsignatures = 1                                                                             // MofN signatures required from Primary addresses list
	RegIDData.Identity.Revocationauthority = "satinder@"                                                                 // Recovaction ID
	RegIDData.Identity.Recoveryauthority = "satinder@"                                                                   // Recovery ID
	RegIDData.Identity.Privateaddress = "zs19eqgy366y5lhnvq8v2v7gaqa9uenjx2ryqf3guy34ffy9apdfcazflwyg8ncxk56dzvfz6xxc94" // Shielded Address

	args := make(kmdgo.APIParams, 1)
	args[0] = RegIDData
	// fmt.Println(args)

	RegID, err := appName.RegisterIdentity(args)
	if err != nil {
		fmt.Printf("Code: %v\n", RegID.Error.Code)
		fmt.Printf("Message: %v\n\n", RegID.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println(RegID.Result)

}

// Example output will be just a txid:
//
// c325219a8965a4e996573ea23df5099bd941927d4e6ff7fda37221aae01635c6
