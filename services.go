package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const servicesPath = "services"

type Service struct {
	AutoDeploy     string                `json:"autoDeploy,omitempty"`
	Branch         string                `json:"branch,omitempty"`
	BuildFilter    *BuildFilter          `json:"buildFilter,omitempty"`
	CreateAt       string                `json:"createdAt,omitempty"`
	EnvVars        []EnvironmentVariable `json:"envVars,omitempty"`
	ID             string                `json:"id,omitempty"`
	Image          *Image                `json:"image,omitempty"`
	ImagePath      string                `json:"imagePath,omitempty"`
	Name           string                `json:"name,omitempty"`
	NotifyOnFail   string                `json:"notifyOnFail,omitempty"`
	OwnerID        string                `json:"ownerId,omitempty"`
	Repo           string                `json:"repo,omitempty"`
	RootDir        string                `json:"rootDir,omitempty"`
	SecretFiles    []SecretFiles         `json:"secretFiles,omitempty"`
	ServiceDetails interface{}           `json:"serviceDetails,omitempty"`
	Slug           string                `json:"slug,omitempty"`
	Suspended      string                `json:"suspended,omitempty"`
	Suspenders     []string              `json:"suspenders,omitempty"`
	Type           string                `json:"type,omitempty"`
	UpdatedAt      string                `json:"updatedAt,omitempty"`
}

type StaticSiteDetails struct {
	BuildCommand               string        `json:"buildCommand"`
	ParentServer               *ParentServer `json:"parentServer"`
	PablicPath                 string        `json:"publicPath"`
	PullRequestPreviewsEnabled string        `json:"pullRequestPreviewsEnabled"`
	URL                        string        `json:"url"`
	Headers                    []Header      `json:"headers,omitempty"`
	Routes                     []Route       `json:"routes,omitempty"`
}

type WebServiceDetails struct {
	Autoscaling                *Autoscaling        `json:"autoscaling,omitempty"`
	Disk                       *Disk               `json:"disk,omitempty"`
	Env                        string              `json:"env,omitempty"`
	EnvSpecificDetails         *EnvSpecificDetails `json:"envSpecificDetails,omitempty"`
	HealthCheckPath            string              `json:"healthCheckPath,omitempty"`
	NumInstances               int64               `json:"numInstances,omitempty"`
	OpenPorts                  []OpenPort          `json:"openPorts,omitempty"`
	ParentServer               *ParentServer       `json:"parentServer,omitempty"`
	Plan                       string              `json:"plan,omitempty"`
	PullRequestPreviewsEnabled string              `json:"pullRequestPreviewsEnabled,omitempty"`
	Region                     string              `json:"region,omitempty"`
	URL                        string              `json:"url,omitempty"`
}

type PrivateServiceDetails struct {
	Autoscaling                *Autoscaling        `json:"autoscaling"`
	Disk                       *Disk               `json:"disk"`
	Env                        string              `json:"env"`
	EnvSpecificDetails         *EnvSpecificDetails `json:"envSpecificDetails"`
	NumInstances               int64               `json:"numInstances"`
	OpenPorts                  []OpenPort          `json:"openPorts"`
	ParentServer               *ParentServer       `json:"parentServer"`
	Plan                       string              `json:"plan"`
	PullRequestPreviewsEnabled string              `json:"pullRequestPreviewsEnabled"`
	Region                     string              `json:"region"`
	Url                        string              `json:"url"`
}

type BackgroundWorkerDetails struct {
	Autoscaling                *Autoscaling        `json:"autoscaling"`
	Disk                       *Disk               `json:"disk"`
	Env                        string              `json:"env"`
	EnvSpecificDetails         *EnvSpecificDetails `json:"envSpecificDetails"`
	NumInstances               int64               `json:"numInstances"`
	ParentServer               *ParentServer       `json:"parentServer"`
	Plan                       string              `json:"plan"`
	PullRequestPreviewsEnabled string              `json:"pullRequestPreviewsEnabled"`
	Region                     string              `json:"region"`
}

type CronJobDetails struct {
	Env                 string              `json:"env"`
	EnvSpecificDetails  *EnvSpecificDetails `json:"envSpecificDetails"`
	Plan                string              `json:"plan"`
	Region              string              `json:"region"`
	Schedule            string              `json:"schedule"`
	LastSuccessfulRunAt string              `json:"lastSuccessfulRunAt"`
}

type BuildFilter struct {
	Paths        []string `json:"paths,omitempty"`
	IgnoredPaths []string `json:"ignoredPaths,omitempty"`
}

type Image struct {
	OwnerId              string `json:"ownerId,omitempty"`
	RegistryCredentialId string `json:"registryCredentialId,omitempty"`
	ImagePath            string `json:"imagePath,omitempty"`
}

type Disk struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	SizeGB    int64  `json:"sizeGB,omitempty"`
}

type Autoscaling struct {
	Enabled  bool                `json:"enabled"`
	Min      int64               `json:"min"`
	Max      int64               `json:"max"`
	Criteria AutoscalingCriteria `json:"criteria"`
}

type AutoscalingCriteria struct {
	CPU    AutoscalingCriteriaObject `json:"cpu,omitempty"`
	Memory AutoscalingCriteriaObject `json:"memory,omitempty"`
}

type AutoscalingCriteriaObject struct {
	Enabled    bool  `json:"enabled,omitempty"`
	Percentage int64 `json:"percentage,omitempty"`
}

type EnvSpecificDetails struct {
	DockerCommand        string              `json:"dockerCommand"`
	DockerContext        string              `json:"dockerContext"`
	DockerfilePath       string              `json:"dockerfilePath"`
	PreDeployCommand     string              `json:"preDeployCommand"`
	RegistryCredential   *RegistryCredential `json:"registryCredential"`
	BuildCommand         string              `json:"buildCommand"`
	StartCommand         string              `json:"startCommand"`
	RegistryCredentialId string              `json:"registryCredentialId,omitempty"`
}

type OpenPort struct {
	Port     int64  `json:"port"`
	Protocol string `json:"protocol"`
}

type ParentServer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Services struct {
	Service `json:"service"`
}

type Header struct {
	Path  string `json:"path,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Route struct {
	Type        string `json:"type,omitempty"`
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
}

type SecretFiles struct {
	Name     string `json:"name,omitempty"`
	Contents string `json:"contents,omitempty"`
}

type DiskData struct {
	Name      string `json:"name,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	SizeGB    int64  `json:"sizeGB,omitempty"`
}

type GetServicesArgs struct {
	Name string
	Type string
}

type Scale struct {
	NumInstances int64 `json:"numInstances,omitempty"`
}

func (c *Client) GetServices(args *GetServicesArgs) ([]Service, error) {
	var services []Services
	parameters := url.Values{}
	url, err := url.Parse(fmt.Sprintf("%s/%s", c.HostURL, servicesPath))
	if err != nil {
		return nil, err
	}
	if args != nil {
		if args.Name != "" {
			parameters.Add("name", args.Name)
		}
		if args.Type != "" {
			parameters.Add("type", args.Type)
		}
	}
	url.RawQuery = parameters.Encode()

	err = c.doRequest(http.MethodGet, url.String(), nil, &services)
	if err != nil {
		return nil, err
	}

	return c.ServicesToSlice(services), nil
}

func (c *Client) GetService(id string) (*Service, error) {
	service := Service{}
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

func (c *Client) CreateService(data Service) (*Service, error) {
	service := Service{}

	services, err := c.GetServices(&GetServicesArgs{Name: data.Name})
	if err != nil {
		return nil, err
	}
	if services != nil {
		return nil, fmt.Errorf("the name `%s` is already in use. Please use a different name", data.Name)
	}

	err = c.doRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.HostURL, servicesPath), data, &service)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (c *Client) UpdateService(id string, data Service) (*Service, error) {
	service := Service{}

	services, err := c.GetServices(&GetServicesArgs{Name: data.Name})
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

	err = c.doRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s", c.HostURL, servicesPath, id), data, &service)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (c *Client) DeleteService(id string) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.HostURL, servicesPath, id), nil, nil)
}

func (c *Client) ServicesToSlice(services []Services) []Service {
	var result []Service
	for _, service := range services {
		envVars, err := c.GetEnvironmentVariables(service.ID)
		if err != nil {
			return nil
		}
		service.Service.EnvVars = envVars
		result = append(result, service.Service)
	}
	return result
}
