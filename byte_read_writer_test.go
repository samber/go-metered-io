package metered

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReadWriter(t *testing.T) {
	is := assert.New(t)

	conn, err := net.Dial("tcp", "tcpbin.com:4242")
	is.NotNil(conn)
	is.Nil(err)

	rw := NewReadWriter(conn)
	is.NotNil(rw)

	n, err := rw.Write([]byte("hello world\n"))
	is.Equal(n, 12)
	is.Nil(err)
	is.Equal(rw.Rx(), uint64(0))
	is.Equal(rw.Tx(), uint64(12))

	resp := make([]byte, 1024)
	n, err = rw.Read(resp)
	is.Equal(n, 12)
	is.Nil(err)
	is.Equal(rw.Rx(), uint64(12))
	is.Equal(rw.Tx(), uint64(12))

	n, err = rw.Write([]byte("hello world\n"))
	is.Equal(n, 12)
	is.Nil(err)
	is.Equal(rw.Rx(), uint64(12))
	is.Equal(rw.Tx(), uint64(24))

	resp = make([]byte, 1024)
	n, err = rw.Read(resp)
	is.Equal(n, 12)
	is.Nil(err)
	is.Equal(rw.Rx(), uint64(24))
	is.Equal(rw.Tx(), uint64(24))

	conn.Close()

	n, err = rw.Write([]byte("hello world\n"))
	is.Equal(n, 0)
	is.Error(err)

	resp = make([]byte, 1024)
	n, err = rw.Read(resp)
	is.Equal(n, 0)
	is.Error(err)
}
