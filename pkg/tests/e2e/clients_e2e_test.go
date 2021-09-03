package e2e

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/jmsilvadev/iaccountapi/apiclient"
	"github.com/jmsilvadev/iaccountapi/apiclient/models"
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

var api = apiclient.NewClient()
var exUUID = uuid.NewString()

func TestE2ESimpleUse(t *testing.T) {

	u := uuid.NewString()
	acc := helperGenerateNewAccount(u)
	apiAcc := apiclient.NewClient()
	na, err := apiAcc.Create(acc)

	if err != nil {
		if apiDOWN := helperIsAPIDown(err.Error); apiDOWN == true {
			t.Fatalf("In E2E tests we need to use a real connection, please provide one ")
		}
	}

	if na.Data.ID != u {
		t.Errorf("Got and Expected are not equals.\n Err: got: %s, expected: %s ",
			na.Data.ID, u)
	}

	na, _ = apiAcc.Fetch(u)
	if na.Data.ID != u {
		t.Errorf("Got and Expected are not equals.\n Err: got: %s, expected: %s ",
			na.Data.ID, u)
	}

	resp, _ := apiAcc.Delete(u, 0)
	if resp != true {
		t.Errorf("Got and Expected are not equals.\n Err: got: %v, expected: %v ",
			resp, true)
	}

}

func TestE2EClientCreateTableDriven(t *testing.T) {
	for _, scenario := range helperCreateScenarios() {
		got, err := api.Create(scenario.input)
		if err != nil {
			if apiDOWN := helperIsAPIDown(err.Error); apiDOWN == true {
				t.Fatalf("In E2E tests we need to use a real connection, please provide one ")
			}
		}
		if scenario.err != "" {
			if err.Error != scenario.err {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n Err: got: %s, expected: %s ",
					scenario.name, err.Error, scenario.err)
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

func TestE2EClientFetchTableDriven(t *testing.T) {
	for _, scenario := range helperFetchScenarios() {
		got, err := api.Fetch(scenario.input)
		if err != nil {
			if apiDOWN := helperIsAPIDown(err.Error); apiDOWN == true {
				t.Fatalf("In E2E tests we need to use a real connection, please provide one ")
			}
		}
		if scenario.err != "" {
			if err.Error != scenario.err {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n Err: got: %s, expected: %s ",
					scenario.name, err.Error, scenario.err)
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

func TestE2EClientDeleteTableDriven(t *testing.T) {
	for _, scenario := range helperDeleteScenarios() {
		got, err := api.Delete(scenario.input.id, scenario.input.ver)
		if err != nil {
			if apiDOWN := helperIsAPIDown(err.Error); apiDOWN == true {
				t.Fatalf("In E2E tests we need to use a real connection, please provide one ")
			}
		}
		if scenario.err != "" {
			if err.Error != scenario.err {
				t.Errorf("Got and Expected are not equals.\n Scenario: %s,\n Err: got: %s, expected: %s ",
					scenario.name, err.Error, scenario.err)
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

func helperGenerateNewAccount(newID string) *models.Account {
	ac := api.NewAccount()

	cls := "Personal"
	opt := true
	country := "GB"
	jnt := true
	status := "confirmed"
	swi := true
	if newID != "" {
		ac.Data.ID = newID
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
	sc :=
		scenariosCreate{
			{
				"Should create an account and generate an uuid When receive a Create request without uuid",
				helperGenerateNewAccount(""),
				"",
			},
			{
				"Should create an account When receive a Create request with an uuid",
				helperGenerateNewAccount(exUUID),
				"",
			},
			{
				"Should not create an account When receive a Create request with an existent uuid",
				helperGenerateNewAccount(exUUID),
				"Account cannot be created as it violates a duplicate constraint",
			},
			{
				"Should not create an account When receive a Create request with an empty account",
				api.NewAccount(),
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
	newID := uuid.NewString()
	return scenariosFetch{
		{
			"Should fetch an account When receive a Fetch request with a valid uuid",
			exUUID,
			exUUID,
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
			newID,
			"",
			"record " + newID + " does not exist",
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
	newID := uuid.NewString()
	return scenariosDelete{
		{
			"Should delete an account When receive a Delete request with a valid uuid",
			&inputDelete{
				id:  exUUID,
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
				id:  newID,
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

func helperIsAPIDown(err string) bool {
	if apiDOWN := strings.Contains(err, "connection refused"); apiDOWN == true {
		return true
	}
	return false
}
