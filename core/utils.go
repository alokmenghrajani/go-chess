package core

import (
	"bytes"
)

func ToLower(b byte) byte {
	return bytes.ToLower([]byte{b})[0]
}

/**
 * I know it's against Go's philosophy, but it's super helpful...
 */
func assert(expr bool, msg string) {
	if !expr {
		panic(msg)
	}
}
