package statuspage

import (
	"context"
	"time"
)

type IncidentService service

type Incident struct {
	Id         string `json:"id"`
	Components []struct {
		Id                 string    `json:"id"`
		PageId             string    `json:"page_id"`
		GroupId            string    `json:"group_id"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		Group              bool      `json:"group"`
		Name               string    `json:"name"`
		Description        string    `json:"description"`
		Position           int       `json:"position"`
		Status             string    `json:"status"`
		Showcase           bool      `json:"showcase"`
		OnlyShowIfDegraded bool      `json:"only_show_if_degraded"`
		AutomationEmail    string    `json:"automation_email"`
		StartDate          string    `json:"start_date"`
	} `json:"components"`
	CreatedAt       time.Time `json:"created_at"`
	Impact          string    `json:"impact"`
	ImpactOverride  string    `json:"impact_override"`
	IncidentUpdates []struct {
		Id                 string `json:"id"`
		IncidentId         string `json:"incident_id"`
		AffectedComponents []struct {
			Code      string `json:"code"`
			Name      string `json:"name"`
			OldStatus string `json:"old_status"`
			NewStatus string `json:"new_status"`
		} `json:"affected_components"`
		Body                 string    `json:"body"`
		CreatedAt            time.Time `json:"created_at"`
		CustomTweet          string    `json:"custom_tweet"`
		DeliverNotifications bool      `json:"deliver_notifications"`
		DisplayAt            time.Time `json:"display_at"`
		Status               string    `json:"status"`
		TweetId              string    `json:"tweet_id"`
		TwitterUpdatedAt     time.Time `json:"twitter_updated_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		WantsTwitterUpdate   bool      `json:"wants_twitter_update"`
	} `json:"incident_updates"`
	Metadata struct {
		Jira struct {
			IssueId string `json:"issue_id"`
		} `json:"jira"`
	} `json:"metadata"`
	MonitoringAt                              time.Time `json:"monitoring_at"`
	Name                                      string    `json:"name"`
	PageId                                    string    `json:"page_id"`
	PostmortemBody                            string    `json:"postmortem_body"`
	PostmortemBodyLastUpdatedAt               time.Time `json:"postmortem_body_last_updated_at"`
	PostmortemIgnored                         bool      `json:"postmortem_ignored"`
	PostmortemNotifiedSubscribers             bool      `json:"postmortem_notified_subscribers"`
	PostmortemNotifiedTwitter                 bool      `json:"postmortem_notified_twitter"`
	PostmortemPublishedAt                     bool      `json:"postmortem_published_at"`
	ResolvedAt                                time.Time `json:"resolved_at"`
	ScheduledAutoCompleted                    bool      `json:"scheduled_auto_completed"`
	ScheduledAutoInProgress                   bool      `json:"scheduled_auto_in_progress"`
	ScheduledFor                              time.Time `json:"scheduled_for"`
	AutoTransitionDeliverNotificationsAtEnd   bool      `json:"auto_transition_deliver_notifications_at_end"`
	AutoTransitionDeliverNotificationsAtStart bool      `json:"auto_transition_deliver_notifications_at_start"`
	AutoTransitionToMaintenanceState          bool      `json:"auto_transition_to_maintenance_state"`
	AutoTransitionToOperationalState          bool      `json:"auto_transition_to_operational_state"`
	ScheduledRemindPrior                      bool      `json:"scheduled_remind_prior"`
	ScheduledRemindedAt                       time.Time `json:"scheduled_reminded_at"`
	ScheduledUntil                            time.Time `json:"scheduled_until"`
	Shortlink                                 string    `json:"shortlink"`
	Status                                    string    `json:"status"`
	UpdatedAt                                 time.Time `json:"updated_at"`
}

func (s *IncidentService) ListScheduledIncidents(ctx context.Context, pageID string) (*[]Incident, error) {
	path := "v1/pages/" + pageID + "/incidents/scheduled"
	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var incidents []Incident
	_, err = s.client.do(ctx, req, &incidents)

	return &incidents, err
}

func (s *IncidentService) ListUnresolvedIncidents(ctx context.Context, pageID string) (*[]Incident, error) {
	path := "v1/pages/" + pageID + "/incidents/unresolved"
	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var incidents []Incident
	_, err = s.client.do(ctx, req, &incidents)

	return &incidents, err
}

// CreateIncidentRequestBody is the create incident request body representation
type CreateIncidentRequestBody struct {
	Incident CreateIncidentParams `json:"incident"`
}

type CreateIncidentParams struct {
	Name           string `json:"name,omitempty"`
	Status         string `json:"status,omitempty"`
	ImpactOverride string `json:"impact_override,omitempty"`
	//ScheduledFor                              time.Time         `json:"scheduled_for,omitempty"`
	//ScheduledUntil                            time.Time         `json:"scheduled_until,omitempty"`
	ScheduledRemindPrior                      bool              `json:"scheduled_remind_prior,omitempty"`
	AutoTransitionToMaintenanceState          bool              `json:"auto_transition_to_maintenance_state,omitempty"`
	AutoTransitionToOperationalState          bool              `json:"auto_transition_to_operational_state,omitempty"`
	ScheduledAutoInProgress                   bool              `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledAutoCompleted                    bool              `json:"scheduled_auto_completed,omitempty"`
	AutoTransitionDeliverNotificationsAtStart bool              `json:"auto_transition_deliver_notifications_at_start,omitempty"`
	AutoTransitionDeliverNotificationsAtEnd   bool              `json:"auto_transition_deliver_notifications_at_end,omitempty"`
	DeliverNotifications                      bool              `json:"deliver_notifications,omitempty"`
	AutoTweetAtBeginning                      bool              `json:"auto_tweet_at_beginning,omitempty"`
	AutoTweetOnCompletion                     bool              `json:"auto_tweet_on_completion,omitempty"`
	AutoTweetOnCreation                       bool              `json:"auto_tweet_on_creation,omitempty"`
	AutoTweetOneHourBefore                    bool              `json:"auto_tweet_one_hour_before,omitempty"`
	BackfillDate                              string            `json:"backfill_date,omitempty"`
	Backfilled                                bool              `json:"backfilled,omitempty"`
	Body                                      string            `json:"body,omitempty"`
	Components                                map[string]string `json:"components,omitempty"`
	ComponentIds                              []string          `json:"component_ids,omitempty"`
	ScheduledAutoTransition                   bool              `json:"scheduled_auto_transition,omitempty"`
}

// CreateIncident creates an incident
func (s *IncidentService) CreateIncident(ctx context.Context, pageID string, incident CreateIncidentParams) (*Incident, error) {
	path := "v1/pages/" + pageID + "/incidents"
	payload := CreateIncidentRequestBody{Incident: incident}
	req, err := s.client.newRequest("POST", path, payload)
	if err != nil {
		return nil, err
	}

	var newIncident Incident
	_, err = s.client.do(ctx, req, &newIncident)

	return &newIncident, err
}

func (s *IncidentService) DeleteIncident(ctx context.Context, pageID string, incidentID string) error {
	path := "v1/pages/" + pageID + "/incidents/" + incidentID
	req, err := s.client.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	_, err = s.client.do(ctx, req, nil)
	return err
}

// UpdateIncidentRequestBody is the update incident request body representation
type UpdateIncidentRequestBody struct {
	Incident UpdateIncidentParams `json:"incident"`
}

type UpdateIncidentParams struct {
	Status       string   `json:"status,omitempty"`
	Body         string   `json:"body,omitempty"`
	ComponentIds []string `json:"component_ids,omitempty"`
}

// UpdateIncident updates an incident
func (s *IncidentService) UpdateIncident(ctx context.Context, pageID string, incident UpdateIncidentParams, incidentID string) (*Incident, error) {
	path := "v1/pages/" + pageID + "/incidents/" + incidentID
	payload := UpdateIncidentRequestBody{Incident: incident}
	req, err := s.client.newRequest("PATCH", path, payload)
	if err != nil {
		return nil, err
	}

	var newIncident Incident
	_, err = s.client.do(ctx, req, &newIncident)

	return &newIncident, err
}
