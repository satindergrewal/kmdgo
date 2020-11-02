package main

import (
	"fmt"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSC`

	_netInfo, _ := appName.RPCResultMap("getnetworkinfo", []interface{}{})
	netInfo := _netInfo.(map[string]interface{})

	fmt.Printf("%T\n", netInfo)
	fmt.Printf("%v\n", netInfo["subversion"])
}
