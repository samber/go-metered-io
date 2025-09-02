
# Metered io.Reader and io.Writer

[![tag](https://img.shields.io/github/tag/samber/go-metered-io.svg)](https://github.com/samber/go-metered-io/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.20.0-%23007d9c)
[![GoDoc](https://godoc.org/github.com/samber/go-metered-io?status.svg)](https://pkg.go.dev/github.com/samber/go-metered-io)
![Build Status](https://github.com/samber/go-metered-io/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/go-metered-io)](https://goreportcard.com/report/github.com/samber/go-metered-io)
[![Coverage](https://img.shields.io/codecov/c/github/samber/go-metered-io)](https://codecov.io/gh/samber/go-metered-io)
[![Contributors](https://img.shields.io/github/contributors/samber/go-metered-io)](https://github.com/samber/go-metered-io/graphs/contributors)
[![License](https://img.shields.io/github/license/samber/go-metered-io)](./LICENSE)

A drop-in replacement to `io.Reader` and `io.Writer` from the standard library with the total number of bytes transferred.

## 🚀 Install

```sh
go get github.com/samber/go-metered-io
```

This library is v1 and follows SemVer strictly. No breaking changes will be made to exported APIs before v2.0.0.

## 💡 Spec

GoDoc: [https://pkg.go.dev/github.com/samber/go-metered-io](https://pkg.go.dev/github.com/samber/go-metered-io)

Byte:
- metered.NewReader
- metered.NewWriter
- metered.NewReadCloser
- metered.NewWriteCloser
- metered.NewReadWriter

String:
- metered.NewStringWriter

Rune:
- metered.NewRuneReader

## Examples

### Metered reader

```go
import "github.com/samber/go-metered-io"

r := metered.NewReader(strings.NewReader("Lorem ipsum dolor sit amet..."))

for {
    buff := make([]byte, 10)

    _, err := r.Read(buff)
    if err != nil {
        break
    }
}

fmt.Printf("Total bytes: %d", r.Rx())
```

### Metered writer

```go
import "github.com/samber/go-metered-io"

var buf bytes.Buffer
w := metered.NewWriter(&buf)

for i := 0 ; i < 10 ; i++ {
    _, err := w.Write("Hello world\n")
    if err != nil {
        break
    }
}

fmt.Printf("Total bytes: %d", w.Tx())
```

## 🤝 Contributing

- Ping me on twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/samber/go-metered-io)
- Fix [open issues](https://github.com/samber/go-metered-io/issues) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=samber/go-metered-io)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2023 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
