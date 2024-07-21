package main

import (
	"fmt"
	"sort"
)

func main() {

	//1、sort升序排序
	// intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	// float8List := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
	// stringList := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}

	// sort.Ints(intList)
	// sort.Float64s(float8List)
	// sort.Strings(stringList)

	// fmt.Println(intList)
	// fmt.Println(float8List)
	// fmt.Println(stringList)

	//2、sort降序排序

	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	floatList := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}

	sort.Sort(sort.Reverse(sort.IntSlice(intList)))

	sort.Sort(sort.Reverse(sort.Float64Slice(floatList)))

	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	fmt.Printf("%v\n%v\n%v\n", intList, floatList, stringList)

}
