package main

import (
	"fmt"
	"reflect"
)

//student结构体
type Student struct {
	Name  string `json:"name1" form:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄:%v 成绩:%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法...")
}

//打印字段
func PrintStructField(s interface{}) {

	//判断参数是不是结构体类型
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	//1、通过类型变量里面的Field可以获取结构体的字段
	field0 := t.Field(0)
	fmt.Printf("%#v \n", field0) //reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4adf20), Tag:"json:\"name\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	fmt.Println("字段名称：", field0.Name)
	fmt.Println("字段类型：", field0.Type)
	fmt.Println("字段Tag：", field0.Tag.Get("json"))
	fmt.Println("字段Tag：", field0.Tag.Get("form"))
	//2、通过类型变量里面的FieldByName可以获取结构体的字段
	fmt.Println("----------------------")
	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println("字段名称：", field1.Name)
		fmt.Println("字段类型：", field1.Type)
		fmt.Println("字段Tag：", field1.Tag.Get("json"))
	}

	//3、通过类型变量里面的NumField获取到该结构体有几个字段

	var fieldCount = t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性")

	//4、通过值变量获取结构体属性对应的值

	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println("----------------------")
	for i := 0; i < fieldCount; i++ {
		fmt.Printf("属性名称:%v 属性值:%v 属性类型:%v 属性Tag:%v\n", t.Field(i).Name, v.Field(i), t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}

}

//打印执行方法
func PrintStructFn(s interface{}) {

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}
	//1、通过类型变量里面的Method可以获取结构体的方法
	method0 := t.Method(0)    //和结构体方法的顺序没有关系，和结构体方法的ASCII有关系
	fmt.Println(method0.Name) //GetInfo
	fmt.Println(method0.Type) //func(main.Student) string

	fmt.Println("--------------------------")
	//2、通过类型变量获取这个结构体有多少个方法

	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name) //Print
		fmt.Println(method1.Type) //func(main.Student)
	}
	fmt.Println("--------------------------")
	//3、通过《值变量》执行方法 （注意需要使用值变量，并且要注意参数） v.Method(0).Call(nil) 或者v.MethodByName("Print").Call(nil)
	// v.Method(1).Call(nil)
	v.MethodByName("Print").Call(nil)

	info1 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info1)
	//4、执行方法传入参数 （注意需要使用《值变量》，并且要注意参数,接收的参数是[]reflect.Value的切片）

	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params) //执行方法传入参数

	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

	// 5、获取方法数量

	fmt.Println("方法数量:", t.NumMethod())

}

func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   15,
		Score: 98,
	}
	// PrintStructField(stu1)
	PrintStructFn(&stu1)

	fmt.Printf("%#v\n", stu1)
}
