package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUsersService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"subscription_id": "1", "page": "2"})
		fmt.Fprint(w, `{"success":true, "response": [{"user_id":2}]}`)
	})

	opt := &UsersOptions{SubscriptionID: "1", ListOptions: ListOptions{Page: 2}}
	users, _, err := client.Users.List(context.Background(), opt)
	if err != nil {
		t.Errorf("Users.List returned error: %v", err)
	}

	want := []*User{{UserID: Int(2)}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Users.List returned %+v, want %+v", users, want)
	}
}

func TestUsersService_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/users/update", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"subscription_id": "1", "quantity": "2", "plan_id": "123"})
		fmt.Fprint(w, `{"success":true, "response": {"subscription_id": 1, "user_id":2}}`)
	})

	opt := &UserUpdateOptions{PlanID: 123}
	resp, _, err := client.Users.Update(context.Background(), 1, 2, opt)
	if err != nil {
		t.Errorf("Users.Update returned error: %v", err)
	}

	want := &User{SubscriptionID: Int(1), UserID: Int(2)}
	if !reflect.DeepEqual(resp, want) {
		t.Errorf("Users.Update returned %+v, want %+v", resp, want)
	}
}

func TestUsersService_Cancel(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/subscription/users_cancel", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"subscription_id": "1"})
		fmt.Fprint(w, `{"success":true}`)
	})

	resp, _, err := client.Users.Cancel(context.Background(), 1)
	if err != nil {
		t.Errorf("Users.Cancel returned error: %v", err)
	}

	want := true
	if !reflect.DeepEqual(resp, want) {
		t.Errorf("Users.Cancel returned %+v, want %+v", resp, want)
	}
}
