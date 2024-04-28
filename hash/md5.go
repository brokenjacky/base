package hash

import (
    "fmt"
    "crypto/md5"
)

func Md5(str string) string {
    data := []byte(str)
    has := md5.Sum(data)
    md5str := fmt.Sprintf("%x", has)
    return md5str
}
