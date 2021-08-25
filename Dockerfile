FROM golang:1.13-alpine
ENV CGO_ENABLED 0
RUN mkdir /form3libs
COPY pkg/. /form3libs
WORKDIR /form3libs
RUN go get -u golang.org/x/lint/golint && go mod download && go mod tidy && go mod vendor
ENTRYPOINT ./run-tests.sh