package main

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
	EnvironmentVariable `json:"envVar"`
}

func (c *Client) GetEnvironmentVariables(serviceId string) ([]EnvironmentVariable, error) {
	environmentVariables := []EnvironmentVariables{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, environmentVariablesPath), nil, &environmentVariables)
	if err != nil {
		return nil, err
	}

	return EnvironmentVariablesToSlice(environmentVariables), nil
}

func (c *Client) UpdateEnvironmentVariables(serviceId string, data []EnvironmentVariable) ([]EnvironmentVariable, error) {
	environmentVariables := []EnvironmentVariables{}
	err := c.doRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s/%s", c.HostURL, servicesPath, serviceId, environmentVariablesPath), data, &environmentVariables)
	if err != nil {
		return nil, err
	}

	return EnvironmentVariablesToSlice(environmentVariables), nil
}

func EnvironmentVariablesToSlice(environmentVariables []EnvironmentVariables) []EnvironmentVariable {
	var result []EnvironmentVariable
	for _, environmentVariable := range environmentVariables {
		result = append(result, environmentVariable.EnvironmentVariable)
	}
	return result
}
