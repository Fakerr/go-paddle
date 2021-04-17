package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestModifiersService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/modifiers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"plan_id": "1"})
		fmt.Fprint(w, `{"success":true, "response": [{"modifier_id":1}]}`)
	})

	opt := &ModifiersOptions{PlanID: 1}
	modifiers, _, err := client.Modifiers.List(context.Background(), opt)
	if err != nil {
		t.Errorf("Modifiers.List returned error: %v", err)
	}

	want := []*Modifier{{ModifierID: Int(1)}}
	if !reflect.DeepEqual(modifiers, want) {
		t.Errorf("Modifiers.List returned %+v, want %+v", modifiers, want)
	}
}

func TestModifiersService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/modifiers/create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"subscription_id": "1", "modifier_amount": "1", "modifier_recurring": "true"})
		fmt.Fprint(w, `{"success":true, "response": {"subscription_id": 1, "modifier_id":1}}`)
	})

	opt := &ModifierCreateOptions{ModifierRecurring: true}
	modifier, _, err := client.Modifiers.Create(context.Background(), 1, 1, opt)
	if err != nil {
		t.Errorf("Modifiers.Create returned error: %v", err)
	}

	want := &Modifier{SubscriptionID: Int(1), ModifierID: Int(1)}
	if !reflect.DeepEqual(modifier, want) {
		t.Errorf("Modifiers.Create returned %+v, want %+v", modifier, want)
	}
}

func TestModifiersService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/modifiers/delete", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"modifier_id": "1"})
		fmt.Fprint(w, `{"success":true}`)
	})

	resp, _, err := client.Modifiers.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Modifiers.Delete returned error: %v", err)
	}

	want := true
	if !reflect.DeepEqual(resp, want) {
		t.Errorf("Modifiers.Delete returned %+v, want %+v", resp, want)
	}
}
