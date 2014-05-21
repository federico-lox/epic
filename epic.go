// Copyright 2014 Federico "Lox" Lucignano. All rights reserved. Use of this source code is governed by the MIT license
//that can be found in the LICENSE file.

/*
Package epic streamlines validating test results, it moves away from the richness of methods that characterizes the
"testing" package and it diverges from other testing frameworks and libraries by avoiding an as rich library of matchers.

Epic goals are to stay true to Go's minimalism and practicality while adding some fun to writing tests and going through
executions results.

Features

* Simple syntax with humor: Only two methods, that's all you'll have to deal with

* 100% integrated with "go test", works side-by-side with the "testing" package

* Clear, well formatted output for failures

* QOTF: For each failure a Quote of The Fail to make you smile at your misery

Installing and updatinsg

Install the epic package with the following command:

	go get gopkg.in/federico-lox/epic.v0

To ensure you're using the latest version, run the following instead:

	go get -u gopkg.in/federico-lox/epic.v0


Remember to also add the package to your tests' imports:

	import "gopkg.in/federico-lox/epic.v0"

Example

Here's the gist of how to use epic in your tests:

	package mytests

	import "time"

	func TestEpicExample(t *testing.T) {
		today := time.Now().Round(time.Minute)
		birthday, _ := time.Parse("2006-01-02 15:04 CEST", "1982-02-28 22:00 CEST")

		// Is it my birthday? If not, fail.
		epic.Win(t, today, birthday)

		// I hope today it's not my birthday! If it is, fail.
		epic.Fail(t, today, birthday)
	}

In the previous example the first validation will fail, here's what that would look like in go test's output:

	### mytests.TestEpicExample in filename.go at line 8
	--- QOTF: Huston, we have a problem!
	--- GOT : 2014-05-19
	--- WANT: 1982-02-28
	--- FAIL: TestEpicExample (0.00 seconds)

Easy to read and fun, isn't it?

XYZ Not Supported

No worries, if you're using epic then you're writing a normal test case using "go test" and the "testing" package, feel
free to leverage those tools when working with epic to cover your needs; epic will never add support for non-generic
cases as it's meant to stay simple, easy to use and fun!

Contributions

Want to contribute? The best way is to open a pull request on Github at https://github.com/federico-lox/epic.

In particular you can help making testing everyone's code more fun by adding new QOTFs.

*/
package epic

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"
)

const (
	indexReset   = -1
	qotfLabel    = "QOTF"
	gotLabel     = "GOT"
	wantLabel    = "WANT"
	notLabel     = "(not) "
	reportFormat = "### %s in %s at line %d\n--- %-4s: %s\n--- %-4s: %v\n--- %-4s: %s%v\n"
)

var (
	extractContext = context
	headlines      = []string{
		"Uh-oh...",
		"Huston, we have a problem!",
		"Oh boy!",
		"Gimme 10 bucks and I'll make all your problems disappear *wink*",
		"To fail or not to fail, that is the question...",
		"Talk me out of it!",
		"I'm outta here!!!",
		"Holy failing tests, Batman!",
		"One can always count on you, uh?",
		"And the winner is...",
		"You're dead to me...",
		"Well done! Now I'll have to hang around until you won't fix this one.",
		"This is *not* happening...",
		"I am going for my coffee, when I'm back you'd better have this cleared!",
		"These are not the tests you're looking for.",
		"Of course this had to fail when I was about to leave for a beer.",
	}
	random *rand.Rand
)

// Win validates "got" against "good" for equality and fails "test" if they differ.
// Use this function every time your test's success is bound to a specific value.
func Win(test *testing.T, got interface{}, good interface{}) {
	if report, ok := validate(got, good, true); !ok {
		fmt.Print(report)
		test.Fail()
	}
}

// Fail validates "got" against "bad" for inequality and fails "test" if they're equal.
// Use this function every time your test's success is bound to any value except a specific one.
func Fail(test *testing.T, got interface{}, bad interface{}) {
	if report, ok := validate(got, bad, false); !ok {
		fmt.Print(report)
		test.Fail()
	}
}

func init() {
	random = rand.New(rand.NewSource(time.Now().Unix()))
}

func context() (string, string, int) {
	pc, file, line, ok := runtime.Caller(3)
	var function string

	if ok {
		function = runtime.FuncForPC(pc).Name()
		file = filepath.Base(file)
	} else {
		function = "???"
		file = "???"
		line = 1
	}

	return function, file, line
}

func validate(got interface{}, expected interface{}, truth bool) (report string, ok bool) {
	if reflect.DeepEqual(got, expected) != truth {
		var not string

		if !truth {
			not = notLabel
		}

		function, file, line := extractContext()
		report = fmt.Sprintf(
			reportFormat,
			function,
			file,
			line,
			qotfLabel,
			headlines[random.Intn(len(headlines))],
			gotLabel,
			got,
			wantLabel,
			not,
			expected,
		)

		ok = false
	} else {
		ok = true
	}

	return
}
