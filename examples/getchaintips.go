package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var ct kmdgo.GetChainTips
	ct, err := appName.GetChainTips()
	if err != nil {
		fmt.Printf("Code: %v\n", ct.Error.Code)
		fmt.Printf("Message: %v\n\n", ct.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("ct value", ct)
	//fmt.Println(ct.Result)
}