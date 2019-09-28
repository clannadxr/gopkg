package fmtutil

import "bytes"

// TrimBom 去除bom头
func TrimBom(in []byte) []byte {
	return bytes.TrimPrefix(in, []byte{239, 187, 191})
}
