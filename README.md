<!--
# SPDX-FileCopyrightText: 2025 Felix Kästner
# SPDX-License-Identifier: Apache-2.0
-->

# str

str: seamlessly work with Strings in Go 🚤

[![Issues](https://img.shields.io/github/issues/felix-kaestner/str?color=29b6f6&style=flat-square)](https://github.com/felix-kaestner/str/issues)
[![License](https://img.shields.io/github/license/felix-kaestner/str?color=29b6f6&style=flat-square)](https://github.com/felix-kaestner/str/blob/main/LICENSE)
[![Go Documentation](https://img.shields.io/badge/go-documentation-blue?color=29b6f6&style=flat-square)](https://pkg.go.dev/github.com/felix-kaestner/str)
[![Go Report](https://goreportcard.com/badge/github.com/felix-kaestner/str?style=flat-square)](https://goreportcard.com/report/github.com/felix-kaestner/str)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/felix-kaestner/str/pulls)

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
