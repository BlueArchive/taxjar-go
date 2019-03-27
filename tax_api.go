package taxjar

import (
	"encoding/json"
)

// TaxRepository defines the interface for working with Tax through the API.
type TaxRepository interface {
	get(taxParams) (Tax, error)
}

// TaxApi implements TaxRepository
type TaxApi struct {
	client *Client
}

func (api TaxApi) get(params taxParams) (Tax, error) {
	taxList := TaxList{}
	data, err := api.client.Post("/taxes", params)
	if err != nil {
		return taxList.Tax, err
	}
	err = json.Unmarshal(data, &taxList)
	return taxList.Tax, err
}

//
// cdodge: Add this function to validate VAT ID
//
func (api TaxApi) ValidateVATNumber(number string) (*VATIDValidation, error) {
	validation := VATIDValidation{}
	data, err := api.client.Get("/validation", &VATIDValidationParams{VAT: number})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &validation)
	return &validation, err
}
