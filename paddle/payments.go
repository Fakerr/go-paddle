package paddle

import (
	"context"
	"encoding/json"
	"net/http"
)

// PaymentsService handles communication with the payments related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/payments/
type PaymentsService service

// Payment represents a Paddle payment.
type Payment struct {
	ID             *int     `json:"id,omitempty"`
	SubscriptionID *int     `json:"subscription_id,omitempty"`
	Amount         *float64 `json:"amount,omitempty"`
	Currency       *string  `json:"currency,omitempty"`
	PayoutDate     *string  `json:"payout_date,omitempty"`
	IsPaid         *int     `json:"is_paid,omitempty"`
	ReceiptUrl     *string  `json:"receipt_url,omitempty"`
	IsOneOffCharge *int     `json:"is_one_off_charge,omitempty"`
}

// paymentTemp is temporary structure that aims to hotfix the invalid IsOneOffCharge type.
// In api declared as int, but returned as boolean (at least on sandbox)
type paymentTemp struct {
	ID             *int        `json:"id,omitempty"`
	SubscriptionID *int        `json:"subscription_id,omitempty"`
	Amount         *float64    `json:"amount,omitempty"`
	Currency       *string     `json:"currency,omitempty"`
	PayoutDate     *string     `json:"payout_date,omitempty"`
	IsPaid         *int        `json:"is_paid,omitempty"`
	ReceiptUrl     *string     `json:"receipt_url,omitempty"`
	IsOneOffCharge interface{} `json:"is_one_off_charge,omitempty"`
}

func (t *Payment) UnmarshalJSON(data []byte) error {
	var temp paymentTemp
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	var sanitizedIsOneOffCharge *int
	if temp.IsOneOffCharge != nil {
		switch v := temp.IsOneOffCharge.(type) {
		case *bool:
			if *v == true {
				trueVal := 1
				sanitizedIsOneOffCharge = &trueVal
			}
		case *int:
			sanitizedIsOneOffCharge = v
		}
	}

	*t = Payment{
		ID:             temp.ID,
		SubscriptionID: temp.SubscriptionID,
		Amount:         temp.Amount,
		Currency:       temp.Currency,
		PayoutDate:     temp.PayoutDate,
		IsPaid:         temp.IsPaid,
		ReceiptUrl:     temp.ReceiptUrl,
		IsOneOffCharge: sanitizedIsOneOffCharge,
	}

	return nil
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
