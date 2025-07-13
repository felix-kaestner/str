// SPDX-FileCopyrightText: 2025 Felix Kästner
// SPDX-License-Identifier: Apache-2.0

package str

import "unsafe"

// UnsafeString returns a string from a byte slice without allocation.
func UnsafeString(b []byte) string {
	// #nosec G103
	return *(*string)(unsafe.Pointer(&b))
}

// UnsafeBytes returns a byte slice from a string without allocation.
func UnsafeBytes(s string) []byte {
	// #nosec G103
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
