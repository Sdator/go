package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	pf "path/filepath"
)

// 获取程序自身路径
var _, baseFile = pf.Split(os.Args[0])

func main() {
	var file string
	flag.StringVar(&file, "f", "(文件名)", "指定一个文件")
	flag.Parse()
	if file == "(文件名)" {
		println("指定一个文件如： -f xxx.bin")
		return
	}
	写文件(file)
}

// 写文件 读写并修改
func 写文件(path string) bool {

	// 打开文件
	文件句柄, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("文件打开失败，错误代码：", err)
		return false
	}
	// 延迟释放文件句柄
	defer 文件句柄.Close()

	// 修改文件读写位置
	// 跳过前面0x20个字节 删除文件头
	文件句柄.Seek(0x20, 0)

	// 从当前位置 读入文件所有字节
	字节流, _ := ioutil.ReadAll(文件句柄)

	文件句柄.Truncate(0)
	文件句柄.Seek(0, 0)

	// if num, err = 文件句柄.Write(字节流); err != nil {
	// 	fmt.Println("文件写入失败，错误代码：", err)
	// 	return false
	// }

	for _, err = 文件句柄.Write(字节流); err != nil; {
		fmt.Println("文件写入失败，错误代码：", err)
		return false
	}

	println("删除成功")

	return true

}

// 处理函数 递归处理函数
func 处理函数(path string, info os.FileInfo, err error) error {

	// 过滤 无需使用的文件和无效的路径
	// 忽略程序自身文件 并 忽略 路径符号 如 . ..
	if !(info.IsDir() || path == baseFile || info.Name() == "file.go") {

		// 打开文件
		文件句柄, _ := os.OpenFile(path, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("文件打开失败，错误代码：", err)
			return err
		}
		// 延迟释放文件句柄
		defer 文件句柄.Close()

		// 修改文件读写位置
		// 跳过前面0x20个字节 删除文件头
		文件句柄.Seek(0x20, 0)

		// 从当前位置 读入文件所有字节
		字节流, _ := ioutil.ReadAll(文件句柄)

		// 抹掉文件
		// 执行后 之前保存的数据流也会被抹掉
		// 同时也要把文件指针清零 不然留下空白
		文件句柄.Truncate(0)
		文件句柄.Seek(0, 0)

		// 输入流 := bufio.NewReader(字节流)
		// 从文件句柄创建输出流
		// 输出流 := bufio.NewWriter()
		// 输出流 := bufio.NewWriter(文件句柄)
		num, _ := 文件句柄.Write(字节流)
		fmt.Println("写入了多少个字节:", num)

		// 输出流.Write(字节流)
		// 读取文件
		// 二进制 := make([]byte, 10)
		// 二进制, _ := 字节流.Peek(10)
		// num, _ := 字节流.Read(二进制)

		// num, _ := 输入流.WriteTo(输出流)

		// fmt.Printf("输入流可读取的字节数：%d\n输出流中已使用的字节数：%d\n输出流中未使用的字节数：%d\n", 输入流.Buffered(), 输出流.Buffered(), 输出流.Available())
		// fmt.Println("传输了多少字节到输出流:", num)

		// 字节流, _ := ioutil.ReadAll(缓冲流)
		// 写入数量, _ := 缓冲流.Write(字节流[:16])
		// 清空文件
		// 输出流.Flush()

	}

	// fmt.Printf("文件路径：%s\tos路径:%s\t是否绝对路径：%s", path, baseFile, )

	// 二进制, _ := 文件句柄.ReadAt(make([]byte, 20))
	//
	// //
	// fmt.Println(info.Name(), 字节流)

	// bufio.NewReaderSize(文件句柄, info.Size())
	// 缓冲 := bufio.NewReader(文件句柄)
	// 读取指定字节
	// fmt.Println(缓冲.Peek(0x10))
	// fmt.Println(缓冲.Buffered())

	// 缓冲 := bufio.NewWriter(文件句柄)
	// fmt.Println(缓冲)

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
