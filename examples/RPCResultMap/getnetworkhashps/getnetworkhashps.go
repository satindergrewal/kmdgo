package main

import (
	"fmt"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSC`

	netHashPS, _ := appName.RPCResultMap("getnetworkhashps", []interface{}{120, -1})

	fmt.Printf("%T\n", netHashPS)
	fmt.Printf("%v\n", netHashPS)
}
