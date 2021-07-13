package gtils

import (
	"testing"
)

func TestMD5V(t *testing.T) {
	password := "123456"
	md5Passwd := MD5V(password)
	t.Log(md5Passwd)
}
