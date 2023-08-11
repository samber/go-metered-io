package metered

import (
	"io"
	"sync/atomic"
)

// Implements io.Writer
type Writer struct {
	writer io.Writer
	tx     uint64
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		writer: writer,
	}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	atomic.AddUint64(&w.tx, uint64(n))
	return
}

func (w *Writer) Tx() uint64 {
	return atomic.LoadUint64(&w.tx)
}

// Implements io.WriterCloser
type WriteCloser struct {
	*Writer
	io.Closer
}

func NewWriteCloser(w io.WriteCloser) *WriteCloser {
	return &WriteCloser{
		Writer: NewWriter(w),
		Closer: w,
	}
}
