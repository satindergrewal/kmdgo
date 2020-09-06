
### Sapling address generation library

This library is based on Zcash's [librustzcash](https://github.com/zcash/librustzcash), and parts of this library are taken from the [zecpaperwallet](https://github.com/adityapk00/zecpaperwallet).

### Compiling library

To generate the linkable static library follow these steps:

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
  // rust_generate_wallet function takes four parameters
  // 1) nohd:             set it to false, if you don't want HD wallet
  // 2) zcount:           the number of sapling addresses you want to generate
  // 3) seed:             the user specified passphrase, which gives the same address everytime if given the same passphrase
  // 4) is_iguana_seed:   set this to true if you want the output to always give a deterministic address based on user specified seed phrase
  bool nohd = false;
  int zcount = 1;
  char *seed = "user specified seed phrase";
  bool is_iguana_seed = true;

  char * from_rust = rust_generate_wallet(nohd, zcount, seed, is_iguana_seed);
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
    "address": "zs1z6tqgxd5fekfya07rhg6p5amwfwkmh3xwp8dcfvwqy8x4vvq0sg473d0lmgz4qevm2l4zzkhfrv",
    "private_key": "secret-extended-key-main1qwteh95uqqqqpqzv60m0kuuwrp2l0me3784kzctd6c3cfnsaflc4nw87p3huh8rp5cxy3kvuv453vsswsgfcf6kpj36az8t5qtt2u0lm2rf2auusny7qzvnxc9hn46erwzrkz9xhnk222qs7grye4qc4ulgh079xcvmmlcczpe9h4rg0385u4jfx2kutfxpx8jvjqlyf8u866c2n0j9sfc956nlwl07qy3a50vd2h6tdg2fsu5gksh25m46r7akwdxfcvc7f28mvx7s8ch3cp",
    "seed": {
      "HDSeed": "fe50eb2add6c3e1ecc550f901fa737cbebab7b7a1dbf6827d4b0fd3521d2f93e",
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
	nohd := C.bool(false)
  zcount := C.uint(1)
  seed := C.CString("user specified seed phrase")
  isIguanaSeed := C.bool(true)

  fromRust := C.CString("")
  defer C.free(unsafe.Pointer(fromRust))
  fromRust = C.rust_generate_wallet(nohd, zcount, seed, isIguanaSeed)
  fmt.Println(C.GoString(fromRust))
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
    "address": "zs1z6tqgxd5fekfya07rhg6p5amwfwkmh3xwp8dcfvwqy8x4vvq0sg473d0lmgz4qevm2l4zzkhfrv",
    "private_key": "secret-extended-key-main1qwteh95uqqqqpqzv60m0kuuwrp2l0me3784kzctd6c3cfnsaflc4nw87p3huh8rp5cxy3kvuv453vsswsgfcf6kpj36az8t5qtt2u0lm2rf2auusny7qzvnxc9hn46erwzrkz9xhnk222qs7grye4qc4ulgh079xcvmmlcczpe9h4rg0385u4jfx2kutfxpx8jvjqlyf8u866c2n0j9sfc956nlwl07qy3a50vd2h6tdg2fsu5gksh25m46r7akwdxfcvc7f28mvx7s8ch3cp",
    "seed": {
      "HDSeed": "fe50eb2add6c3e1ecc550f901fa737cbebab7b7a1dbf6827d4b0fd3521d2f93e",
      "path": "m/32'/133'/0'"
    }
  }
]
```