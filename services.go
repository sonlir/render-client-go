package render

import (
	"fmt"
	"net/http"
	"time"
)

const servicesPath = "services"

type Service struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"`
	Repo           string         `json:"repo"`
	Name           string         `json:"name"`
	AutoDeploy     string         `json:"autoDeploy"`
	Brach          string         `json:"branch"`
	BuildFilter    BuildFilter    `json:"buildFilter"`
	CreateAt       time.Time      `json:"createdAt"`
	NotifyOnFail   string         `json:"notifyOnFail"`
	Image          Image          `json:"image"`
	OwnerId        string         `json:"ownerId"`
	Slug           string         `json:"slug"`
	RootDir        string         `json:"rootDir"`
	Suspended      string         `json:"suspended"`
	Suspenders     []string       `json:"suspenders"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	ServiceDetails ServiceDetails `json:"serviceDetails"`
	ImagePath      string         `json:"imagePath"`
}

type ServiceDetails struct {
	Autoscaling                Autoscaling        `json:"autoscaling"`
	BuildPlan                  string             `json:"buildPlan"`
	BuildCommand               string             `json:"buildCommand"`
	DockerCommand              string             `json:"dockerCommand"`
	DockerContext              string             `json:"dockerContext"`
	DockerfilePath             string             `json:"dockerfilePath"`
	Env                        string             `json:"env"`
	EnvSpecificDetails         EnvSpecificDetails `json:"envSpecificDetails"`
	HealthCheckPath            string             `json:"healthCheckPath"`
	PublishPath                string             `json:"publishPath"`
	NumInstances               int                `json:"numInstances"`
	OpenPorts                  []OpenPort         `json:"openPorts"`
	ParentServer               ParentServer       `json:"parentServer"`
	Plan                       string             `json:"plan"`
	PullRequestPreviewsEnabled string             `json:"pullRequestPreviewsEnabled"`
	Region                     string             `json:"region"`
	Url                        string             `json:"url"`
	Disk                       Disk               `json:"disk"`
	Schedule                   string             `json:"schedule"`
	LastSuccessfulRunAt        string             `json:"lastSuccessfulRunAt"`
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
	Type           string                `json:"type,omitempty"`
	Name           string                `json:"name,omitempty"`
	OwnerID        string                `json:"ownerId,omitempty"`
	Repo           string                `json:"repo,omitempty"`
	AutoDeploy     string                `json:"autoDeploy,omitempty"`
	Branch         string                `json:"branch,omitempty"`
	Image          Image                 `json:"image,omitempty"`
	BuildFilter    BuildFilter           `json:"buildFilter,omitempty"`
	RootDir        string                `json:"rootDir,omitempty"`
	EnvVars        []EnvironmentVariable `json:"envVars,omitempty"`
	SecretFiles    []SecretFiles         `json:"secretFiles,omitempty"`
	ServiceDetails ServiceDetailsData    `json:"serviceDetails,omitempty"`
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

func (c *Client) GetServices() (*[]Services, error) {
	services := []Services{}
	err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.HostURL, servicesPath), nil, &services)
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
