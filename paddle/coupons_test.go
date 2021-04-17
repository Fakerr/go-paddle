package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCouponsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/product/list_coupons", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"product_id": "1"})
		fmt.Fprint(w, `{"success":true, "response": [{"coupon":"1"}]}`)
	})

	coupons, _, err := client.Coupons.List(context.Background(), 1)
	if err != nil {
		t.Errorf("Coupons.List returned error: %v", err)
	}

	want := []*Coupon{{Coupon: String("1")}}
	if !reflect.DeepEqual(coupons, want) {
		t.Errorf("Coupons.List returned %+v, want %+v", coupons, want)
	}
}

func TestCouponsService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.1/product/create_coupon", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"coupon_type": "a", "discount_type": "a", "discount_amount": "10", "coupon_code": "1"})
		fmt.Fprint(w, `{"success":true, "response": {"coupon_code": ["1"]}}`)
	})

	options := &CouponCreateOptions{CouponCode: "1"}
	codes, _, err := client.Coupons.Create(context.Background(), "a", "a", 10, options)
	if err != nil {
		t.Errorf("Coupons.Create returned error: %v", err)
	}

	want := &CouponCodes{CouponCode: []string{"1"}}
	if !reflect.DeepEqual(codes, want) {
		t.Errorf("Coupons.Create returned %+v, want %+v", codes, want)
	}
}

func TestCouponsService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/product/delete_coupon", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"coupon_code": "a", "product_id": "1"})
		fmt.Fprint(w, `{"success":true}`)
	})

	options := &CouponDeleteOptions{ProductID: 1}
	resp, _, err := client.Coupons.Delete(context.Background(), "a", options)
	if err != nil {
		t.Errorf("Coupons.Delete returned error: %v", err)
	}

	want := true
	if !reflect.DeepEqual(resp, want) {
		t.Errorf("Coupons.Delete returned %+v, want %+v", resp, want)
	}
}

func TestCouponsService_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.1/product/update_coupon", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"coupon_code": "a"})
		fmt.Fprint(w, `{"success":true, "response": {"updated": 1}}`)
	})

	options := &CouponUpdateOptions{CouponCode: "a"}
	update, _, err := client.Coupons.Update(context.Background(), options)
	if err != nil {
		t.Errorf("Coupons.Update returned error: %v", err)
	}

	want := 1
	if !reflect.DeepEqual(*update, want) {
		t.Errorf("Coupons.Update returned %+v, want %+v", *update, want)
	}
}
