package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPaymentsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/payments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"subscription_id": "1"})
		fmt.Fprint(w, `{"success":true, "response": [{"id":1, "subscription_id":1}]}`)
	})

	opt := &PaymentsOptions{SubscriptionID: 1}
	payments, _, err := client.Payments.List(context.Background(), opt)
	if err != nil {
		t.Errorf("Payments.List returned error: %v", err)
	}

	want := []*Payment{{ID: Int(1), SubscriptionID: Int(1)}}
	if !reflect.DeepEqual(payments, want) {
		t.Errorf("Payments.List returned %+v, want %+v", payments, want)
	}
}

func TestPaymentsService_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/payments_reschedule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"payment_id": "1", "date": "2021-11-11"})
		fmt.Fprint(w, `{"success":true}`)
	})

	ok, _, err := client.Payments.Update(context.Background(), 1, "2021-11-11")
	if err != nil {
		t.Errorf("Payments.Update returned error: %v", err)
	}

	want := true
	if !reflect.DeepEqual(ok, want) {
		t.Errorf("Payments.List returned %+v, want %+v", ok, want)
	}
}
