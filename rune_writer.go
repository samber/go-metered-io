package metered

import (
	"io"
	"sync/atomic"
)

// Implements io.RuneWriter
type RuneReader struct {
	reader io.RuneReader
	rx     uint64
}

func NewRuneReader(rw io.RuneReader) *RuneReader {
	return &RuneReader{
		reader: rw,
	}
}

func (rw *RuneReader) ReadRune() (r rune, n int, err error) {
	r, n, err = rw.reader.ReadRune()
	atomic.AddUint64(&rw.rx, uint64(n))
	return
}

func (rw *RuneReader) Tx() uint64 {
	return atomic.LoadUint64(&rw.rx)
}
