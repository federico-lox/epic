package epic

import (
	"fmt"
	"testing"
)

const (
	fakeFunction = "a"
	fakeFile     = "x"
	fakeLine     = -1
	fakeQotf     = "Hello!"
)

var (
	origQotf          = qotf
	validateTestCases = []struct {
		got    int
		want   int
		truth  bool
		ok     bool
		report string
	}{
		{1, 1, true, true, ""},
		{1, 2, false, true, ""},
		{1, 2, true, false, fmt.Sprintf(reportFormat, fakeFunction, fakeFile, fakeLine, fakeQotf, 1, "", 2)},
		{1, 1, false, false, fmt.Sprintf(reportFormat, fakeFunction, fakeFile, fakeLine, fakeQotf, 1, notLabel, 1)},
	}
)

func mockGlobals() {
	qotf = []string{fakeQotf}
	extractContext = func() (string, string, int) {
		return fakeFunction, fakeFile, fakeLine
	}
}

func restoreGlobals() {
	qotf = origQotf
	extractContext = context
}

func TestValidate(test *testing.T) {
	mockGlobals()
	defer restoreGlobals()

	for _, testCase := range validateTestCases {
		report, ok := validate(testCase.got, testCase.want, testCase.truth)

		if report != testCase.report || ok != testCase.ok {
			test.Errorf("Result mismatch\n%v != %v\n%s%s", ok, testCase.ok, report, testCase.report)
		}
	}
}
