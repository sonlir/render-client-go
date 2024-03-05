package render

import (
	"fmt"
	"net/http"
)

const ownersPath = "owners"

type Owner struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

type Owners struct {
	Owner Owner `json:"owner"`
}

func (c *Client) GetOwners() (*[]Owners, error) {
	owners := []Owners{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.HostURL, ownersPath), nil, &owners)
	if err != nil {
		return nil, err
	}

	return &owners, nil
}

func (c *Client) GetOwner(id string) (*Owner, error) {
	owner := Owner{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.HostURL, ownersPath, id), nil, &owner)
	if err != nil {
		return nil, err
	}

	return &owner, nil
}
