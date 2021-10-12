package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

// MyFuncOpts ok
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

// MyFunc ok
func MyFunc(opts MyFuncOpts) int {
	fmt.Println(opts)
	return 0
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndReminder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func divAndReminder1(numerator int, denominator int) (result int, reminder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, reminder, err
	}
	result, reminder = numerator/denominator, numerator%denominator
	return result, reminder, err
}

func divAndReminder2(numerator int, denominator int) (result int, reminder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, reminder = numerator/denominator, numerator%denominator
	return
}

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	fmt.Println("==> Example 5-1:")
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})

	fmt.Println("==> Variadic Input Parameter and Slices:")
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	fmt.Println("==> Multiple Return Values")
	result, reminder, err := divAndReminder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, reminder)

	fmt.Println("==> Ignoring Returned Values")
	result1, _, err := divAndReminder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result1)

	fmt.Println("==> Named Return Values")
	x, y, z := divAndReminder1(5, 2)
	fmt.Println(x, y, z)

	fmt.Println("==> Blank Returns")
	result2, reminder2, err := divAndReminder2(5, 2)
	fmt.Println(result2, reminder2, err)

	fmt.Println("==> Function Are Values")
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	fmt.Println(opMap)

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression: ", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator: ", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	fmt.Println("==> Function Type Declarations")
	type opFuncType func(int, int) int
	var opMapType = map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	fmt.Println(opMapType)

	fmt.Println("==> Anonymous Functions")
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}

	fmt.Println("==> Passing Functions as Parameters")
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Println(people)

	fmt.Println("sort by last name:")
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	fmt.Println("sort by age:")
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	fmt.Println("==> Returning Functions from Functions")
	twoBase := makeMult(2)
	treeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), treeBase(i))
	}

	fmt.Println("==> defer")
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	fmt.Println("Go Is Call By Value")
	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	slice := []int{1, 2, 3}
	modSlice(slice)
	fmt.Println(slice)

}
