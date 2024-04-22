package render

import (
	"fmt"
	"net/http"
	"net/url"
)

type RegistryCredentialRequest struct {
	Registry  string `json:"registry"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	AuthToken string `json:"authToken"`
	OwnerId   string `json:"ownerId"`
}
type RegistryCredentialResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Registry string `json:"registry"`
	Username string `json:"username"`
}

type GetRegistryCredentialsArgs struct {
	Name string
}

func (c *Client) GetRegistryCredentials(args *GetRegistryCredentialsArgs) ([]RegistryCredentialResponse, error) {
	var registryCredentials []RegistryCredentialResponse
	parameters := url.Values{}
	url, err := url.Parse(fmt.Sprintf("%s/%s", c.HostURL, registrycredentialsPath))
	if err != nil {
		return nil, err
	}
	if args != nil {
		if args.Name != "" {
			parameters.Add("name", args.Name)
		}
	}

	url.RawQuery = parameters.Encode()

	err = c.doRequest(http.MethodGet, url.String(), nil, &registryCredentials)
	if err != nil {
		return nil, err
	}

	return registryCredentials, nil
}

func (c *Client) GetRegistryCredential(id string) (*RegistryCredentialResponse, error) {
	var registryCredential *RegistryCredentialResponse
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.HostURL, registrycredentialsPath, id), nil, &registryCredential)
	if err != nil {
		return nil, err
	}

	return registryCredential, nil
}

func (c *Client) CreateRegistryCredential(data RegistryCredentialRequest) (*RegistryCredentialResponse, error) {
	var registryCredential *RegistryCredentialResponse

	registryCredentials, err := c.GetRegistryCredentials(&GetRegistryCredentialsArgs{Name: data.Name})
	if err != nil {
		return nil, err
	}
	if registryCredentials != nil {
		return nil, fmt.Errorf("the name `%s` is already in use. Please use a different name", data.Name)
	}

	err = c.doRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.HostURL, registrycredentialsPath), data, &registryCredential)
	if err != nil {
		return nil, err
	}

	return registryCredential, nil
}

func (c *Client) UpdateRegistryCredential(id string, data RegistryCredentialRequest) (*RegistryCredentialResponse, error) {
	var registryCredential *RegistryCredentialResponse

	registryCredentials, err := c.GetRegistryCredentials(&GetRegistryCredentialsArgs{Name: data.Name})
	if err != nil {
		return nil, err
	}
	if registryCredentials != nil && registryCredentials[0].ID != id {
		return nil, fmt.Errorf("the name `%s` is already in use. Please use a different name", data.Name)
	}

	err = c.doRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s", c.HostURL, registrycredentialsPath, id), data, &registryCredential)
	if err != nil {
		return nil, err
	}

	return registryCredential, nil
}

func (c *Client) DeleteRegistryCredential(id string) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.HostURL, registrycredentialsPath, id), nil, nil)
}
