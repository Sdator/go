package main

import (
	"file"
	"fmt"
)

type 装备 struct {
	武器 string
	腰带 string
}

type 单位 struct {
	名字 string
	等级 int
	a  装备
}

func main() {
	file.OpenFile("hello")

	// 李 := 单位{名字: "李", 等级: 999}
	// 套 := 装备{武器: "双截棍", 腰带: "长老腰带"}
	李 := 单位{名字: "李", 等级: 999, 武器: "双截棍", 腰带: "长老腰带"}

	fmt.Println(李)
}
