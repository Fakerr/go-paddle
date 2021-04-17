package paddle

import (
	"context"
	"net/http"
)

// UsersService handles communication with the user related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/users
type UsersService service

// User represents a Paddle user.
type User struct {
	SubscriptionID     *int                `json:"subscription_id,omitempty"`
	PlanID             *int                `json:"plan_id,omitempty"`
	UserID             *int                `json:"user_id,omitempty"`
	UserEmail          *string             `json:"user_email,omitempty"`
	MarketingConsent   *bool               `json:"marketing_consent,omitempty"`
	UpdateURL          *string             `json:"update_url,omitempty"`
	CancelURL          *string             `json:"cancel_url,omitempty"`
	State              *string             `json:"state,omitempty"`
	SignupDate         *string             `json:"signup_date,omitempty"`
	LastPayment        *UserPayment        `json:"last_payment,omitempty"`
	NextPayment        *UserPayment        `json:"next_payment,omitempty"`
	PaymentInformation *PaymentInformation `json:"payment_information,omitempty"`
	PausedAt           *string             `json:"paused_at,omitempty"`
	PausedFrom         *string             `json:"paused_from,omitempty"`
}

type UserPayment struct {
	Amount   *float64 `json:"amount,omitempty"`
	Currency *string  `json:"currency,omitempty"`
	Date     *string  `json:"date,omitempty"`
}

type PaymentInformation struct {
	PaymentMethod  *string `json:"payment_method,omitempty"`
	CardType       *string `json:"card_type,omitempty"`
	LastFourDigits *string `json:"last_four_digits,omitempty"`
	ExpiryDate     *string `json:"expiry_date,omitempty"`
}

type UsersResponse struct {
	Success  bool    `json:"success"`
	Response []*User `json:"response"`
}

// UsersOptions specifies the optional parameters to the
// UsersService.List method.
type UsersOptions struct {
	// SubscriptionID filters users based on their susbscription id.
	SubscriptionID string `url:"subscription_id,omitempty"`

	// PlanID filters users by the plan id.
	PlanID string `url:"plan_id,omitempty"`

	// State filters users based on the state. Possible values are: active,
	// past_due, trialing, paused, deleted. Returns all active, past_due,
	// trialing and paused subscription plans if not specified.
	State string `url:"state,omitempty"`

	ListOptions
}

// List all users subscribed to any of your subscription plans
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/users/listusers
func (s *UsersService) List(ctx context.Context, options *UsersOptions) ([]*User, *http.Response, error) {
	u := "2.0/subscription/users"
	req, err := s.client.NewRequest("POST", u, options)
	if err != nil {
		return nil, nil, err
	}

	usersResponse := new(UsersResponse)
	response, err := s.client.Do(ctx, req, usersResponse)
	if err != nil {
		return nil, response, err
	}

	return usersResponse.Response, response, nil
}

type UserUpdate struct {
	SubscriptionID  int     `url:"subscription_id,omitempty"`
	Quantity        int     `url:"quantity,omitempty"`
	Currency        string  `url:"currency,omitempty"`
	RecurringPrice  float64 `url:"recurring_price,omitempty"`
	BillImmediately bool    `url:"bill_immediately,omitempty"`
	PlanID          int     `url:"plan_id,omitempty"`
	Prorate         bool    `url:"prorate,omitempty"`
	KeepModifiers   bool    `url:"keep_modifiers,omitempty"`
	Passthrough     string  `url:"passthrough,omitempty"`
	Pause           bool    `url:"pause,omitempty"`
}

type UserUpdateOptions struct {
	Currency        string
	RecurringPrice  float64
	BillImmediately bool
	PlanID          int
	Prorate         bool
	KeepModifiers   bool
	Passthrough     string
	Pause           bool
}

type UserUpdateResponse struct {
	Success  bool  `json:"success"`
	Response *User `json:"response"`
}

// Update the quantity, price, and/or plan of a user’s subscription
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/users/updateuser
func (s *UsersService) Update(ctx context.Context, subscriptionID, quantity int, options *UserUpdateOptions) (*User, *http.Response, error) {
	u := "2.0/subscription/users/update"

	update := &UserUpdate{
		SubscriptionID: subscriptionID,
		Quantity:       quantity,
	}
	if options != nil {
		update.Currency = options.Currency
		update.RecurringPrice = options.RecurringPrice
		update.BillImmediately = options.BillImmediately
		update.PlanID = options.PlanID
		update.Prorate = options.Prorate
		update.KeepModifiers = options.KeepModifiers
		update.Passthrough = options.Passthrough
		update.Pause = options.Pause
	}
	req, err := s.client.NewRequest("POST", u, update)
	if err != nil {
		return nil, nil, err
	}

	userUpdateResponse := new(UserUpdateResponse)
	response, err := s.client.Do(ctx, req, userUpdateResponse)
	if err != nil {
		return nil, response, err
	}

	return userUpdateResponse.Response, response, nil
}

type UserCancel struct {
	SubscriptionID int `url:"subscription_id,omitempty"`
}

type UserCancelResponse struct {
	Success bool `json:"success"`
}

// Cancel the specified user’s subscription
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/users/canceluser
func (s *UsersService) Cancel(ctx context.Context, subscriptionID int) (bool, *http.Response, error) {
	u := "2.0/subscription/users_cancel"

	cancel := &UserCancel{
		SubscriptionID: subscriptionID,
	}
	req, err := s.client.NewRequest("POST", u, cancel)
	if err != nil {
		return false, nil, err
	}

	userCancelResponse := new(UserCancelResponse)
	response, err := s.client.Do(ctx, req, userCancelResponse)
	if err != nil {
		return false, response, err
	}

	return userCancelResponse.Success, response, nil
}
