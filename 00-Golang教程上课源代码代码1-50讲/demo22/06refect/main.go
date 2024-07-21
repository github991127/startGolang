package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue(x interface{}) {
	// *x = 120  //invalid indirect of x (type interface {})

	// v, _ := x.(*int)
	// *v = 120 // invalid memory address or nil pointer dereferenc

	v := reflect.ValueOf(x)
	// fmt.Println(v.Kind()) //ptr
	// fmt.Println(v.Elem().Kind()) //int64
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("你好 go语言")
	}
}
func main() {
	var a int64 = 100
	reflectSetValue(&a)

	fmt.Println(a)

	var b string = "你好golang"
	reflectSetValue(&b)
	fmt.Println(b)

}
