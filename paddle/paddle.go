package paddle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://vendors.paddle.com/api/"
	sandboxBaseURL = "https://sandbox-vendors.paddle.com/api/"

	checkoutBaseURL        = "https://checkout.paddle.com/api/"
	sandboxCheckoutBaseURL = "https://sandbox-checkout.paddle.com/api/"

	vendorIdAttribute       = "vendor_id"
	vendorAuthCodeAttribute = "vendor_auth_code"
)

// A Client manages communication with the Paddle API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// The vendor ID identifies your seller account. This can be found in Developer Tools > Authentication.
	VendorID *string

	// The vendor auth code is a private API key for authenticating API requests.
	// This key should never be used in client side code or shared publicly. This can be found in Developer Tools > Authentication.
	VendorAuthCode *string

	// Base URL for API requests. BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Paddle API.
	Users         *UsersService
	Plans         *PlansService
	Modifiers     *ModifiersService
	Payments      *PaymentsService
	OneOffCharges *OneOffChargesService
	Webhooks      *WebhooksService
	OrderDetails  *OrderDetailsService
	UserHistory   *UserHistoryService
	Prices        *PricesService
	Coupons       *CouponsService
	Products      *ProductsService
	RefundPayment *RefundPaymentService
	PayLink       *PayLinkService
}

type service struct {
	client *Client
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve. (minimum: 1)
	Page int `url:"page,omitempty"`

	// Number of subscription records to return per page. (minimum: 1, maximum: 200)
	ResultsPerPage int `url:"results_per_page,omitempty"`
}

// NewClient returns a new Paddle API client. It requires a vendor_id
// and a vendor_auth_code arguments. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(vendorID, vendorAuthCode string, httpClient *http.Client) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return getClient(httpClient, baseURL, vendorID, vendorAuthCode)
}

// NewSandboxClient returns a new Paddle API client for the sandbox environment.
// It requires a vendor_id and a vendor_auth_code arguments. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewSandboxClient(vendorID, vendorAuthCode string, httpClient *http.Client) *Client {
	baseURL, _ := url.Parse(sandboxBaseURL)
	return getClient(httpClient, baseURL, vendorID, vendorAuthCode)
}

// getCLient creates and returns a Paddle API client.
func getClient(httpClient *http.Client, baseURL *url.URL, vendorID, vendorAuthCode string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{
		client:         httpClient,
		BaseURL:        baseURL,
		VendorID:       String(vendorID),
		VendorAuthCode: String(vendorAuthCode),
	}

	c.common.client = c
	c.Users = (*UsersService)(&c.common)
	c.Plans = (*PlansService)(&c.common)
	c.Modifiers = (*ModifiersService)(&c.common)
	c.Payments = (*PaymentsService)(&c.common)
	c.OneOffCharges = (*OneOffChargesService)(&c.common)
	c.Webhooks = (*WebhooksService)(&c.common)
	c.Coupons = (*CouponsService)(&c.common)
	c.Products = (*ProductsService)(&c.common)
	c.RefundPayment = (*RefundPaymentService)(&c.common)
	c.PayLink = (*PayLinkService)(&c.common)
	return c
}

// NewCheckoutClient returns a new Paddle API client for checkouts.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewCheckoutClient(httpClient *http.Client) *Client {
	baseURL, _ := url.Parse(checkoutBaseURL)
	return getCheckoutClient(httpClient, baseURL)
}

// NewSandboxCheckoutClient returns a new Paddle API client for the sandbox checkout enivronement.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewSandboxCheckoutClient(httpClient *http.Client) *Client {
	baseURL, _ := url.Parse(sandboxCheckoutBaseURL)
	return getCheckoutClient(httpClient, baseURL)
}

// getCheckoutCLient creates and returns a checkout Paddle API client.
func getCheckoutClient(httpClient *http.Client, baseURL *url.URL) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}

	c.common.client = c
	c.OrderDetails = (*OrderDetailsService)(&c.common)
	c.UserHistory = (*UserHistoryService)(&c.common)
	c.Prices = (*PricesService)(&c.common)
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is url form encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, options interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	payload, err := newPayload(c.VendorID, c.VendorAuthCode, options)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), payload)
	if err != nil {
		return nil, err
	}

	if payload.Size() > 0 {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

// newPayload encodes opt into ``URL encoded'' form and return a *strings.Reader. opt
// must be a struct whose fields may contain "url" tags.
// Client's VendorID and VendorAuthCode will be attached to the payload.
func newPayload(vendorID, vendorAuthCode *string, opt interface{}) (*strings.Reader, error) {
	data, err := query.Values(opt)
	if err != nil {
		return nil, err
	}

	if vendorID != nil && vendorAuthCode != nil {
		data.Set(vendorIdAttribute, *vendorID)
		data.Set(vendorAuthCodeAttribute, *vendorAuthCode)
	}

	payload := strings.NewReader(data.Encode())
	return payload, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	if err := checkResponse(resp, data); err != nil {
		return resp, err
	}

	if v != nil {
		if err := json.Unmarshal(data, v); err != nil {
			return resp, fmt.Errorf("Unmarshal error %s\n", err)
		}
	}
	return resp, nil
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// An unsuccessful call to the Dashboard API will return a 200 response containing
// a field success set to false. Additionally an error object will be returned,
// containing a code referencing the error, and a message in a human-readable format.
type ErrorResponse struct {
	response *http.Response // HTTP response that caused this error

	Success    bool  `json:"success"`
	ErrorField Error `json:"error"`
}

// Check wether or not the API response contains an error
func checkResponse(r *http.Response, data []byte) error {
	errorResponse := &ErrorResponse{response: r}
	if data != nil {
		if err := json.Unmarshal(data, errorResponse); err != nil {
			return err
		}
	}
	if !errorResponse.Success {
		return fmt.Errorf("Error: %v, %s",
			errorResponse.ErrorField.Code,
			errorResponse.ErrorField.Message)
	}
	return nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Float64 is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
