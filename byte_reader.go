package metered

import (
	"io"
	"sync/atomic"
)

// Implements io.Reader
type Reader struct {
	reader io.Reader
	rx     uint64
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		reader: reader,
	}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	atomic.AddUint64(&r.rx, uint64(n))
	return
}

func (r *Reader) Rx() uint64 {
	return atomic.LoadUint64(&r.rx)
}

// Implements io.ReadCloser
type ReadCloser struct {
	*Reader
	io.Closer
}

func NewReadCloser(r io.ReadCloser) *ReadCloser {
	return &ReadCloser{
		Reader: NewReader(r),
		Closer: r,
	}
}
