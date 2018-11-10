package main

import "strings"
import "fmt"
import "github.com/satindergrewal/kmdgo/kmdutil"

func main() {
    dir := kmdutil.AppDataDir("komodo", false)
    fmt.Println(dir)
}