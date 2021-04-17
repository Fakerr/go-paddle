package paddle

import (
	"context"
	"net/http"
)

// RefundPaymentService handles communication with the payments refund related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/payments/
type RefundPaymentService service

type RefundPaymentResponse struct {
	Success  bool `json:"success"`
	Response *struct {
		RefundRequestID *int `json:"refund_request_id"`
	} `json:"response"`
}

type RefundPayment struct {
	OrderID string  `url:"order_id,omitempty"`
	Amount  float64 `url:"amount,omitempty"`
	Reason  string  `url:"reason,omitempty"`
}

// RefundPaymentOptions specifies the optional parameters to the
// RefundPayment.Refund method.
type RefundPaymentOptions struct {
	Amount float64
	Reason string
}

// Request a refund for a one-time or subscription payment, either in full or partial
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/payments/refundpayment
func (s *RefundPaymentService) Refund(ctx context.Context, orderID string, options *RefundPaymentOptions) (*int, *http.Response, error) {
	u := "2.0/payment/refund"

	refund := &RefundPayment{
		OrderID: orderID,
	}
	if options != nil {
		refund.Amount = options.Amount
		refund.Reason = options.Reason
	}
	req, err := s.client.NewRequest("POST", u, refund)
	if err != nil {
		return nil, nil, err
	}

	refundPaymentsResponse := new(RefundPaymentResponse)
	response, err := s.client.Do(ctx, req, refundPaymentsResponse)
	if err != nil {
		return nil, response, err
	}

	return refundPaymentsResponse.Response.RefundRequestID, response, nil
}
