package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var mpool kmdgo.GetRawMempool
	mpool, err := appName.GetRawMempool(true)
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
}