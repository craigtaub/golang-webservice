# Golang simple Webservice

Run with

    go run service.go

API endpoint:

    `http://localhost:8080/return-jsonarray/`

API processing endpoint

    `http://localhost:8080/fetch-jsonarray/`

Setup:

    export GOPATH=<root directory>
    go get github.com/stretchr/testify

Tests:

    go run service.go &
    go test handlers
