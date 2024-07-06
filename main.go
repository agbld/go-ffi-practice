package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lreed_solomon
#include "reed_solomon.c"
#include <stdlib.h>  // Include for C.free
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {
    // Convert Go strings to C strings
    input := "Hello, world!"
    fmt.Println("Input:", input)
    cInput := C.CString(input)
    defer C.free(unsafe.Pointer(cInput))

    // Prepare output buffer for encoded data
    cEncoded := make([]byte, len(input)+1)
    cEncodedPtr := (*C.char)(unsafe.Pointer(&cEncoded[0]))

    // Call C function to encode
    C.encode(cInput, cEncodedPtr)

    // Convert encoded C string back to Go string
    encoded := C.GoString(cEncodedPtr)
    fmt.Println("Encoded:", encoded)

    // Prepare output buffer for decoded data
    cDecoded := make([]byte, len(input)+1)
    cDecodedPtr := (*C.char)(unsafe.Pointer(&cDecoded[0]))

    // Call C function to decode
    C.decode(cEncodedPtr, cDecodedPtr)

    // Convert decoded C string back to Go string
    decoded := C.GoString(cDecodedPtr)
    fmt.Println("Decoded:", decoded)
}
