package statuspage

import (
	"context"
	"time"
)

type MaintenanceService service

func (s *MaintenanceService) ListActiveMaintenance(ctx context.Context, pageID string) (*[]Incident, error) {
	path := "v1/pages/" + pageID + "/incidents/active_maintenance"
	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var maintenance []Incident
	_, err = s.client.do(ctx, req, &maintenance)

	return &maintenance, err
}

// CreateMaintenanceRequestBody is the create maintenance request body representation
type CreateMaintenanceRequestBody struct {
	Incident CreateMaintenanceParams `json:"incident"`
}

type CreateMaintenanceParams struct {
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	//ImpactOverride                            string            `json:"impact_override,omitempty"`
	ScheduledFor                              time.Time         `json:"scheduled_for,omitempty"`
	ScheduledUntil                            time.Time         `json:"scheduled_until,omitempty"`
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

// CreateMaintenance creates a maintenance event
func (s *MaintenanceService) CreateMaintenance(ctx context.Context, pageID string, incident CreateMaintenanceParams) (*Incident, error) {
	path := "v1/pages/" + pageID + "/incidents"
	payload := CreateMaintenanceRequestBody{Incident: incident}
	req, err := s.client.newRequest("POST", path, payload)
	if err != nil {
		return nil, err
	}

	var newIncident Incident
	_, err = s.client.do(ctx, req, &newIncident)

	return &newIncident, err
}
