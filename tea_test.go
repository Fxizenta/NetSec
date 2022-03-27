package main

import (
	"fmt"
	"github.com/ywj11407/tea"
	"io"
	"os"
	"testing"
)

func BenchmarkTEAByteAppendandInit(b *testing.B) {
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
		var resultLength = len(buf)
		var mol = resultLength % 8
		if mol != 0 {
			resultLength = resultLength + 8 - mol
			for i := 0; i < 8-mol; i++ {
				buf = append(buf, byte(0))
			}
		}
		_ = tea.ValidateKey3(key)
		_ = make([]uint32, 2)
		_ = make([]uint32, 2)
		_ = make([]byte, resultLength)
		var _ = resultLength
		var _ = 0
		var _ = 0
	}
}

func BenchmarkTEAEnRound(b *testing.B) {
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
	var resultLength = len(buf)
	var mol = resultLength % 8
	if mol != 0 {
		resultLength = resultLength + 8 - mol
		for i := 0; i < 8-mol; i++ {
			buf = append(buf, byte(0))
		}
	}
	k := tea.ValidateKey3(key)
	v := make([]uint32, 2)
	o := make([]uint32, 2)
	result := make([]byte, resultLength)
	var convertTimes = resultLength
	var next = 0
	var times = 0
	rounds := 16

	for ; times < convertTimes; times += 8 {
		next = times + 4
		v[0] = tea.Byte2int(buf, times)
		v[1] = tea.Byte2int(buf, next)
		for i := 0; i < b.N; i++ {
			o = tea.En(v, k, rounds)
		}
		tea.Int2byte(o[0], result, times)
		tea.Int2byte(o[1], result, next)
	}

}
