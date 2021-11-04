package statuspage

import (
	"context"
)

// ComponentService handles communication with the page related methods
// of the Statuspage API.
//
// Statuspage API docs: https://developer.statuspage.io/#tag/pages
type ComponentService service

// Component is the Statuspage API component representation
type Component struct {
	ID                 *string    `json:"id,omitempty"`
	PageID             *string    `json:"page_id,omitempty"`
	GroupID            *string    `json:"group_id,omitempty"`
	CreatedAt          *Timestamp `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp `json:"updated_at,omitempty"`
	Group              *bool      `json:"group,omitempty"`
	Name               *string    `json:"name,omitempty"`
	Description        *string    `json:"description,omitempty"`
	Position           *int32     `json:"position,omitempty"`
	Status             *string    `json:"status,omitempty"`
	Showcase           *bool      `json:"showcase,omitempty"`
	OnlyShowIfDegraded *bool      `json:"only_show_if_degraded,omitempty"`
	AutomationEmail    *string    `json:"automation_email,omitempty"`
}

func (c Component) String() string {
	return Stringify(c)
}

// GetComponent returns component information for a given page and component id
func (s *ComponentService) GetComponent(ctx context.Context, pageID string, componentID string) (*Component, error) {
	path := "v1/pages/" + pageID + "/components/" + componentID
	req, err := s.client.newRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	var component Component
	_, err = s.client.do(ctx, req, &component)

	return &component, err
}

// ListComponents returns a list of all components for a given page id
func (s *ComponentService) ListComponents(ctx context.Context, pageID string) (*[]Component, error) {
	path := "v1/pages/" + pageID + "/components"
	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var components []Component
	_, err = s.client.do(ctx, req, &components)

	return &components, err
}

// DeleteComponent deletes a component for a given page and component id
func (s *ComponentService) DeleteComponent(ctx context.Context, pageID string, componentID string) error {
	path := "v1/pages/" + pageID + "/components/" + componentID
	req, err := s.client.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.do(ctx, req, nil)
	return err
}

// UpdateComponentParams are the parameters that can be changed using the update component API endpoint
type UpdateComponentParams struct {
	Description        string    `json:"description,omitempty"`
	Status             string    `json:"status,omitempty"`
	Name               string    `json:"name,omitempty"`
	OnlyShowIfDegraded bool      `json:"only_show_if_degraded,omitempty"`
	GroupID            string    `json:"group_id,omitempty"`
	Showcase           bool      `json:"showcase,omitempty"`
	StartDate          Timestamp `json:"start_date,omitempty"`
}

// UpdateComponentRequestBody is the update component request body representation
type UpdateComponentRequestBody struct {
	Component UpdateComponentParams `json:"component"`
}

// UpdateComponent updates a component for a given page and component id
func (s *ComponentService) UpdateComponent(ctx context.Context, pageID string, componentID string, component UpdateComponentParams) (*Component, error) {
	path := "v1/pages/" + pageID + "/components/" + componentID
	payload := UpdateComponentRequestBody{Component: component}
	req, err := s.client.newRequest("PATCH", path, payload)
	if err != nil {
		return nil, err
	}

	var updatedComponent Component
	_, err = s.client.do(ctx, req, &updatedComponent)

	return &updatedComponent, err
}
