package fmtutil

// TrimBoom 去除boom头
func TrimBoom(in []byte) []byte {
	if len(in) > 3 && in[0] == 239 {
		return in[3:]
	}
	return in
}
