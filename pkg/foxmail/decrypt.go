package foxmail

import (
	"bytes"
	"encoding/hex"
)

func DecryptPassword(strHash string, bVersion6 bool) (string, error) {
	//var strPlainPassword []byte
	// 第一步：定义不同版本的秘钥
	//vector<int> a(8), b, c;
	var a []byte
	var fc byte
	if bVersion6 {
		a = []byte{'~', 'd', 'r', 'a', 'G', 'o', 'n', '~'}
		fc = 0x5A
	} else {
		a = []byte{'~', 'F', '@', '7', '%', 'm', '$', '~'}
		fc = 0x71
	}

	// 第二步：以字节为单位将16进制密文转成10进制
	b, err := hex.DecodeString(strHash)
	if err != nil {
		return "", err
	}

	// 第三步：将第一个元素替换成与指定数异或后的结果
	c := b
	c[0] = c[0] ^ fc

	// 第四步：不断扩容拷贝自身
	for len(b) > len(a) {
		a = bytes.Repeat(a, 2)
	}

	count := len(b)
	d := make([]byte, count)
	for i := 1; i < count; i++ {
		d[i-1] = b[i] ^ a[i-1]
	}

	e := make([]byte, count)
	for i := 0; i < count; i++ {
		if int(d[i]-c[i]) < 0 {
			e[i] = d[i] + 255 - c[i]
		} else {
			e[i] = d[i] - c[i]
		}
		//fmt.Println(e[i])
	}
	//最后一个字符去掉
	return string(e[:count-1]), nil
}
