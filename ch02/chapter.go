package main

import (
	"fmt"
	"math/cmplx"
)

// Using const start
const xc int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const zc = 20 * 10

// Using const stop

const xz = 10

const typedX int = 10

func ineffectual() {
	l := 10
	l = 10
	fmt.Println(l)
	l = 30
}

func main() {
	fmt.Println("==> Literals")
	s1 := "Greetings and\n \"Salutations\""
	fmt.Println(s1)

	s2 := `Greetings and
 "Salutations"`
	fmt.Println(s2)

	fmt.Println("==> Booleans")
	var flag bool // no value assigned, set to false
	var isAwesome = true
	fmt.Println(flag)
	fmt.Println(isAwesome)

	fmt.Println("==> Integer operators")
	var x2 int = 10
	x2 *= 2
	fmt.Println(x2)

	fmt.Println("==> Complex types")
	// Complex numbers
	fmt.Println("--> Example 2-1 Start")
	fmt.Println("---------------------")
	x := complex(2.5, 3.1)
	e := complex(10.2, 2)
	fmt.Println(x + e)
	fmt.Println(x - e)
	fmt.Println(x * e)
	fmt.Println(x / e)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
	fmt.Println("---------------------")
	fmt.Println("--> Example 2-1 End")

	// Explicit Type Conversions
	fmt.Println("==> Explicit Type Conversion")
	fmt.Println("--> Example 2-2 Start")
	fmt.Println("---------------------")
	var x1 int = 10
	var y1 float64 = 30.2
	var z1 float64 = float64(x1) + y1
	fmt.Println(z1)
	var d1 int = x1 + int(y1)
	fmt.Println(d1)
	fmt.Println("--> Example 2-2 End")
	fmt.Println("---------------------")

	x1 = 30
	fmt.Println(&x1)
	x1 = 40
	fmt.Println(&x1)
	{
		x1 := 10
		fmt.Printf("scope -> %v -> [%v]\n", x1, &x1)
	}
	fmt.Println("ok -> ", x1)

	fmt.Println("==> var Versus :=")
	var test1 int = 10
	fmt.Println(test1)

	// test1 := 30
	// no new variables on left side of :=

	test2 := 20
	fmt.Println(test2)

	test2, test3 := 40, 50
	fmt.Println(test2)
	fmt.Println(test3)

	var test4 = 60
	fmt.Println(test4)

	var test5 int
	test5 = 70

	fmt.Println(test5)

	test6 := 80
	fmt.Println(test6)

	test6 = 90
	fmt.Println(test6)

	fmt.Println("==> Using const")
	fmt.Println("--> Example 2-3 Start")
	const yc = "hello"

	fmt.Println(xc)
	fmt.Println(yc)

	// xc = xc + 1 // cannot assign to xc (constant 10 of type int64) compiler (UnassignableOperand)
	// yc = "bye" // cannot assign to yc (untyped string constant "hello") compiler (UnassignableOperand)

	fmt.Println("--> Example 2-3 End")

	fmt.Println("==> Typed and Untyped Constants")
	var y int = xz
	var z float64 = xz
	var d byte = xz

	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)

	var i int = typedX
	fmt.Println(i)

	fmt.Println("==> Unused Variables")
	ineffectual()

	const cc1 float64 = 11.4

	var v1 float64 = float64(typedX)
	fmt.Println(v1)

	const typedY float64 = 70.2
	var v2 float64 = typedY
	fmt.Println(v2)

	const cc2 = 800
	var v3 float64 = cc2
	fmt.Println(v3)
}
