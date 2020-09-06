package main

/*
#cgo CFLAGS: -I../
#cgo LDFLAGS: -L../ -lsaplinglib -lpthread -ldl -framework Security
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
