package render

import (
	"fmt"
	"net/http"
)

const environmentVariablesPath = "env-vars"

type EnvironmentVariable struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type EnvironmentVariables struct {
	EnvVar EnvironmentVariable `json:"envVar"`
}

func (c *Client) GetEnvironmentVariables(serviceId string) (*[]EnvironmentVariables, error) {
	environmentVariables := []EnvironmentVariables{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, environmentVariablesPath), nil, &environmentVariables)
	if err != nil {
		return nil, err
	}

	return &environmentVariables, nil
}

func (c *Client) UpdateEnvironmentVariables(serviceId string, data []EnvironmentVariable) (*[]EnvironmentVariables, error) {
	environmentVariable := []EnvironmentVariables{}
	err := c.doRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, environmentVariablesPath), data, &environmentVariable)
	if err != nil {
		return nil, err
	}

	return &environmentVariable, nil
}
