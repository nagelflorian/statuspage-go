package statuspage

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPage_marshall(t *testing.T) {
	testJSONMarshal(t, &Page{}, "{}")

	u := &Page{
		ID:                       String("a"),
		CreatedAt:                &Timestamp{referenceTime},
		UpdatedAt:                &Timestamp{referenceTime},
		Name:                     String("b"),
		PageDescription:          String("c"),
		Headline:                 String("d"),
		Branding:                 String("e"),
		Subdomain:                String("f"),
		Domain:                   String("g"),
		URL:                      String("h"),
		SupportURL:               String("i"),
		HiddenFromSearch:         Bool(true),
		AllowPageSubscribers:     Bool(false),
		AllowIncidentSubscribers: Bool(true),
		AllowEmailSubscribers:    Bool(false),
		AllowSmsSubscribers:      Bool(true),
		AllowRssAtomFeeds:        Bool(false),
		AllowWebhookSubscribers:  Bool(true),
		NotificationsFromEmail:   String("j"),
		NotificationsEmailFooter: String("k"),
		ActivityScore:            Int64(1),
		TwitterUsername:          String("l"),
		ViewersMustBeTeamMembers: Bool(false),
		IPRestrictions:           String("m"),
		City:                     String("n"),
		State:                    String("o"),
		Country:                  String("p"),
		TimeZone:                 String("q"),
		CSSBodyBackgroundColor:   String("r"),
		CSSFontColor:             String("s"),
		CSSLightFontColor:        String("t"),
		CSSGreens:                String("u"),
		CSSYellows:               String("v"),
		CSSOranges:               String("w"),
		CSSReds:                  String("x"),
		CSSBlues:                 String("y"),
		CSSBorderColor:           String("z"),
		CSSGraphColor:            String("aa"),
		CSSLinkColor:             String("ab"),
		FaviconLogo: &PageLogo{
			UpdatedAt: &Timestamp{referenceTime},
			Size:      Int64(2),
			URL:       String("ac"),
		},
		TransactionalLogo: &PageLogo{
			UpdatedAt: &Timestamp{referenceTime},
			Size:      Int64(2),
			URL:       String("ad"),
		},
		HeroCover: &PageLogo{
			UpdatedAt: &Timestamp{referenceTime},
			Size:      Int64(2),
			URL:       String("ae"),
		},
		EmailLogo: &PageLogo{
			UpdatedAt: &Timestamp{referenceTime},
			Size:      Int64(2),
			URL:       String("af"),
		},
		TwitterLogo: &PageLogo{
			UpdatedAt: &Timestamp{referenceTime},
			Size:      Int64(2),
			URL:       String("ag"),
		},
	}

	want := `{
		"id": "a",
		"created_at": "2006-01-02T15:04:05Z",
		"updated_at": "2006-01-02T15:04:05Z",
		"name": "b",
		"page_description": "c",
		"headline": "d",
		"branding": "e",
		"subdomain": "f",
		"domain": "g",
		"url": "h",
		"support_url": "i",
		"hidden_from_search": true,
		"allow_page_subscribers": false,
		"allow_incident_subscribers": true,
		"allow_email_subscribers": false,
		"allow_sms_subscribers": true,
		"allow_rss_atom_feeds": false,
		"allow_webhook_subscribers": true,
		"notifications_from_email": "j",
		"notifications_email_footer": "k",
		"activity_score": 1,
		"twitter_username":"l",
		"viewers_must_be_team_members":false,
		"ip_restrictions":"m",
		"city":"n",
		"state":"o",
		"country":"p",
		"time_zone":"q",
		"css_body_background_color":"r",
		"css_font_color":"s",
		"css_light_font_color":"t",
		"css_greens":"u",
		"css_yellows":"v",
		"css_oranges":"w",
		"css_reds":"x",
		"css_blues":"y",
		"css_border_color":"z",
		"css_graph_color":"aa",
		"css_link_color":"ab",
		"favicon_logo": {
			"updated_at":"2006-01-02T15:04:05Z",
			"size":2,
			"url":"ac"
		},
		"transactional_logo": {
			"updated_at":"2006-01-02T15:04:05Z",
			"size":2,
			"url":"ad"
		},
		"hero_cover": {
			"updated_at":"2006-01-02T15:04:05Z",
			"size":2,
			"url":"ae"
		},
		"email_logo": {
			"updated_at":"2006-01-02T15:04:05Z",
			"size":2,
			"url":"af"
		},
		"twitter_logo": {
			"updated_at":"2006-01-02T15:04:05Z",
			"size":2,
			"url":"ag"
		}
	}`

	testJSONMarshal(t, u, want)
}

func TestPageService_GetPage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/pages/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":"1"}`)
	})

	page, err := client.Page.GetPage(context.Background(), "1")
	if err != nil {
		t.Errorf("PageService.GetPage returned error: %v", err)
	}

	want := &Page{ID: String("1")}
	if !reflect.DeepEqual(page, want) {
		t.Errorf("PageService.GetPage returned %+v, want %+v", page, want)
	}
}

func TestPageService_ListPages(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/pages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":"1"}, {"id":"2"}]`)
	})

	page, err := client.Page.ListPages(context.Background())
	if err != nil {
		t.Errorf("PageService.ListPages returned error: %v", err)
	}

	want := &[]Page{
		{ID: String("1")},
		{ID: String("2")},
	}
	if !reflect.DeepEqual(page, want) {
		t.Errorf("PageService.ListPages returned %+v, want %+v", page, want)
	}
}
