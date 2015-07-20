package core

import (
	"bytes"
)

func ToLower(b byte) byte {
	return bytes.ToLower([]byte{b})[0]
}
