package main

import (
	"fmt"
	"github.com/satindergrewal/kmdgo/kmdutil"
)

func main() {
    dir := kmdutil.AppDataDir("komodo", false)
    fmt.Println(dir)
}