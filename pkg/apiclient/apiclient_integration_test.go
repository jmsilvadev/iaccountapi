package apiclient_test

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/jmsilvadev/form3libs/apiclient"
	"github.com/jmsilvadev/form3libs/apiclient/internal/entities"
	"github.com/jmsilvadev/form3libs/apiclient/internal/requests"
	"github.com/jmsilvadev/form3libs/apiclient/models"
)

type scenariosCreate []struct {
	name  string
	input *models.Account
	err   string
}

type scenariosFetch []struct {
	name   string
	input  string
	expect string
	err    string
}

type scenariosDelete []struct {
	name   string
	input  *inputDelete
	expect bool
	err    string
}

type inputDelete struct {
	id  string
	ver int
}

var accountService *apiclient.AccountService
var exUuid string
var req requests.Client

func init() {
	accountService = apiclient.NewClient()
}

func TestIntegrationClientCreateTableDriven(t *testing.T) {
	for _, scenario := range helperCreateScenarios() {
		got, err := accountService.Create(scenario.input)
		if apiDOWN := helperIsAPIDown(err); apiDOWN == true {
			t.Fatalf("In integration tests we need to use a real connection, please provide one ")
		}
		if scenario.err != "" {
			if err.Error != scenario.err {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n Err: got: status %v message %s, expected: %s ",
					scenario.name, err.StatusCode, err.Error, scenario.err)
			} else {
				t.Logf("Testing scenario: %s  PASSED\n", scenario.name)
			}
		} else {
			if *got.Data.Version != *new(int64) {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n Response: got: %v, expected: %v ",
					scenario.name, got, scenario.input)
			} else {
				t.Logf("Testing scenario: %s  PASSED\n", scenario.name)
			}
		}
	}
}

func TestIntegrationClientFetchTableDriven(t *testing.T) {
	for _, scenario := range helperFetchScenarios() {
		got, err := accountService.Fetch(scenario.input)
		if apiDOWN := helperIsAPIDown(err); apiDOWN == true {
			t.Fatalf("In integration tests we need to use a real connection, please provide one ")
		}
		if scenario.err != "" {
			if err.Error != scenario.err {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n Err: got: status %v message %s, expected: %s ",
					scenario.name, err.StatusCode, err.Error, scenario.err)
			} else {
				t.Logf("Testing scenario: %s  PASSED\n", scenario.name)
			}
		} else {
			if got.Data.ID != scenario.expect {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n Response: got: %v, expected: %v ",
					scenario.name, got, scenario.input)
			} else {
				t.Logf("Testing scenario: %s  PASSED\n", scenario.name)
			}
		}
	}
}

func TestIntegrationClientDeleteTableDriven(t *testing.T) {
	for _, scenario := range helperDeleteScenarios() {
		got, err := accountService.Delete(scenario.input.id, scenario.input.ver)
		if apiDOWN := helperIsAPIDown(err); apiDOWN == true {
			t.Fatalf("In integration tests we need to use a real connection, please provide one ")
		}
		if scenario.err != "" {
			if err.Error != scenario.err {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n Err: got: status %v message %s, expected: %s ",
					scenario.name, err.StatusCode, err.Error, scenario.err)
			} else {
				t.Logf("Testing scenario: %s  PASSED\n", scenario.name)
			}
		} else {
			if got == false {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n Response: got: %v, expected: %v ",
					scenario.name, got, scenario.input)
			} else {
				t.Logf("Testing scenario: %s  PASSED\n", scenario.name)
			}
		}
	}
}

func helperGenerateNewAccount(newUuid string) *models.Account {
	ac := accountService.NewAccount()

	cls := "Personal"
	opt := true
	country := "GB"
	jnt := true
	status := "confirmed"
	swi := true
	if newUuid != "" {
		ac.Data.ID = newUuid
	}
	ac.Data.OrganisationID = uuid.NewString()
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

func helperCreateScenarios() scenariosCreate {

	exUuid = uuid.NewString()

	sc :=
		scenariosCreate{
			{
				"Should create an account and generate an uuid When receive a Create request without uuid",
				helperGenerateNewAccount(""),
				"",
			},
			{
				"Should create an account When receive a Create request with an uuid",
				helperGenerateNewAccount(exUuid),
				"",
			},
			{
				"Should not create an account When receive a Create request with an existent uuid",
				helperGenerateNewAccount(exUuid),
				"Account cannot be created as it violates a duplicate constraint",
			},
			{
				"Should not create an account When receive a Create request with an empty account",
				accountService.NewAccount(),
				"validation failure list:\nvalidation failure list:\nvalidation failure list:\naccount_classification in body should be one of [Personal Business]\ncountry in body should match '^[A-Z]{2}$'\nname in body is required\nstatus in body should be one of [pending confirmed failed]\ntype in body is required",
			},
			{
				"Should not create an account When receive a Create request with a Nil",
				nil,
				"Invalid Account",
			},
			{
				"Should not create an account When receive a Create request with an invalid uuid",
				helperGenerateNewAccount("oooo-error"),
				"Please provide a valid id, invalid UUID length: 10",
			},
			{
				"Should not create an account When receive a Create request with an invalid organization uuid",
				&models.Account{
					Data: &models.AccountData{
						Attributes:     &models.AccountAttributes{},
						ID:             uuid.NewString(),
						OrganisationID: "oooo-error",
					},
				},
				"Please provide a valid organization id, invalid UUID length: 10",
			},
		}
	return sc
}

func helperFetchScenarios() scenariosFetch {
	newId := uuid.NewString()
	return scenariosFetch{
		{
			"Should fetch an account When receive a Fetch request with a valid uuid",
			exUuid,
			exUuid,
			"",
		},
		{
			"Should thrown an error When receive a Fetch request with an Invalid uuid",
			"invalid-uuid",
			"",
			"invalid UUID length: 12",
		},
		{
			"Should thrown an error When receive a Fetch request with a non-existent uuid",
			newId,
			"",
			"record " + newId + " does not exist",
		},
		{
			"Should thrown an error When receive a Fetch request with an empty uuid",
			"",
			"",
			"invalid UUID length: 0",
		},
		{
			"Should thrown an error When receive a Fetch request with an invalid uuid",
			"oooo-error",
			"",
			"invalid UUID length: 10",
		},
	}
}

func helperDeleteScenarios() scenariosDelete {
	newId := uuid.NewString()
	return scenariosDelete{
		{
			"Should delete an account When receive a Delete request with a valid uuid",
			&inputDelete{
				id:  exUuid,
				ver: 0,
			},
			true,
			"",
		},
		{
			"Should thrown an error When receive a Delete request with an Invalid uuid",
			&inputDelete{
				id:  "invalid-uuid",
				ver: 0,
			},
			false,
			"Please provide a valid id",
		},
		{
			"Should thrown an error When receive a Delete request with a non-existent uuid",
			&inputDelete{
				id:  newId,
				ver: 0,
			},
			false,
			"No Content",
		},
		{
			"Should thrown an error When receive a Delete request with an empty uuid",
			&inputDelete{
				id:  "",
				ver: 0,
			},
			false,
			"Please provide a valid id",
		},
	}
}

func helperIsAPIDown(err *entities.ErrorMessage) bool {
	if err != nil {
		if apiDOWN := strings.Contains(err.Error, "connection refused"); apiDOWN == true {
			return true
		}
	}
	return false
}
