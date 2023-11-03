package cryptox

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Md5[M string | []byte](s M) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func Sha256[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func Sha224[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(s)))
}

func Sha512[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}

func Sha384[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum384([]byte(s)))
}
