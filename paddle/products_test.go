package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestProductsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/product/get_products", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"success":true, "response": {"total":1, "count":1, "products":[{"id": 1}]}}`)
	})

	products, _, err := client.Products.List(context.Background())
	if err != nil {
		t.Errorf("Products.List returned error: %v", err)
	}

	want := &OneTimeProducts{Total: Int(1), Count: Int(1), Products: []*OneTimeProduct{{ID: Int(1)}}}
	if !reflect.DeepEqual(products, want) {
		t.Errorf("Products.List returned %+v, want %+v", products, want)
	}
}
