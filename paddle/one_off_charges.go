package paddle

import (
	"context"
	"fmt"
	"net/http"
)

// OneOffChargesService handles communication with the one-off charges related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/one-off-charges/
type OneOffChargesService service

// OneOffCharge represents a Paddle one-off charge.
type OneOffCharge struct {
	InvoiceID      *int     `json:"invoice_id,omitempty"`
	SubscriptionID *int     `json:"subscription_id,omitempty"`
	Amount         *float64 `json:"amount,omitempty"`
	Currency       *string  `json:"currency,omitempty"`
	PaymentDate    *string  `json:"payment_date,omitempty"`
	ReceiptUrl     *string  `json:"receipt_url,omitempty"`
	OrderID        *string  `json:"order_id,omitempty"`
	Status         *string  `json:"status,omitempty"`
}

type OneOffChargeCreate struct {
	Amount     float64 `url:"amount,omitempty"`
	ChargeName string  `url:"charge_name,omitempty"`
}

type OneOffChargeResponse struct {
	Success  bool          `json:"success"`
	Response *OneOffCharge `json:"response"`
}

// Make an immediate one-off charge on top of an existing user subscription
//
// Paddle API docs: https://developer.paddle.com/api-reference/subscription-api/one-off-charges/createcharge
func (s *OneOffChargesService) Create(ctx context.Context, subscriptionID int, amount float64, chargeName string) (*OneOffCharge, *http.Response, error) {
	u := fmt.Sprintf("2.0/subscription/%d/charge", subscriptionID)

	OneOffChargeCreate := &OneOffChargeCreate{
		Amount:     amount,
		ChargeName: chargeName,
	}

	req, err := s.client.NewRequest("POST", u, OneOffChargeCreate)
	if err != nil {
		return nil, nil, err
	}

	oneOffChargeResponse := new(OneOffChargeResponse)
	response, err := s.client.Do(ctx, req, oneOffChargeResponse)
	if err != nil {
		return nil, response, err
	}

	return oneOffChargeResponse.Response, response, nil
}
