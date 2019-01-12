package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	// Define appName type from kmdgo package
	var appName kmdgo.AppType
		
	// Define appname variable. The name value must be the matching value of it's data directory name.
	// Example Komodo's data directory is `komodo`, VerusCoin's data directory is `VRSC` and so on.
	appName = `komodo`
	
	// define the variable with GetBlockHash struct from pacakge kmdgo
	var gb kmdgo.GetBlockHash

	// Get output of the GetBlockHash() method and store it to GetBlockHash struct type variable
	gb, err := appName.GetBlockHash(1000)
	if err != nil {
		log.Println("Error happened:", err)
	}
	
	// Can print and use the struct variable outputs in further code logic. Check GetBlockHash struct in package file.
	//fmt.Println(gb)
	fmt.Println(gb.Result)

	fmt.Println(gb.Error)
	//fmt.Println(gb.ID)
}