package structbytesperf

import (
	"testing"
	"unsafe"
)

func BenchmarkUnsafeMarshal(b *testing.B) {
	page := &Page{
		id:       100,
		flags:    0x01,
		count:    1000,
		overflow: 0x01,
	}

  var expect [SizeOfPage]byte
	for i := 0; i < b.N; i++ {
		expect = *(*[SizeOfPage]byte)(unsafe.Pointer(page))
	}
  _ = string(expect[:])
}
