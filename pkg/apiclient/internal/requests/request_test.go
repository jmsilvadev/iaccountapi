package requests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jmsilvadev/iaccountapi/apiclient/internal/utils/mocks"
)

type scenarios []struct {
	method      string
	queryString string
	body        []byte
	statusCode  int
}

var client *http.Client

const endpoint = "/v1/organisation/accounts/"

func TestSendRequestTableDriven(t *testing.T) {

	client = mocks.MockHTTPClient(func(req *http.Request) *http.Response {
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

	req := Client{
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

func TestSimulateTransportInternalError(t *testing.T) {
	req := Client{
		BaseURL:    "http://wrong-base-url",
		HTTPClient: client,
	}

	_, err := req.SendRequest("GET", "", nil)

	if err == nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}

}

func TestSimulateRequesInternalError(t *testing.T) {
	req := Client{
		BaseURL:    "http://wrong-base-url",
		HTTPClient: client,
	}

	_, err := req.SendRequest("NONEXISTENT", "", nil)

	fmt.Println(err)

	if err == nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}

}
