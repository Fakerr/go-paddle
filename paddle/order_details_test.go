package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestOrderDetailsService_Get(t *testing.T) {
	client, mux, _, teardown := checkoutSetup()
	defer teardown()

	mux.HandleFunc("/1.0/order", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testSoftFormValues(t, r, values{"checkout_id": "1"})
		fmt.Fprint(w, `{"success":true, "response": {
		"state": "xyz",
		"checkout":{"checkout_id": "1"},
		"order": {"order_id": 1, "completed": {"date": "123"}, "customer": {"email": "abc"}},
		"Lockers": [{"locker_id": 1}]}}`)
	})

	order, _, err := client.OrderDetails.Get(context.Background(), "1")
	if err != nil {
		t.Errorf("OrderDetails.Get returned error: %v", err)
	}

	want := &OrderDetails{
		State:    String("xyz"),
		Checkout: &Checkout{CheckoutID: String("1")},
		Order: &Order{
			OrderID:   Int(1),
			Completed: &OrderCompleted{Date: String("123")},
			Customer:  &Customer{Email: String("abc")},
		},
		Lockers: []*Locker{{LockerID: Int(1)}},
	}

	if !reflect.DeepEqual(order, want) {
		t.Errorf("OrderDetails.Get returned %+v, want %+v", order, want)
	}
}
