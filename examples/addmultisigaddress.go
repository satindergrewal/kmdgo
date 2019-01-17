/*
$ komodo-cli getnewaddress
RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr
$ komodo-cli getnewaddress
RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s
$ komodo-cli addmultisigaddress 2 '["RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr","RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s"]'
bWDExTNrQZ4kSRRXwUUgHibyYuzPLS6FgP
*/

package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var mtsig kmdgo.AddMultiSigAddress

	// This example says requires 2 of 2 signatures to create tx from the multisig address
	nreq := 2
	// Get the address from command `getnewaddress` to test
	addrs := `["RLJBn63c4Fkc4csnybinhZRWhtpy8ZYnsr","RS6eYaKKqGCVysYj9BFZT4fczM4s9oo59s"]`

	mtsig, err := appName.AddMultiSigAddress(nreq, addrs)
	if err != nil {
		fmt.Printf("Code: %v\n", mtsig.Error.Code)
		fmt.Printf("Message: %v\n\n", mtsig.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("mtsig value", mtsig)
	fmt.Println("-------")
	fmt.Println(mtsig.Result)
}