package csvutil

import (
	"io/ioutil"

	"github.com/clannadxr/gopkg/fmtutil"
	"github.com/gocarina/gocsv"
)

// UnmarshalFromFile 从文件unmarshal到struct
func UnmarshalFromFile(filename string, out interface{}) (err error) {
	csvDataBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	// 去除bom
	csvDataBytes = fmtutil.TrimBom(csvDataBytes)
	return gocsv.UnmarshalBytes(csvDataBytes, out)
}
