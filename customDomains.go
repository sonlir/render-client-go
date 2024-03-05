package render

import (
	"fmt"
	"net/http"
	"time"
)

const customDomainsPath = "custom-domains"

type CustomDomain struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	DomainType         string    `json:"domainType"`
	PublicSuffix       string    `json:"publicSuffix"`
	RedirectForName    string    `json:"redirectForName"`
	VerificationStatus string    `json:"verificationStatus"`
	CreatedAt          time.Time `json:"createdAt"`
	Server             Server    `json:"server"`
}

type Server struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CustomDomains struct {
	CustomDomain CustomDomain `json:"customDomain"`
}

type CustomDomainData struct {
	Name string `json:"name"`
}

func (c *Client) GetCustomDomains(serviceId string) (*[]CustomDomains, error) {
	customDomains := []CustomDomains{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, customDomainsPath), nil, &customDomains)
	if err != nil {
		return nil, err
	}

	return &customDomains, nil
}

func (c *Client) GetCustomDomain(serviceId, idOrName string) (*CustomDomain, error) {
	customDomain := CustomDomain{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, customDomainsPath, idOrName), nil, &customDomain)
	if err != nil {
		return nil, err
	}

	return &customDomain, nil
}

func (c *Client) CreateCustomDomain(serviceId string, data CustomDomainData) (*[]CustomDomain, error) {
	customDomain := []CustomDomain{}
	err := c.doRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, customDomainsPath), data, &customDomain)
	if err != nil {
		return nil, err
	}

	return &customDomain, nil
}

func (c *Client) DeleteCustomDomain(serviceId, idOrName string) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, customDomainsPath, idOrName), nil, nil)
}
