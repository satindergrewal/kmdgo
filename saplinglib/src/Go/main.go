package main

/*
#cgo CFLAGS: -I../
#cgo LDFLAGS: -L../../dist/darwin -lsaplinglib -lpthread -ldl -framework Security
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
	// rust_generate_wallet function takes four parameters
	// 1) nohd:				set it to false, if you don't want HD wallet
	// 2) zcount:			the number of sapling addresses you want to generate
	// 3) seed:				the user specified passphrase, which gives the same address everytime if given the same passphrase
	// 4) isIguanaSeed:		set this to true if you want the output to always give a deterministic address based on user specified seed phrase
	nohd := C.bool(false)
	zcount := C.uint(1)
	seed := C.CString("user specified seed phrase")
	isIguanaSeed := C.bool(true)

	fromRust := C.CString("")
	defer C.free(unsafe.Pointer(fromRust))
	fromRust = C.rust_generate_wallet(nohd, zcount, seed, isIguanaSeed)
	fmt.Println(C.GoString(fromRust))
}
