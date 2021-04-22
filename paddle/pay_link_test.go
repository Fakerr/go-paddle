package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPayLinkService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/product/generate_pay_link", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"product_id": "1"})
		fmt.Fprint(w, `{"success":true, "response": {"url":"abc"}}`)
	})

	payLink := &PayLinkCreate{ProductID: 1}
	url, _, err := client.PayLink.Create(context.Background(), payLink)
	if err != nil {
		t.Errorf("PayLink.Create returned error: %v", err)
	}

	want := "abc"
	if !reflect.DeepEqual(*url, want) {
		t.Errorf("PayLink.Create returned %+v, want %+v", *url, want)
	}
}
