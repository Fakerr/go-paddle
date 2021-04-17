package paddle

import (
	"context"
	"fmt"
	"net/http"
)

// PricesService handles communication with the prices related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/checkout-api/prices
type PricesService service

// Prices represents a Paddle order details.
type Prices struct {
	CustomerCountry *string    `json:"customer_country,omitempty"`
	Products        []*Product `json:"products,omitempty"`
}

type Product struct {
	ProductID                  *int           `json:"product_id,omitempty"`
	ProductTitle               *string        `json:"product_title,omitempty"`
	Currency                   *string        `json:"currency,omitempty"`
	VendorSetPricesIncludedTax *bool          `json:"vendor_set_prices_included_tax,omitempty"`
	Price                      *Price         `json:"price,omitempty"`
	ListPrice                  *Price         `json:"list_price,omitempty"`
	AppliedCoupon              *AppliedCoupon `json:"applied_coupon,omitempty"`
}

type Price struct {
	Gross *float64 `json:"gross,omitempty"`
	Net   *float64 `json:"net,omitempty"`
	Tax   *float64 `json:"tax,omitempty"`
}

type AppliedCoupon struct {
	Code     *string  `json:"code,omitempty"`
	Discount *float64 `json:"discount,omitempty"`
}

type PricesResponse struct {
	Success  bool    `json:"success"`
	Response *Prices `json:"response"`
}

type PricesOptions struct {
	CustomerCountry string
	CustomerIP      string
	Coupons         string
}

// Retrieve prices for one or multiple products or plans
//
// Paddle API docs: https://developer.paddle.com/api-reference/checkout-api/prices/getprices
func (s *PricesService) Get(ctx context.Context, productIDs string, options *PricesOptions) (*Prices, *http.Response, error) {
	u := fmt.Sprintf("2.0/prices?product_ids=%v", productIDs)

	if options != nil {
		if options.CustomerCountry != "" {
			u = fmt.Sprintf("%s&customer_country=%v", u, options.CustomerCountry)
		}
		if options.CustomerIP != "" {
			u = fmt.Sprintf("%s&customer_ip=%v", u, options.CustomerIP)
		}
		if options.Coupons != "" {
			u = fmt.Sprintf("%s&coupons=%v", u, options.Coupons)
		}
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	pricesResponse := new(PricesResponse)
	response, err := s.client.Do(ctx, req, pricesResponse)
	if err != nil {
		return nil, response, err
	}

	return pricesResponse.Response, response, nil
}
