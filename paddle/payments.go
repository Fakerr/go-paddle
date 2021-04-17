package paddle

import (
	"context"
	"net/http"
)

// PaymentsService handles communication with the payments related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/payments/
type PaymentsService service

// Payment represents a Paddle payment.
type Payment struct {
	ID             *int    `json:"id,omitempty"`
	SubscriptionID *int    `json:"subscription_id,omitempty"`
	Amount         *int    `json:"amount,omitempty"`
	Currency       *string `json:"currency,omitempty"`
	PayoutDate     *string `json:"payout_date,omitempty"`
	IsPaid         *int    `json:"is_paid,omitempty"`
	ReceiptUrl     *string `json:"receipt_url,omitempty"`
	IsOneOffCharge *int    `json:"is_one_off_charge,omitempty"`
}

type PaymentsResponse struct {
	Success  bool       `json:"success"`
	Response []*Payment `json:"response"`
}

// PaymentsOptions specifies the optional parameters to the
// Payments.List method.
type PaymentsOptions struct {
	// Payments for a specific subscription.
	SubscriptionID int `url:"subscription_id,omitempty"`
	// The product/plan ID (single or comma-separated values)
	Plan int `url:"plan,omitempty"`
	// Payment is paid (0 = No, 1 = Yes)
	IsPaid int `url:"is_paid,omitempty"`
	// Payments starting from (date in format YYYY-MM-DD)
	From string `url:"from,omitempty"`
	// Payments up to (date in format YYYY-MM-DD)
	To string `url:"to,omitempty"`
	// Non-recurring payments created from the
	IsOneOffCharge bool `url:"is_one_off_charge,omitempty"`
}

// List all paid and upcoming (unpaid) payments
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/payments/listpayments
func (s *PaymentsService) List(ctx context.Context, options *PaymentsOptions) ([]*Payment, *http.Response, error) {
	u := "2.0/subscription/payments"
	req, err := s.client.NewRequest("POST", u, options)
	if err != nil {
		return nil, nil, err
	}

	paymentsResponse := new(PaymentsResponse)
	response, err := s.client.Do(ctx, req, paymentsResponse)
	if err != nil {
		return nil, response, err
	}

	return paymentsResponse.Response, response, nil
}

type PaymentUpdate struct {
	PaymentID int    `url:"payment_id,omitempty"`
	Date      string `url:"date,omitempty"`
}

type PaymentUpdateResponse struct {
	Success bool `json:"success"`
}

// Change the due date of the upcoming subscription payment
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/payments/updatepayment
func (s *PaymentsService) Update(ctx context.Context, paymentID int, date string) (bool, *http.Response, error) {
	u := "2.0/subscription/payments_reschedule"

	update := &PaymentUpdate{
		PaymentID: paymentID,
		Date:      date,
	}

	req, err := s.client.NewRequest("POST", u, update)
	if err != nil {
		return false, nil, err
	}

	paymentUpdateResponse := new(PaymentUpdateResponse)
	response, err := s.client.Do(ctx, req, paymentUpdateResponse)
	if err != nil {
		return false, response, err
	}

	return paymentUpdateResponse.Success, response, nil
}
