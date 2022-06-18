package structbytesperf

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func BenchmarkBinary(b *testing.B) {
	page := &Page{
		id:       100,
		flags:    0x01,
		count:    1000,
		overflow: 0x01,
	}

	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		if err := binary.Write(&buf, binary.LittleEndian, page); err != nil {
			b.Error(err)
		}
	}
	_ = buf.String()
}
