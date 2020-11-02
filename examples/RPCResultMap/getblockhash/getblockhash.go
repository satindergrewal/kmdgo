package main

import (
	"fmt"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSC`
	_info, _ := appName.RPCResultMap("getinfo", []interface{}{})
	info := _info.(map[string]interface{})
	blockHash, _ := appName.RPCResultMap("getblockhash", []interface{}{info["blocks"]})
	fmt.Printf("%T\n", blockHash)
	fmt.Printf("%v\n", blockHash.(string))
}
