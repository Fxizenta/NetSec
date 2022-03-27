package main

var m string

//func BenchmarkAESNewCipher(b *testing.B) {
//	f, err := os.Open("./test")
//	if err != nil {
//		fmt.Println("err = ", err)
//		return
//	}
//	// 关闭文件
//	defer f.Close()
//	buf := make([]byte, 1024*8) // 8k大小
//	// n 代表从文件读取内容长度
//	_, err1 := f.Read(buf)
//	if err1 != nil && err1 != io.EOF {
//		// 文件出错，同时没到结尾
//		fmt.Println("err1 = ", err1)
//		return
//	}
//	key := []byte("hgfedcba87654321")
//	for i := 0; i < b.N; i++ {
//		block, err := aes.NewCipher(key)
//		if err != nil {
//			print("error")
//		}
//		buf = padding(buf, block.BlockSize())
//		_ = cipher.NewCBCEncrypter(block, key)
//	}
//
//}
//
//func BenchmarkAESNewCBCEncrypter(b *testing.B) {
//	f, err := os.Open("./test")
//	if err != nil {
//		fmt.Println("err = ", err)
//		return
//	}
//	// 关闭文件
//	defer f.Close()
//	buf := make([]byte, 1024*8) // 8k大小
//	// n 代表从文件读取内容长度
//	_, err1 := f.Read(buf)
//	if err1 != nil && err1 != io.EOF {
//		// 文件出错，同时没到结尾
//		fmt.Println("err1 = ", err1)
//		return
//	}
//	key := []byte("hgfedcba87654321")
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		print("error")
//	}
//	buf = padding(buf, block.BlockSize())
//	for i := 0; i < b.N; i++ {
//
//		_ = cipher.NewCBCEncrypter(block, key)
//	}
//
//}
//
//func BenchmarkAESCryptBlocks(b *testing.B) {
//	f, err := os.Open("./test")
//	if err != nil {
//		fmt.Println("err = ", err)
//		return
//	}
//	// 关闭文件
//	defer f.Close()
//	buf := make([]byte, 1024*8) // 8k大小
//	// n 代表从文件读取内容长度
//	_, err1 := f.Read(buf)
//	if err1 != nil && err1 != io.EOF {
//		// 文件出错，同时没到结尾
//		fmt.Println("err1 = ", err1)
//		return
//	}
//	key := []byte("hgfedcba87654321")
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		print("error")
//	}
//	buf = padding(buf, block.BlockSize())
//	blockMode := cipher.NewCBCEncrypter(block, key)
//	for i := 0; i < b.N; i++ {
//		blockMode.CryptBlocks(buf, buf)
//	}
//
//}
