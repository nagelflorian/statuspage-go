package statuspage

import (
	"context"
	"time"
)

// PageService handles communication with the page related methods
// of the Statuspage API.
//
// Statuspage API docs: https://developer.statuspage.io/#tag/pages
type PageService service

// PageLogo is the Statuspage API page logo representation
type PageLogo struct {
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Size      int32      `json:"size,omitempty"`
	URL       string     `json:"url,omitempty"`
}

// Page is the Statuspage API page representation
type Page struct {
	ID                       string     `json:"id"`
	CreatedAt                *time.Time `json:"created_at,omitempty"`
	UpdatedAt                *time.Time `json:"updated_at,omitempty"`
	Name                     string     `json:"name,omitempty"`
	PageDescription          string     `json:"page_description,omitempty"`
	Headline                 string     `json:"headline,omitempty"`
	Branding                 string     `json:"branding,omitempty"`
	Subdomain                string     `json:"subdomain,omitempty"`
	Domain                   string     `json:"domain,omitempty"`
	URL                      string     `json:"url,omitempty"`
	SupportURL               string     `json:"support_url,omitempty"`
	HiddenFromSearch         *bool      `json:"hidden_from_search"`
	AllowPageSubscribers     *bool      `json:"allow_page_subscribers"`
	AllowIncidentSubscribers *bool      `json:"allow_incident_subscribers"`
	AllowEmailSubscribers    *bool      `json:"allow_email_subscribers"`
	AllowSmsSubscribers      *bool      `json:"allow_sms_subscribers"`
	AllowRssAtomFeeds        *bool      `json:"allow_rss_atom_feeds"`
	AllowWebhookSubscribers  *bool      `json:"allow_webhook_subscribers"`
	NotificationsFromEmail   string     `json:"notifications_from_email,omitempty"`
	NotificationsEmailFooter string     `json:"notifications_email_footer,omitempty"`
	ActivityScore            int32      `json:"activity_score,omitempty"`
	TwitterUsername          string     `json:"twitter_username,omitempty"`
	ViewersMustBeTeamMembers *bool      `json:"viewers_must_be_team_members"`
	IPRestrictions           string     `json:"ip_restrictions,omitempty"`
	City                     string     `json:"city,omitempty"`
	State                    string     `json:"state,omitempty"`
	Country                  string     `json:"country,omitempty"`
	TimeZone                 string     `json:"time_zone,omitempty"`
	CSSBodyBackgroundColor   string     `json:"css_body_background_color,omitempty"`
	CSSFontColor             string     `json:"css_font_color,omitempty"`
	CSSLightFontColor        string     `json:"css_light_font_color,omitempty"`
	CSSGreens                string     `json:"css_greens,omitempty"`
	CSSYellows               string     `json:"css_yellows,omitempty"`
	CSSOranges               string     `json:"css_oranges,omitempty"`
	CSSReds                  string     `json:"css_reds,omitempty"`
	CSSBlues                 string     `json:"css_blues,omitempty"`
	CSSBorderColor           string     `json:"css_border_color,omitempty"`
	CSSGraphColor            string     `json:"css_graph_color,omitempty"`
	CSSLinkColor             string     `json:"css_link_color,omitempty"`
	FaviconLogo              PageLogo   `json:"favicon_logo,omitempty"`
	TransactionalLogo        PageLogo   `json:"transactional_logo,omitempty"`
	HeroCover                PageLogo   `json:"hero_cover,omitempty"`
	EmailLogo                PageLogo   `json:"email_logo,omitempty"`
	TwitterLogo              PageLogo   `json:"twitter_logo,omitempty"`
}

// ListPages returns a list of all pages
func (s *PageService) ListPages(ctx context.Context) (*[]Page, error) {
	path := "v1/pages"
	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var pages []Page
	_, err = s.client.do(ctx, req, &pages)

	return &pages, err
}

// UpdatePageParams are the parameters that can be changed using the update page API endpoint
type UpdatePageParams struct {
	Name                     string `json:"name,omitempty"`
	Domain                   string `json:"domain,omitempty"`
	Subdomain                string `json:"subdomain,omitempty"`
	URL                      string `json:"url,omitempty"`
	Branding                 string `json:"branding,omitempty"`
	CSSBodyBackgroundColor   string `json:"css_body_background_color,omitempty"`
	CSSFontColor             string `json:"css_font_color,omitempty"`
	CSSLightFontColor        string `json:"css_light_font_color,omitempty"`
	CSSGreens                string `json:"css_greens,omitempty"`
	CSSYellows               string `json:"css_yellows,omitempty"`
	CSSOranges               string `json:"css_oranges,omitempty"`
	CSSReds                  string `json:"css_reds,omitempty"`
	CSSBlues                 string `json:"css_blues,omitempty"`
	CSSBorderColor           string `json:"css_border_color,omitempty"`
	CSSGraphColor            string `json:"css_graph_color,omitempty"`
	CSSLinkColor             string `json:"css_link_color,omitempty"`
	HiddenFromSearch         *bool  `json:"hidden_from_search,omitempty"`
	ViewersMustBeTeamMembers *bool  `json:"viewers_must_be_team_members,omitempty"`
	AllowPageSubscribers     *bool  `json:"allow_page_subscribers,omitempty"`
	AllowIncidentSubscribers *bool  `json:"allow_incident_subscribers,omitempty"`
	AllowEmailSubscribers    *bool  `json:"allow_email_subscribers,omitempty"`
	AllowSmsSubscribers      *bool  `json:"allow_sms_subscribers,omitempty"`
	AllowRssAtomFeeds        *bool  `json:"allow_rss_atom_feeds,omitempty"`
	AllowWebhookSubscribers  *bool  `json:"allow_webhook_subscribers,omitempty"`
	NotificationsFromEmail   string `json:"notifications_from_email,omitempty"`
	TimeZone                 string `json:"time_zone,omitempty"`
	NotificationsEmailFooter string `json:"notifications_email_footer,omitempty"`
}

// UpdatePageRequestBody is the update page request body representation
type UpdatePageRequestBody struct {
	Page UpdatePageParams `json:"page"`
}

// UpdatePage updates page information for a given page id
func (s *PageService) UpdatePage(ctx context.Context, pageID string, page UpdatePageParams) (*Page, error) {
	path := "v1/pages/" + pageID
	payload := UpdatePageRequestBody{Page: page}
	req, err := s.client.newRequest("PATCH", path, payload)
	if err != nil {
		return nil, err
	}

	var updatedPage Page
	_, err = s.client.do(ctx, req, &updatedPage)

	return &updatedPage, err
}

// GetPage returns the page information for a given page id
func (s *PageService) GetPage(ctx context.Context, pageID string) (*Page, error) {
	path := "v1/pages/" + pageID
	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var page Page
	_, err = s.client.do(ctx, req, &page)

	return &page, err
}
