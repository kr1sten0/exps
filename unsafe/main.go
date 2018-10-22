package main

import (
	"fmt"
	"os"
	"unsafe"
)

type page struct {
	id    uint16
	count int32
	ptr   uintptr
}

type meta struct {
	version  uint16
	checksum uint64
	magic    uint16
}

func main() {
	// new buffer with os page size
	buf := make([]byte, os.Getpagesize())

	// get reference of page in buf
	p := (*page)(unsafe.Pointer(&buf[0]))
	p.id = 1
	p.count = 2
	// leave ptr as zero value

	// get reference of meta section on page
	m := (*meta)(unsafe.Pointer(&p.ptr))
	m.version = 3
	m.checksum = 4
	m.magic = 5

	// view buffer data format
	fmt.Printf("buf size: %d\n", len(buf)) // 4kB
	fmt.Printf("%02x\n", buf)

	// view `p.ptr`
	// - 00050003 -> 327683
	// - 00000003 -> 3
	fmt.Printf("%+v\n", p)
}

// `version` -> `checksum` -> `magic`
// BUF DUMP: "01000000 02000000 03000000 00000000 04000000 00000000 05000000 00000000 0000..."
//
// `version` -> `magic` -> `checksum`
// BUF DUMP: "01000000 02000000 03000500 00000000 04000000 00000000 00000000 00000000..."
