package main

import (
	"fmt"
	"github.com/satindergrewal/kmdgo/kmdutil"
)

func main() {
	appName := "VRSC"
	dir := kmdutil.AppDataDir(appName, false)
	fmt.Println(dir)

	rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(appName)
	fmt.Printf("RPC User: %s\nRPC Password: %s\nRPC Port: %s\n", rpcuser, rpcpass, rpcport)
}