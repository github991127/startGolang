package main

import "fmt"

// 条件语句
func f1() {
	var score = 75
	if score >= 90 { // 可将score写成条件语句中的局部变量。if score := 75; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}

// 循环语句
func f2() {
	//普通for
	for i := 1; i <= 4; i++ { //i := 1，定义语句可省略
		fmt.Println(i)
	}

	//双层for
lable1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				//break lable1 //跳出到lable位置，可以跳多层
				continue lable1 //继续到lable位置，可以跳多层
				goto lable2     //传送到lable位置，可以跳语句
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}
lable2:
	//while-for
	i := 1
	for i <= 4 { //i <= 10，判断语句可省略，默认为true
		fmt.Println(i)
		i++ // fmt.Println(i++)错误，i++是语句
	}
	//for-in
	var str = "你好golang"
	for k, v := range str {
		fmt.Printf("key=%v val=%c\n", k, v) //key=0 val=你		key=3 val=好		key=6 val=g		key=7 val=o
	}

	var arr = []string{"php", "java", "node", "golang"}
	for _, val := range arr {
		fmt.Println(val)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}

// 选择语句
func f3() {
	//普通case
	var score = "D"
	switch score { //可将score写成条件语句中的局部变量。switch score := "D"; score {
	case "A", "B", "C":
		fmt.Println("及格")
	case "D":
		fmt.Println("不及格")
	}
	//表达式case
	var age = 18
	switch {
	case age < 24:
		fmt.Println("好好学习")
		//fallthrough //强制执行下一条
	case age >= 24 && age <= 60:
		fmt.Println("好好赚钱")
	case age > 60:
		fmt.Println("注意身体")
	default:
		fmt.Println("输入错误")
	}

}

func main() {
	//f1() // 条件语句
	f2() // 循环语句
	//f3() // 选择语句
}
