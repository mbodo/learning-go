package main

import (
	"fmt"
	"math/rand"
)

func example41() {
	fmt.Println("example 4-1:")
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func example42() {
	fmt.Println("example 4-2:")
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
}

func example43() {
	fmt.Println("example 4-3:")
	x := 10
	fmt.Println(x)
	// fmt := "oops"
	// fmt.Println(x)
}

func example44() {
	fmt.Println("example 4-4:")
	fmt.Println(true)
	true := 10
	fmt.Println(true)
}

func example45() {
	fmt.Println("example 4-5:")
	n := rand.Intn(10)
	if n > 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That is a good number", n)
	}
}

func example46() {
	fmt.Println("example 4-6:")
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That is a good number", n)
	}
}

func example47() {
	fmt.Println("example 4-7:")
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That is a good number", n)
	}
	// fmt.Println(n) // undeclared name
}

func example48() {
	fmt.Println("example 4-8:")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func example49() {
	fmt.Println("example 4-9:")
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

func example410() {
	fmt.Println("example 4-10:")
	fmt.Println("Infinite loop")
	for {
		fmt.Println("Hello")
		break
	}
}

func example411() {
	fmt.Println("example 4-11:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func example412() {
	fmt.Println("example 4-12:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		} else if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}

func example413() {
	fmt.Println("example 4-13:")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

func example414() {
	fmt.Println("example 4-14:")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		fmt.Println(v)
	}
}

func example415() {
	fmt.Println("example 4-15:")
	m := map[string]int{
		"a": 1,
		"c": 2,
		"b": 3,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

func example416() {
	fmt.Println("example 4-16:")
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}
	fmt.Println()
}

func example417() {
	fmt.Println("example 4-17:")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
}

func example418() {
	fmt.Println("example 4-18:")
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}

func example419() {
	fmt.Println("example 4-19:")
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			worldLen := len(word)
			fmt.Println(word, "is exactly the right lenght", worldLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}
}

func example420() {
	fmt.Println("example 4-20:")
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break
		default:
			fmt.Println(i, "is boring")
		}
	}
}

func example420Alt() {
	fmt.Println("example 4-20:")
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
}

func example421() {
	fmt.Println("example 4-21:")
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println("is exactly a right lenght")
		}
	}
}

func example422() {
	fmt.Println("example 4-22:")
	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("That's too low")
	case n > 5:
		fmt.Println("That's too big:", n)
	default:
		fmt.Println("That's a good number:", n)
	}
}

func example423() {
	fmt.Println("example 4-23:")
	a := 10
	// goto skip
	b := 20
	// skip:
	c := 30
	fmt.Println(a, b, c)
	// if c > a {
	// goto inner
	// }
	if a < b {
		// inner:
		// fmt.Println("a is less than b")
	}
}

func example424() {
	fmt.Println("example 4-24:")
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}

func main() {
	//if
	example41()
	example42()
	// Enable shadow in Makefile
	example43()
	example44()
	example45()
	example46()
	example47()
	// for, Four Ways
	example48()
	example49()
	example410()
	example411()
	example412()
	example413()
	example414()
	example415()
	example416()
	example417()
	example418()
	example419()
	example420()
	example420Alt()
	example421()
	example422()
	example423()
	example424()
}
