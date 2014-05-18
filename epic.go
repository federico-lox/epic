/*
Package epic contains two simple functions to validate Go tests
*/
package epic

import (
	"reflect"
	"runtime"
	"testing"
)

const reporterFormat = "Test failed\nwhere:\tline %d, %s\ngot:\t%v\nwant:\t%v"

func Win(test *testing.T, got interface{}, want interface{}) {
	// TODO: detect -fatal flag and use Fatalf instead
	reporter := test.Errorf

	if !reflect.DeepEqual(got, want) {
		_, file, line, _ := runtime.Caller(1)
		reporter(reporterFormat, line, file, got, want)
	}
}

func Fail(test *testing.T, got interface{}, want interface{}) {
	// TODO: detect -fatal flag and use Fatalf instead
	reporter := test.Errorf

	if reflect.DeepEqual(got, want) {
		_, file, line, _ := runtime.Caller(1)
		reporter(reporterFormat, line, file, got, want)
	}
}
