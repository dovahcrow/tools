package misc

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func init() {

}

func ReadString(re io.Reader) string {
	bBody, _ := ioutil.ReadAll(re)
	return string(bBody)
}

func ReadByte(re io.Reader) []byte {
	bBody, _ := ioutil.ReadAll(re)
	return bBody
}



func Md5Encode(in string) (out string) {
	h := md5.New()
	h.Write([]byte(in))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func JsMarshal(in interface{}) (out string) {
	b, err := json.Marshal(in)
	ErrHandle(err, `n`, `jsmarshal`)
	return string(b)
}
