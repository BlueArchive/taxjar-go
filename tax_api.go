package taxjar

import (
	"encoding/json"
)

// TaxRepository defines the interface for working with Tax through the API.
type TaxRepository interface {
	get(taxParams) (Tax, string, error)
}

// TaxApi implements TaxRepository
type TaxApi struct {
	client *Client
}

func (api TaxApi) get(params taxParams) (Tax, string, error) {
	taxList := TaxList{}
	data, err := api.client.Post("/taxes", params)
	if err != nil {
		return taxList.Tax, string(data), err
	}
	err = json.Unmarshal(data, &taxList)
	return taxList.Tax, string(data), err
}
