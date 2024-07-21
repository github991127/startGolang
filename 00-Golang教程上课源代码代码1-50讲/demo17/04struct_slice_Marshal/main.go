package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Student struct {
	Id     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []Student
}

func main() {
	c := Class{
		Title:    "001班",
		Students: make([]Student, 0),
	}

	for i := 1; i <= 10; i++ {
		s := Student{
			Id:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		c.Students = append(c.Students, s)
	}

	// fmt.Println(c)

	strByte, err := json.Marshal(c)

	if err != nil {
		fmt.Println(err)
	} else {

		strJson := string(strByte)
		fmt.Println(strJson)
	}
}
