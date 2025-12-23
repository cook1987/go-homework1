package main

import (
	"fmt"
	"math"
)

// 题目 1
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  int
	height int
}

type Circle struct {
	r int
}

func (c *Rectangle) Area() float64 {
	return float64(c.width * c.height)
}
func (c *Circle) Area() float64 {
	nr := float64(c.r)
	return math.Pi * nr * nr
}

func (c *Rectangle) Perimeter() float64 {
	return float64(2 * c.width * c.height)
}
func (c *Circle) Perimeter() float64 {
	nr := float64(c.r)
	return math.Pi * nr * 2
}

// 题目 2
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	person     Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Println("姓名：", e.person.Name, "，年龄：", e.person.Age, "，工号：", e.EmployeeID)
}

// func main() {
// 	fmt.Println("----------------题目 1-------------")
// 	sh := Rectangle{width: 4, height: 3}
// 	ci := Circle{r: 7}

// 	fmt.Println("矩形的面积：", sh.Area())
// 	fmt.Println("圆形的面积：", ci.Area())

// 	fmt.Println("矩形的周长：", sh.Perimeter())
// 	fmt.Println("圆形的周长：", ci.Perimeter())

// 	fmt.Println("----------------题目 2-------------")

// 	e := Employee{person: Person{Name: "张三", Age: 23}, EmployeeID: 99}
// 	e.PrintInfo()
// }
