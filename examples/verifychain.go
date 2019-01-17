package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var vfy kmdgo.VerifyChain

	checklevel := 1187540 // (numeric, optional, 0-4, default=3)
	numblocks := 288 // (numeric, optional, default=288, 0=all)

	vfy, err := appName.VerifyChain(checklevel, numblocks)
	if err != nil {
		fmt.Printf("Code: %v\n", vfy.Error.Code)
		fmt.Printf("Message: %v\n\n", vfy.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("vfy value", vfy)
	fmt.Println("-------")
	fmt.Println(vfy.Result)
}