记录
当使用var定义了变量后没有定义类型 系统会自动判定
当使用:=进行赋值可以忽略定义和类型 系统会自动判定 不能重复对已存在的变量进行赋值 必须包含一个未定义的变量才能使用


//=========
// 变量
//=========
	// 多重定义
	func main() {
		// 可以进行多个变量赋值
		var (
			name   = "naveen"
			age    = 29
			height int
		)
	}
	// 普通定义不使用
	var name string
	var a, b int
	name="1321"

	//让系统自动判断类型
	var name = "54654"

	// 连续定义 使用了:= 可省略定义var和类型
	
	name, age := "naveen", 29 


	// 空白符 省略一个返回值
	x, _ := Getxy()



//=========
// 函数
//=========
	//参数定义
	func test(num int) {...}
	// 忽略类型定义
	func test(num int,a int) {...}
	等同
	func test(num,a int) {...}	//如果两个变量的类型相同只需在最后一个变量后面定义一次类型即可

	// 数组的
	func test(num [5]int) {...}

	// 带返回值
	func test(num int) bool {...}

	// 多个返回值
	func test(num int) (bool,int) {...}

	// 命名返回值
	// 自动返回变量 定义方式和定义参数一样
	func test(num int) (abc int) { return }
	func test(a, b int) (aa, bb int) {
		//系统会自动在内部定义返回值中的变量 也就是 var aa,bb int
		aa = 12 + a
		bb = 1231
		return
	}



//=========
//	for循环
//=========

	//传统的方法
	for i := 0; i < len(a); i++ { // looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}

	//枚举数组用
	for i, v := range a {//range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}

	// 数组会返回 索引和值 如果只想取一个的时候使用
	//使用 _ 符号作为变量名称 可以忽略某个变量的传值
	for _, v := range a {
		// ignores index
	}

	// 多维数组的枚举方法
	func printarray(a [3][2]string) {
		for _, v1 := range a {
			for _, v2 := range v1 {
				fmt.Printf("%s ", v2)
			}
			fmt.Printf("\n")
		}
	}

//=========
//	数组创建和赋值
//=========
	//先定义在赋值
	var darr[3]int
	darr[0]=1
	darr[1]=2
	darr[2]=33

	// 定义赋值同时进行 没有在前面定义变量 系统会自动识别
	var darr = [3]int{1, 2, 3}

	// 创建一个没有规定长度的数组  
	// 使用了 := 进行赋值系统会自动计算 并省略了关键字var定义 已经定义过的无法使用这种方式 必须存在一个未定义的
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}

	//数组赋值方法1  类似json
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, 		// 最后一句要加上逗号
	}

	//赋值方法2	传统方式
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"



//=========
//	切片
//=========
只要切片在内存中，数组就不能被垃圾回收
	数组名称[:]		获取整个数组的引用
	数组名称[n:n]	获取数组部分数据的引用	
	len(fruitslice)		获取长度
	cap(fruitslice)		获取容量	一般用于切片	容量 等于 底层数组长度-切片开始索引
	cars = append(cars, "Toyota")		给切片扩容新的数据  切片的容量会自动扩大一倍	
	i := make([]int, 5, 5)				创建一个空的切片 长度、容量

	//从一个数组中分割一部分数据
	func main() {
		a := [5]int{76, 77, 78, 79, 80}	//创建变量
		c := a[:]
		var b []int = a[1:4] 			// 创建切片 从数组索引1开始到4结束 取出来就是77 78 79
		fmt.Println(b)
	}

	func main() {
		c := []int{6, 7, 8} 	// 直接创建切片
		fmt.Println(c)
	}


	// 新定义的切片
	var names []string //zero value of a slice is nil
	// 由于切片没有赋值 默认是nil 可以用来判断切片是否为空
    if names == nil {
        fmt.Println("slice is nil going to append")
        names = append(names, "John", "Sebastian", "Vinay")
        fmt.Println("names contents:",names)
	}
	

	func main() {
		// 定义数组
		darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
		// 切数组2-5部分数据到新的变量中 形成引用关系
		dslice := darr[2:5]
		// 枚举切片
		for i := range dslice {
			// 递增数值 由于是引用关系 这些数值改变会返回到初始数组中
			dslice[i]++
		}
	}

	func main() {
		//定义了一个数组 有7个
		fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
		//引用了1-3索引
		fruitslice := fruitarray[1:3]
		//这时候长度是2 容量是 7-1=6 从索引1开始计算一直算到底层数组最后一个索引
		fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
		//这里对切片进行了重新定义 使他从原来的1-3变成了1-3 再加上6 实际上这里是对长度进行了编排  
		// fruitslice{"orange", "grape"} 后面还有一部分其实是隐藏的也就是没有引用到实际是存在的 "mango", "water melon", "pine apple", "chikoo"}
		//这个时候他在切片中的索引位置并非原来数组的位置 然后自身从0开始算起 比如 chikoo在切片的空间中索引是5  但在原来的空间中索引是6
		//所以这里的 cap(fruitslice) 等于6 获取索引-1就是 fruitslice[0-5] 反射到原始数组中就是  fruitarray[1-6]
		fruitslice = fruitslice[:cap(fruitslice)] // re-slicing furitslice till its capacity
		fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
	}

	//切片组合
	func main() {
		veggies := []string{"potatoes", "tomatoes", "brinjal"}
		fruits := []string{"oranges", "apples"}
		//使用了 变量名... 来合并切片
		food := append(veggies, fruits...)
		fmt.Println("food:",food)
	}

	// 切片作为参数传递到函数中
	// 在函数中如果修改了切片也能会反射到外面的原始数组中
	func subtactOne(numbers []int) {
		for i := range numbers {
			numbers[i] -= 2
		}
	}
	func main() {
		nos := []int{8, 7, 6}
		subtactOne(nos)                               // function modifies the slice
		fmt.Println("slice after function call", nos) // modifications are visible outside
	}


	//多维切片 和多维数组差不多
	func main() {  
		// 定义了一个 2维切片 
		pls := [][]string {
			   {"C", "C++"},
			   {"JavaScript"},
			   {"Go", "Rust"},
			   }
		//枚举切片 和数组一样
	   for _, v1 := range pls {
		   for _, v2 := range v1 {
			   fmt.Printf("%s ", v2)
		   }
		   fmt.Printf("\n")
	   }
   }


   // 切片拷贝释放原始数组
	func countries() []string {
		// 创建了一个切片
		countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
		// 取切片开头到最后两个数据之间数据
		neededCountries := countries[:len(countries)-2]
		//创建了一个空的切片 容量为上一个切片的长度
		countriesCpy := make([]string, len(neededCountries))
		//复制neededCountries切片中的数据到countriesCpy中
		//这样我们只需要关注新的切片countriesCpy 他不会和原来的切片数组有引用关系, 原始数组就能被释放
		copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
		return countriesCpy
	}
	func main() {
		countriesNeeded := countries()
		fmt.Println(countriesNeeded)
	}