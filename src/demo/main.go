package main

import "fmt"

func main() {
	// 变量定义  定义的变量必须要使用不然编译器报错
	//var a int
	// 变量推断类型
	// var b = 10
	// //自定义定义并推断  省略var和定义类型
	// c := 10
	// //连续赋值
	// c, d := 1001, 555
	// var e, f int = 4, 5
	// var (
	// 	aa = 12
	// 	bb = "sadsad"
	// 	cc int
	// )
	// complex(1, 73) 实际上等于 (1+73i)
	// a := complex(1, 73)
	// a := 1
	// b := 5.7
	// fmt.Println("aaaa", a+int(b))
	// %T输出变量类型
	// %d %f 输出整数
	// %v 输出变量值
	// fmt.Printf("aaaa %T %d", a, unsafe.Sizeof(a))
	// 自定义类型 别名
	// type abc int
	// var a abc = 6666
	// fmt.Printf("类型： %T  长度：%d \n", a, unsafe.Sizeof(a))
	// fmt.Println("aaaa", math.Sqrt(4))
	// fmt.Printf("type %T value %v", a, a)
	// 数组定义
	// var _ [3]int
	// a := [3]int{12, 78, 50}
	//===
	darr := [...]int{1, 2, 3, 4, 5}
	t := darr[1:2]
	fmt.Println(t, cap(t))
	t = t[:3]
}

func abc(a, b int) (aa int) {
	aa = 12 + a
	return
}
