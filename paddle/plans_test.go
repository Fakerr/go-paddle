package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPlansService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/plans", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"plan": "1"})
		fmt.Fprint(w, `{"success":true, "response": [{"id":1, "initial_price":{"USD": "79.00"}}]}`)
	})

	opt := &PlansOptions{PlanID: 1}
	plans, _, err := client.Plans.List(context.Background(), opt)
	if err != nil {
		t.Errorf("Plans.List returned error: %v", err)
	}

	want := []*Plan{{ID: Int(1), InitialPrice: map[string]interface{}{"USD": "79.00"}}}
	if !reflect.DeepEqual(plans, want) {
		t.Errorf("Plans.List returned %+v, want %+v", plans, want)
	}
}

func TestPlansService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/plans_create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"plan_name": "a", "plan_length": "1", "plan_type": "month", "plan_trial_days": "10"})
		fmt.Fprint(w, `{"success":true, "response": {"product_id":1}}`)
	})

	opt := &PlanCreateOptions{PlanTrialDays: 10}
	product, _, err := client.Plans.Create(context.Background(), "a", "month", 1, opt)
	if err != nil {
		t.Errorf("Plans.Create returned error: %v", err)
	}

	want := &Product{ProductID: Int(1)}
	if !reflect.DeepEqual(product, want) {
		t.Errorf("Plans.Create returned %+v, want %+v", product, want)
	}
}
