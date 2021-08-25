package models

import (
	"encoding/json"
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
			"id": "0f42ba70-e942-4a09-83d0-e8bd0a93f187",
			"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"type": "accounts"
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
