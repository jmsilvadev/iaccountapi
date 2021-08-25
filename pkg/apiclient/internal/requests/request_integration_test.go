package requests_test

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/jmsilvadev/form3libs/apiclient/internal/requests"
)

type scenarios []struct {
	method      string
	queryString string
	body        []byte
	statusCode  int
}

const endpoint = "/v1/organisation/accounts/"

func TestIntegrationSendRequestTableDriven(t *testing.T) {

	var req requests.Client
	req.BaseURL = "http://localhost:8080"
	if os.Getenv("ACCOUNTAPI_URL") != "" {
		req.BaseURL = os.Getenv("ACCOUNTAPI_URL")
	}
	req.HTTPClient = &http.Client{}

	newUuid := uuid.NewString()
	newBody := []byte(`{"data":{"attributes":{"account_classification":"Personal","account_matching_opt_out":true,"alternative_names":["Name"],"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","joint_account":true,"name":["Name"],"status":"confirmed","switched":true},"id":"` + newUuid + `","organisation_id":"fc61e1cc-3ebf-4d3c-a682-1bb3898ee7b9","type":"accounts","version":0}}`)
	sc := scenarios{
		{
			method:      "POST",
			queryString: endpoint,
			body:        newBody,
			statusCode:  201,
		},
		{
			method:      "GET",
			queryString: endpoint + newUuid,
			body:        []byte{},
			statusCode:  200,
		},
		{
			method:      "DELETE",
			queryString: endpoint + newUuid + "?version=0",
			body:        []byte{},
			statusCode:  204,
		},
		{
			method:      "cmjs:/:8989:",
			queryString: endpoint,
			body:        []byte{},
			statusCode:  204,
		},
	}

	for _, scenario := range sc {
		got, err := req.SendRequest(scenario.method, scenario.queryString, scenario.body)
		if apiDOWN := helperIsAPIDown(err); apiDOWN == true {
			t.Fatalf("In integration tests we need to use a real connection, please provide one ")
		}
		if err != nil {
			if err.Error() != `net/http: invalid method "cmjs:/:8989:"` {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n httpReponse: got: %v, expected: %v ",
					scenario.method, got.StatusCode, scenario.statusCode)
			} else {
				t.Logf("Testing scenario: %s  PASSED\n", scenario.method)
			}
		} else {
			if got.StatusCode != scenario.statusCode {
				t.Logf("Testing scenario: %s  FAILED\n", scenario.method)
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n httpReponse: got: %v, expected: %v ",
					scenario.method, got.StatusCode, scenario.statusCode)
			} else {
				t.Logf("Testing scenario: %s  PASSED\n", scenario.method)
			}
		}
	}
}

func TestIntegrationSendRequestWithAPIDown(t *testing.T) {

	var req requests.Client
	req.BaseURL = "http://localhost"
	req.HTTPClient = &http.Client{}

	expect := "Get http://localhost: dial tcp 127.0.0.1:80: connect: connection refused"
	_, err := req.SendRequest("GET", "", nil)
	if err.Error() == expect {
		t.Logf("Testing scenario: %s  PASSED\n", "GET")
	} else {
		t.Logf("Testing scenario: %s  FAIL\n", "GET")
		t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n httpReponse: got: %v, expected: %v ",
			"GET", err.Error(), expect)
	}
}

func TestIntegrationSendRequestWithNonExistentbaseURL(t *testing.T) {

	var req requests.Client
	req.BaseURL = "http://fake-base-url:53"
	req.HTTPClient = &http.Client{}

	_, err := req.SendRequest("GET", "", nil)
	if err != nil {
		t.Logf("Testing scenario: %s  PASSED\n", "GET")
	} else {
		t.Logf("Testing scenario: %s  FAIL\n", "GET")
		t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n httpReponse: got: %v, expected: %v ",
			"GET", err, "non nil value")
	}
}

func helperIsAPIDown(err error) bool {
	if err != nil {
		if apiDOWN := strings.Contains(err.Error(), "connection refused"); apiDOWN == true {
			return true
		}
	}
	return false
}
