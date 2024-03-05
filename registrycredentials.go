package render

import (
	"fmt"
	"net/http"
)

const registrycredentialsPath = "registrycredentials"

type RegistryCredential struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Registry string `json:"registry"`
	Username string `json:"username"`
}

type RegistryCredentialData struct {
	Registry  string `json:"registry"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	AuthToken string `json:"authToken"`
	OwnerId   string `json:"ownerId"`
}

func (c *Client) GetRegistryCredentials() (*[]RegistryCredential, error) {
	registryCredentials := []RegistryCredential{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.HostURL, registrycredentialsPath), nil, &registryCredentials)
	if err != nil {
		return nil, err
	}

	return &registryCredentials, nil
}

func (c *Client) GetRegistryCredential(id string) (*RegistryCredential, error) {
	registryCredential := RegistryCredential{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.HostURL, registrycredentialsPath, id), nil, &registryCredential)
	if err != nil {
		return nil, err
	}

	return &registryCredential, nil
}

func (c *Client) CreateRegistryCredential(data RegistryCredentialData) (*RegistryCredential, error) {
	registryCredential := RegistryCredential{}
	err := c.doRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.HostURL, registrycredentialsPath), data, &registryCredential)
	if err != nil {
		return nil, err
	}

	return &registryCredential, nil
}

func (c *Client) UpdateRegistryCredential(id string, data RegistryCredentialData) (*RegistryCredential, error) {
	registryCredential := RegistryCredential{}
	err := c.doRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s", c.HostURL, registrycredentialsPath, id), data, &registryCredential)
	if err != nil {
		return nil, err
	}

	return &registryCredential, nil
}

func (c *Client) DeleteRegistryCredential(id string) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.HostURL, registrycredentialsPath, id), nil, nil)
}
