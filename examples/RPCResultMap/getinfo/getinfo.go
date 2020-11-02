package main

import (
	"fmt"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSC`
	_info, _ := appName.RPCResultMap("getinfo", []interface{}{})
	fmt.Printf("%T\n", _info)
	fmt.Printf("%v\n", _info.(map[string]interface{})["blocks"])
}
