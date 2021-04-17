package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestOneOffChargesService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/1/charge", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"amount": "10", "charge_name": "xyz"})
		fmt.Fprint(w, `{"success":true, "response": {"invoice_id":1, "subscription_id":1}}`)
	})

	oneOffCharge, _, err := client.OneOffCharges.Create(context.Background(), 1, 10, "xyz")
	if err != nil {
		t.Errorf("OneOffCharges.Create returned error: %v", err)
	}

	want := &OneOffCharge{InvoiceID: Int(1), SubscriptionID: Int(1)}
	if !reflect.DeepEqual(oneOffCharge, want) {
		t.Errorf("OneOffCharges.Create returned %+v, want %+v", oneOffCharge, want)
	}
}
