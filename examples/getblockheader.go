package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var bh kmdgo.GetBlockHeader
	bh, err := appName.GetBlockHeader("0a203fce865faa53e1414b26c3020341a17d4e2136565d9b747f83169c938ac5",true)
	if err != nil {
		fmt.Printf("Code: %v\n", bh.Error.Code)
		fmt.Printf("Message: %v\n\n", bh.Error.Message)
		log.Fatalln("Err happened", err)
	}

	//fmt.Println("bh value", bh)
	//fmt.Println(bh.Result)
	fmt.Println(bh.Result.Solution)
}