package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
)

//= How to Handle Errors: The Basics start

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

//= How to Handle Errors: The Basics end

//= Use Strings for Simple Errors start

func doubleEvenErr(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func doubleEvenErrf(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

//= Use Strings for Simple Errors end

//= Sentinel Errors start

// Sentinel type
type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	// ErrFoo constant
	ErrFoo = Sentinel("foo error")
	// ErrBar constant
	ErrBar = Sentinel("bar error")
)

//= Sentinel Errors end

//= Errors Are Values start

// Status type
type Status int

const (
	// InvalidLogin constant
	InvalidLogin Status = iota + 1
	// NotFound constant
	NotFound
)

// StatusErr struct
type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func login(uid, pwd string) (string, error) {
	if uid == "" || pwd == "" {
		return "", errors.New("missing parameters")
	}
	return "mylogin", nil
}

func getData(file string) ([]byte, error) {
	if file == "" {
		return []byte{}, errors.New("file name is empty")
	}
	return []byte{'d', 'a', 't', 'a'}, nil
}

// LoginAndGetData function
func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	id, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	fmt.Println("User ID:", id)
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintln("unable to get data"),
		}
	}
	fmt.Println(data)
	return data, nil
}

// GenerateError function
func GenerateError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

// GenerateErrorWithNil function
func GenerateErrorWithNil(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}

// GenerateErrorWithErrorInterface function
func GenerateErrorWithErrorInterface(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

//= Errors Are Values end

//= Wrapping Errors start

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

// StatusErrWrap struct
type StatusErrWrap struct {
	Status  Status
	Message string
	Err     error
}

func (sew StatusErrWrap) Error() string {
	return sew.Message
}

func (sew StatusErrWrap) Unwrap() error {
	return sew.Err
}

// LoginAndGetDataUnwrap function
func LoginAndGetDataUnwrap(uid, pwd, file string) ([]byte, error) {
	id, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErrWrap{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			Err:     err,
		}
	}
	fmt.Println("User ID:", id)
	data, err := getData(file)
	if err != nil {
		return nil, StatusErrWrap{
			Status:  NotFound,
			Message: fmt.Sprintln("unable to get data"),
			Err:     err,
		}
	}
	fmt.Println(data)
	return data, nil
}

//= Wrapping Errors end

//= Is and As start

// MyErr struct
type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

// Is method
func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, me2)
	}
	return false
}

// ResourceErr struct
type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

// Is method
func (re ResourceErr) Is(target error) bool {
	fmt.Println("Is Resource -> ", re.Resource)
	fmt.Println("Is Code -> ", re.Code)
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		fmt.Println("ignoreResource -> ", ignoreResource)
		fmt.Println("ignoreCode -> ", ignoreCode)
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		fmt.Println("matchResource -> ", matchResource)
		fmt.Println("matchCode -> ", matchCode)
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}

//= Is and As end

//= Wrapping Errors with defer start

func doThing1(val1 int) (int, error) {
	return 0, fmt.Errorf("doThing1 val1 error")
}

func doThing2(val2 string) (string, error) {
	return "", fmt.Errorf("doThing2 val2 error")
}

func doThings3(val3 int, val4 string) string {
	return val4
}

// DoSomeThings function
func DoSomeThings(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()
	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	return doThings3(val3, val4), nil
}

//= Wrapping Errors with defer end

//= panic and recover start

func doPanic(msg string) {
	panic(msg)
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

//= panic and recover end

func main() {
	fmt.Println("==> How to Handle Errors: The Basics")
	numerator := 20
	denominator := 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)

	fmt.Println("==> Use Strings for Simple Errors")
	var doubled int
	doubled, err = doubleEvenErr(3)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1) // Disable to see next code
	}
	fmt.Println(doubled)

	doubled, err = doubleEvenErrf(5)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1) // Disable to see next code
	}
	fmt.Println(doubled)

	fmt.Println("==> Sentinel Errors")
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err = zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so -> err: ", err)
	}
	fmt.Println(ErrFoo, " -> ", ErrBar)

	fmt.Println("==> Errors Are Values")
	data, err = LoginAndGetData("", "", "")
	fmt.Println(err)
	data, err = LoginAndGetData("1", "pass", "data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("data: ", string(data))

	err = GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)

	err = GenerateErrorWithNil(true)
	fmt.Println(err != nil)
	err = GenerateErrorWithNil(false)
	fmt.Println(err != nil)

	err = GenerateErrorWithErrorInterface(true)
	fmt.Println(err != nil)
	err = GenerateErrorWithErrorInterface(false)
	fmt.Println(err != nil)

	fmt.Println("==> Wrapping Errors")
	err = fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

	_, err = LoginAndGetDataUnwrap("", "", "")
	if err != nil {
		fmt.Println("Error: ", err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println("Unwrapped: ", wrappedErr)
		}
	}

	fmt.Println("==> Is and As")
	err = fileChecker("not_here.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}

	err = ResourceErr{
		Resource: "Database",
		Code:     123,
	}

	err2 := ResourceErr{
		Resource: "Network",
		Code:     456,
	}

	if errors.Is(err, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", err)
	}

	if errors.Is(err2, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", err2)
	}

	if errors.Is(err, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 triggered:", err)
	}

	if errors.Is(err2, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 triggered:", err2)
	}

	if errors.Is(err, ResourceErr{Resource: "Database", Code: 123}) {
		fmt.Println("Database is broken and Code 123 triggered", err)
	}

	fmt.Println("==> Wrapping Errors with defer")
	fmt.Println(DoSomeThings(0, ""))

	fmt.Println("panic and recover")
	// doPanic(os.Args[0])
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
