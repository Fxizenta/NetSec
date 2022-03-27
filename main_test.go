package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func BenchmarkAES(b *testing.B) {
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
	m := string(buf)
	for i := 0; i < b.N; i++ {
		AES(m, "hgfedcba87654321")
	}
}

func BenchmarkTEA(b *testing.B) {
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
	m := string(buf)
	for i := 0; i < b.N; i++ {
		TEA(m, "123adsf")
	}
}

func BenchmarkDES(b *testing.B) {
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
	m := string(buf)
	for i := 0; i < b.N; i++ {
		DES(m, "12345678")
	}
}

func BenchmarkSM4(b *testing.B) {
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
	m := string(buf)
	for i := 0; i < b.N; i++ {
		SM4(m, "1234567890abcdef")
	}
}
