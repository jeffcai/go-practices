package main

import (
	"fmt"
	"math"
	"strconv"
)

type student struct {
	number int
	name   string
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func TotalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.Perimeter()
	}
	return perimeter
}

var ch chan int = make(chan int)

func foo(id int) { //id: 这个routine的标号
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

	v1 := 100

	if v2 := v1; v2 == 100 {
		fmt.Println("inside if ...")
	}

	switch {
	case v1 == 100:
		fmt.Println("v2 ...")
	default:
		fmt.Println("default ...")
	}

	a1 := []int{1, 2, 3}
	fmt.Println(a1)
	for index, a := range a1 {
		fmt.Println(strconv.Itoa(index) + ">" + strconv.Itoa(a))
	}

	m1 := map[int]string{1: "a", 2: "b"}
	for key, val := range m1 {
		fmt.Println(strconv.Itoa(key) + ">" + val)
	}

	fmt.Println(add(1, 2))

	pt := &v1
	fmt.Println(pt)
	fmt.Println(*pt)
	(*pt)++
	fmt.Println(*pt)
	fmt.Println(v1)

	stu := student{number: 1, name: "cjf"}
	fmt.Println(stu)

	r := Rect{width: 1, height: 2}
	fmt.Println("Rect area: " + strconv.FormatFloat(r.Area(), 'G', -1, 64) + " perimeter: " + strconv.FormatFloat(r.Perimeter(), 'G', -1, 64))

	c := Circle{radius: 1}
	fmt.Println("Circle area: " + strconv.FormatFloat(c.Area(), 'G', -1, 64) + " perimeter: " + strconv.FormatFloat(c.Perimeter(), 'G', -1, 64))

	fmt.Println("Total area: " + strconv.FormatFloat(TotalArea(r, c), 'G', -1, 64) + " perimeter: " + strconv.FormatFloat(TotalPerimeter(r, c), 'G', -1, 64))

}

func add(x, y int) int {
	return x + y
}

func Append(x, y string) string {
	return x + y
}
