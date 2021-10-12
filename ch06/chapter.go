package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	msg string
}

func stringp(s string) *string {
	return &s
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func failedUpdate2(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

// MakeFoo example
func MakeFoo() (Foo, error) {
	f := Foo{
		msg: "Hello!",
	}
	return f, nil
}

func main() {
	fmt.Println("==> A Quick Pointer Primer")
	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string

	fmt.Println(pointerX)
	fmt.Println(pointerY)
	fmt.Println(pointerZ)

	fmt.Println("Dereferencing")
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)

	var pointerToZ *int
	fmt.Println(pointerToZ == nil)
	// Dereferencing nil pointer
	// fmt.Println(*pointerToZ)

	var n = new(int)
	fmt.Println(n == nil)
	fmt.Println(*n)

	t := &Foo{
		msg: "ok",
	}
	fmt.Printf("%p \n", t)

	var d string
	e := &d
	fmt.Println(e)

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p1 := person{
		FirstName: "Pat",
		// Won't compile
		// MiddleName: "Perry",
		LastName: "Peterson",
	}
	fmt.Println(p1)

	p2 := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Peterson",
	}
	fmt.Println(p2)

	fmt.Println("==> Pointers Indicate Mutable Parameters")
	var f *int
	failedUpdate(f)
	fmt.Println(f)

	f1 := 10
	failedUpdate2(&f1)
	fmt.Println(f1)
	update(&f1)
	fmt.Println(f1)

	fmt.Println("==> Pointers Are a Last Resort")
	foo, err := MakeFoo()
	if err != nil {
		fmt.Println("Problem with creating a Foo")
	}
	fmt.Println(foo)

	jp := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "",
		Age:  0,
	}
	err1 := json.Unmarshal([]byte(`{"name":"Bob", "age": 30}`), &jp)
	fmt.Println(err1)
	fmt.Println(jp.Name)
	fmt.Println(jp.Age)
}
