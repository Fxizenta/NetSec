package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	_ "crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/tjfoc/gmsm/sm4"
	"github.com/ywj11407/tea"
	"io"
	"log"
	"os"
)

func main() {
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
	timebegin := TimeUse()
	fmt.Print("time begin:", timebegin)
	TEA(m, "123adsf")
	timeend := TimeUse()
	print("time end:", timeend)
	print("total use time: ", timeend-timebegin)

	timebegin2 := TimeUse()
	fmt.Print("time begin:", timebegin2)
	DES(m, "12345678")
	timeend2 := TimeUse()
	print("time end:", timeend2)

	timebegin3 := TimeUse()
	fmt.Print("time begin:", timebegin3)
	AES(m, "hgfedcba87654321")
	timeend3 := TimeUse()
	fmt.Print("time end:", timeend3)

	timebegin4 := TimeUse()
	fmt.Print("time begin:", timebegin4)
	SM4(m, "1234567890abcdef")
	timeend4 := TimeUse()
	fmt.Print("time end:", timeend4, "\n")
	print("tea total use time: ", timeend-timebegin, "\n")
	print("des total use time: ", timeend2-timebegin2, "\n")
	print("aes total use time: ", timeend3-timebegin3, "\n")
	print("sm4 total use time: ", timeend4-timebegin4, "\n")

}

func TimeUse() float64 {
	info, _ := cpu.Times(false)
	time := info[0].Total()
	return time
}

//-----------------------------------tea-------------------------------
func TEA(m string, sec string) {
	_ = tea.Encrypt(m, sec, 16)
	//fmt.Println("scontent:", scode)
}

//-----------------------------------tea-------------------------------

//--------------------------------------des-------------------------------
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func Encrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func DES(m string, k string) {
	key := []byte(k)

	//desCipher, err := des.NewCipher(key)
	//if err != nil {
	//	panic(err)
	//}
	//inputData := []byte(m)
	_, err := Encrypt(m, key)
	if err != nil {
		log.Fatal(err)
	}
	//
	//out := make([]byte, len(inputData))
	//desCipher.Encrypt(out, inputData)
	//fmt.Printf("Encrypted data : %#v\n", out)

}

//-------------------------------------------------------des--------------------

//-------------------------------------------------------sm4-----------------------------
func SM4(m string, k string) {
	key := []byte(k)
	//fmt.Printf("key = %v\n", key)
	data := []byte(m)
	err := sm4.WriteKeyToPemFile("key.pem", key, nil)
	if err != nil {
		print("WriteKeyToPem error")
	}
	key, err = sm4.ReadKeyFromPemFile("key.pem", nil)
	//fmt.Printf("key = %v\n", key)
	//fmt.Printf("data = %x\n", data)
	//ecbMsg, err := sm4.Sm4Ecb(key, data, true)
	//if err != nil {
	//	print("sm4 enc error:%s", err)
	//	return
	//}
	//fmt.Printf("ecbMsg = %x\n", ecbMsg)
	//iv := []byte("0000000000000000")
	//err = sm4.SetIV(iv)
	//fmt.Printf("err = %v\n", err)
	//ecbDec, err := sm4.Sm4Ecb(key, ecbMsg, false)
	//if err != nil {
	//	print("sm4 dec error:%s", err)
	//	return
	//}
	//fmt.Printf("ecbDec = %x\n", ecbDec)
	//if !testCompare(data, ecbDec) {
	//	print("sm4 self enc and dec failed")
	//}
	_, err = sm4.Sm4Cbc(key, data, true)
	if err != nil {
		print("sm4 enc error:%s", err)
	}
	//fmt.Printf("cbcMsg = %x\n", cbcMsg)
	//_, err = sm4.Sm4Cbc(key, cbcMsg, false)
	//if err != nil {
	//	print("sm4 dec error:%s", err)
	//	return
	//}
	//fmt.Printf("cbcDec = %x\n", cbcDec)
	//if !testCompare(data, cbcDec) {
	//	print("sm4 self enc and dec failed")
	//}

	//cbcMsg, err = sm4.Sm4CFB(key, data, true)
	//if err != nil {
	//	print("sm4 enc error:%s", err)
	//}
	//fmt.Printf("cbcCFB = %x\n", cbcMsg)
	//
	//cbcCfb, err := sm4.Sm4CFB(key, cbcMsg, false)
	//if err != nil {
	//	print("sm4 dec error:%s", err)
	//	return
	//}
	//fmt.Printf("cbcCFB = %x\n", cbcCfb)
	//
	//cbcMsg, err = sm4.Sm4OFB(key, data, true)
	//if err != nil {
	//	print("sm4 enc error:%s", err)
	//}
	//fmt.Printf("cbcOFB = %x\n", cbcMsg)
	//
	//cbcOfc, err := sm4.Sm4OFB(key, cbcMsg, false)
	//if err != nil {
	//	print("sm4 dec error:%s", err)
	//	return
	//}
	//fmt.Printf("cbcOFB = %x\n", cbcOfc)
}

func testCompare(key1, key2 []byte) bool {
	if len(key1) != len(key2) {
		return false
	}
	for i, v := range key1 {
		if i == 1 {
			//fmt.Println("type of v", reflect.TypeOf(v))
		}
		a := key2[i]
		if a != v {
			return false
		}
	}
	return true
}

//------------------------------------------------------sm4------------------

//--------------------------------------------------------aes-------------------
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// 加密
func encryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

func AES(m string, k string) {
	d := []byte(m)
	key := []byte(k)
	//fmt.Println("加密前:", string(d))
	_, err := encryptAES(d, key)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println("加密后:", string(x1))
}

//----------------------------------------------aes--------------------------------------
