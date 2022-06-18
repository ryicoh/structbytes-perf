package structbytesperf

import "unsafe"

type Page struct {
	id       uint64
	flags    uint16
	count    uint16
	overflow uint32
}

const SizeOfPage = int(unsafe.Sizeof(Page{}))
