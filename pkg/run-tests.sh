#!/bin/sh

if [ ! -d tests/coverage/ ]; then
    mkdir tests/coverage/
fi

if [ ! -d tests/outputs/ ]; then
    mkdir tests/outputs/
fi

go test -v ./... -cover -coverprofile=tests/coverage/cover.out > tests/outputs/output_test.log

go tool cover -html=tests/coverage/cover.out -o tests/coverage/coverage.html

CHECKFAIL=$(cat tests/outputs/output_test.log | grep FAIL)

cat tests/outputs/output_test.log

if [ -z "${CHECKFAIL}" ]; then
    echo "####################"
    echo "# All Tests PASSED #"
    echo "####################"
else 
    echo "################"
    echo "# Tests FAILED #"
    echo "################"
    exit 1
fi