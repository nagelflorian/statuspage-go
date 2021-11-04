package statuspage

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestComponent_marshall(t *testing.T) {
	testJSONMarshal(t, &Component{}, "{}")

	u := &Component{
		ID:                 String("a"),
		PageID:             String("b"),
		GroupID:            String("c"),
		CreatedAt:          &Timestamp{referenceTime},
		UpdatedAt:          &Timestamp{referenceTime},
		Group:              Bool(true),
		Name:               String("d"),
		Description:        String("e"),
		Position:           Int32(1),
		Status:             String("g"),
		Showcase:           Bool(false),
		OnlyShowIfDegraded: Bool(true),
		AutomationEmail:    String("h"),
	}
	want := `{
		"id": "a",
		"page_id":"b",
		"group_id":"c",
		"created_at": "2006-01-02T15:04:05Z",
		"updated_at": "2006-01-02T15:04:05Z",
		"group":true,
		"name":"d",
		"description":"e",
		"position":1,
		"status":"g",
		"showcase":false,
		"only_show_if_degraded":true,
		"automation_email":"h"
	}`
	testJSONMarshal(t, u, want)
}

func TestComponentService_GetComponent(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/pages/1/components/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":"2"}`)
	})

	component, err := client.Component.GetComponent(context.Background(), "1", "2")
	if err != nil {
		t.Errorf("ComponentService.GetComponent returned error: %v", err)
	}

	want := &Component{ID: String("2")}
	if !reflect.DeepEqual(component, want) {
		t.Errorf("ComponentService.GetComponent returned %+v, want %+v", component, want)
	}
}

func TestComponentService_ListComponent(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/pages/1/components", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":"1"}, {"id":"2"}]`)
	})

	components, err := client.Component.ListComponents(context.Background(), "1")
	if err != nil {
		t.Errorf("ComponentService.ListComponents returned error: %v", err)
	}

	want := &[]Component{
		{ID: String("1")},
		{ID: String("2")},
	}
	if !reflect.DeepEqual(components, want) {
		t.Errorf("ComponentService.ListComponents returned %+v, want %+v", components, want)
	}
}

func TestComponentService_DeleteComponent(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/pages/1/components/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, `{}`)
	})

	err := client.Component.DeleteComponent(context.Background(), "1", "2")
	if err != nil {
		t.Errorf("ComponentService.DeleteComponent returned error: %v", err)
	}
}

func TestComponentService_UpdateComponent(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/pages/1/components/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		fmt.Fprint(w, `{"id":"2", "status": "major_outage"}`)
	})

	componentParams := UpdateComponentParams{
		Status: "major_outage",
	}
	updatedComponent, err := client.Component.UpdateComponent(context.Background(), "1", "2", componentParams)
	if err != nil {
		t.Errorf("ComponentService.UpdateComponent returned error: %v", err)
	}

	want := &Component{ID: String("2"), Status: String("major_outage")}
	if !reflect.DeepEqual(updatedComponent, want) {
		t.Errorf("ComponentService.UpdateComponent returned %+v, want %+v", updatedComponent, want)
	}
}
