package datahub

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAllDomains - Returns all user's Domain
func (c *Client) GetAllDomains(authToken *string) (*[]Domain, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Domains", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	Domains := []Domain{}
	err = json.Unmarshal(body, &Domains)
	if err != nil {
		return nil, err
	}

	return &Domains, nil
}

// GetDomain - Returns a specifc Domain
func (c *Client) GetDomain(DomainID string, authToken *string) (*Domain, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Domains/%s", c.HostURL, DomainID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	Domain := Domain{}
	err = json.Unmarshal(body, &Domain)
	if err != nil {
		return nil, err
	}

	return &Domain, nil
}

// CreateDomain - Create new Domain
func (c *Client) CreateDomain(DomainItems []DomainItem, authToken *string) (*Domain, error) {
	rb, err := json.Marshal(DomainItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/Domains", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	Domain := Domain{}
	err = json.Unmarshal(body, &Domain)
	if err != nil {
		return nil, err
	}

	return &Domain, nil
}

// UpdateDomain - Updates an Domain
func (c *Client) UpdateDomain(DomainID string, DomainItems []DomainItem, authToken *string) (*Domain, error) {
	rb, err := json.Marshal(DomainItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/Domains/%s", c.HostURL, DomainID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	Domain := Domain{}
	err = json.Unmarshal(body, &Domain)
	if err != nil {
		return nil, err
	}

	return &Domain, nil
}

// DeleteDomain - Deletes an Domain
func (c *Client) DeleteDomain(DomainID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/Domains/%s", c.HostURL, DomainID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "Deleted Domain" {
		return errors.New(string(body))
	}

	return nil
}
