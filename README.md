epic
====
[![Build Status](https://travis-ci.org/federico-lox/epic.svg?branch=master)][travisCI]
[![Build Status](https://drone.io/github.com/federico-lox/epic/status.png)][drone.io]
[![Coverage Status](https://coveralls.io/repos/federico-lox/epic/badge.png?branch=master)][coveralls]
[![GoDoc](https://godoc.org/github.com/federico-lox/epic?status.png)][godoc]

Epic streamlines validating test results, it moves away from the richness of methods that characterizes the
`testing` package and it diverges from other testing frameworks and libraries by avoiding an as rich library of matchers.

Epic goals are to stay true to Go's minimalism and practicality while adding some fun to writing tests and going through
executions results.

Features
--------

* Simple syntax with humor: Only two methods, that's all you'll have to deal with
* 100% integrated with "go test", works side-by-side with the "testing" package
* Clear, well formatted output for failures
* No magic, ~100 LoC of idiomatic code that you can read and **understand**
* QOTF: For each failure a Quote of The Fail to make you smile at your misery

Documentation
-------------

Epic's documentation is available at [godoc.org][godoc]


[travisCI]: https://travis-ci.org/federico-lox/epic "Build status (Go 1.0, 1.1, 1.2, tip at TravisCI)"
[drone.io]: https://drone.io/github.com/federico-lox/epic/latest "Build status (Go latest stable and coverage at Drone.io)"
[coveralls]: https://coveralls.io/r/federico-lox/epic?branch=master "Code coverage"
[godoc]: https://godoc.org/github.com/federico-lox/epic "Package documentation"
