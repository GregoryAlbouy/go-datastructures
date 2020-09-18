package testx

import (
	"reflect"
	"testing"
)

// type test struct {
// 	fun   TestFunc
// 	cases []Testcase
// }

// Tester interface
// type Tester interface {
// 	Check(t *testing.T, tc Testcase, got interface{})
// }

// Testcase struct
type Testcase struct {
	Desc     string
	Input    interface{}
	Expected interface{}
}

// TestFunc type
// type TestFunc = func(...interface{}) interface{}

// // Args type
// type Args []interface{}

// // New returns a test struct
// func New(testcases []Testcase) Tester {
// 	return &test{nil, testcases}
// }

// // New returns a test struct
// func New(tf TestFunc, testcases []Testcase) Tester {
// 	return &test{tf, testcases}
// }

// func Func(f interface{}) (TestFunc, error) {
// 	if reflect.ValueOf(f).Kind() != reflect.Func {
// 		return nil, errors.New("f is not a function")
// 	}

// 	return func(args ...interface{}) interface{} {
// 		return f(args...)
// 	}, nil
// }

// Check func
func Check(t *testing.T, tc Testcase, got interface{}) {
	if !reflect.DeepEqual(tc.Expected, got) {
		t.Errorf("%s: expected %+v, got %+v", tc.Desc, tc.Expected, got)
	}
}
