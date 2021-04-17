package paddle

import (
	"context"
	"net/http"
)

// PlansService handles communication with the plans related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/plans
type PlansService service

// Plan represents a Paddle plan.
type Plan struct {
	ID             *int                   `json:"id,omitempty"`
	Name           *string                `json:"name,omitempty"`
	BillingType    *string                `json:"billing_type,omitempty"`
	BillingPeriod  *int                   `json:"billing_period,omitempty"`
	TrialDays      *int                   `json:"trial_days,omitempty"`
	InitialPrice   map[string]interface{} `json:"initial_price,omitempty"`
	RecurringPrice map[string]interface{} `json:"recurring_price,omitempty"`
}

type PlansResponse struct {
	Success  bool    `json:"success"`
	Response []*Plan `json:"response"`
}

// PlansOptions specifies the optional parameters to the
// PLansService.List method.
type PlansOptions struct {
	// PlanID filters Products/Plans based on their id.
	PlanID int `url:"plan,omitempty"`
}

// List all of the available subscription plans in your account
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/plans/listplans
func (s *PlansService) List(ctx context.Context, options *PlansOptions) ([]*Plan, *http.Response, error) {
	u := "2.0/subscription/plans"
	req, err := s.client.NewRequest("POST", u, options)
	if err != nil {
		return nil, nil, err
	}

	plansResponse := new(PlansResponse)
	response, err := s.client.Do(ctx, req, plansResponse)
	if err != nil {
		return nil, response, err
	}

	return plansResponse.Response, response, nil
}

type PlanCreate struct {
	PlanName          string `url:"plan_name,omitempty"`
	PlanLength        int    `url:"plan_length,omitempty"`
	PlanType          string `url:"plan_type,omitempty"`
	PlanTrialDays     int    `url:"plan_trial_days,omitempty"`
	MainCurrencyCode  string `url:"main_currency_code,omitempty"`
	RecurringPriceUsd string `url:"recurring_price_usd,omitempty"`
	RecurringPriceGbp string `url:"recurring_price_gbp,omitempty"`
	RecurringPriceEur string `url:"recurring_price_eur,omitempty"`
}

type PlanCreateOptions struct {
	PlanTrialDays     int
	MainCurrencyCode  string
	RecurringPriceUsd string
	RecurringPriceGbp string
	RecurringPriceEur string
}

type PlanCreateResponse struct {
	Success  bool     `json:"success"`
	Response *Product `json:"response"`
}

// Create a new subscription plan with the supplied parameters
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/plans/createplan
func (s *PlansService) Create(ctx context.Context, planName, planType string, planLength int, options *PlanCreateOptions) (*Product, *http.Response, error) {
	u := "2.0/subscription/plans_create"

	create := &PlanCreate{
		PlanName:   planName,
		PlanLength: planLength,
		PlanType:   planType,
	}
	if options != nil {
		create.PlanTrialDays = options.PlanTrialDays
		create.MainCurrencyCode = options.MainCurrencyCode
		create.RecurringPriceUsd = options.RecurringPriceUsd
		create.RecurringPriceGbp = options.RecurringPriceGbp
		create.RecurringPriceEur = options.RecurringPriceEur
	}
	req, err := s.client.NewRequest("POST", u, create)
	if err != nil {
		return nil, nil, err
	}

	planCreateResponse := new(PlanCreateResponse)
	response, err := s.client.Do(ctx, req, planCreateResponse)
	if err != nil {
		return nil, response, err
	}

	return planCreateResponse.Response, response, nil
}
