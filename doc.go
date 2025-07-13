// SPDX-FileCopyrightText: 2025 Felix Kästner
// SPDX-License-Identifier: Apache-2.0

// Package str provides unsafe, zero-allocation conversions between Go's
// string and []byte types. It is designed for use in performance-sensitive
// situations where minimizing memory allocations is essential.
//
// # Background
//
// Go's standard conversions between string and []byte are are straightforward,
// but they always allocate new memory and copy data. This ensures immutability and
// type safety, but can introduce overhead in high-performance code. This package
// provides unsafe alternatives to these conversions that avoid these allocations,
// allowing direct access to the underlying data.
//
// # Disclaimer
// Only use these functions if you fully understand the implications of them.
package str
