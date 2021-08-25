package apiclient

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jmsilvadev/form3libs/apiclient/internal/entities"
	"github.com/jmsilvadev/form3libs/apiclient/internal/requests"
	"github.com/jmsilvadev/form3libs/apiclient/models"
)

const endpoint = "/v1/organisation/accounts/"

var toJSON = func(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

var fromJSON = func(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}

var readBody = func(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

// APIClient is an interface to give flexibility to use any http client
type APIClient interface {
	SendRequest(method string, endpoint string, body []byte) (*http.Response, error)
}

// AccountService is the point receiver to agregate the methods to communicate with Form3API
type AccountService struct {
	req APIClient
}

// getDefaultClient returns the protected default http client to communicate with the Form3API
func getDefaultClient() APIClient {
	baseURL := "http://localhost:8080"
	if os.Getenv("ACCOUNTAPI_URL") != "" {
		baseURL = os.Getenv("ACCOUNTAPI_URL")
	}
	return &requests.Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// setNewClient injects a custom new client into apiclient
func setNewClient(baseURL string, api APIClient) (*AccountService, error) {
	if strings.TrimSpace(baseURL) == "" {
		return nil, errors.New("Invalid custom api URL")
	}
	if api == nil {
		return nil, errors.New("Invalid custom api client")
	}
	return &AccountService{
		req: api,
	}, nil
}

// NewClient returns a reference to AccountService Object to manipulate all methods inside of the pkg
func NewClient() *AccountService {
	return &AccountService{
		req: getDefaultClient(),
	}
}

// NewAccount returns a valid and empty representation of a model from form3API
func (acc AccountService) NewAccount() *models.Account {
	return &models.Account{
		Data: &models.AccountData{
			Attributes: &models.AccountAttributes{
				AccountClassification:   new(string),
				AccountMatchingOptOut:   new(bool),
				AccountNumber:           "",
				AlternativeNames:        []string{},
				BankID:                  "",
				BankIDCode:              "",
				BaseCurrency:            "",
				Bic:                     "",
				Country:                 new(string),
				Iban:                    "",
				JointAccount:            new(bool),
				Name:                    []string{},
				SecondaryIdentification: "",
				Status:                  new(string),
				Switched:                new(bool),
			},
			ID:             uuid.NewString(),
			OrganisationID: uuid.NewString(),
			Type:           "",
		},
	}
}

// Fetch method returning a model from the representation of an account in Form3API
func (acc AccountService) Fetch(accountUUID string) (*entities.Account, *entities.ErrorMessage) {
	if _, err := uuid.Parse(accountUUID); err != nil {
		return nil, entities.ToErrorMessage(406, err)
	}
	resp, err := acc.req.SendRequest("GET", endpoint+accountUUID, nil)
	if err != nil {
		return nil, entities.ToErrorMessage(406, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, acc.decodeResponseError(resp)
	}
	return acc.decodeResponse(resp)
}

// Delete method removes an account from a Form3API
func (acc AccountService) Delete(accountUUID string, version int) (bool, *entities.ErrorMessage) {
	if _, err := uuid.Parse(accountUUID); err != nil || strings.TrimSpace(accountUUID) == "" {
		return false, entities.ToErrorMessage(406, errors.New("Please provide a valid id"))
	}
	resp, err := acc.req.SendRequest("DELETE", endpoint+accountUUID+"?version="+strconv.Itoa(version), nil)
	if err != nil {
		return false, entities.ToErrorMessage(406, err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, acc.decodeResponseError(resp)
	}
	return true, nil
}

// Create method cretates an account in Form3API and returning a model from the representation of an account in Form3API
func (acc AccountService) Create(account *models.Account) (*entities.Account, *entities.ErrorMessage) {
	if account == nil {
		return nil, entities.ToErrorMessage(406, errors.New("Invalid Account"))
	}

	if _, err := uuid.Parse(account.Data.ID); err != nil {
		return nil, entities.ToErrorMessage(406, errors.New("Please provide a valid id, "+err.Error()))
	}

	if _, err := uuid.Parse(account.Data.OrganisationID); err != nil {
		return nil, entities.ToErrorMessage(406, errors.New("Please provide a valid organization id, "+err.Error()))
	}

	body, err := toJSON(account)
	if err != nil {
		return nil, entities.ToErrorMessage(406, err)
	}
	resp, err := acc.req.SendRequest("POST", endpoint, body)
	if err != nil {
		return nil, entities.ToErrorMessage(406, err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, acc.decodeResponseError(resp)
	}
	return acc.decodeResponse(resp)
}

func (acc AccountService) readData(resp *http.Response) ([]byte, error) {
	if resp.Body == nil {
		return nil, errors.New("Empty Body")
	}
	defer resp.Body.Close()
	content, err := readBody(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (acc AccountService) decodeResponse(resp *http.Response) (*entities.Account, *entities.ErrorMessage) {
	var account *entities.Account
	content, err := acc.readData(resp)
	if err != nil {
		return nil, entities.ToErrorMessage(resp.StatusCode, err)
	}
	err = fromJSON(content, &account)
	if err != nil {
		return nil, entities.ToErrorMessage(resp.StatusCode, err)
	}
	return account, nil
}

func (acc AccountService) decodeResponseError(resp *http.Response) *entities.ErrorMessage {
	content, err := acc.readData(resp)
	if err != nil {
		return entities.ToErrorMessage(resp.StatusCode, err)
	}

	if string(content) == "" {
		return entities.ToErrorMessage(resp.StatusCode, errors.New("No Content"))
	}

	var contentError *entities.ErrorMessage
	err = fromJSON(content, &contentError)
	if err != nil {
		return entities.ToErrorMessage(resp.StatusCode, err)
	}
	contentError.StatusCode = resp.StatusCode
	return contentError
}
