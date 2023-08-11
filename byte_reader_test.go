package metered

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReader(t *testing.T) {
	is := assert.New(t)

	r := NewReader(strings.NewReader("foobar"))

	buff1 := make([]byte, 1)
	n, err := r.Read(buff1)
	is.Equal(n, 1)
	is.Nil(err)
	is.Equal(r.Rx(), uint64(1))

	buff2 := make([]byte, 10)
	n, err = r.Read(buff2)
	is.Equal(n, 5)
	is.Nil(err)
	is.Equal(r.Rx(), uint64(6))

	buff3 := make([]byte, 10)
	n, err = r.Read(buff3)
	is.Equal(n, 0)
	is.ErrorIs(io.EOF, err)
	is.Equal(r.Rx(), uint64(6))
}
