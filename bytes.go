package pool

import (
	"bytes"
	"sync"
)

var bytesPool = &sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func BytesBufferGet() *bytes.Buffer {
	return bytesPool.Get().(*bytes.Buffer)
}

func BytesBufferPut(buf *bytes.Buffer) {
	buf.Reset()
	bytesPool.Put(buf)
}
