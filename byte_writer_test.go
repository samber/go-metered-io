package metered

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWriter(t *testing.T) {
	is := assert.New(t)

	var buf bytes.Buffer
	w := NewWriter(&buf)

	n, err := w.Write([]byte("foobar"))
	is.Equal(n, 6)
	is.Nil(err)
	is.Equal(w.Tx(), uint64(6))

	n, err = w.Write([]byte(""))
	is.Equal(n, 0)
	is.Nil(err)
	is.Equal(w.Tx(), uint64(6))

	n, err = w.Write([]byte("foobar"))
	is.Equal(n, 6)
	is.Nil(err)
	is.Equal(w.Tx(), uint64(12))
}
