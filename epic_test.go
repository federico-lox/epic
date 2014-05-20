package epic

import (
	"fmt"
	"testing"
)

const (
	fakeFunction = "a"
	fakeFile     = "x"
	fakeLine     = -1
	fakeHeadline = "Hello!"
)

var (
	origHeadlines     = headlines
	validateTestCases = []struct {
		got    int
		want   int
		truth  bool
		ok     bool
		report string
	}{
		{1, 1, true, true, ""},
		{1, 2, false, true, ""},
		{1, 2, true, false, fmt.Sprintf(
			reportFormat,
			fakeFunction,
			fakeFile,
			fakeLine,
			qotfLabel,
			fakeHeadline,
			gotLabel,
			1,
			wantLabel,
			"",
			2,
		)},
		{1, 1, false, false, fmt.Sprintf(
			reportFormat,
			fakeFunction,
			fakeFile,
			fakeLine,
			qotfLabel,
			fakeHeadline,
			gotLabel,
			1,
			wantLabel,
			notLabel,
			1,
		)},
	}
)

func mockHeadlines(items ...string) {
	headlines = []string{fakeHeadline}
}

func restoreHeadlines() {
	headlines = origHeadlines
}

func fakeContext() (string, string, int) {
	return fakeFunction, fakeFile, fakeLine
}

func TestValidate(test *testing.T) {
	mockHeadlines()
	defer restoreHeadlines()
	extractContext = fakeContext

	for _, testCase := range validateTestCases {
		report, ok := validate(testCase.got, testCase.want, testCase.truth)

		if report != testCase.report || ok != testCase.ok {
			test.Errorf("Result mismatch\n%v != %v\n%s%s", ok, testCase.ok, report, testCase.report)
		}
	}
}
