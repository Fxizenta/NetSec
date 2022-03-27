package main

import (
	"bytes"
	"fmt"
	"github.com/tjfoc/gmsm/sm4"
	"io"
	"os"
	"testing"
)

func pkcs7Padding(src []byte) []byte {
	padding := sm4.BlockSize - len(src)%sm4.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func xor(in, iv []byte) (out []byte) {
	if len(in) != len(iv) {
		return nil
	}

	out = make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i] ^ iv[i]
	}
	return
}

func BenchmarkSM4PaddingAndInit(b *testing.B) {
	f, err := os.Open("./test")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 关闭文件
	defer f.Close()
	buf := make([]byte, 1024*8) // 8k大小
	// n 代表从文件读取内容长度
	_, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {
		// 文件出错，同时没到结尾
		fmt.Println("err1 = ", err1)
		return
	}
	key := []byte("1234567890abcdef")
	for i := 0; i < b.N; i++ {
		if len(key) != sm4.BlockSize {
			//	return nil, errors.New("SM4: invalid key size " + strconv.Itoa(len(key)))
			//
		}
		var inData []byte
		inData = pkcs7Padding(buf)
		var IV = make([]byte, 16)
		iv := make([]byte, sm4.BlockSize)
		copy(iv, IV)
		_ = make([]byte, len(inData))
		_, err := sm4.NewCipher(key)
		if err != nil {
			print(err)
		}
	}
}

func BenchmarkSM4Encrypt(b *testing.B) {
	f, err := os.Open("./test")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 关闭文件
	defer f.Close()
	buf := make([]byte, 1024*8) // 8k大小
	// n 代表从文件读取内容长度
	_, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {
		// 文件出错，同时没到结尾
		fmt.Println("err1 = ", err1)
		return
	}
	key := []byte("1234567890abcdef")
	if len(key) != sm4.BlockSize {
	}
	var inData []byte

	inData = pkcs7Padding(buf)
	var IV = make([]byte, 16)
	iv := make([]byte, sm4.BlockSize)
	copy(iv, IV)
	out := make([]byte, len(inData))
	c, err := sm4.NewCipher(key)
	if err != nil {

	}
	for i := 0; i < len(inData)/16; i++ {
		in_tmp := xor(inData[i*16:i*16+16], iv)
		out_tmp := make([]byte, 16)
		for i := 0; i < b.N; i++ {
			c.Encrypt(out_tmp, in_tmp)
		}
		copy(out[i*16:i*16+16], out_tmp)
		iv = out_tmp
	}

}
