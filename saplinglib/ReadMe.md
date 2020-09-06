
### Sapling address generation library

This library is based on Zcash's [librustzcash](https://github.com/zcash/librustzcash), and parts of this library are taken from the [zecpaperwallet](https://github.com/adityapk00/zecpaperwallet).

To generate the linkable static library follow these stesp:

#### Step 0:

You **must** [install Rust](https://www.rust-lang.org/tools/install) on your machine. You need to:

```shell
curl https://sh.rustup.rs -sSf | sh
```

#### Step 1:

To compile a static linker library for your machine, just execute the following command:

```shell
make
```

It will generate a `libsaplinglib.a` file.


### C Example Code

You can use following C example to use the library:

Save the following contents in `csapling.c` file:

```C
#include <stdio.h>

#include "saplinglib.h"

int main() {
	// rust_generate_wallet function takes two parameters
	// parameter 1: the number of sapling addresses you want to generate
	// parameter 2: the user specified passphrase, which gives the same address everytime if given the same passphrase
	char * from_rust = rust_generate_wallet(1,"user-provided-passphrase");
	char *stri = from_rust;
	printf("%s", stri);
	rust_free_string(from_rust);
  	
  	return 0;
}
```

Assuming you saved the `csapling.c` in the same directory where the library file `libsaplinglib.a` is located, compile it using the following command:

```shell
gcc csapling.c -I./src -L./ -lsaplinglib -lpthread -ldl -o csapling -framework Security
```

It will generate a binary named `csapling` which if you execute will give the following output:

```json
➜  saplinglib ./csapling
[
  {
    "num": 0,
    "address": "zs1zjenpjczvcn2sm3yllqyr3zcvzzu9l3sfttnh97zwxpz6wrgem697pjerx8c6gn840fjv9lsks8",
    "private_key": "secret-extended-key-main1qw96tl5tqqqqpqqgdszc6p628dzxdaxzan8a8yk8jsdzdzzklcqu8u2eaqyv78nkhk02dfsn2f4t770ecs5huan4hralwzpmnu26u4rn3k59deq3x8asgtac5tlmrururs9uhp65298m5l7p20zwxy4ytvzvlshe8r6hn2qqg7s3lq37jg0kqdh32ksk9lmhzgxdkm2tcy6wz80ymsency0ym2ldhay4e47yjjnfntkjz5jvmmqvfyy9xtdfjpc7hufdpgmt4aldlnqlkjx8f",
    "seed": {
      "HDSeed": "92035c7f9b4db895f55ab6b00385451585a0c0470898fd4202717de46ee08e9c",
      "path": "m/32'/133'/0'"
    }
  }
]
```


### Go Language Example Code

You can use following Go language example to use the library:

Save the following contents in `gosapling.go` file:


```golang
package main

/*
#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L./ -lsaplinglib -lpthread -ldl -framework Security
#include "saplinglib.h"
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	foo := C.CString("")
	defer C.free(unsafe.Pointer(foo))
	foo = C.rust_generate_wallet(C.uint(1), C.CString("user-provided-passphrase"))
	// fmt.Printf("%T", foo)
	fmt.Println(C.GoString(foo))
}
```

Assuming you saved the `gosapling.go` in the same directory where the library file `libsaplinglib.a` is located, compile it using the following command:

```shell
go build -o gosapling gosapling.go
```

It will generate a binary named `gosapling` which if you execute will give the following output:


```json
➜  saplinglib ./gosapling
[
  {
    "num": 0,
    "address": "zs1zjenpjczvcn2sm3yllqyr3zcvzzu9l3sfttnh97zwxpz6wrgem697pjerx8c6gn840fjv9lsks8",
    "private_key": "secret-extended-key-main1qw96tl5tqqqqpqqgdszc6p628dzxdaxzan8a8yk8jsdzdzzklcqu8u2eaqyv78nkhk02dfsn2f4t770ecs5huan4hralwzpmnu26u4rn3k59deq3x8asgtac5tlmrururs9uhp65298m5l7p20zwxy4ytvzvlshe8r6hn2qqg7s3lq37jg0kqdh32ksk9lmhzgxdkm2tcy6wz80ymsency0ym2ldhay4e47yjjnfntkjz5jvmmqvfyy9xtdfjpc7hufdpgmt4aldlnqlkjx8f",
    "seed": {
      "HDSeed": "92035c7f9b4db895f55ab6b00385451585a0c0470898fd4202717de46ee08e9c",
      "path": "m/32'/133'/0'"
    }
  }
]
```