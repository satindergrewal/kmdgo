package main

import (
	"fmt"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSCTEST`

	// var txid string =
	rawtx, err := appName.RPCResultMap("getrawtransaction", []interface{}{"509872ef75e303fba19719adfa38f00c4f3d68854138616e2b980fc8d8d3c514", 1})
	if err != nil {
		fmt.Printf("Error getting transaction details: %v\n", err)
	}
	// blockDetails := _blockDetails.(map[string]interface{})
	fmt.Printf("%T\n", rawtx)
	fmt.Printf("%+v\n", rawtx)
}
