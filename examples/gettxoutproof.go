package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var pf kmdgo.GetTxOutProof

	txids := `["0660e479396b71715ad29d6fdbf24632bb22a3719240230a77b44b83ef6fd8f2", "9d2e64dd827e30e38c06e060b51e76e38c6bc01c612ea82d6a7b239cc713ae68"]`

	pf, err := appName.GetTxOutProof(txids)
	if err != nil {
		fmt.Printf("Code: %v\n", pf.Error.Code)
		fmt.Printf("Message: %v\n\n", pf.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("pf value", pf)
	fmt.Println("-------")
	fmt.Println(pf.Result)
}