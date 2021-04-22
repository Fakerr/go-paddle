package paddle

import (
	"context"
	"net/http"
)

// PayLinkService handles communication with the pay link related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/pay-links
type PayLinkService service

type PayLinkCreate struct {
	ProductID               int    `url:"product_id,omitempty"`
	Title                   string `url:"title,omitempty"`
	WebhookURL              string `url:"webhook_url,omitempty"`
	Prices                  string `url:"prices,omitempty"`
	RecurringPrices         string `url:"recurring_prices,omitempty"`
	TrialDays               int    `url:"trial_days,omitempty"`
	CustomMessage           string `url:"custom_message,omitempty"`
	CouponCode              string `url:"coupon_code,omitempty"`
	Discountable            int    `url:"discountable,omitempty"`
	ImageURL                string `url:"image_url,omitempty"`
	ReturnURL               string `url:"return_url,omitempty"`
	QuantityVariable        int    `url:"quantity_variable,omitempty"`
	Quantity                int    `url:"quantity,omitempty"`
	Expires                 string `url:"expires,omitempty"`
	Affiliates              string `url:"affiliates,omitempty"`
	RecurringAffiliateLimit int    `url:"recurring_affiliate_limit,omitempty"`
	MarketingConsent        int    `url:"marketing_consent,omitempty"`
	CustomerEmail           string `url:"customer_email,omitempty"`
	CustomerCountry         string `url:"customer_country,omitempty"`
	CustomerPostcode        string `url:"customer_postcode,omitempty"`
	Passthrough             string `url:"passthrough,omitempty"`
	VatNumber               string `url:"vat_number,omitempty"`
	VatCompanyName          string `url:"vat_company_name,omitempty"`
	VatStreet               string `url:"vat_street,omitempty"`
	VatCity                 string `url:"vat_city,omitempty"`
	VatState                string `url:"vat_state,omitempty"`
	VatCountry              string `url:"vat_country,omitempty"`
	VatPostcode             string `url:"vat_postcode,omitempty"`
}

type PayLinkCreateResponse struct {
	Success  bool `json:"success"`
	Response *struct {
		URL *string `json:"url"`
	} `json:"response"`
}

// Generate a link with custom attributes set for a one-time or subscription checkout
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/pay-links/createpaylink
func (s *PayLinkService) Create(ctx context.Context, payLink *PayLinkCreate) (*string, *http.Response, error) {
	u := "2.0/product/generate_pay_link"

	req, err := s.client.NewRequest("POST", u, payLink)
	if err != nil {
		return nil, nil, err
	}

	payLinkCreateResponse := new(PayLinkCreateResponse)
	response, err := s.client.Do(ctx, req, payLinkCreateResponse)
	if err != nil {
		return nil, response, err
	}

	return payLinkCreateResponse.Response.URL, response, nil
}
