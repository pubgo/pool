package pool

import (
	"bytes"
	"strings"
	"sync"
)

var stringsPool = &sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func StringsBufferGet() *strings.Builder {
	return stringsPool.Get().(*strings.Builder)
}

func StringsBufferPut(buf *strings.Builder) {
	buf.Reset()
	stringsPool.Put(buf)
}
