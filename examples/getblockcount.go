package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var bc kmdgo.GetBlockCount
	bc, err := appName.GetBlockCount()
	if err != nil {
		fmt.Printf("Code: %v\n", bc.Error.Code)
		fmt.Printf("Message: %v\n\n", bc.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("bc value", bc)
	fmt.Println(bc.Result)
	fmt.Println(bc.Error.Code)
	fmt.Println(bc.Error.Message)
	fmt.Println(bc.ID)
}