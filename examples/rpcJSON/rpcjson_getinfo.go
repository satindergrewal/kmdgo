package main

import (
	"fmt"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSC`
	_info, _ := appName.RPCJSON("getinfo", []interface{}{})
	fmt.Println(_info)
}
