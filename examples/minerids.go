package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var sinfo kmdgo.MinerIDs

	block_height := `1187540`

	sinfo, err := appName.MinerIDs(block_height)
	if err != nil {
		fmt.Printf("Code: %v\n", sinfo.Error.Code)
		fmt.Printf("Message: %v\n\n", sinfo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	//fmt.Println("sinfo value", sinfo)
	fmt.Println("-------")
	fmt.Println(sinfo.Result)
	fmt.Println("-------")
	fmt.Println("Numnotaries: ", sinfo.Result.Numnotaries)
	fmt.Printf("\n\n\n\n")

	for i, v := range sinfo.Result.Mined {
		fmt.Println(i)
		fmt.Println("Notaryid: ", v.Notaryid)
		fmt.Println("KMDaddress: ", v.KMDaddress)
		fmt.Println("Pubkey: ", v.Pubkey)
		fmt.Println("Blocks: ", v.Blocks)
	}
}