package main

import (
	"fmt"
	// "io/ioutil"
	// 缓冲
	"bufio"
	"log"
	"os"
	pf "path/filepath"
)

// 获取程序自身路径
var _, baseFile = pf.Split(os.Args[0])

// 定义一个类型
// type aFiles os.File

func main() {
	// fmt.Println(os.Args)
	读文件(".")
	// fmt.Println()
}

// 方法
// delByte test
// func (file *aFiles) delByte() bool {
// 	file.Read([]byte{0, 0, 0, 0})

// 	return true
// }

// 处理函数 递归处理函数
func 处理函数(path string, info os.FileInfo, err error) error {

	// fmt.Printf("文件路径：%s\tos路径:%s\t是否绝对路径：%s", path, baseFile, )

	// 过滤 无需使用的文件和无效的路径
	// 忽略程序自身文件 并 忽略 路径符号 如 . ..
	if path != baseFile && !info.IsDir() && info.Name() != "file.go" {
		文件句柄, err := os.OpenFile(path, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("文件打开失败，错误代码：", err)
			return err
		}
		defer 文件句柄.Close()

		// 二进制, _ := 文件句柄.ReadAt(make([]byte, 20))
		// 字节流, _ := ioutil.ReadAll(文件句柄)
		文件句柄.Seek(0x20, 0)
		// bufio.NewReaderSize(文件句柄, info.Size())
		// 缓冲 := bufio.NewReader(文件句柄)
		// 读取指定字节
		// fmt.Println(缓冲.Peek(0x10))
		// fmt.Println(缓冲.Buffered())

		缓冲 := bufio.NewWriter(文件句柄)
		fmt.Println(缓冲)

		// 缓冲.WriteTo()
		// buf := make([]byte, 0x20)
		// a, _ := 缓冲.Read(buf)

		// ioutil.WriteFile(path+".test", 缓冲, 0777)

		// 文件句柄.Write()

		// 多次读取指针会递增
		// fmt.Println(缓冲.ReadByte())
		// fmt.Println(缓冲.ReadByte())
		// 截断只会保留前面的
		// 文件句柄.Truncate(0x20)

		// fmt.Println(字节流[:0x22])

		// 方法一 创建一个数组再生成一个切片
		// var str [0x20]byte
		// aa := str[:]
		// 方法二 使用make直接生成指定大小的切片
		// num, _ := 文件句柄.Write(make([]byte, 0x20))

		// fmt.Println(num)

	}

	return err
}

// 读文件 读写并修改
func 读文件(path string) bool {
	//递归操作
	err := pf.Walk(path, 处理函数)
	if err != nil {
		fmt.Printf("列举文件出错，错误代码：%v", err)
		log.Fatal(err)

		return false
	}
	return true
}
