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
	
	// define the variable with CoinSupply struct from pacakge kmdgo
	var cs kmdgo.CoinSupply

	// Get output of the CoinSupply() method and store it to CoinSupply struct type variable
	cs, err := appName.CoinSupply(100)
	if err != nil {
		log.Println("err happened", err)
	}
	
	// Can print and use the struct variable outputs in further code logic. Check CoinSupply struct in package file.
	fmt.Println(cs)
	fmt.Println(cs.Result)
	
	fmt.Println("Result:", cs.Result.Result)
	fmt.Println("Coin:", cs.Result.Coin)
	fmt.Println("Height:", cs.Result.Height)
	fmt.Printf("Supply: %0.8f\n", cs.Result.Supply)
	fmt.Printf("Zfunds: %0.8f\n", cs.Result.Zfunds)
	fmt.Printf("Sprout: %0.8f\n", cs.Result.Sprout)
	fmt.Printf("Total: %0.8f\n", cs.Result.Total)
	fmt.Println(cs.Error)
	fmt.Println(cs.ID)
}