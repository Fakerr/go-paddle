package paddle

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWebhooksService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/2.0/alert/webhooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"page": "1", "alerts_per_page": "2"})
		fmt.Fprint(w, `{"success":true, "response": {"current_page":1, "data":[{"id": 1, "fields": {"order_id": 1}}]}}`)
	})

	opt := &WebhookEventOptions{ListOptions: ListOptions{Page: 1}, AlertsPerPage: "2"}
	event, _, err := client.Webhooks.Get(context.Background(), opt)
	if err != nil {
		t.Errorf("Webhooks.Get returned error: %v", err)
	}

	want := &WebhookEvent{CurrentPage: Int(1), Data: []*EventData{{ID: Int(1), Fields: &EventField{OrderID: Int(1)}}}}
	if !reflect.DeepEqual(event, want) {
		t.Errorf("Webhooks.Get returned %+v, want %+v", event, want)
	}
}
