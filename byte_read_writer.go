package metered

import (
	"io"
	"sync/atomic"
)

// Implements io.ReadWriter
type ReadWriter struct {
	rw io.ReadWriter
	rx uint64
	tx uint64
}

func NewReadWriter(rw io.ReadWriter) *ReadWriter {
	return &ReadWriter{
		rw: rw,
	}
}

func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	n, err = rw.rw.Read(p)
	atomic.AddUint64(&rw.rx, uint64(n))
	return
}

func (rw *ReadWriter) Write(p []byte) (n int, err error) {
	n, err = rw.rw.Write(p)
	atomic.AddUint64(&rw.tx, uint64(n))
	return
}

func (rw *ReadWriter) Rx() uint64 {
	return atomic.LoadUint64(&rw.rx)
}

func (rw *ReadWriter) Tx() uint64 {
	return atomic.LoadUint64(&rw.tx)
}

// Implements io.ReadWriteCloser
type ReadWriteCloser struct {
	*ReadWriter
	io.Closer
}

func NewReadWriteCloser(rw io.ReadWriteCloser) *ReadWriteCloser {
	return &ReadWriteCloser{
		ReadWriter: NewReadWriter(rw),
		Closer:     rw,
	}
}
