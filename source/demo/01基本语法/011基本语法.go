package main

import (
	"fmt"
	"strconv"
	"strings"
	//pname "fmt"// 导入，包名命名为pname
	//_ "fmt"// 匿名导入，添加_，表示不暂时使用包名
	//. "fmt"// 方法全导入，不需要包名调用
	//"bao/calc" // 自定义包，需要导入项目根目录下的相对路径，若其中有Init函数，则会自动调用
	//"github.com/shopspring/decimal"//第三方包

	//https://pkg.go.dev查询包文档并写入导入语句
	//命令行go mod init 【项目名称】自动生成go.mod文件
	//命令行go mod tidy 自动导入依赖包
)

// 声明与输出
func getUserinfo() (string, int) {
	return "匿名变量", 3
}
func f1() {
	//	Print无空格，不换行 Println有空格，换行
	fmt.Println("Hi", "golang") //Hi golang
	fmt.Print("Hi", "golang\n") //Higolang

	//	Printf不换行
	var a int = 2
	b := 3
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v，b的类型是%T\n", b, b)

	// 变量声明-批量
	var (
		username = "张三"
		age      = 1
	)
	fmt.Println(username, age)

	// 变量声明-批量短，（函数外部禁止使用短声明，如全局变量）
	username2, age2 := "张三", 2
	fmt.Println(username2, age2)

	// 匿名变量_。在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量
	var _, age3 = getUserinfo() // 匿名变量不占用命名空间，不会分配内存，不存在重复声明
	fmt.Println(age3)

	// 常量声明const
	const pi = 3.14159
	fmt.Println(pi)

	// 常量声明-批量，如果省略了值则表示和上面一行的值相同。iota代表从0开始自增int
	//const (
	//	n1 = 100
	//	n2
	//	n3
	//	n4
	//)
	const (
		n1 = iota
		n2
		n3
		n4
	)
	fmt.Println(n1, n2, n3, n4)

}

// 数据类型
func f2() {
	// float
	f1 := 3.13513
	fmt.Printf("%.2f,%T\n", f1, f1)
	f1 = 3.13513e2 // 科学计数法
	fmt.Printf("%v,%T\n", f1, f1)

	// str-多行
	str1234 := `str1
	str2
		str3
	str4
	`
	fmt.Println(str1234, len(str1234))
	// 分割字符串strings.Split
	var str1 = "123-456-789"
	arr := strings.Split(str1, "-")
	fmt.Println(arr) //[123 456 789]

	// 连接字符串strings.Join(a[]string, sep string)
	str2 := strings.Join(arr, "*")
	fmt.Println(str2)                       //123*456*789
	arr = []string{"php", "java", "golang"} //切片
	fmt.Println(arr)                        //[php java golang]
	str3 := strings.Join(arr, "-")
	fmt.Printf("%v,%T\n", str3, str3)
	// 字符串判断strings
	flag := strings.Contains(str1, str2) //包含
	flag = strings.HasPrefix(str1, str2) //前缀
	flag = strings.HasSuffix(str1, str2) //后缀
	i := strings.Index(str1, str2)       //子串出现的位置，无则返回-1
	fmt.Println(flag, i)

	s := "你好 golang"

	// 输出字符串
	//for i := 0; i < len(s); i++ { //按byte字节
	//	fmt.Printf("%v(%c) ", s[i], s[i])
	//}
	//fmt.Println()
	//for _, v := range s { //按rune字符
	//	fmt.Printf("%v(%c) ", v, v)
	//}

	// 修改字符串
	byteStr := []byte(s)
	byteStr[9] = 'S'
	fmt.Println(string(byteStr))

	runeStr := []rune(s)
	runeStr[0] = '大'
	fmt.Println(string(runeStr))

}

// 类型转换
func f3() {
	// Others-String
	// Sprintf 使用中需要注意转换的格式 int 为%d  float 为%f  bool 为%t   byte 为%c
	var i int = 20
	var f float64 = 12.456
	var t bool = true
	var b byte = 'a'

	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("值：%v 类型：%T\n", str1, str1)

	str2 := fmt.Sprintf("%.2f", f)
	fmt.Printf("值：%v 类型：%T\n", str2, str2)

	str3 := fmt.Sprintf("%t", t)
	fmt.Printf("值：%v 类型：%T\n", str3, str3)

	str4 := fmt.Sprintf("%c", b)
	fmt.Printf("值：%v 类型：%T\n", str4, str4)

	//strconv
	str5 := strconv.FormatInt(int64(i), 10) //参数1：int64 的数值;参数2：传值int类型的进制
	fmt.Printf("值：%v 类型：%T\n", str5, str5)

	str6 := strconv.FormatFloat(float64(f), 'f', 4, 32)
	fmt.Printf("值：%v 类型：%T\n", str6, str6)
	//参数 1：要转换的值
	//参数 2：格式化类型 'f'（-ddd.dddd）、
	//	 'b'（-ddddp±ddd，指数为二进制）、
	//	 'e'（-d.dddde±dd，十进制指数）、
	//	 'E'（-d.ddddE±dd，十进制指数）、
	//	 'g'（指数很大时用'e'格式，否则'f'格式）、
	//	 'G'（指数很大时用'E'格式，否则'f'格式）。
	//参数 3: 保留的小数点 -1（不对小数点格式化）
	//参数 4：格式化的类型 传入 64  32

	//String-Others
	//ParseInt，参数1：string数据，参数2：进制，参数3：位数 32 64 16
	str7 := "123456"
	num, _ := strconv.ParseInt(str7, 10, 64)
	fmt.Printf("%v--%T\n", num, num)

	//ParseFloat，参数1：string数据，参数2：位数 32 64
	str8 := "123456.333"
	num2, _ := strconv.ParseFloat(str8, 64)
	fmt.Printf("%v--%T\n", num2, num2)
}

// 数据运算
func f4() {
	fmt.Println(10 / 3)
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3) //以前者除数正负号为准
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)
	year := 2024
	fmt.Println(year%4 == 0 && year%100 != 0 || year%400 == 0)

	a := 5                     //  101
	b := 2                     //  010
	fmt.Println("a&b=", a&b)   // 000   值0
	fmt.Println("a|b=", a|b)   // 111  值7
	fmt.Println("a^b=", a^b)   // 111  值7。^异或，不同为真
	fmt.Println("a<<b=", a<<b) // 10100  值20   5乘2的2次方
	fmt.Println("a>>b=", a>>b) // 1 值1    5除2的2次方
}
func main() {
	//f1() // 声明与输出
	//f2() // 数据类型
	f3() // 类型转换
	//f4() //数据运算
}
