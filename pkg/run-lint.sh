#!/bin/sh

if [ ! -d tests/coverage/ ]; then
    mkdir tests/coverage/
fi

if [ ! -d tests/outputs/ ]; then
    mkdir tests/outputs/
fi

GOVET=$(go vet ./...)
if [ -z "${GOVET}" ]; then
    echo "Vet Success"
else 
    echo "Vet Fail"
    echo $GOVET
    exit 1
fi

GOFMT=$(go fmt ./...)
if [ -z "${GOFMT}" ]; then
    echo "Fmt Success"
else 
    echo "Fmt Fail"
    echo $GOFMT
    exit 1
fi

golint ./... > tests/outputs/output_lint.log
GOLINT=$(cat tests/outputs/output_lint.log)

if [ -z "${GOLINT}" ]; then
    echo "Lint Success"
else 
    cat tests/outputs/output_lint.log
    echo "#############"
    echo "# Lint Fail #"
    echo "#############"
    exit 1
fi