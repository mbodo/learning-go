package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Person struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Counter struct
type Counter struct {
	total       int
	lastUpdated time.Time
}

// Increment the counter
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

// IntTree struct
type IntTree struct {
	val         int
	left, right *IntTree
}

// Insert method
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

// Contains method
func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

// Adder type
type Adder struct {
	start int
}

// AddTo method
func (a Adder) AddTo(val int) int {
	return a.start + val
}

// Score struct
type Score struct {
	val int
}

// HighScore type
type HighScore Score

// Employee struct
type Employee struct {
	Name string
	ID   string
}

// Description method
func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

// Manager struct
type Manager struct {
	Employee
	Reports []Employee
}

// FindNewEmplpoyees struct
func (m Manager) FindNewEmplpoyees() []Employee {
	return []Employee{{Name: "Bob Bobson", ID: "12345"}}
}

// Inner struct
type Inner struct {
	X int
}

// IntPrinter method
func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

// Double method
func (i Inner) Double() string {
	return i.IntPrinter(i.X * 2)
}

// Outer struct
type Outer struct {
	Inner
	X int
}

// IntPrinter method
func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

// LogicProvider struct
type LogicProvider struct{}

// Process method
func (lp LogicProvider) Process(data string) string {
	return data
}

// ProcessLogic interface
type ProcessLogic interface {
	Process(data string) string
}

// Client struct
type Client struct {
	L ProcessLogic
}

// Program func
func (c Client) Program() {
	// get data from somewhere
	m := "call business logic in LogicProvider"
	fmt.Println(m)

	c.L.Process(m)
}

// Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Closer interface
type Closer interface {
	Close() error
}

// ReadCloser interface
type ReadCloser interface {
	Reader
	Closer
}

// MyInt int type
type MyInt int

func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		fmt.Println("i is nil, type of j is interface{}", j)
	case int:
		fmt.Println("j is of type int", j)
	case MyInt:
		fmt.Println("j is of type MyInt", j)
	case io.Reader:
		fmt.Println("j is of type io.Reader", j)
	case string:
		fmt.Println("j is a string", j)
	case bool, rune:
		fmt.Println("i is either a bool or rune, so j is of type interface{}", j)
	default:
		fmt.Println("no idea what i is, so j is of type interface{}", j)
	}
}

// Implicit Interfaces Make Dependency Injection Easier

// LogOutput function
func LogOutput(message string) {
	fmt.Println(message)
}

// SimpleDataStore struct
type SimpleDataStore struct {
	userData map[string]string
}

// UserNameForID method
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// NewSimpleDataStore function
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

// DataStore interface
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// Logger interface
type Logger interface {
	Log(message string)
}

// LoggerAdapter function type
type LoggerAdapter func(message string)

// Log function
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// SimpleLogic struct
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

// SayHello method
func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknow user")
	}
	return "Hello, " + name, nil
}

// SayGoodbye method
func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

// NewSimpleLogic function
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// Logic interface
type Logic interface {
	SayHello(userID string) (string, error)
}

// Controller struct
type Controller struct {
	l     Logger
	logic Logic
}

// SayHello method
func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	c.l.Log("message: " + message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// NewController function
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	fmt.Println("==> Methods")
	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}

	output := p.String()
	fmt.Println(output)

	fmt.Println("==> Pointer Recievers and Value Recievers")
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	doUpdateWrong(c)
	fmt.Println("in main:", c.String())
	doUpdateRight(&c)
	fmt.Println("in main:", c.String())

	fmt.Println("==> Code Your Methods for nil Instances")
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  //true
	fmt.Println(it.Contains(12)) //false

	fmt.Println("==> Methods Are Functions Too")
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5)) // prints 15
	f1 := myAdder.AddTo
	fmt.Println(f1(10)) // prints 20
	fmt.Println("Method expression:")
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15)) //prints 25

	fmt.Println("==> Type Declarations Aren’t Inheritance")
	var i int = 300
	var s Score = Score{
		val: 100,
	}
	var hs HighScore = HighScore{
		val: 200,
	}
	//hs = s
	//s = i
	s = Score{
		val: i,
	}
	hs = HighScore{
		val: 400,
	}
	fmt.Println(s)
	fmt.Println(hs)

	fmt.Println("==> iota Is for Enumerations—Sometimes")
	type MailCategory int

	const (
		Uncategorized MailCategory = iota
		Personal
		Spam
		Social
		Advertisements
	)
	fmt.Println(Uncategorized, Personal, Spam, Social, Advertisements)

	type BitField int

	const (
		Field1 BitField = 1 << iota // assigned 1
		Field2
		Field3
		Field4
	)

	fmt.Println(Field1)
	fmt.Println(Field2)
	fmt.Println(Field3)
	fmt.Println(Field4)

	fmt.Println("==> Use Embedding for Composition")
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)
	fmt.Println(o.Inner.X)

	fmt.Println("==> Embedding Is Not Inheritance")
	// var eFail Employee = m
	var eOK Employee = m.Employee
	fmt.Println(eOK)

	o = Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.Double())

	fmt.Println("==> A Quick Lesson on Interfaces")
	type Stringer interface {
		String() string
	}

	fmt.Println("==> Interfaces Are Type-Safe Duck Typing")
	cl := Client{
		L: LogicProvider{},
	}
	cl.Program()

	fmt.Println("==> Embedding and Interfaces")

	fmt.Println("==> Interfaces and nil")
	var s1 *string
	fmt.Println(s1 == nil) // prints true
	var i1 interface{}
	fmt.Println(i1 == nil) // prints true
	i1 = s1
	fmt.Println(i1 == nil) // prints false

	fmt.Println("==> The Empty Interface Says Nothing")
	var i2 interface{}
	i2 = 20
	i2 = "hello"
	i2 = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}
	fmt.Println(i2)

	fmt.Println("==> Type Assertions and Type Switches")
	var i3 interface{}
	var mine MyInt = 20
	i3 = mine
	i4 := i3.(MyInt)
	fmt.Println(i4 + 1)

	i5, ok := i3.(int)
	if !ok {
		fmt.Printf("unexpected type for %v \n", i3)
	} else {
		fmt.Println(i5 + 1)
	}

	doThings(i3)

	fmt.Println("==> Implicit Interfaces Make Dependency Injection Easier")
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	ctr := NewController(l, logic)
	http.HandleFunc("/hello", ctr.SayHello)
	http.ListenAndServe(":8080", nil)
}
