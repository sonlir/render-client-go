package render

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
	ServiceDetails ServiceDetails        `json:"serviceDetails,omitempty"`
	Slug           string                `json:"slug,omitempty"`
	Suspended      string                `json:"suspended,omitempty"`
	Suspenders     []string              `json:"suspenders,omitempty"`
	Type           string                `json:"type,omitempty"`
	UpdatedAt      string                `json:"updatedAt,omitempty"`
}

type ServiceDetails struct {
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
	Schedule                   string              `json:"schedule"`
	LastSuccessfulRunAt        string              `json:"lastSuccessfulRunAt"`
	BuildCommand               string              `json:"buildCommand"`
	PublicPath                 string              `json:"publicPath"`
	Headers                    []Header            `json:"headers,omitempty"`
	Routes                     []Route             `json:"routes,omitempty"`
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
	Enabled  bool                 `json:"enabled"`
	Min      int64                `json:"min,omitempty"`
	Max      int64                `json:"max,omitempty"`
	Criteria *AutoscalingCriteria `json:"criteria,omitempty"`
}

type AutoscalingCriteria struct {
	CPU    *AutoscalingCriteriaObject `json:"cpu,omitempty"`
	Memory *AutoscalingCriteriaObject `json:"memory,omitempty"`
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

	if data.Type == "web_service" || data.Type == "private_service" || data.Type == "background_worker" {
		if data.ServiceDetails.Autoscaling != nil {
			err = c.doRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s/autoscaling", c.HostURL, servicesPath, service.ID), data.ServiceDetails.Autoscaling, nil)
			if err != nil {
				return nil, err
			}
		}
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

	if data.Type == "web_service" || data.Type == "private_service" || data.Type == "background_worker" {
		if data.ServiceDetails.NumInstances != service.ServiceDetails.NumInstances {
			scale := Scale{NumInstances: data.ServiceDetails.NumInstances}
			err = c.doRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/scale", c.HostURL, servicesPath, id), scale, nil)
			if err != nil {
				return nil, err
			}
		}
		if data.ServiceDetails.Autoscaling != nil {
			err = c.doRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s/autoscaling", c.HostURL, servicesPath, id), data.ServiceDetails.Autoscaling, nil)
			if err != nil {
				return nil, err
			}
		}
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
