package main

import (
	"errors"
	"fmt"
	"go-functions/simplemath"
	"io"
	"math"
	"net/http"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	// fmt.Printf("%f\n", simplemath.Add(2, 3))
	// res, err := simplemath.Divide(2, 3)
	// if err != nil {
	// 	fmt.Printf("An error occurred: %s", err.Error())
	// } else {
	// 	println(res)
	// }
	// fmt.Printf("%f\n", simplemath.Sum(2, 3, 4, 5, 6))

	// sv := simplemath.NewSemanticVersion(1, 2, 3)
	// println(sv.String())

	// var tripper http.RoundTripper = &RoundTripCounter{}
	// r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	// _, _ = tripper.RoundTrip(r)

	// func() {
	// 	println("My first anonymous function")
	// }()

	// a := func(name string) string {
	// 	fmt.Printf("My name is %s\n", name)
	// 	return name
	// }

	// fmt.Printf("My name again is %s\n", a("András"))

	// mathExpr := mathExpression3(SubtractExpr)

	// println(mathExpr(2, 3))

	// fmt.Printf("Result is: %f", double(2, 3, mathExpression3(AddExpr)))
	// fmt.Printf("Result is: %f", double(2, 3, mathExpression3(SubtractExpr)))

	// p2 := powerOfTwo()
	// println(p2())
	// println(p2())

	// bad state in anonymous functions
	// var funcs []func() int64

	// for i := 0; i < 10; i++ {
	// 	cleanI := i
	// 	funcs = append(funcs, func() int64 {
	// 		return int64(math.Pow(float64(cleanI), 2))
	// 	})
	// }

	// for _, f := range funcs {
	// 	println(f())
	// }

	// continue on error
	if err := ReadFullFile(); err != nil {
		println(err.Error())
	}
}

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}

// return functions from functions
func mathExpression() func(float64, float64) float64 {
	return func(f float64, f2 float64) float64 {
		return f + f2
	}
}

func mathExpression2() func(float64, float64) float64 {
	return simplemath.Add
}

func mathExpression3(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		panic("invalid expression")
	}
}

// functions as function parameter
func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}

// stateful function
func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

// error handling
func readSomething() error {
	var r io.Reader = BadReader{errors.New("my nonsense reader")}
	if _, err := r.Read([]byte("test something")); err != nil {
		fmt.Printf("an error occurred %s", err)
		return err
	}

	return nil
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

// continue on error, recover form panic, panicking on recovery
type SimpleReader struct {
	count int
}

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	if br.count == 2 {
		panic(errCatastrophicReader)
	}
	if br.count > 3 {
		return 0, io.EOF
	}
	br.count += 1
	return br.count, nil
}

var errCatastrophicReader = errors.New("something very bad happened")

func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("a panic occurred but it is ok")
		} else if p != nil {
			panic("an unexopected error occurred")
		}
	}()
	for {
		value, readerErr := r.Read([]byte("something"))
		if readerErr == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			err = readerErr
			return err
		}

		println(value)
	}

	return nil
}

// defer functions
func (br *SimpleReader) Close() error {
	println("closing reader")
	return nil
}
