package main

func main() {

	//1、整型和整型之间的转换
	// var a int8 = 20
	// var b int16 = 40
	// fmt.Println(int16(a) + b)

	//2、浮点型和浮点型之间的转换
	// var a float32 = 20
	// var b float64 = 40
	// fmt.Println(float64(a) + b)

	//3、整型和浮点型之间的转换
	// var a float32 = 20.23
	// var b int = 40
	// fmt.Println(a + float32(b))

	//注意：转换的时候建议从低位转换成高位，高位转换成低位的时候如果转换不成功就会溢出，和我们想的结果不一样。

}
