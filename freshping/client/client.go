package freshping_client

import (
	"fmt"

	req "github.com/imroc/req/v3"
)

type Client struct {
	c *req.Client
}

type CheckListResponse struct {
	Results []Check `json:"results"`
}

type Check struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	PerformanceStatus  string `json:"performance_status"`
	AlertNote          string `json:"alert_note"`
	Location           string `json:"location"`
	AlertUsers         []int  `json:"alert_users"`
	AlertContacts      []int  `json:"alert_contacts"`
	MonitoringInterval int    `json:"monitoring_interval"`
	URL                string `json:"url"`
	RequestTimeout     int    `json:"request_timeout"`
	BasicAuthUsername  string `json:"basic_auth_username"`
	BasicAuthPassword  string `json:"basic_auth_password"`
	CommandString      string `json:"command_string"`
	SuccessString      string `json:"success_string"`
	ErrorString        string `json:"error_string"`
}

type UserListResponse struct {
	Users    []User    `json:"users"`
	Contacts []Contact `json:"contacts"`
}

type User struct {
	ID                        int    `json:"id"`
	Name                      string `json:"name"`
	Email                     string `json:"email"`
	Role                      string `json:"role"`
	DisableWeeklyReportEmails bool   `json:"disable_weekly_report_emails"`
	DisableAlertEmails        bool   `json:"disable_alert_emails"`
}

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ClientOption func(c *Client)

func New(options ...ClientOption) *Client {
	c := &Client{
		c: req.C(),
	}
	c.c.SetBaseURL("https://api.freshping.io")

	for _, o := range options {
		o(c)
	}

	return c
}

func (c *Client) GetChecks() ([]Check, error) {
	res, err := c.r().Get("/api/v1/checks/")

	if err != nil {
		return nil, fmt.Errorf("error retrieving checks: %q", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("error retrieving checks: %q", res.Status)
	}
	var response CheckListResponse

	err = res.UnmarshalJson(&response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
	}

	return response.Results, nil
}

func (c *Client) GetCheck(id int64) (*Check, error) {
	res, err := c.r().Get(fmt.Sprintf("/api/v1/checks/%d", id))

	if err != nil {
		return nil, fmt.Errorf("error retrieving check: %q", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("error retrieving check: %q", res.Status)
	}
	var check Check

	err = res.UnmarshalJson(&check)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
	}

	return &check, nil
}

func (c *Client) GetUsers() (*UserListResponse, error) {
	res, err := c.r().Get("/api/v1/users/")

	if err != nil {
		return nil, fmt.Errorf("error retrieving users: %q", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("error retrieving users: %q", res.Status)
	}
	var response UserListResponse

	err = res.UnmarshalJson(&response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
	}

	return &response, nil
}

func (c *Client) r() *req.Request {
	return c.c.R()
}

func (c *Client) setAuth(apiKey string, subdomain string) *Client {
	c.c.SetCommonBasicAuth(apiKey, subdomain)
	return c
}

func WithAuth(apiKey string, subdomain string) ClientOption {
	return func(c *Client) {
		c.setAuth(apiKey, subdomain)
	}
}
