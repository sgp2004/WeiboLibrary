package wb_util

import "crypto/md5"
import "io"
import "fmt"

func CreateMD5String(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	h1 := fmt.Sprintf("%x", h.Sum(nil))
	return h1
}
