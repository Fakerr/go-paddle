package paddle

import (
	"context"
	"net/http"
)

// ModifiersService handles communication with the modifers related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/modifiers/
type ModifiersService service

// Modifier represents a Paddle modifier.
type Modifier struct {
	ModifierID     *int    `json:"modifier_id,omitempty"`
	SubscriptionID *int    `json:"subscription_id,omitempty"`
	Amount         *string `json:"amount,omitempty"`
	Currency       *string `json:"currency,omitempty"`
	IsRecurring    *bool   `json:"is_recurring,omitempty"`
	Description    *string `json:"description,omitempty"`
}

type ModifiersResponse struct {
	Success  bool        `json:"success"`
	Response []*Modifier `json:"response"`
}

// ModifiersOptions specifies the optional parameters to the
// ModifiersService.List method.
type ModifiersOptions struct {
	// SubscriptionID filters modifiers based on the subscription id.
	SubscriptionID int `url:"subscription_id,omitempty"`
	// PlanID filters modifiers based on the plan id.
	PlanID int `url:"plan_id,omitempty"`
}

// List all subscription modifiers
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/modifiers/listmodifiers
func (s *ModifiersService) List(ctx context.Context, options *ModifiersOptions) ([]*Modifier, *http.Response, error) {
	u := "2.0/subscription/modifiers"
	req, err := s.client.NewRequest("POST", u, options)
	if err != nil {
		return nil, nil, err
	}

	modifiersResponse := new(ModifiersResponse)
	response, err := s.client.Do(ctx, req, modifiersResponse)
	if err != nil {
		return nil, response, err
	}

	return modifiersResponse.Response, response, nil
}

type ModifierCreate struct {
	SubscriptionID      int     `url:"subscription_id,omitempty"`
	ModifierAmount      float64 `url:"modifier_amount,omitempty"`
	ModifierRecurring   bool    `url:"modifier_recurring,omitempty"`
	ModifierDescription string  `url:"modifier_description,omitempty"`
}

type ModifierCreateOptions struct {
	ModifierRecurring   bool
	ModifierDescription string
}

type ModifierCreateResponse struct {
	Success  bool      `json:"success"`
	Response *Modifier `json:"response"`
}

// Create a subscription modifier to dynamically change the subscription payment amount
//
// Paddle API docs:  https://developer.paddle.com/api-reference/subscription-api/modifiers/createmodifier
func (s *ModifiersService) Create(ctx context.Context, subscriptionID int, modifierAmount float64, options *ModifierCreateOptions) (*Modifier, *http.Response, error) {
	u := "2.0/subscription/modifiers/create"

	create := &ModifierCreate{
		SubscriptionID: subscriptionID,
		ModifierAmount: modifierAmount,
	}
	if options != nil {
		create.ModifierRecurring = options.ModifierRecurring
		create.ModifierDescription = options.ModifierDescription
	}
	req, err := s.client.NewRequest("POST", u, create)
	if err != nil {
		return nil, nil, err
	}

	modifierCreateResponse := new(ModifierCreateResponse)
	response, err := s.client.Do(ctx, req, modifierCreateResponse)
	if err != nil {
		return nil, response, err
	}

	return modifierCreateResponse.Response, response, nil
}

type ModifierDelete struct {
	ModifierID int `url:"modifier_id,omitempty"`
}

type ModifierDeleteResponse struct {
	Success bool `json:"success"`
}

// Delete an existing subscription modifier
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/modifiers/deletemodifier
func (s *ModifiersService) Delete(ctx context.Context, modifierID int) (bool, *http.Response, error) {
	u := "2.0/subscription/modifiers/delete"

	delete := &ModifierDelete{
		ModifierID: modifierID,
	}
	req, err := s.client.NewRequest("POST", u, delete)
	if err != nil {
		return false, nil, err
	}

	modifierDeleteResponse := new(ModifierDeleteResponse)
	response, err := s.client.Do(ctx, req, modifierDeleteResponse)
	if err != nil {
		return false, response, err
	}

	return modifierDeleteResponse.Success, response, nil
}
