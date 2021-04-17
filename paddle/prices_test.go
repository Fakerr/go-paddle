package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPricesService_Get(t *testing.T) {
	client, mux, _, teardown := checkoutSetup()
	defer teardown()

	mux.HandleFunc("/2.0/prices", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testSoftFormValues(t, r, values{"product_ids": "1", "customer_country": "tn"})
		fmt.Fprint(w, `{"success":true, "response": {
                        "customer_country": "tn",
                        "products": [{
                              "product_id": 1,
                              "price": {"gross": 10},
                              "applied_coupon": {"code": "1"}
                         }]}}`)
	})

	options := &PricesOptions{CustomerCountry: "tn"}
	prices, _, err := client.Prices.Get(context.Background(), "1", options)
	if err != nil {
		t.Errorf("Prices.Get returned error: %v", err)
	}

	want := &Prices{
		CustomerCountry: String("tn"),
		Products: []*Product{{
			ProductID:     Int(1),
			Price:         &Price{Gross: Float64(10)},
			AppliedCoupon: &AppliedCoupon{Code: String("1")},
		}},
	}
	if !reflect.DeepEqual(prices, want) {
		t.Errorf("Prices.Get returned %+v, want %+v", prices, want)
	}
}
