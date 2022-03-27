package main

import (
	"crypto/des"
	"fmt"
	"io"
	"os"
	"testing"
)

func BenchmarkDESNewCipher(b *testing.B) {
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
	key := []byte("12345678")
	for i := 0; i < b.N; i++ {
		_, err := des.NewCipher(key)
		if err != nil {
			panic(err)
		}
		_ = make([]byte, len(buf))
	}
}

func BenchmarkDESEncrypt(b *testing.B) {
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
	key := []byte("12345678")
	desCipher, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	out := make([]byte, len(buf))
	for i := 0; i < b.N; i++ {
		desCipher.Encrypt(out, buf)
	}
}
