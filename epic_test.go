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
	origSequence      = headlinesSequence
	origIndex         = headlinesIndex
	validateTestCases = []struct {
		got      int
		want     int
		truth    bool
		expected string
	}{
		{1, 1, true, ""},
		{1, 2, false, ""},
		{1, 2, true, fmt.Sprintf(
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
		{1, 1, false, fmt.Sprintf(
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
	headlinesIndex = indexReset
	headlinesSequence = []int{0}
}

func restoreHeadlines() {
	headlines = origHeadlines
	headlinesSequence = origSequence
	headlinesIndex = origIndex
}

func TestValidateHeadlne(test *testing.T) {
	mockHeadlines()
	defer restoreHeadlines()

	if got := headline(); got != fakeHeadline {
		test.Errorf("Got %v, want %v", got, fakeHeadline)
	}
}

func fakeContext() (string, string, int) {
	return fakeFunction, fakeFile, fakeLine
}

func TestValidate(test *testing.T) {
	mockHeadlines()
	defer restoreHeadlines()
	extractContext = fakeContext

	for _, testCase := range validateTestCases {
		report := validate(testCase.got, testCase.want, testCase.truth)

		if report != testCase.expected {
			test.Errorf("Report mismatch\n%s%s", report, testCase.expected)
		}
	}
}
