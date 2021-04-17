package paddle

import (
	"context"
	"fmt"
	"net/http"
)

// OrderDetailsService handles communication with the order_details related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/checkout-api/order-details/
type OrderDetailsService service

// OrderDetails represents a Paddle order details.
type OrderDetails struct {
	State    *string   `json:"state,omitempty"`
	Checkout *Checkout `json:"checkout,omitempty"`
	Order    *Order    `json:"order,omitempty"`
	Lockers  []*Locker `json:"lockers,omitempty"`
}

type Checkout struct {
	CheckoutID *string `json:"checkout_id,omitempty"`
	ImageURL   *string `json:"image_url,omitempty"`
	Title      *string `json:"title,omitempty"`
}

type Order struct {
	OrderID                    *int            `json:"order_id,omitempty"`
	Total                      *string         `json:"total,omitempty"`
	TotalTax                   *string         `json:"total_tax,omitempty"`
	Currency                   *string         `json:"currency,omitempty"`
	FormattedTotal             *string         `json:"formatted_total,omitempty"`
	FormattedTax               *string         `json:"formatted_tax,omitempty"`
	CouponCode                 *string         `json:"coupon_code,omitempty"`
	ReceiptUrl                 *string         `json:"receipt_url,omitempty"`
	CustomerSuccessRedirectURL *string         `json:"customer_success_redirect_url,omitempty"`
	HasLocker                  *bool           `json:"rhas_locker,omitempty"`
	IsSubscription             *bool           `json:"is_subscription,omitempty"`
	ProductID                  *int            `json:"product_id,omitempty"`
	SubscriptionID             *int            `json:"subscription_id,omitempty"`
	SubscriptionOrderID        *string         `json:"subscription_order_id,omitempty"`
	Quantity                   *int            `json:"quantity,omitempty"`
	Completed                  *OrderCompleted `json:"completed,omitempty"`
	Customer                   *Customer       `json:"customer,omitempty"`
}

type OrderCompleted struct {
	Date         *string `json:"date,omitempty"`
	TimeZone     *string `json:"time_zone,omitempty"`
	TimeZoneType *int    `json:"time_zone_type,omitempty"`
}

type Customer struct {
	Email            *string `json:"email,omitempty"`
	MarketingConsent *bool   `json:"marketing_consent,omitempty"`
}

type Locker struct {
	LockerID     *int    `json:"locker_id,omitempty"`
	ProductID    *int    `json:"product_id,omitempty"`
	ProductName  *string `json:"product_name,omitempty"`
	LicenseCode  *string `json:"license_code,omitempty"`
	Instructions *string `json:"instructions,omitempty"`
	Download     *string `json:"download,omitempty"`
}

type OrderDetailsResponse struct {
	Success  bool          `json:"success"`
	Response *OrderDetails `json:"response"`
}

// Get information about an order after a transaction completes
//
// Paddle API docs: https://developer.paddle.com/api-reference/checkout-api/order-details/getorder
func (s *OrderDetailsService) Get(ctx context.Context, checkoutID string) (*OrderDetails, *http.Response, error) {
	u := fmt.Sprintf("1.0/order?checkout_id=%v", checkoutID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	orderResponse := new(OrderDetailsResponse)
	response, err := s.client.Do(ctx, req, orderResponse)
	if err != nil {
		return nil, response, err
	}

	return orderResponse.Response, response, nil
}
