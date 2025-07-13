# str

str: seamlessly work with Strings in Go 🚤

<p>
    <a href="https://github.com/felix-kaestner/str/issues">
        <img alt="Issues" src="https://img.shields.io/github/issues/felix-kaestner/str?color=29b6f6&style=flat-square">
    </a>
    <a href="https://github.com/felix-kaestner/str/blob/main/LICENSE">
        <img alt="License" src="https://img.shields.io/github/license/felix-kaestner/str?color=29b6f6&style=flat-square">
    </a>
    <a href="https://pkg.go.dev/github.com/felix-kaestner/str">
        <img alt="Go Documentation" src="https://img.shields.io/badge/go-documentation-blue?color=29b6f6&style=flat-square">
    </a>
    <a href="https://goreportcard.com/report/github.com/felix-kaestner/str">
        <img alt="Go Report" src="https://goreportcard.com/badge/github.com/felix-kaestner/str?style=flat-square">
    </a>
    <!-- <a href="https://codecov.io/gh/felix-kaestner/str">
        <img src="https://img.shields.io/codecov/c/github/felix-kaestner/str?style=flat-square&token=KK7ZG7A90X"/>
    </a> -->
</p>

## Quickstart

Unsafe, zero-allocation conversions between `string` and `[]byte`.

```go
package main

import "github.com/felix-kaestner/str"

func main() {
    json := `{"foo":{"bar":42}}`
    // Shares the same underlying memory
    var buf []byte = str.UnsafeBytes(json)
}
```

## Installation

Install with the `go get` command:

```sh
go get -u github.com/felix-kaestner/str
```

---

Released under the [Apache 2.0 License](LICENSE).
