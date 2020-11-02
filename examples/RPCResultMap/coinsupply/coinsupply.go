package main

import (
	"fmt"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSCTEST`

	_cs, _ := appName.RPCResultMap("coinsupply", []interface{}{})
	cs := _cs.(map[string]interface{})

	fmt.Printf("%T\n", cs)
	fmt.Printf("%v\n", cs["total"])
}
