package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type WebServiceDeployResponse struct {
	Service WebServiceResponse `json:"service"`
}

type WebServiceRequest struct {
	Type           string                   `json:"type"`
	Name           string                   `json:"name"`
	OwnerID        string                   `json:"ownerId"`
	Repo           *string                  `json:"repo,omitempty"`
	AutoDeploy     *string                  `json:"autoDeploy,omitempty"`
	Branch         *string                  `json:"branch,omitempty"`
	Image          *Image                   `json:"image,omitempty"`
	BuildFilter    *BuildFilter             `json:"buildFilter,omitempty"`
	RootDir        *string                  `json:"rootDir,omitempty"`
	EnvVars        []EnvironmentVariable    `json:"envVars,omitempty"`
	SecretFiles    []SecretFiles            `json:"secretFiles,omitempty"`
	ServiceDetails WebServiceDetailsRequest `json:"serviceDetails"`
}

type WebServiceDetailsRequest struct {
	Disk                       *DiskRequest        `json:"disk,omitempty"`
	Env                        string              `json:"env"`
	EnvSpecificDetails         *EnvSpecificDetails `json:"envSpecificDetails,omitempty"`
	HealthCheckPath            *string             `json:"healthCheckPath,omitempty"`
	NumInstances               int64               `json:"numInstances"`
	Plan                       *string             `json:"plan,omitempty"`
	PullRequestPreviewsEnabled *string             `json:"pullRequestPreviewsEnabled,omitempty"`
	Region                     *string             `json:"region,omitempty"`
}

type WebServiceResponse struct {
	ID             string                    `json:"id"`
	AutoDeploy     string                    `json:"autoDeploy"`
	Branch         *string                   `json:"branch,omitempty"`
	BuildFilter    *BuildFilter              `json:"buildFilter,omitempty"`
	CreateAt       string                    `json:"createdAt"`
	ImagePath      *string                   `json:"imagePath,omitempty"`
	Name           string                    `json:"name"`
	NotifyOnFail   string                    `json:"notifyOnFail"`
	OwnerID        string                    `json:"ownerId"`
	Repo           *string                   `json:"repo,omitempty"`
	RootDir        *string                   `json:"rootDir,omitempty"`
	Slug           string                    `json:"slug"`
	Suspended      string                    `json:"suspended"`
	Suspenders     []string                  `json:"suspenders"`
	Type           string                    `json:"type"`
	UpdatedAt      string                    `json:"updatedAt"`
	EnvVars        []EnvironmentVariable     `json:"envVars,omitempty"`
	ServiceDetails WebServiceDetailsResponse `json:"serviceDetails"`
}

type WebServiceDetailsResponse struct {
	Autoscaling                *Autoscaling       `json:"autoscaling,omitempty"`
	Disk                       *DiskResponse      `json:"disk,omitempty"`
	Env                        string             `json:"env"`
	EnvSpecificDetails         EnvSpecificDetails `json:"envSpecificDetails"`
	HealthCheckPath            string             `json:"healthCheckPath"`
	NumInstances               int64              `json:"numInstances"`
	OpenPorts                  []OpenPort         `json:"openPorts"`
	ParentServer               *ParentServer      `json:"parentServer,omitempty"`
	Plan                       string             `json:"plan"`
	PullRequestPreviewsEnabled string             `json:"pullRequestPreviewsEnabled"`
	Region                     string             `json:"region"`
	URL                        string             `json:"url"`
	BuildPlan                  string             `json:"buildPlan"`
}

type Autoscaling struct {
	Enabled  bool                 `json:"enabled"`
	Min      int64                `json:"min,omitempty"`
	Max      int64                `json:"max,omitempty"`
	Criteria *AutoscalingCriteria `json:"criteria,omitempty"`
}

type AutoscalingCriteria struct {
	CPU    AutoscalingCriteriaObject `json:"cpu"`
	Memory AutoscalingCriteriaObject `json:"memory"`
}

type AutoscalingCriteriaObject struct {
	Enabled    bool  `json:"enabled"`
	Percentage int64 `json:"percentage"`
}

type OpenPort struct {
	Port     int64  `json:"port"`
	Protocol string `json:"protocol"`
}

type ParentServer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WebServices struct {
	WebServiceResponse `json:"service"`
}

type Scale struct {
	NumInstances int64 `json:"numInstances"`
}

func (c *Client) GetWebServices(name *string) ([]WebServiceResponse, error) {
	var services []WebServices
	parameters := url.Values{}
	url, err := url.Parse(fmt.Sprintf("%s/%s", c.HostURL, servicesPath))
	if err != nil {
		return nil, err
	}
	if name != nil {
		parameters.Add("name", *name)
	}
	parameters.Add("type", "web_service")
	url.RawQuery = parameters.Encode()

	err = c.doRequest(http.MethodGet, url.String(), nil, &services)
	if err != nil {
		return nil, err
	}

	return c.WebServicesToSlice(services), nil
}

func (c *Client) GetWebService(id string) (*WebServiceResponse, error) {
	service := WebServiceResponse{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.HostURL, servicesPath, id), nil, &service)
	if err != nil {
		return nil, err
	}

	service.EnvVars, err = c.GetEnvironmentVariables(id)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (c *Client) CreateWebService(data WebServiceRequest, autoscaling *Autoscaling) (*WebServiceResponse, error) {
	deploy := WebServiceDeployResponse{}
	service := WebServiceResponse{}

	services, err := c.GetWebServices(&data.Name)
	if err != nil {
		return nil, err
	}
	if services != nil {
		return nil, fmt.Errorf("the name `%s` is already in use. Please use a different name", data.Name)
	}

	data.Type = "web_service"

	err = c.doRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.HostURL, servicesPath), data, &deploy)
	if err != nil {
		return nil, err
	}

	if autoscaling != nil {
		err = c.doRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s/autoscaling", c.HostURL, servicesPath, deploy.Service.ID), deploy.Service.ServiceDetails.Autoscaling, &autoscaling)
		if err != nil {
			return nil, err
		}
		deploy.Service.ServiceDetails.Autoscaling = autoscaling
	}

	service = deploy.Service

	return &service, nil
}

func (c *Client) UpdateWebService(id string, data WebServiceRequest, autoscaling *Autoscaling) (*WebServiceResponse, error) {
	service := WebServiceResponse{}

	services, err := c.GetWebServices(&data.Name)
	if err != nil {
		return nil, err
	}
	if services != nil && services[0].ID != id {
		return nil, fmt.Errorf("the name `%s` is already in use. Please use a different name", data.Name)
	}

	envVars, err := c.UpdateEnvironmentVariables(id, data.EnvVars)
	if err != nil {
		return nil, err
	}

	service.EnvVars = envVars
	data.Type = "web_service"

	err = c.doRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s", c.HostURL, servicesPath, id), data, &service)
	if err != nil {
		return nil, err
	}

	if data.ServiceDetails.NumInstances != service.ServiceDetails.NumInstances {
		scale := Scale{NumInstances: data.ServiceDetails.NumInstances}
		err = c.doRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/scale", c.HostURL, servicesPath, id), scale, nil)
		if err != nil {
			return nil, err
		}
	}
	if autoscaling != nil {
		err = c.doRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s/autoscaling", c.HostURL, servicesPath, service.ID), service.ServiceDetails.Autoscaling, &autoscaling)
		if err != nil {
			return nil, err
		}
		service.ServiceDetails.Autoscaling = autoscaling
	}

	return &service, nil
}

func (c *Client) DeleteWebService(id string) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.HostURL, servicesPath, id), nil, nil)
}

func (c *Client) WebServicesToSlice(services []WebServices) []WebServiceResponse {
	var result []WebServiceResponse
	for _, service := range services {
		envVars, err := c.GetEnvironmentVariables(service.ID)
		if err != nil {
			return nil
		}
		service.WebServiceResponse.EnvVars = envVars
		result = append(result, service.WebServiceResponse)
	}
	return result
}
