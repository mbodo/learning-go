package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("==> Arrays")
	var x [3]int
	fmt.Println(x)

	var y = [3]int{10, 20, 30}
	fmt.Println(y)

	// sparse array
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(z)

	var y1 = [...]int{10, 20, 30}
	fmt.Println(y1)
	fmt.Println(y == y1)

	x[0] = 10
	fmt.Println(x[0])

	// x[3] = 30
	// index 3 (constant of type int) is out of bounds

	fmt.Println("len(x): ", len(x))

	var zero1 [0]int
	fmt.Println(zero1)

	fmt.Println("==> Slices")
	var s = []int{10, 20, 30}
	fmt.Println(s)

	var s1 = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(s1)

	s[0] = 10
	// fmt.Println(s[3])
	// panic: runtime error: index out of range [3] with length 3

	var s2 []int
	fmt.Println(s2 == nil)

	// fmt.Println(s == s1)
	// cannot compare s == s1 (operator == not defined for []int)

	var s2x = []int{}
	fmt.Printf("Zero-lenght slice: s2x -> [%v] -> len(%v) -> cap(%v) \n", s2x, len(s2x), cap(s2x))

	var s3 []int
	s3 = append(s3, 10)
	fmt.Println(s3)

	var s4 = []int{10, 20, 30}
	s4 = append(s4, 40)
	fmt.Println(s4)

	s4 = append(s4, 50, 60, 70)
	fmt.Println(s4)

	var s5 = []int{80, 90, 100}
	s4 = append(s4, s5...)
	fmt.Println(s4)

	// Function cap
	fmt.Println("Function cap:")
	var c []int
	fmt.Println(c, len(c), cap(c))
	c = append(c, 10)
	fmt.Println(c, len(c), cap(c))
	c = append(c, 20)
	fmt.Println(c, len(c), cap(c))
	c = append(c, 30)
	fmt.Println(c, len(c), cap(c))
	c = append(c, 40)
	fmt.Println(c, len(c), cap(c))
	c = append(c, 50)
	fmt.Println(c, len(c), cap(c))

	// Function make
	fmt.Println("Function make:")
	m := make([]int, 5)
	fmt.Printf("%v -> len(%v) -> cap(%v) \n", m, len(m), cap(m))

	m1 := make([]int, 5, 10)
	fmt.Printf("%v -> len(%v) -> cap(%v) \n", m1, len(m1), cap(m1))
	var e1 = []int{10, 20, 30, 40, 50}
	m1 = append(m1, e1...)
	fmt.Printf("%v -> len(%v) -> cap(%v) \n", m1, len(m1), cap(m1))

	m2 := make([]int, 0, 10)
	fmt.Printf("%v -> len(%v) -> cap(%v) \n", m2, len(m2), cap(m2))
	m2 = append(m2, 10, 20, 30, 40)
	fmt.Printf("%v -> len(%v) -> cap(%v) \n", m2, len(m2), cap(m2))
	m2 = append(m2, 50, 60, 70, 80, 90, 100)
	fmt.Printf("%v -> len(%v) -> cap(%v) \n", m2, len(m2), cap(m2))
	m2 = append(m2, 110)
	fmt.Printf("%v -> len(%v) -> cap(%v) \n", m2, len(m2), cap(m2))

	var m3 = []int{}
	fmt.Printf("Zero slice [%v] -> len(%v) -> cap(%v)\n", m3, len(m3), cap(m3))

	var m4 []int
	fmt.Printf("nil slice [%v] -> len(%v) -> cap(%v)\n", m4, len(m4), cap(m4))

	// Slicing slices
	fmt.Println("Slicing slices")
	ss1 := []int{1, 2, 3, 4}
	ss2 := ss1[:2]
	ss3 := ss1[1:]
	ss4 := ss1[1:3]
	ss5 := ss1[:]
	fmt.Println("ss1: ", ss1)
	fmt.Println("ss2: ", ss2)
	fmt.Println("ss3: ", ss3)
	fmt.Println("ss4: ", ss4)
	fmt.Println("ss5: ", ss5)

	ss7 := []int{1, 2, 3, 4}
	ss8 := ss7[:2]
	ss9 := ss7[1:]
	ss7[1] = 20
	ss8[0] = 10
	ss9[1] = 30
	fmt.Println("ss7: ", ss7)
	fmt.Println("ss8: ", ss8)
	fmt.Println("ss9: ", ss9)

	ss10 := make([]int, 0, 5)
	ss10 = append(ss10, 1, 2, 3, 4)
	fmt.Printf("ss10 -> [%v] -> len(%v) -> cap(%v)\n", ss10, len(ss10), cap(ss10))
	ss11 := ss10[:2]
	fmt.Printf("ss11 -> [%v] -> len(%v) -> cap(%v)\n", ss11, len(ss11), cap(ss11))
	ss12 := ss10[2:]
	fmt.Printf("ss12 -> [%v] -> len(%v) -> cap(%v)\n", ss12, len(ss12), cap(ss12))
	fmt.Println(cap(ss10), cap(ss11), cap(ss12))
	fmt.Printf("[%s]\n", " -> ss11 = append(ss11, 30, 40, 50)")
	ss11 = append(ss11, 30, 40, 50)
	fmt.Printf("ss11 -> [%v] -> len(%v) -> cap(%v)\n", ss11, len(ss11), cap(ss11))
	fmt.Printf(" -- [%v]\n", ss10)
	fmt.Printf("[%s]\n", " -> ss10 = append(ss10, 60)")
	ss10 = append(ss10, 60)
	fmt.Printf("ss10 -> [%v] -> len(%v) -> cap(%v)\n", ss10, len(ss10), cap(ss10))
	fmt.Printf("[%s]\n", " -> ss12 = append(ss12, 70)")
	ss12 = append(ss12, 70)
	fmt.Printf("ss12 -> [%v] -> len(%v) -> cap(%v)\n", ss12, len(ss12), cap(ss12))
	fmt.Printf("ss10 -> [%v] -> len(%v) -> cap(%v)\n", ss10, len(ss10), cap(ss10))
	fmt.Printf("ss11 -> [%v] -> len(%v) -> cap(%v)\n", ss11, len(ss11), cap(ss11))
	fmt.Printf("ss12 -> [%v] -> len(%v) -> cap(%v)\n", ss12, len(ss12), cap(ss12))

	fmt.Println("Full slice expression:")
	ss13 := make([]int, 0, 5)
	ss13 = append(ss13, 1, 2, 3, 4)
	fmt.Printf("ss13 -> [%v] -> len(%v) -> cap(%v)\n", ss13, len(ss13), cap(ss13))
	ss14 := ss13[:2:2]
	fmt.Printf("ss14 -> [%v] -> len(%v) -> cap(%v)\n", ss14, len(ss14), cap(ss14))
	ss15 := ss13[2:4:4]
	fmt.Printf("ss15 -> [%v] -> len(%v) -> cap(%v)\n", ss15, len(ss15), cap(ss15))
	fmt.Println(cap(ss13), cap(ss14), cap(ss15))
	ss14 = append(ss14, 30, 40, 50)
	fmt.Printf("ss14 -> [%v] -> len(%v) -> cap(%v)\n", ss14, len(ss14), cap(ss14))
	ss13 = append(ss13, 60)
	fmt.Printf("ss13 -> [%v] -> len(%v) -> cap(%v)\n", ss13, len(ss13), cap(ss13))
	ss15 = append(ss15, 70)
	fmt.Printf("ss15 -> [%v] -> len(%v) -> cap(%v)\n", ss15, len(ss15), cap(ss15))
	fmt.Printf("ss13 -> [%v] -> len(%v) -> cap(%v)\n", ss13, len(ss13), cap(ss13))
	fmt.Printf("ss14 -> [%v] -> len(%v) -> cap(%v)\n", ss14, len(ss14), cap(ss14))
	fmt.Printf("ss15 -> [%v] -> len(%v) -> cap(%v)\n", ss15, len(ss15), cap(ss15))

	fmt.Println("Converting Arrays to Slices:")
	x111 := [4]int{5, 6, 7, 8}
	y111 := x111[:2]
	z111 := x111[2:]
	x111[0] = 10
	fmt.Printf("x111 -> [%v] -> len(%v) -> cap(%v)\n", x111, len(x111), cap(x111))
	fmt.Printf("y111 -> [%v] -> len(%v) -> cap(%v)\n", y111, len(y111), cap(y111))
	fmt.Printf("z111 -> [%v] -> len(%v) -> cap(%v)\n", z111, len(z111), cap(z111))

	fmt.Println("Converting Arrays to Slices - Full slice expression:")
	x121 := [4]int{5, 6, 7, 8}
	y121 := x121[:2]
	z121 := x121[2:]
	x111[0] = 10
	fmt.Printf("x121 -> [%v] -> len(%v) -> cap(%v)\n", x121, len(x121), cap(x121))
	fmt.Printf("y121 -> [%v] -> len(%v) -> cap(%v)\n", y121, len(y121), cap(y121))
	fmt.Printf("z121 -> [%v] -> len(%v) -> cap(%v)\n", z121, len(z121), cap(z121))

	fmt.Println("Slice copy:")
	x211 := []int{1, 2, 3, 4}
	y211 := make([]int, 4)
	fmt.Printf("%v\n", y211)
	num := copy(y211, x211)
	fmt.Println(y211, num)

	x221 := []int{1, 2, 3, 4}
	y221 := make([]int, 2)
	num1 := copy(y221, x221)
	fmt.Println(y221, num1)

	x222 := []int{1, 2, 3, 4}
	fmt.Printf("x222[:3] %v\n", x222[:3])
	fmt.Printf("x222[1:] %v\n", x222[1:])
	num2 := copy(x222[:3], x222[1:])
	fmt.Println(x222, num2)

	fmt.Println("Array copy")
	x232 := []int{1, 2, 3, 4}
	d232 := [4]int{5, 6, 7, 8}
	y232 := make([]int, 2)
	copy(y232, d232[:])
	fmt.Println(y232)
	copy(d232[:], x232)
	fmt.Println(d232)

	fmt.Println("==> Strings and Runes and Bytes")
	var s111 string = "Hello there"
	var b111 byte = s111[6]
	fmt.Println(b111)
	var s112 string = s111[4:7]
	var s113 string = s111[:5]
	var s114 string = s111[6:]
	fmt.Println(s112)
	fmt.Println(s113)
	fmt.Println(s114)

	var s211 string = "Hello öl"
	var s221 string = s211[4:7]
	var s231 string = s211[:5]
	var s241 string = s211[6:]
	fmt.Printf("%s -> len(s211) <- %v \n", s221, len(s211))
	fmt.Println(s231)
	fmt.Println(s241)

	var a111 rune = 'x'
	var s311 string = string(a111)
	fmt.Println(s311)
	var b211 byte = 'y'
	var s321 string = string(b211)
	fmt.Println(s321)

	// var x422 = 65
	// var y422 = string(x422)
	// fmt.Println(y422)

	var s511 string = "Hello öl"
	var bs []byte = []byte(s511)
	var rs []rune = []rune(s511)
	fmt.Println(bs)
	fmt.Println(rs)

	fmt.Println("==> Maps:")
	var nilMap map[string]int
	fmt.Println(nilMap)
	fmt.Println(len(nilMap))
	fmt.Println(nilMap == nil)

	// empty literal map
	totalWins := map[string]int{}
	fmt.Println(totalWins)
	fmt.Println(len(totalWins))
	fmt.Println(totalWins == nil)

	//non-empty literal map
	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)
	fmt.Println(len(teams))
	fmt.Println(teams == nil)

	ages := make(map[int][]string, 10)
	fmt.Println(ages)
	fmt.Println(len(ages))
	fmt.Println(ages == nil)

	// Reading and Writing a Map
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3

	m111 := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m111["hello"]
	fmt.Println(v, ok)

	v, ok = m111["world"]
	fmt.Println(v, ok)

	v, ok = m111["goodbye"]
	fmt.Println(v, ok)

	// Redeclare
	m111 = map[string]int{
		"goodbye": 7,
	}

	fmt.Println(m111)

	fmt.Println("Deleting from Maps")
	m111 = map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(m111)

	delete(m111, "hello")
	fmt.Println(m111)

	fmt.Println("Using Maps and Sets")
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])

	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	fmt.Println("With struct")
	strSet := map[int]struct{}{}
	vals = []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		strSet[v] = struct{}{}
	}
	if _, ok := strSet[5]; ok {
		fmt.Println("5 is in the set")
	}

	// Structs
	fmt.Println("Structs:")
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	fmt.Println("Zero value struct [fred] -> ", fred)

	bob := person{}
	fmt.Println("Zero value struct literal [bob] -> ", bob)

	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println("Julia -> ", julia)

	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println("Beth -> ", beth)
	fmt.Println("beth.name -> ", beth.name)
	fmt.Println("beth.age -> ", beth.age)
	fmt.Println("beth.pet -> ", beth.pet)

	fmt.Println("Anonymous Structs")
	var petowner struct {
		name string
		age  int
		pet  string
	}

	petowner.name = "bob"
	petowner.age = 50
	petowner.pet = "dog"

	fmt.Println("petowner.name -> ", petowner.name)
	fmt.Println("petowner.age -> ", petowner.age)
	fmt.Println("petowner.pet -> ", petowner.pet)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println("pet.name -> ", pet.name)
	fmt.Println("pet.kind -> ", pet.kind)

	fmt.Println("Comparing and Converting Structs")
	type firstPerson struct {
		name string
		age  int
	}

	type secondPerson struct {
		name string
		age  int
	}

	first := firstPerson{
		name: "first",
		age:  1,
	}

	fmt.Println(first)

	second := secondPerson{
		name: "second",
		age:  2,
	}

	fmt.Println(second)

	//fmt.Println("Compare first == second")
	//fmt.Println(first == second)
	fmt.Println("To convert second -> ", reflect.TypeOf(second), ", ", second)
	converted := firstPerson(second)
	fmt.Println("Convert from secondPerson to firstPerson -> ", reflect.TypeOf(converted), ", ", converted)

	type thirdPerson struct {
		age  int
		name string
	}

	third := thirdPerson{
		age:  3,
		name: "third",
	}

	fmt.Println(third)
	// converted3 := thirdPerson(first)

	type fourthPerson struct {
		firstName string
		age       int
	}

	fourth := fourthPerson{
		firstName: "fourth",
		age:       4,
	}

	fmt.Println(fourth)
	// converted4 = fourthPerson(first)

	type fifthPerson struct {
		name          string
		age           int
		favoriteColor string
	}

	fifth := fifthPerson{
		name:          "fifth",
		age:           5,
		favoriteColor: "green",
	}

	fmt.Println(fifth)
	// converted5 := fifthPerson(first)

	particular := firstPerson{
		name: "Bob",
		age:  50,
	}

	var anonymous struct {
		name string
		age  int
	}

	anonymous = particular
	fmt.Println("Compare & convert with anonymous struct -> ", anonymous == particular)

	var t [0]person
	fmt.Println(t)

	t = [0]person{}
	fmt.Println(t)

	p := []person{}
	fmt.Println(p)

}
