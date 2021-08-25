package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jmsilvadev/form3libs/apiclient/internal/requests"
)

type scenarios []struct {
	method      string
	queryString string
	body        []byte
	statusCode  int
}

const endpoint = "/v1/organisation/accounts/"

var client = MockHTTPClient(func(req *http.Request) *http.Response {
	if req.Method == "GET" {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	}

	if req.Method == "DELETE" {
		return &http.Response{
			StatusCode: 204,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	}

	if req.Method == "POST" {
		return &http.Response{
			StatusCode: 201,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	}

	return &http.Response{
		StatusCode: 404,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
		Header:     make(http.Header),
	}
})

func TestRoundTrip(t *testing.T) {
	req := requests.Client{
		BaseURL:    "http://localhost:8080/v1/organisation/accounts/",
		HTTPClient: client,
	}

	_, err := req.SendRequest("GET", "", nil)

	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}

}

func TestSimulateTransportInternalError(t *testing.T) {
	req := requests.Client{
		BaseURL:    "http://wrong-base-url",
		HTTPClient: client,
	}

	_, err := req.SendRequest("GET", "", nil)

	if err == nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}

}

func TestFakeMockAPIClient(t *testing.T) {

	req := FakeMockAPIClient{
		BaseURL:    "http://localhost:8080",
		HTTPClient: client,
	}

	sc := scenarios{
		{
			method:      "GET",
			queryString: endpoint,
			body:        []byte{},
			statusCode:  200,
		},
		{
			method:      "DELETE",
			queryString: endpoint,
			body:        []byte{},
			statusCode:  204,
		},
		{
			method:      "POST",
			queryString: endpoint,
			body:        []byte{},
			statusCode:  201,
		},
		{
			method:      "PUT",
			queryString: endpoint,
			body:        []byte{},
			statusCode:  404,
		},
	}

	for _, scenario := range sc {
		got, err := req.SendRequest(scenario.method, scenario.queryString, scenario.body)
		if err != nil {
			t.Error("Got an unexpected error: " + err.Error())
		}
		if got.StatusCode != scenario.statusCode {
			t.Logf("Testing scenario: %s  FAILED\n", scenario.method)
			t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n httpReponse: got: %v, expected: %v ",
				scenario.method, got.StatusCode, scenario.statusCode)
		} else {
			t.Logf("Testing scenario: %s  PASSED\n", scenario.method)
		}

	}
}

func TestFakeMockErrorAPIClient(t *testing.T) {
	req := FakeMockErrorAPIClient{
		BaseURL:    "http://localhost:8080",
		HTTPClient: &http.Client{},
	}

	_, err := req.SendRequest("GET", "", nil)

	if err.Error() != "Fake Error" {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "Fake Error")
	}

	_, err = req.SendRequest("POST", "", nil)

	if err.Error() != "Fake Error" {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err.Error(), "Fake Error")
	}

	_, err = req.SendRequest("DELETE", "", nil)

	if err.Error() != "Fake Error" {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err.Error(), "Fake Error")
	}

	_, err = req.SendRequest("PUT", "", nil)

	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err.Error(), nil)
	}

}
