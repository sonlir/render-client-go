package render

import (
	"fmt"
	"net/http"
	"time"
)

const servicesPath = "services"

type Service struct {
	AutoDeploy     string         `json:"autoDeploy"`
	Branch         string         `json:"branch"`
	BuildFilter    BuildFilter    `json:"buildFilter"`
	CreateAt       time.Time      `json:"createdAt"`
	ID             string         `json:"id"`
	Image          Image          `json:"image"`
	ImagePath      string         `json:"imagePath"`
	Name           string         `json:"name"`
	NotifyOnFail   string         `json:"notifyOnFail"`
	OwnerId        string         `json:"ownerId"`
	Repo           string         `json:"repo"`
	RootDir        string         `json:"rootDir"`
	ServiceDetails ServiceDetails `json:"serviceDetails"`
	Slug           string         `json:"slug"`
	Suspended      string         `json:"suspended"`
	Suspenders     []string       `json:"suspenders"`
	Type           string         `json:"type"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

type ServiceDetails struct {
	Autoscaling                Autoscaling        `json:"autoscaling"`
	BuildCommand               string             `json:"buildCommand"`
	BuildPlan                  string             `json:"buildPlan"`
	Disk                       Disk               `json:"disk"`
	DockerCommand              string             `json:"dockerCommand"`
	DockerContext              string             `json:"dockerContext"`
	DockerfilePath             string             `json:"dockerfilePath"`
	Env                        string             `json:"env"`
	EnvSpecificDetails         EnvSpecificDetails `json:"envSpecificDetails"`
	HealthCheckPath            string             `json:"healthCheckPath"`
	LastSuccessfulRunAt        string             `json:"lastSuccessfulRunAt"`
	NumInstances               int                `json:"numInstances"`
	OpenPorts                  []OpenPort         `json:"openPorts"`
	ParentServer               ParentServer       `json:"parentServer"`
	Plan                       string             `json:"plan"`
	PublishPath                string             `json:"publishPath"`
	PullRequestPreviewsEnabled string             `json:"pullRequestPreviewsEnabled"`
	Region                     string             `json:"region"`
	Schedule                   string             `json:"schedule"`
	Url                        string             `json:"url"`
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
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Autoscaling struct {
	Enabled  bool                `json:"enabled"`
	Min      int                 `json:"min"`
	Max      int                 `json:"max"`
	Criteria AutoscalingCriteria `json:"criteria"`
}

type AutoscalingCriteria struct {
	CPU    AutoscalingCriteriaObject `json:"cpu,omitempty"`
	Memory AutoscalingCriteriaObject `json:"memory,omitempty"`
}

type AutoscalingCriteriaObject struct {
	Enabled    bool `json:"enabled,omitempty"`
	Percentage int  `json:"percentage,omitempty"`
}

type EnvSpecificDetails struct {
	DockerCommand      string             `json:"dockerCommand"`
	DockerContext      string             `json:"dockerContext"`
	DockerfilePath     string             `json:"dockerfilePath"`
	PreDeployCommand   string             `json:"preDeployCommand"`
	RegistryCredential RegistryCredential `json:"registryCredential"`
	BuildCommand       string             `json:"buildCommand"`
	StartCommand       string             `json:"startCommand"`
}

type OpenPort struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type ParentServer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Services struct {
	Service Service `json:"service"`
}

type ServiceData struct {
	AutoDeploy     string                `json:"autoDeploy,omitempty"`
	Branch         string                `json:"branch,omitempty"`
	BuildFilter    BuildFilter           `json:"buildFilter,omitempty"`
	EnvVars        []EnvironmentVariable `json:"envVars,omitempty"`
	Image          Image                 `json:"image,omitempty"`
	Name           string                `json:"name,omitempty"`
	OwnerID        string                `json:"ownerId,omitempty"`
	Repo           string                `json:"repo,omitempty"`
	RootDir        string                `json:"rootDir,omitempty"`
	SecretFiles    []SecretFiles         `json:"secretFiles,omitempty"`
	ServiceDetails ServiceDetailsData    `json:"serviceDetails,omitempty"`
	Type           string                `json:"type,omitempty"`
}

type ServiceDetailsData struct {
	BuildCommand               string                 `json:"buildCommand,omitempty"`
	Headers                    []Header               `json:"headers,omitempty"`
	PublishPath                string                 `json:"publishPath,omitempty"`
	PullRequestPreviewsEnabled string                 `json:"pullRequestPreviewsEnabled,omitempty"`
	Routes                     []Route                `json:"routes,omitempty"`
	Disk                       DiskData               `json:"disk,omitempty"`
	Env                        string                 `json:"env,omitempty"`
	EnvSpecificDetails         EnvSpecificDetailsData `json:"envSpecificDetails,omitempty"`
	HealthCheckPath            string                 `json:"healthCheckPath,omitempty"`
	NumInstances               int                    `json:"numInstances,omitempty"`
	Plan                       string                 `json:"plan,omitempty"`
	Region                     string                 `json:"region,omitempty"`
	Schedule                   string                 `json:"schedule,omitempty"`
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
	SizeGB    int    `json:"sizeGB,omitempty"`
}

type EnvSpecificDetailsData struct {
	DockerCommand        string `json:"dockerCommand,omitempty"`
	DockerContext        string `json:"dockerContext,omitempty"`
	DockerfilePath       string `json:"dockerfilePath,omitempty"`
	PreDeployCommand     string `json:"preDeployCommand,omitempty"`
	RegistryCredentialId string `json:"registryCredentialId,omitempty"`
	BuildCommand         string `json:"buildCommand,omitempty"`
	StartCommand         string `json:"startCommand,omitempty"`
}

func (c *Client) GetServices(serviceType string) (*[]Services, error) {
	services := []Services{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s?type=%s", c.HostURL, servicesPath, serviceType), nil, &services)
	if err != nil {
		return nil, err
	}

	return &services, nil
}

func (c *Client) GetService(id string) (*Service, error) {
	service := Service{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.HostURL, servicesPath, id), nil, &service)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (c *Client) CreateService(data ServiceData) (*Service, error) {
	service := Service{}
	err := c.doRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.HostURL, servicesPath), data, &service)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (c *Client) UpdateService(id string, data ServiceData) (*Service, error) {
	service := Service{}
	err := c.doRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s", c.HostURL, servicesPath, id), data, &service)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (c *Client) DeleteService(id string) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.HostURL, servicesPath, id), nil, nil)
}
