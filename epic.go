/*
Package epic contains two simple functions to validate Go tests
*/
package epic

import (
	"reflect"
	"runtime"
	"testing"
)

const reporterFormat = "Test failed\nwhere:\tline %d, %s\ngot:\t%v\n%s:\t%v"

func Win(test *testing.T, got interface{}, good interface{}) {
	validate(test, got, good, true)
}

func Fail(test *testing.T, got interface{}, bad interface{}) {
	validate(test, got, bad, false)
}

func validate(test *testing.T, got interface{}, flag interface{}, expected bool) {
	// TODO: detect -fatal flag and use Fatalf instead
	reporter := test.Errorf
	var label string

	if expected {
		label = "good"
	} else {
		label = "bad"
	}

	if reflect.DeepEqual(got, flag) != expected {
		_, file, line, _ := runtime.Caller(2)
		reporter(reporterFormat, line, file, got, label, flag)
	}
}
