package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRefundPaymentService_Refund(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/payment/refund", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"order_id": "1", "amount": "10"})
		fmt.Fprint(w, `{"success":true, "response": {"refund_request_id": 1}}`)
	})

	opt := &RefundPaymentOptions{Amount: float64(10)}
	request_id, _, err := client.RefundPayment.Refund(context.Background(), "1", opt)
	if err != nil {
		t.Errorf("RefundPayment.Refund returned error: %v", err)
	}

	want := 1
	if !reflect.DeepEqual(*request_id, want) {
		t.Errorf("RefundPayment.Refund returned %+v, want %+v", *request_id, want)
	}
}
