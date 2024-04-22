package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type Owner struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

type Owners struct {
	Owner `json:"owner"`
}

type GetOwnersArgs struct {
	Name  string
	Email string
}

func (c *Client) GetOwners(args *GetOwnersArgs) ([]Owner, error) {
	var owners []Owners
	parameters := url.Values{}
	url, err := url.Parse(fmt.Sprintf("%s/%s", c.HostURL, ownersPath))
	if err != nil {
		return nil, err
	}
	if args != nil {
		if args.Name != "" {
			parameters.Add("name", args.Name)
		}
		if args.Email != "" {
			parameters.Add("email", args.Email)
		}
	}

	url.RawQuery = parameters.Encode()

	err = c.doRequest(http.MethodGet, url.String(), nil, &owners)
	if err != nil {
		return nil, err
	}

	return OwnersToSlice(owners), nil
}

func (c *Client) GetOwner(id string) (*Owner, error) {
	var owner *Owner
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.HostURL, ownersPath, id), nil, &owner)
	if err != nil {
		return nil, err
	}

	return owner, nil
}

func OwnersToSlice(owners []Owners) []Owner {
	var result []Owner
	for _, owner := range owners {
		result = append(result, owner.Owner)
	}
	return result
}
