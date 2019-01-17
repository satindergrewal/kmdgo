package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var sinfo kmdgo.GetTxOutSetInfo

	sinfo, err := appName.GetTxOutSetInfo()
	if err != nil {
		fmt.Printf("Code: %v\n", sinfo.Error.Code)
		fmt.Printf("Message: %v\n\n", sinfo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("sinfo value", sinfo)
	fmt.Println("-------")
	fmt.Println(sinfo.Result)
	fmt.Println("-------")
	fmt.Println("Height: ", sinfo.Result.Height)
	fmt.Println("Bestblock: ", sinfo.Result.Bestblock)
	fmt.Println("Transactions: ", sinfo.Result.Transactions)
	fmt.Println("Txouts: ", sinfo.Result.Txouts)
	fmt.Println("BytesSerialized: ", sinfo.Result.BytesSerialized)
	fmt.Println("HashSerialized: ", sinfo.Result.HashSerialized)
	fmt.Println("TotalAmount: ", sinfo.Result.TotalAmount)
}