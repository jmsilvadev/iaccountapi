package entities

//AccountData represents an account in the form3 org section.
//See https://api-docs.form3.tech/api.html#organisation-accounts for
//more information about fields.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
	CreatedOn      *string            `json:"created_on,omitempty"`
	ModifiedOn     *string            `json:"modified_on,omitempty"`
}

//AccountAttributes represents an account in the form3 org section.
//See https://api-docs.form3.tech/api.html#organisation-accounts for
//more information about fields.
type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

//Account represents an account in the form3 org section.
//See https://api-docs.form3.tech/api.html#organisation-accounts for
//more information about fields.
type Account struct {
	Data *AccountData `json:"data"`
}

// ErrorMessage represents an error returned by form3 API
type ErrorMessage struct {
	StatusCode int
	Error      string `json:"error_message"`
}

// ToErrorMessage transforms an error in a ErrorMessage struct
func ToErrorMessage(status int, err error) *ErrorMessage {
	return &ErrorMessage{
		StatusCode: status,
		Error:      err.Error(),
	}
}
