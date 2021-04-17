package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserHistoryService_Get(t *testing.T) {
	client, mux, _, teardown := checkoutSetup()
	defer teardown()

	mux.HandleFunc("/2.0/user/history", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testSoftFormValues(t, r, values{"email": "abc", "product_id": "1"})
		fmt.Fprint(w, `{"success":true, "response": {"message": "abc"}}`)
	})

	options := &UserHistoryOptions{ProductID: Int64(1)}
	history, _, err := client.UserHistory.Get(context.Background(), "abc", options)
	if err != nil {
		t.Errorf("UserHistory.Get returned error: %v", err)
	}

	want := &UserHistory{Message: String("abc")}
	if !reflect.DeepEqual(history, want) {
		t.Errorf("UserHistory.Get returned %+v, want %+v", history, want)
	}
}
