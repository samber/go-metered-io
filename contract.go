package metered

import "io"

var _ io.Reader = (*Reader)(nil)
var _ io.ReadCloser = (*ReadCloser)(nil)

var _ io.Writer = (*Writer)(nil)
var _ io.WriteCloser = (*WriteCloser)(nil)

var _ io.ReadWriter = (*ReadWriter)(nil)
var _ io.ReadWriteCloser = (*ReadWriteCloser)(nil)

var _ io.RuneReader = (*RuneReader)(nil)

var _ io.StringWriter = (*StringWriter)(nil)
