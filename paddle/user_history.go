package paddle

import (
	"context"
	"fmt"
	"net/http"
)

// UserHistoryService handles communication with the user history related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/checkout-api/user-history/
type UserHistoryService service

// UserHistory represents a Paddle plan.
type UserHistory struct {
	Message  *string `json:"message,omitempty"`
	Callback *string `json:"callback,omitempty"`
}

type UserHistoryResponse struct {
	Success  bool         `json:"success"`
	Response *UserHistory `json:"response"`
}

// UserHistoryOptions specifies the optional parameters to the
// UserHistoryService.Get method.
type UserHistoryOptions struct {
	VendorID  *int64
	ProductID *int64
}

// Send the customer an order history and license recovery email
//
// Paddle API docs: https://developer.paddle.com/api-reference/checkout-api/user-history/getuserhistory
func (s *UserHistoryService) Get(ctx context.Context, email string, options *UserHistoryOptions) (*UserHistory, *http.Response, error) {
	u := fmt.Sprintf("2.0/user/history?email=%v", email)

	if options != nil {
		if options.VendorID != nil {
			u = fmt.Sprintf("%s&vendor_id=%d", u, *options.VendorID)
		}
		if options.ProductID != nil {
			u = fmt.Sprintf("%s&product_id=%d", u, *options.ProductID)
		}
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	userHistoryResponse := new(UserHistoryResponse)
	response, err := s.client.Do(ctx, req, userHistoryResponse)
	if err != nil {
		return nil, response, err
	}

	return userHistoryResponse.Response, response, nil
}
