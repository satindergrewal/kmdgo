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
	foo := C.CString("")
	defer C.free(unsafe.Pointer(foo))
	foo = C.rust_generate_wallet(C.uint(1), C.CString("myverysecretandstrongpassphrase_noneabletobrute"))
	// fmt.Printf("%T", foo)
	fmt.Println(C.GoString(foo))
}
