# kmdgo


To use this package do

```shell
go get -u github.com/satindergrewal/kmdgo
```


Example code:

```go
//main.go
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
```

Output:

```shell
$ go run main.go
/home/satinder/.komodo/VRSC
RPC User: user1773837506
RPC Password: passbae43ecd576o8ariotkgjhasdfiyosidulrkhdf9390bf03b68
RPC Port: 27486
```