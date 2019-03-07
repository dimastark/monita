package client

import (
	"fmt"
	"net/http"

	"gopkg.in/resty.v1"
)

// MonitaClient represents client to monita API
type MonitaClient struct {
	*resty.Client

	baseURL string
}

// Options represents data for NewClient constructor
type Options struct {
	BaseURL string
	Token   string
}

// NewClient returns new Monita Client
func NewClient(options Options) *MonitaClient {
	r := resty.New()

	if options.Token != "" {
		r.SetAuthToken(options.Token)
	}

	return &MonitaClient{r, options.BaseURL}
}

// SetAdminSecret sets secret for admin handlers
func (mc *MonitaClient) SetAdminSecret(secret string) *MonitaClient {
	mc.SetCookie(&http.Cookie{
		Name:  "ADMIN_SECRET",
		Value: secret,
	})

	return mc
}

func (mc *MonitaClient) request(method, path string, body, result interface{}, successCode int) (interface{}, error) {
	url := fmt.Sprintf("%s/%s", mc.baseURL, path)

	prepared := mc.R()

	if body != nil {
		prepared.SetBody(body)
	}

	if result != nil {
		prepared.SetResult(result)
	}

	res, err := prepared.Execute(method, url)

	if err != nil {
		return nil, err
	}

	if statusCode := res.StatusCode(); statusCode != successCode {
		return nil, fmt.Errorf(res.String())
	}

	return res.Result(), nil
}
