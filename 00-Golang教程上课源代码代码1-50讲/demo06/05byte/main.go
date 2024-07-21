package main

func main() {

	//1、golang中定义字符  字符属于int类型

	// var a = 'a'
	// fmt.Printf("值：%v  类型:%T", a, a) //值：97  类型:int32
	//当我们直接输出 byte（字符）的时候输出的是这个字符对应的码值

	//2、原样输出字符
	// var a = 'a'
	// fmt.Printf("值：%c  类型:%T", a, a) //值：a  类型:int32

	//3、定义一个字符串输出字符串里面的字符

	// var str = "this"
	// fmt.Printf("值：%v 原样输出%c  类型:%T", str[2], str[2], str[2])

	//4、一个汉子占用 3个字节(utf-8), 一个字母占用一个字节

	// unsafe.Sizeof() 没法查看string类型数据所占用的存储空间

	// var str = "this" //占用4个字节
	// fmt.Println(len(str))

	// var str = "你好go" //8
	// fmt.Println(len(str))

	// 5、定义一个字符 字符的值是汉子

	//golang中汉子使用的是utf8编码 编码后的值就是int类型

	// var a = '国'
	// fmt.Printf("值：%v  类型:%T\n", a, a) //Unicode编码10进制  值：22269  类型:int32
	// fmt.Printf("值：%c  类型:%T", a, a)

	//6、通过循环输出字符串里面的字符
	// s := "你好 golang"
	// for i := 0; i < len(s); i++ {   //byte
	// 	fmt.Printf("%v(%c) ", s[i], s[i])
	// }

	// s := "你好 golang"
	// for _, v := range s { //rune
	// 	fmt.Printf("%v(%c) ", v, v)
	// }

	//7、修改字符串
	// s1 := "big"
	// byteStr := []byte(s1)
	// byteStr[0] = 'p'
	// fmt.Println(string(byteStr))

	// s1 := "你好golang"
	// runeStr := []rune(s1)
	// runeStr[0] = '大'
	// fmt.Println(string(runeStr))

}
