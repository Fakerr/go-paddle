package paddle

import (
	"context"
	"net/http"
)

// CouponsService handles communication with the coupons related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/coupons/listcoupons
type CouponsService service

// Coupon represents a Paddle plan.

type Coupon struct {
	Coupon           *string  `json:"coupon,omitempty"`
	Description      *string  `json:"description,omitempty"`
	DiscountType     *string  `json:"discount_type,omitempty"`
	DiscountAmount   *float64 `json:"discount_amount,omitempty"`
	DiscountCurrency *string  `json:"discount_currency,omitempty"`
	AllowedUses      *int     `json:"allowed_uses,omitempty"`
	TimesUsed        *int     `json:"times_used,omitempty"`
	IsRecurring      *bool    `json:"is_recurring,omitempty"`
	Expires          *string  `json:"expires,omitempty"`
}

type CouponsResponse struct {
	Success  bool      `json:"success"`
	Response []*Coupon `json:"response"`
}

// CouponsOptions specifies the optional parameters to the
// CouponsService.List method.
type CouponsOptions struct {
	ProductID int `url:"product_id,omitempty"`
}

// List all coupons valid for a specified one-time product or subscription plan
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/coupons/listcoupons
func (s *CouponsService) List(ctx context.Context, productID int) ([]*Coupon, *http.Response, error) {
	u := "2.0/product/list_coupons"

	options := &CouponsOptions{
		ProductID: productID,
	}

	req, err := s.client.NewRequest("POST", u, options)
	if err != nil {
		return nil, nil, err
	}

	couponsResponse := new(CouponsResponse)
	response, err := s.client.Do(ctx, req, couponsResponse)
	if err != nil {
		return nil, response, err
	}

	return couponsResponse.Response, response, nil
}

type CouponCreate struct {
	CouponCode     string  `url:"coupon_code,omitempty"`
	CouponPrefix   string  `url:"coupon_prefix,omitempty"`
	NumCoupons     int     `url:"num_coupons,omitempty"`
	Description    string  `url:"description,omitempty"`
	CouponType     string  `url:"coupon_type,omitempty"`
	ProductIds     string  `url:"product_ids,omitempty"`
	DiscountType   string  `url:"discount_type,omitempty"`
	DiscountAmount float64 `url:"discount_amount,omitempty"`
	Currency       string  `url:"currency,omitempty"`
	AllowedUses    int     `url:"allowed_uses,omitempty"`
	Expires        string  `url:"expires,omitempty"`
	Recurring      int     `url:"recurring,omitempty"`
	Group          string  `url:"group,omitempty"`
}

type CouponCreateOptions struct {
	CouponCode   string
	CouponPrefix string
	NumCoupons   int
	Description  string
	ProductIds   string
	Currency     string
	AllowedUses  int
	Expires      string
	Recurring    int
	Group        string
}

type CouponCreateResponse struct {
	Success  bool         `json:"success"`
	Response *CouponCodes `json:"response"`
}

type CouponCodes struct {
	CouponCode []string `json:"coupon_code,omitempty"`
}

// Create a new coupon for the given product or a checkout
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/coupons/createcoupon
func (s *CouponsService) Create(ctx context.Context, couponType, discountType string, discountAmount float64, options *CouponCreateOptions) (*CouponCodes, *http.Response, error) {
	u := "2.1/product/create_coupon"

	create := &CouponCreate{
		CouponType:     couponType,
		DiscountType:   discountType,
		DiscountAmount: discountAmount,
	}
	if options != nil {
		create.CouponCode = options.CouponCode
		create.CouponPrefix = options.CouponPrefix
		create.NumCoupons = options.NumCoupons
		create.Description = options.Description
		create.ProductIds = options.ProductIds
		create.Currency = options.Currency
		create.AllowedUses = options.AllowedUses
		create.Expires = options.Expires
		create.Recurring = options.Recurring
		create.Group = options.Group
	}
	req, err := s.client.NewRequest("POST", u, create)
	if err != nil {
		return nil, nil, err
	}

	couponCreateResponse := new(CouponCreateResponse)
	response, err := s.client.Do(ctx, req, couponCreateResponse)
	if err != nil {
		return nil, response, err
	}

	return couponCreateResponse.Response, response, nil
}

type CouponDelete struct {
	CouponCode *string `url:"coupon_code,omitempty"`
	ProductID  *int    `url:"product_id,omitempty"`
}

type CouponDeleteOptions struct {
	ProductID int
}

type CouponDeleteResponse struct {
	Success bool `json:"success"`
}

// Delete a given coupon and prevent it from being further used
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/coupons/deletecoupon
func (s *CouponsService) Delete(ctx context.Context, couponCode string, options *CouponDeleteOptions) (bool, *http.Response, error) {
	u := "2.0/product/delete_coupon"

	delete := &CouponDelete{
		CouponCode: String(couponCode),
	}
	if options != nil {
		delete.ProductID = Int(options.ProductID)
	}
	req, err := s.client.NewRequest("POST", u, delete)
	if err != nil {
		return false, nil, err
	}

	couponDeleteResponse := new(CouponDeleteResponse)
	response, err := s.client.Do(ctx, req, couponDeleteResponse)
	if err != nil {
		return false, response, err
	}

	return couponDeleteResponse.Success, response, nil
}

type CouponUpdateOptions struct {
	CouponCode     string  `url:"coupon_code,omitempty"`
	Group          string  `url:"group,omitempty"`
	NewCouponCode  string  `url:"new_coupon_code,omitempty"`
	NewGroup       string  `url:"new_group,omitempty"`
	ProductIds     string  `url:"product_ids,omitempty"`
	Expires        string  `url:"expires,omitempty"`
	AllowedUses    int     `url:"allowed_uses,omitempty"`
	Currency       string  `url:"currency,omitempty"`
	DiscountAmount float64 `url:"discount_amount,omitempty"`
	Recurring      int     `url:"recurring,omitempty"`
}

type CouponUpdateResponse struct {
	Success  bool `json:"success"`
	Response *struct {
		Updated *int `json:"updated"`
	} `json:"response"`
}

// Update an existing coupon in your account
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/coupons/updatecoupon
func (s *CouponsService) Update(ctx context.Context, options *CouponUpdateOptions) (*int, *http.Response, error) {
	u := "2.1/product/update_coupon"
	req, err := s.client.NewRequest("POST", u, options)
	if err != nil {
		return nil, nil, err
	}

	couponUpdateResponse := new(CouponUpdateResponse)
	response, err := s.client.Do(ctx, req, couponUpdateResponse)
	if err != nil {
		return nil, response, err
	}

	return couponUpdateResponse.Response.Updated, response, nil
}
