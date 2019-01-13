package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var dif kmdgo.GetDifficulty
	dif, err := appName.GetDifficulty()
	if err != nil {
		fmt.Printf("Code: %v\n", dif.Error.Code)
		fmt.Printf("Message: %v\n\n", dif.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("dif value", dif)
	fmt.Printf("%0.8f\n", dif.Result)
}