#!/bin/bash

# To run in the github.com/dlespiau/covertool/examples/calc directory

function fatal() {
    echo $1
    exit 1
}

echo "• Build the coverage-instructed version of calc"
go test -o calc -covermode count &> /dev/null

echo "• Run the unit tests"
go test -covermode count -coverprofile unit-tests.cov | grep coverage

# Run calc with some combination of arguments to reach code paths not tested
# by unit tests.
# One would also check calc returns a proper error message and exit status but
# I've omitted those for simplicity.

echo "• Cover the sub() function"
result="$(./calc -test.coverprofile=sub.cov sub 1 2)"
[[ $result != "-1"* ]] && fatal "expected -1 got $result"
echo "$result" | grep coverage

echo "• Cover the error path taken when not enough arguments are provided"
./calc -test.coverprofile=error1.cov foo | grep coverage

echo "• Cover the error path taken when providing an unknown operation"
./calc -test.coverprofile=error2.cov mul 3 4 | grep coverage

# time to merge all coverage profiles

go install github.com/dlespiau/covertool

echo "• Merge all coverage profiles and report the total coverage"
covertool merge -o all.cov unit-tests.cov sub.cov error1.cov error2.cov
covertool coverage all.cov
