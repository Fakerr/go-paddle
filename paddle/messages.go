package paddle

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
)

// ValidatePayload validates an incoming Paddle Webhook event request
// and returns the (map[string]string) payload.
// The Content-Type header of the payload needs to be "application/x-www-form-urlencoded".
// If the Content-Type is different then an error is returned.
// publicKey is the Paddle public key.
//
// Example usage:
//
//    func PaddleWebhookHandler(w http.ResponseWriter, r *http.Request) {
//	payload, err := paddle.ValidatePayload(r, []byte(config.PaddleWebHookPublicKey))
//	if err != nil { ... }
//      // Process payload...
//    }
func ValidatePayload(r *http.Request, publicKey []byte) (map[string]string, error) {
	payload := map[string]string{}

	ct := r.Header.Get("Content-Type")
	if ct != "application/x-www-form-urlencoded" {
		return nil, fmt.Errorf("Webhook request has unsupported Content-Type %q", ct)
	}

	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	// Get the p_signature parameter.
	p_signature := r.Form.Get("p_signature")

	// Remove the p_signature parameter from Form
	r.Form.Del("p_signature")

	// Verify signature to make sure the request was sent by Paddle
	if err := validateSignature(r.Form, p_signature, publicKey); err != nil {
		return nil, err
	}

	// Construct payload from fields sent in the request
	for k := range r.Form {
		payload[k] = r.Form.Get(k) // r.Form is a map[string][]string
	}

	return payload, nil
}

// validateSignature validates the signature for the given payload.
// The signature is included on each webhook with the attribute p_signature.
// payload is the Form payload sent by Paddle Webhooks.
// publicKey is the Paddle public key.
//
// Paddle Reference: https://paddle.com/docs/reference-verifying-webhooks/
func validateSignature(form url.Values, p_signature string, publicKey []byte) error {
	// Find PEM public key block.
	der, _ := pem.Decode(publicKey)
	if der == nil {
		return errors.New("Could not parse public key pem")
	}

	// Parse public key in PKIX, ASN.1 DER form.
	pub, err := x509.ParsePKIXPublicKey(der.Bytes)
	if err != nil {
		return errors.New("Could not parse public key pem der")
	}

	signingKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return errors.New("Not the correct key format")
	}

	// base64 decode p_signature
	sig, err := base64.StdEncoding.DecodeString(p_signature)
	if err != nil {
		return err
	}

	// ksort() and serialize the Form
	sha1Sum := sha1.Sum(phpserialize(form))

	err = rsa.VerifyPKCS1v15(signingKey, crypto.SHA1, sha1Sum[:], sig)
	if err != nil {
		return err
	}
	return nil
}

// php serialize Form in sorted order
func phpserialize(form url.Values) []byte {
	var keys []string
	for k := range form {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	serialized := fmt.Sprintf("a:%d:{", len(keys))
	for _, k := range keys {
		serialized += fmt.Sprintf("s:%d:\"%s\";s:%d:\"%s\";", len(k), k, len(form.Get(k)), form.Get(k))
	}
	serialized += "}"

	return []byte(serialized)
}

// ParsePayload parses the alert payload. For recognized alert types, a
// value of the corresponding struct type will be returned.
// An error will be returned for unrecognized alert types.
//
// Example usage:
//
//    func PaddleWebhookHandler(w http.ResponseWriter, r *http.Request) {
//       payload, err := paddle.ValidatePayload(r, s.webhookSecretKey)
//       if err != nil { ... }
//       alert, err := paddle.ParsePayload(payload)
//       if err != nil { ... }
//       switch alert := alert.(type) {
//       case *paddle.SubscriptionCreatedAlert:
//           processSubscriptionCreatedAlert(alert)
//       case *paddle.SubscriptionCanceledAlert:
//           processSubscriptionCanceledAlert(alert)
//       ...
//       }
//     }
//
func ParsePayload(payload map[string]string) (interface{}, error) {
	var parsedPayload interface{}
	alert_type := payload["alert_name"]

	switch alert_type {
	case "subscription_created":
		parsedPayload = &SubscriptionCreatedAlert{}
	case "subscription_updated":
		parsedPayload = &SubscriptionUpdatedAlert{}
	case "subscription_cancelled":
		parsedPayload = &SubscriptionCancelledAlert{}
	case "subscription_payment_succeeded":
		parsedPayload = &SubscriptionPaymentSucceededAlert{}
	case "subscription_payment_failed":
		parsedPayload = &SubscriptionPaymentFailedAlert{}
	case "subscription_payment_refunded":
		parsedPayload = &SubscriptionPaymentRefundedAlert{}
	case "payment_succeeded":
		parsedPayload = &PaymentSucceededAlert{}
	case "payment_refunded":
		parsedPayload = &PaymentRefundedAlert{}
	case "locker_processed":
		parsedPayload = &LockerProcessedAlert{}
	case "payment_dispute_created":
		parsedPayload = &PaymentDisputeCreatedAlert{}
	case "payment_dispute_closed":
		parsedPayload = &PaymentDisputeClosedAlert{}
	case "high_risk_transaction_created":
		parsedPayload = &HighRiskTransactionCreatedAlert{}
	case "high_risk_transaction_updated":
		parsedPayload = &HighRiskTransactionUpdatedAlert{}
	case "transfer_created":
		parsedPayload = &TransferCreatedAlert{}
	case "transfer_paid":
		parsedPayload = &TransferPaidAlert{}
	case "new_audience_member":
		parsedPayload = &NewAudienceMemberAlert{}
	case "update_audience_member":
		parsedPayload = &UpdateAudienceMemberAlert{}
	case "invoice_paid":
		parsedPayload = &InvoicePaidAlert{}
	case "invoice_sent":
		parsedPayload = &InvoiceSentAlert{}
	case "invoice_overdue":
		parsedPayload = &InvoiceOverdueAlert{}
	default:
		return nil, fmt.Errorf("unknown alert_type: %v", alert_type)
	}

	// Marshal payload and unmarshal it
	j, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(j, &parsedPayload); err != nil {
		return nil, err
	}

	return parsedPayload, nil
}
