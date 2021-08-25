package entities

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestTransformAccount(t *testing.T) {
	accountJSON := `
	{
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
	}
	`
	errMsg := &Account{}
	err := json.Unmarshal([]byte(accountJSON), errMsg)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestTransformErrorMessage(t *testing.T) {
	errorJSON := `{"error_message": "validation failure"}`
	errMsg := &ErrorMessage{}
	err := json.Unmarshal([]byte(errorJSON), errMsg)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestToErrorMessage(t *testing.T) {
	err := ToErrorMessage(500, errors.New("Test Error"))
	if err.StatusCode != 500 && err.Error != "Test Error" {
		t.Errorf("Got and Expected are not equals.\n got: status %v message %s, expected: %v ",
			err.StatusCode, err.Error, "Test Error")
	}
}
