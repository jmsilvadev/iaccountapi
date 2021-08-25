package apiclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jmsilvadev/form3libs/apiclient/internal/utils/mocks"
	"github.com/jmsilvadev/form3libs/apiclient/models"
)

const baseURL = "http://localhost:8080"

var accountFake *models.Account
var accountService *AccountService
var exUUID = "0f42ba70-e942-4a09-83d0-e8bd0a93f187"

func init() {
	accountService, _ = setNewClient(baseURL, &mocks.FakeMockAPIClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	})
	accountFake = HelperGetNewAccount()
}
func TestNewAccount(t *testing.T) {
	got := accountService.NewAccount()
	if got.Data.OrganisationID == "" {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			got, "")
	} else {
		t.Logf("Testing TestNewAccount: PASSED\n")
	}

	if got.Data.ID == "" {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			got, "A non empty value")
	} else {
		t.Logf("Testing TestNewAccount: PASSED\n")
	}
}

func TestFetch(t *testing.T) {
	got, _ := accountService.Fetch(exUUID)
	if got.Data.ID == exUUID {
		t.Logf("Testing Fetch: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			got.Data.ID, exUUID)
	}
}

func TestCreate(t *testing.T) {
	got, _ := accountService.Create(accountFake)
	if got.Data.ID == exUUID {
		t.Logf("Testing Create: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			got.Data.ID, exUUID)
	}

	toJSON = func(v interface{}) ([]byte, error) {
		return nil, errors.New("Fake Error")
	}

	_, err := accountService.Create(accountFake)
	if err.Error == "Fake Error" {
		t.Logf("Testing Fetch: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "Fake Error")
	}

	toJSON = func(v interface{}) ([]byte, error) {
		return json.Marshal(v)
	}

}

func TestDelete(t *testing.T) {
	got, _ := accountService.Delete(exUUID, 0)
	if got == true {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			got, false)
	}
}

func TestReadData(t *testing.T) {
	_, err := accountService.readData(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(mocks.FakeResponseBody)),
	})
	if err == nil {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "nil value")
	}

	_, err = accountService.readData(&http.Response{
		StatusCode: 200,
	})
	if err.Error() == "Empty Body" {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "nil value")
	}

	readBody = func(r io.Reader) ([]byte, error) {
		return nil, errors.New("Fake Error")
	}

	_, err = accountService.readData(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(mocks.FakeResponseBody)),
	})

	if err.Error() == "Fake Error" {
		t.Logf("Testing Fetch: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "Fake Error")
	}

	readBody = func(r io.Reader) ([]byte, error) {
		return ioutil.ReadAll(r)
	}
}

func TestDecodeResponse(t *testing.T) {
	_, err := accountService.decodeResponse(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(mocks.FakeResponseBody)),
	})
	if err == nil {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "nil value")
	}

	_, err = accountService.decodeResponse(&http.Response{
		StatusCode: 200,
	})
	if err.Error == "Empty Body" {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "nil value")
	}

	_, err = accountService.decodeResponse(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`"forcing-error""forcing-error"`)),
	})
	if err.Error == `invalid character '"' after top-level value` {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "nil value")
	}
}

func TestGetDeaultClient(t *testing.T) {
	got := getDefaultClient()
	if got == nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			got, "non nil value")
	} else {
		t.Logf("Testing TestNewAccount: PASSED\n")
	}
}

func TestSetNewClient(t *testing.T) {
	got, _ := setNewClient(baseURL, &mocks.FakeMockAPIClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	})

	if got == nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			got, "Non nil value")
	} else {
		t.Logf("Testing TestNewAccount: PASSED\n")
	}

	_, err := setNewClient(baseURL, nil)
	expect := "Invalid custom api client"
	if err.Error() != expect {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err.Error(), expect)
	} else {
		t.Logf("Testing TestNewAccount: PASSED\n")
	}

	_, err = setNewClient("", nil)
	expect = "Invalid custom api URL"
	if err.Error() != expect {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err.Error(), expect)
	} else {
		t.Logf("Testing TestNewAccount: PASSED\n")
	}
}

func TestDecodeResponseError(t *testing.T) {
	err := accountService.decodeResponseError(&http.Response{
		StatusCode: 400,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
			"error_message": "id is not a valid uuid"
		}`)),
	})
	if err.Error == "id is not a valid uuid" && err.StatusCode == 400 {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "nil value")
	}

	err = accountService.decodeResponseError(&http.Response{
		StatusCode: 204,
		Body:       ioutil.NopCloser(bytes.NewBufferString("")),
	})
	if err.Error == "No Content" {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: status %v message %s, expected: %v ",
			err.StatusCode, err.Error, "nil value")
	}

	err = accountService.decodeResponseError(&http.Response{
		StatusCode: 200,
	})
	if err.Error == "Empty Body" && err.StatusCode == 200 {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n status %v message %s, expected: %v ",
			err.StatusCode, err.Error, "nil value")
	}

	err = accountService.decodeResponseError(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`"forcing-error""forcing-error"`)),
	})
	if err.Error == `invalid character '"' after top-level value` && err.StatusCode == 200 {
		t.Logf("Testing Delete: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n status %v message %s, expected: %v ",
			err.StatusCode, err.Error, "nil value")
	}
}

func TestFetchError(t *testing.T) {
	accountService, _ = setNewClient(baseURL, &mocks.FakeMockErrorAPIClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	})
	accountFake = HelperGetNewAccount()
	_, err := accountService.Fetch(exUUID)
	if err.Error == "Fake Error" {
		t.Logf("Testing Fetch: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "Fake Error")
	}
}

func TestCreateError(t *testing.T) {
	_, err := accountService.Create(accountFake)
	if err.Error == "Fake Error" {
		t.Logf("Testing Fetch: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "Fake Error")
	}
}

func TestDeleteError(t *testing.T) {
	_, err := accountService.Delete(exUUID, 0)
	if err.Error == "Fake Error" {
		t.Logf("Testing Fetch: PASSED\n")
	} else {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, "Fake Error")
	}
}

func HelperGetNewAccount() *models.Account {
	ac := accountService.NewAccount()
	cls := "Personal"
	opt := true
	country := "GB"
	jnt := true
	status := "confirmed"
	swi := true
	ac.Data.ID = exUUID
	ac.Data.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	ac.Data.Type = "accounts"
	ac.Data.Attributes.AccountClassification = &cls
	ac.Data.Attributes.AccountMatchingOptOut = &opt
	ac.Data.Attributes.BankID = "400300"
	ac.Data.Attributes.Country = &country
	ac.Data.Attributes.JointAccount = &jnt
	ac.Data.Attributes.Name = []string{"Name"}
	ac.Data.Attributes.AlternativeNames = []string{"Name"}
	ac.Data.Attributes.Status = &status
	ac.Data.Attributes.Switched = &swi
	ac.Data.Attributes.BankIDCode = "GBDSC"
	ac.Data.Attributes.BaseCurrency = "GBP"
	ac.Data.Attributes.Bic = "NWBKGB22"
	ac.Data.Attributes.AccountNumber = ""

	return ac
}
