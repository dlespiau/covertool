
[![Build Status](https://travis-ci.org/dlespiau/covertool.svg?branch=master)](https://travis-ci.org/dlespiau/covertool)

# `covertool`

`covertool` is a command line utility to manipulate go coverage reports. It can:

- **Merge** reports together into one single report. This is useful to gather
  reports produced by different runs and consolidate the results.

# Coverage beyond unit tests

Code coverage shouldn't be for unit tests only. There are many interesting
things to say about getting coverage data for integration (end to end) tests.
One may even want to gather coverage for a running service in production to get
hot paths or detect unused part of an application.

I call this coverage-instrumented go binaries. The full story can be read in
this [blog post](https://damien.lespiau.name/posts/2017-01-29-building-and-using-coverage-instrumented-programs-with-go/).


This repository contains support packages and tools to produce
and use coverage-instrumented Go programs.

Package [cover](https://github.com/dlespiau/covertool/tree/master/pkg/cover)
can be used to build instrumented programs.

Package [exit](https://github.com/dlespiau/covertool/tree/master/pkg/exit)
is an atexit implementation.

The `covertool` utility can merge profiles produced by different runs of the
same binary and display the resulting code coverage:

```
$ go install github.com/dlespiau/covertool
$ covertool merge -o all.go unit-tests.cov usecase1.cov usecase2.cov error1.cov error2.cov ...
$ covertool report all.go
coverage: 92.9% of statements
```

Finally, the `example/calc` directory contains a fully working example:

```
$ cd $GOPATH/src/github.com/dlespiau/covertool/examples/calc
$ ./run-tests.sh 
• Build the coverage-instrumented version of calc

• Run the unit tests
PASS
coverage: 7.1% of statements
ok  	github.com/dlespiau/covertool/examples/calc	0.003s

• Cover the sub() function
• Result: coverage: 57.1% of statements

• Cover the error path taken when not enough arguments are provided
expected 3 arguments, got 1
• Result: coverage: 21.4% of statements

• Cover the error path taken when providing an unknown operation
unknown operation: mul
• Result: coverage: 50.0% of statements

• Merge all coverage profiles and report the total coverage
coverage: 92.9% of statements
```
