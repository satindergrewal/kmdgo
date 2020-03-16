// This file is just for testing with temporary code
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("%T\n", os.Getenv("RPCURL"))

	fmt.Printf("%s lives in %s.\n", os.Getenv("RPCURL"), os.Getenv("BURROW"))

	key := "RPCURL"
	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Printf("%s not set\n", key)
	} else {
		fmt.Printf("%s=%s\n", key, val)
	}
}
