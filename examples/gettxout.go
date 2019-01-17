package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var txout kmdgo.GetTxOut

	txid := `c2abe4ae3b4965ba4d55048f355d9ddc803bf1b6123a10902ae4114260f516fb`

	txout, err := appName.GetTxOut(txid, 1)
	if err != nil {
		fmt.Printf("Code: %v\n", txout.Error.Code)
		fmt.Printf("Message: %v\n\n", txout.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("txout value", txout)
	fmt.Println(txout.Result)
	fmt.Println("----------")
	fmt.Println("Bestblock: ", txout.Result.Bestblock)
	fmt.Println("Confirmations: ", txout.Result.Confirmations)
	fmt.Println("Rawconfirmations: ", txout.Result.Rawconfirmations)
	fmt.Println("Value: ", txout.Result.Value)
	//fmt.Println("ScriptPubKey: ", txout.Result.ScriptPubKey)
	fmt.Println("ScriptPubKey Asm: ", txout.Result.ScriptPubKey.Asm)
	fmt.Println("ScriptPubKey Hex: ", txout.Result.ScriptPubKey.Hex)
}