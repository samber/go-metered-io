package metered

import (
	"io"
	"sync/atomic"
)

// Implements io.StringWriter
type StringWriter struct {
	writer io.StringWriter
	tx     uint64
}

func NewStringWriter(sw io.StringWriter) *StringWriter {
	return &StringWriter{
		writer: sw,
	}
}

func (sw *StringWriter) WriteString(p string) (n int, err error) {
	n, err = sw.writer.WriteString(p)
	atomic.AddUint64(&sw.tx, uint64(n))
	return
}

func (sw *StringWriter) Tx() uint64 {
	return atomic.LoadUint64(&sw.tx)
}
