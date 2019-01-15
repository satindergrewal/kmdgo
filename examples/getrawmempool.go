package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var mpool kmdgo.GetRawMempoolTrue
	mpool, err := appName.GetRawMempoolTrue(true)
	if err != nil {
		fmt.Printf("Code: %v\n", mpool.Error.Code)
		fmt.Printf("Message: %v\n\n", mpool.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("mpool value", mpool)
	fmt.Println()
	fmt.Println(mpool.Result)

	fmt.Println("-----")
	for i, v := range mpool.Result {
		fmt.Println(i)
		fmt.Println(v)
		fmt.Println(v.Size)
	}

	fmt.Printf("\n\n\n\n")

	var rmem kmdgo.GetRawMempoolFalse
	rmem, err = appName.GetRawMempoolFalse(false)
	if err != nil {
		fmt.Printf("Code: %v\n", rmem.Error.Code)
		fmt.Printf("Message: %v\n\n", rmem.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("rmem value", rmem)
	fmt.Println()
	fmt.Println(rmem.Result)

	fmt.Println("-----")
	fmt.Println(len(rmem.Result))
	for i, v := range rmem.Result {
		fmt.Println(i)
		fmt.Println(v)
		//fmt.Println(v.Size)
	}
}