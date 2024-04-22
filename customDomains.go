package render

import (
	"fmt"
	"net/http"
)

type CustomDomain struct {
	ID                 *string `json:"id,omitempty"`
	Name               string  `json:"name"`
	DomainType         *string `json:"domainType,omitempty"`
	PublicSuffix       *string `json:"publicSuffix,omitempty"`
	RedirectForName    *string `json:"redirectForName,omitempty"`
	VerificationStatus *string `json:"verificationStatus,omitempty"`
	CreatedAt          *string `json:"createdAt,omitempty"`
	Server             *Server `json:"server,omitempty"`
}

type Server struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CustomDomains struct {
	CustomDomain `json:"customDomain"`
}

func (c *Client) GetCustomDomains(serviceId string) ([]CustomDomain, error) {
	var customDomains []CustomDomains
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, customDomainsPath), nil, &customDomains)
	if err != nil {
		return nil, err
	}

	return CustomDomainsToSlice(customDomains), nil
}

func (c *Client) GetCustomDomain(serviceId, idOrName string) (*CustomDomain, error) {
	var customDomain CustomDomain
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, customDomainsPath, idOrName), nil, &customDomain)
	if err != nil {
		return nil, err
	}

	return &customDomain, nil
}

func (c *Client) CreateCustomDomain(serviceId string, data CustomDomain) ([]CustomDomain, error) {
	var customDomain []CustomDomain
	err := c.doRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, customDomainsPath), data, &customDomain)
	if err != nil {
		return nil, err
	}

	return customDomain, nil
}

func (c *Client) DeleteCustomDomain(serviceId, idOrName string) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, customDomainsPath, idOrName), nil, nil)
}

func CustomDomainsToSlice(customDomains []CustomDomains) []CustomDomain {
	var result []CustomDomain
	for _, customDomain := range customDomains {
		result = append(result, customDomain.CustomDomain)
	}
	return result
}
