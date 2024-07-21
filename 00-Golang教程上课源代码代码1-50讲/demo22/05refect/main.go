package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	// v.Kind() //获取种类
	kind := v.Kind()

	switch kind {
	case reflect.Int64:
		fmt.Printf("int类型的原始值%v,计算后的值是%v \n", v.Int(), v.Int()+10)
	case reflect.Float32:
		fmt.Printf("Float32类型的原始值%v\n", v.Float())
	case reflect.Float64:
		fmt.Printf("Float64类型的原始值%v\n", v.Float())
	case reflect.String:
		fmt.Printf("string类型的原始值%v\n", v.String())
	default:
		fmt.Printf("还没有判断这个类型\n")
	}

}
func main() {
	var a int64 = 100
	var b float32 = 12.3
	var c string = "你好golang"
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
}
