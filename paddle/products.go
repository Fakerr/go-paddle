package paddle

import (
	"context"
	"net/http"
)

// ProductsService handles communication with the products related
// methods of the Paddle API.
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/products
type ProductsService service

// Product represents a Paddle plan.
type OneTimeProducts struct {
	Total    *int              `json:"total,omitempty"`
	Count    *int              `json:"count,omitempty"`
	Products []*OneTimeProduct `json:"products,omitempty"`
}

type OneTimeProduct struct {
	ID          *int                      `json:"id,omitempty"`
	Name        *string                   `json:"name,omitempty"`
	Description *string                   `json:"description,omitempty"`
	BasePrice   *float64                  `json:"base_price,omitempty"`
	SalePrice   *string                   `json:"sale_price,omitempty"`
	Screenshots *[]map[string]interface{} `json:"screenshots,omitempty"`
	Icon        *string                   `json:"icon,omitempty"`
	Currency    *string                   `json:"currency,omitempty"`
}

type ProductsResponse struct {
	Success  bool             `json:"success"`
	Response *OneTimeProducts `json:"response"`
}

// List all published one-time products in your account
//
// Paddle API docs: https://developer.paddle.com/api-reference/product-api/products/getproducts
func (s *ProductsService) List(ctx context.Context) (*OneTimeProducts, *http.Response, error) {
	u := "2.0/product/get_products"
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	productsResponse := new(ProductsResponse)
	response, err := s.client.Do(ctx, req, productsResponse)
	if err != nil {
		return nil, response, err
	}

	return productsResponse.Response, response, nil
}
