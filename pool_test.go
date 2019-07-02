package pool_test

import (
	"github.com/pubgo/errors"
	"github.com/pubgo/pool"
	"testing"
)

func TestBytesBufferGet(t *testing.T) {
	buf := pool.BytesBufferGet()
	defer pool.BytesBufferPut(buf)

	buf.WriteString("hello")
	errors.T(buf.String() != "hello", "error")
}
