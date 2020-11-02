package main

import (
	"fmt"
	"strconv"

	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `VRSC`

	var blockNum uint64 = 0
	_blockDetails, _ := appName.RPCResultMap("getblock", []interface{}{strconv.FormatUint(blockNum, 10), 2})
	blockDetails := _blockDetails.(map[string]interface{})
	// fmt.Printf("%T\n", blockDetails)
	// fmt.Printf("%+v\n", blockDetails)
	fmt.Printf("%v\n", blockDetails["time"].(float64))
	fmt.Printf("%v\n", blockDetails["hash"].(string))

	txns := blockDetails["tx"].([]interface{})
	for txIndex, _txid := range txns {
		fmt.Println(txIndex)
		fmt.Printf("%T\n", _txid)

		txData := _txid.(map[string]interface{})
		fmt.Printf("%v\n", txData["txid"].(string))

		// fmt.Printf("%T\n", txData["vout"])
		// // fmt.Printf("%v\n", txData["vout"].([]interface{}))

		// vOutData := txData["vout"].([]interface{})
		// fmt.Printf("%T\n", vOutData[0])
		// fmt.Println(vOutData[0].(map[string]interface{})["scriptPubKey"].(map[string]interface{})["addresses"].([]interface{})[0].(string))
	}
}
