package mocks

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// FakeResponseBody is to use in tests
const FakeResponseBody = `{
	"data": {
		"attributes": {
			"account_classification": "Personal",
			"account_matching_opt_out": true,
			"alternative_names": null,
			"bank_id": "400300",
			"bank_id_code": "GBDSC",
			"base_currency": "GBP",
			"bic": "NWBKGB22",
			"country": "GB",
			"joint_account": true,
			"name": [
				"Name"
			],
			"status": "confirmed",
			"switched": true
		},
		"created_on": "2021-08-20T10:12:40.832Z",
		"id": "0f42ba70-e942-4a09-83d0-e8bd0a93f187",
		"modified_on": "2021-08-20T10:12:40.832Z",
		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		"type": "accounts",
		"version": 0
	},
	"links": {
		"self": "/v1/organisation/accounts/0f42ba70-e942-4a09-83d0-e8bd0a93f187"
	}
}`

// FakeMockAPIClient  mock to return custom responses
type FakeMockAPIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// FakeMockErrorAPIClient mock to return Fake Errors
type FakeMockErrorAPIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// RoundTripFunc returns a transport layer mock in a http client
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip returns a transport layer mock in a http client
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.String() != "http://localhost:8080/v1/organisation/accounts/" {
		return nil, errors.New("Internal Error")
	}
	return f(req), nil
}

// MockHTTPClient mock to return Fake HTTP Transport
func MockHTTPClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// SendRequest mock to return custom responses
func (cli *FakeMockAPIClient) SendRequest(method string, endpoint string, body []byte) (*http.Response, error) {

	switch method {
	case "GET":
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(FakeResponseBody)),
		}, nil
	case "DELETE":
		return &http.Response{
			StatusCode: 204,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`DELETED`)),
		}, nil
	case "POST":
		return &http.Response{
			StatusCode: 201,
			Body:       ioutil.NopCloser(bytes.NewBufferString(FakeResponseBody)),
		}, nil
	default:
		return &http.Response{
			StatusCode: 404,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Not Found")),
		}, nil
	}
}

// SendRequest mock to return Fake Errors
func (cli *FakeMockErrorAPIClient) SendRequest(method string, endpoint string, body []byte) (*http.Response, error) {

	switch method {
	case "GET":
		return nil, errors.New("Fake Error")
	case "DELETE":
		return nil, errors.New("Fake Error")
	case "POST":
		return nil, errors.New("Fake Error")
	default:
		return nil, nil
	}
}
