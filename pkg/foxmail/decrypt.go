package foxmail

import (
	"bytes"
	"encoding/hex"
)

func DecryptPassword(encpwd string, v6 bool) (string, error) {
	// 第一步：定义不同版本的秘钥
	var a []byte
	var fc byte
	if v6 {
		a = []byte{'~', 'd', 'r', 'a', 'G', 'o', 'n', '~'}
		fc = 0x5A
	} else {
		a = []byte{'~', 'F', '@', '7', '%', 'm', '$', '~'}
		fc = 0x71
	}

	// 第二步：以字节为单位将16进制密文转成10进制
	b, err := hex.DecodeString(encpwd)
	if err != nil {
		return "", err
	}

	// 第三步：将第一个元素替换成与指定数异或后的结果
	c := b
	c[0] = c[0] ^ fc

	// 第四步：不断扩容拷贝自身
	if len(b) > len(a) {
		n := len(b) / len(a)
		if len(b)%len(a) > 0 {
			n += 1
		}
		a = bytes.Repeat(a, n)
	}

	count := len(b)
	d := make([]int, count)
	for i := 1; i < count; i++ {
		d[i-1] = int(b[i]) ^ int(a[i-1])
	}

	e := make([]byte, count)
	for i := 0; i < count; i++ {
		ec := int(c[i])
		if d[i]-ec < 0 {
			e[i] = byte(d[i] + 255 - ec)
		} else {
			e[i] = byte(d[i] - ec)
		}
	}
	//最后一个字符去掉
	return string(e[:count-1]), nil
}
