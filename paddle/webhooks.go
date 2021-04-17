package paddle

import (
	"context"
	"net/http"
)

// WebhooksService handles communication with the webhooks related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/alert-api/webhooks/
type WebhooksService service

// WebhookEvent represents a Paddle plan.
type WebhookEvent struct {
	CurrentPage   *int         `json:"current_page,omitempty"`
	TotalPages    *int         `json:"total_pages,omitempty"`
	AlertsPerPage *int         `json:"alerts_per_page,omitempty"`
	TotalAlerts   *int         `json:"total_alerts,omitempty"`
	QueryHead     *string      `json:"query_head,omitempty"`
	QueryTail     *string      `json:"query_tail,omitempty"`
	Data          []*EventData `json:"data,omitempty"`
}

type EventData struct {
	ID        *int        `json:"id,omitempty"`
	AlertName *string     `json:"alert_name,omitempty"`
	Status    *string     `json:"status,omitempty"`
	CreatedAt *string     `json:"created_at,omitempty"`
	UpdatedAt *string     `json:"updated_at,omitempty"`
	Attempts  *int        `json:"attempts,omitempty"`
	Fields    *EventField `json:"fields,omitempty"`
}

type EventField struct {
	OrderID          *int    `json:"order_id,omitempty"`
	Amount           *string `json:"amount,omitempty"`
	Currency         *string `json:"currency,omitempty"`
	Email            *string `json:"email,omitempty"`
	MarketingConsent *int    `json:"marketing_consent,omitempty"`
}

type WebhookEventResponse struct {
	Success  bool          `json:"success"`
	Response *WebhookEvent `json:"response"`
}

// WebhookEventOptions specifies the optional parameters to the
// WebhooksService.Get method.
type WebhookEventOptions struct {
	// Number of webhook alerts to return per page. Returns 10 alerts by default.
	AlertsPerPage string `url:"alerts_per_page,omitempty"`
	// The date and time (UTC - Coordinated Universal Time) at which the webhook occurred before (end date). In the format: YYYY-MM-DD HH:MM:SS
	QueryHead string `url:"query_head,omitempty"`
	// The date and time (UTC - Coordinated Universal Time) at which the webhook occurred after (start date). In the format: YYYY-MM-DD HH:MM:SS
	QueryTail string `url:"query_tail,omitempty"`

	ListOptions
}

// Retrieve past events and alerts that Paddle has sent to webhooks on your account
//
// Paddle API docs: https://developer.paddle.com/api-reference/alert-api/webhooks/webhooks
func (s *WebhooksService) Get(ctx context.Context, options *WebhookEventOptions) (*WebhookEvent, *http.Response, error) {
	u := "2.0/alert/webhooks"
	req, err := s.client.NewRequest("POST", u, options)
	if err != nil {
		return nil, nil, err
	}

	eventResponse := new(WebhookEventResponse)
	response, err := s.client.Do(ctx, req, eventResponse)
	if err != nil {
		return nil, response, err
	}

	return eventResponse.Response, response, nil
}
