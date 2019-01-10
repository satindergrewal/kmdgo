package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var bh kmdgo.GetBestBlockhash
	bh, err := appName.GetBestBlockhash()
	if err != nil {
		log.Println("err happened", err)
	}

	fmt.Println("bh value", bh)
	fmt.Println(bh.Result)
	fmt.Println(bh.Error.Code)
	fmt.Println(bh.Error.Message)
	fmt.Println(bh.ID)
}