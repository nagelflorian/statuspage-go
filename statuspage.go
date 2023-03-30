package statuspage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const version = "1.0.0"
const hostURL = "api.statuspage.io"

// A Client manages communication with the Statuspage API.
type Client struct {
	httpClient *http.Client

	// Base URL for API requests. Defaults to the public Statuspage API.
	BaseURL   *url.URL
	UserAgent string
	Token     string
	Version   string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Statuspage API.
	Page        *PageService
	Component   *ComponentService
	Incidents   *IncidentService
	Maintenance *MaintenanceService
}

type service struct {
	client *Client
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "OAuth "+c.Token)

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 400 {
		if v == nil {
			return resp, nil
		}

		err = json.NewDecoder(resp.Body).Decode(v)
		return resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	return nil, fmt.Errorf("response %s: %d – %s", resp.Status, resp.StatusCode, string(body))
}

// NewClient returns a new Statuspage API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(token string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL := &url.URL{Host: hostURL, Scheme: "https"}

	c := &Client{
		BaseURL:    baseURL,
		httpClient: http.DefaultClient,
		Token:      token,
		Version:    version,
	}

	c.common.client = c
	c.Page = (*PageService)(&c.common)
	c.Component = (*ComponentService)(&c.common)
	c.Incidents = (*IncidentService)(&c.common)
	c.Maintenance = (*MaintenanceService)(&c.common)

	return c
}
