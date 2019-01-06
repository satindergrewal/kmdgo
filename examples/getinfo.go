package main

import (
	"fmt"
    "github.com/satindergrewal/kmdgo"
)

func main() {
    // Define appname variable. The name value must be the matching value of it's data directory name.
    // Example Komodo's data directory is `komodo`, VerusCoin's data directory is `VRSC` and so on.
    appName := `komodo`

    // define the variable with GetInfo struct from pacakge kmdgo
	var rval kmdgo.GetInfo

    // Get output of the ResultGetinfo function and store it to GetInfo struct variable
    rval = kmdgo.ResultGetInfo(appName)

    // Can print and use the struct variable outputs in further code logic. Check GetInfo struct in package file.
    fmt.Println(rval)
    fmt.Println(rval.Result)
    
    fmt.Println("Version:", rval.Result.Version)
    fmt.Println("Balance:", rval.Result.Balance)
    fmt.Println("Blocks:", rval.Result.Blocks)
    fmt.Println("Name:", rval.Result.Name)
    fmt.Println("Connections:", rval.Result.Connections)
    fmt.Println("Difficulty:", rval.Result.Difficulty)
    fmt.Println("Magic:", rval.Result.Magic)
    fmt.Println(rval.Error)
    fmt.Println(rval.ID)
}