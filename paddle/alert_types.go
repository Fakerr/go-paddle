package paddle

// Sent when an order is processed for a product or plan with webhook fulfillment enabled
// Paddle reference: https://developer.paddle.com/webhook-reference/product-fulfillment/fulfillment-webhook
type FulfillmentWebhook struct {
	EventTime          *string `json:"event_time"`
	PCountry           *string `json:"p_country"`
	PCoupon            *string `json:"p_coupon"`
	PCouponSavings     *string `json:"p_coupon_savings"`
	PCurrency          *string `json:"p_currency"`
	PEarnings          *string `json:"p_earnings"`
	POrderID           *string `json:"p_order_id"`
	PPaddleFee         *string `json:"p_paddle_fee"`
	PPrice             *string `json:"p_price"`
	PProductID         *string `json:"p_product_id"`
	PQuantity          *string `json:"p_quantity"`
	PSaleGross         *string `json:"p_sale_gross"`
	PTaxAmount         *string `json:"p_tax_amount"`
	PUsedPriceOverride *string `json:"p_used_price_override"`
	Passthrough        *string `json:"passthrough"`
	Quantity           *string `json:"quantity"`
}

// Fired when a new subscription is created, and a customer has successfully subscribed.
// Paddle Reference: https://developer.paddle.com/webhook-reference/subscription-alerts/subscription-created
type SubscriptionCreatedAlert struct {
	AlertName          *string `json:"alert_name"`
	AlertID            *string `json:"alert_id"`
	CancelURL          *string `json:"cancel_url"`
	CheckoutID         *string `json:"checkout_id"`
	Currency           *string `json:"currency"`
	Email              *string `json:"email"`
	EventTime          *string `json:"event_time"`
	MarketingConsent   *int    `json:"marketing_consent"`
	NextBillDate       *string `json:"next_bill_date"`
	Passthrough        *string `json:"passthrough"`
	Quantity           *string `json:"quantity"`
	Source             *string `json:"source"`
	Status             *string `json:"status"`
	SubscriptionID     *string `json:"subscription_id"`
	SubscriptionPlanID *string `json:"subscription_plan_id"`
	UnitPrice          *string `json:"unit_price"`
	UserID             *string `json:"user_id"`
	UpdateURL          *string `json:"update_url"`
}

// Fired when the plan, price, quantity, status of an existing subscription changes, or if the payment date is rescheduled manually.
// Paddle reference: https://developer.paddle.com/webhook-reference/subscription-alerts/subscription-updated
type SubscriptionUpdatedAlert struct {
	AlertName             *string `json:"alert_name"`
	AlertID               *string `json:"alert_id"`
	CancelURL             *string `json:"cancel_url"`
	CheckoutID            *string `json:"checkout_id"`
	Email                 *string `json:"email"`
	EventTime             *string `json:"event_time"`
	MarketingConsent      *int    `json:"marketing_consent"`
	NewPrice              *string `json:"new_price"`
	NewQuantity           *string `json:"new_quantity"`
	NewUnitPrice          *string `json:"new_unit_price"`
	NextBillDate          *string `json:"next_bill_date"`
	OldPrice              *string `json:"old_price"`
	OldQuantity           *string `json:"old_quantity"`
	OldUnitPrice          *string `json:"old_unit_price"`
	Currency              *string `json:"currency"`
	Passthrough           *string `json:"passthrough"`
	Status                *string `json:"status"`
	SubscriptionID        *string `json:"subscription_id"`
	SubscriptionPlanID    *string `json:"subscription_plan_id"`
	UserID                *string `json:"user_id"`
	UpdateURL             *string `json:"update_url"`
	OldNextBillDate       *string `json:"old_next_bill_date"`
	OldStatus             *string `json:"old_status"`
	OldSubscriptionPlanID *string `json:"old_subscription_plan_id"`
	PausedAt              *string `json:"paused_at"`
	PausedFrom            *string `json:"paused_from"`
	PausedReason          *string `json:"paused_reason"`
}

// The subscription canceled alert is triggered whenever a user cancel a subscription
// Paddle Reference: https://developer.paddle.com/webhook-reference/subscription-alerts/subscription-cancelled
type SubscriptionCancelledAlert struct {
	AlertName                 *string `json:"alert_name"`
	AlertID                   *string `json:"alert_id"`
	CancellationEffectiveDate *string `json:"cancellation_effective_date"`
	CheckoutID                *string `json:"checkout_id"`
	Currency                  *string `json:"currency"`
	Email                     *string `json:"email"`
	EventTime                 *string `json:"event_time"`
	MarketingConsent          *int    `json:"marketing_consent"`
	Passthrough               *string `json:"passthrough"`
	Quantity                  *string `json:"quantity"`
	Status                    *string `json:"status"`
	SubscriptionID            *string `json:"subscription_id"`
	SubscriptionPlanID        *string `json:"subscription_plan_id"`
	UnitPrice                 *string `json:"unit_price"`
	UserID                    *string `json:"user_id"`
}

// Fired when a subscription payment is received successfully.
// Paddle reference: https://developer.paddle.com/webhook-reference/subscription-alerts/subscription-payment-succeeded
type SubscriptionPaymentSucceededAlert struct {
	AlertName             *string `json:"alert_name"`
	AlertID               *string `json:"alert_id"`
	BalanceCurrency       *string `json:"balance_currency"`
	BalanceEarnings       *string `json:"balance_earnings"`
	BalanceFee            *string `json:"balance_fee"`
	BalanceGross          *string `json:"balance_gross"`
	BalanceTax            *string `json:"balance_tax"`
	CheckoutID            *string `json:"checkout_id"`
	Country               *string `json:"country"`
	Coupon                *string `json:"coupon"`
	Currency              *string `json:"currency"`
	CustomerName          *string `json:"customer_name"`
	Earnings              *string `json:"earnings"`
	Email                 *string `json:"email"`
	EventTime             *string `json:"event_time"`
	Fee                   *string `json:"fee"`
	InitialPayment        *int    `json:"initial_payment"`
	Instalments           *string `json:"instalments"`
	MarketingConsent      *int    `json:"marketing_consent"`
	NextBillDate          *string `json:"next_bill_date"`
	NextPaymentAmount     *string `json:"next_payment_amount"`
	OrderID               *string `json:"order_id"`
	Passthrough           *string `json:"passthrough"`
	PaymentMethod         *string `json:"payment_method"`
	PaymentTax            *string `json:"payment_tax"`
	PlanName              *string `json:"plan_name"`
	Quantity              *string `json:"quantity"`
	ReceiptURL            *string `json:"receipt_url"`
	SaleGross             *string `json:"sale_gross"`
	Status                *string `json:"status"`
	SubscriptionID        *string `json:"subscription_id"`
	SubscriptionPaymentID *string `json:"subscription_payment_id"`
	SubscriptionPlanID    *string `json:"subscription_plan_id"`
	UnitPrice             *string `json:"unit_price"`
	UserID                *string `json:"user_id"`
}

// Fired when a payment for an existing subscription fails.
// Paddle reference! https://developer.paddle.com/webhook-reference/subscription-alerts/subscription-payment-failed
type SubscriptionPaymentFailedAlert struct {
	AlertName             *string `json:"alert_name"`
	AlertID               *string `json:"alert_id"`
	Amount                *string `json:"amount"`
	CancelURL             *string `json:"cancel_url"`
	CheckoutID            *string `json:"checkout_id"`
	Currency              *string `json:"currency"`
	Email                 *string `json:"email"`
	EventTime             *string `json:"event_time"`
	MarketingConsent      *int    `json:"marketing_consent"`
	NextRetryDate         *string `json:"next_retry_date"`
	Passthrough           *string `json:"passthrough"`
	Quantity              *string `json:"quantity"`
	Status                *string `json:"status"`
	SubscriptionID        *string `json:"subscription_id"`
	SubscriptionPlanID    *string `json:"subscription_plan_id"`
	UnitPrice             *string `json:"unit_price"`
	UpdateURL             *string `json:"update_url"`
	SubscriptionPaymentID *string `json:"subscription_payment_id"`
	Instalments           *string `json:"instalments"`
	OrderID               *string `json:"order_id"`
	UserID                *string `json:"user_id"`
	AttemptNumber         *string `json:"attempt_number"`
}

// Fired when a refund for an existing subscription payment is issued.
// Paddle reference: https://developer.paddle.com/webhook-reference/subscription-alerts/subscription-payment-refunded
type SubscriptionPaymentRefundedAlert struct {
	AlertName               *string `json:"alert_name"`
	AlertID                 *string `json:"alert_id"`
	Amount                  *string `json:"amount"`
	BalanceCurrency         *string `json:"balance_currency"`
	BalanceEarningsDecrease *string `json:"balance_earnings_decrease"`
	BalanceFeeRefund        *string `json:"balance_fee_refund"`
	BalanceGrossRefund      *string `json:"balance_gross_refund"`
	BalanceTaxRefund        *string `json:"balance_tax_refund"`
	CheckoutID              *string `json:"checkout_id"`
	Currency                *string `json:"currency"`
	EarningsDecrease        *string `json:"earnings_decrease"`
	Email                   *string `json:"email"`
	EventTime               *string `json:"event_time"`
	FeeRefund               *string `json:"fee_refund"`
	GrossRefund             *string `json:"gross_refund"`
	InitialPayment          *int    `json:"initial_payment"`
	Instalments             *string `json:"instalments"`
	MarketingConsent        *int    `json:"marketing_consent"`
	OrderID                 *string `json:"order_id"`
	Passthrough             *string `json:"passthrough"`
	Quantity                *string `json:"quantity"`
	RefundReason            *string `json:"refund_reason"`
	RefundType              *string `json:"refund_type"`
	Status                  *string `json:"status"`
	SubscriptionID          *string `json:"subscription_id"`
	SubscriptionPaymentID   *string `json:"subscription_payment_id"`
	SubscriptionPlanID      *string `json:"subscription_plan_id"`
	TaxRefund               *string `json:"tax_refund"`
	UnitPrice               *string `json:"unit_price"`
	UserID                  *string `json:"user_id"`
}

// Fired when a payment is made into your Paddle account.
// Paddle reference: https://developer.paddle.com/webhook-reference/one-off-purchase-alerts/payment-succeeded
type PaymentSucceededAlert struct {
	AlertName         *string `json:"alert_name"`
	AlertID           *string `json:"alert_id"`
	BalanceCurrency   *string `json:"balance_currency"`
	BalanceEarnings   *string `json:"balance_earnings"`
	BalanceFee        *string `json:"balance_fee"`
	BalanceGross      *string `json:"balance_gross"`
	BalanceTax        *string `json:"balance_tax"`
	CheckoutID        *string `json:"checkout_id"`
	Country           *string `json:"country"`
	Coupon            *string `json:"coupon"`
	Currency          *string `json:"currency"`
	CustomerName      *string `json:"customer_name"`
	Earnings          *string `json:"earnings"`
	Email             *string `json:"email"`
	EventTime         *string `json:"event_time"`
	Fee               *string `json:"fee"`
	IP                *string `json:"ip"`
	MarketingConsent  *int    `json:"marketing_consent"`
	OrderID           *string `json:"order_id"`
	Passthrough       *string `json:"passthrough"`
	PaymentMethod     *string `json:"payment_method"`
	PaymentTax        *string `json:"payment_tax"`
	ProductID         *string `json:"product_id"`
	ProductName       *string `json:"product_name"`
	Quantity          *string `json:"quantity"`
	ReceiptURL        *string `json:"receipt_url"`
	SaleGross         *string `json:"sale_gross"`
	UsedPriceOverride *string `json:"used_price_override"`
}

// Fired when a payment is refunded.
// Paddle reference: https://developer.paddle.com/webhook-reference/one-off-purchase-alerts/payment-refunded
type PaymentRefundedAlert struct {
	AlertName               *string `json:"alert_name"`
	AlertID                 *string `json:"alert_id"`
	Amount                  *string `json:"amount"`
	BalanceCurrency         *string `json:"balance_currency"`
	BalanceEarningsDecrease *string `json:"balance_earnings_decrease"`
	BalanceFeeRefund        *string `json:"balance_fee_refund"`
	BalanceGrossRefund      *string `json:"balance_gross_refund"`
	BalanceTaxRefund        *string `json:"balance_tax_refund"`
	CheckoutID              *string `json:"checkout_id"`
	Currency                *string `json:"currency"`
	EarningsDecrease        *string `json:"earnings_decrease"`
	Email                   *string `json:"email"`
	EventTime               *string `json:"event_time"`
	FeeRefund               *string `json:"fee_refund"`
	GrossRefund             *string `json:"gross_refund"`
	MarketingConsent        *int    `json:"marketing_consent"`
	OrderID                 *string `json:"order_id"`
	Passthrough             *string `json:"passthrough"`
	Quantity                *string `json:"quantity"`
	RefundReason            *string `json:"refund_reason"`
	RefundType              *string `json:"refund_type"`
	TaxRefund               *string `json:"tax_refund"`
}

// Fired when an order is created after a successful payment event.
// Paddle reference: https://developer.paddle.com/webhook-reference/one-off-purchase-alerts/order-processing-completed
type LockerProcessedAlert struct {
	AlertName        *string `json:"alert_name"`
	AlertID          *string `json:"alert_id"`
	CheckoutID       *string `json:"checkout_id"`
	CheckoutRecovery *string `json:"checkout_recovery"`
	Coupon           *string `json:"coupon"`
	Download         *string `json:"download"`
	Email            *string `json:"email"`
	EventTime        *string `json:"event_time"`
	Instructions     *string `json:"instructions"`
	Licence          *string `json:"licence"`
	MarketingConsent *int    `json:"marketing_consent"`
	OrderID          *string `json:"order_id"`
	ProductID        *string `json:"product_id"`
	Quantity         *string `json:"quantity"`
	Source           *string `json:"source"`
}

// Fired when a dispute/chargeback is raised for a card transaction.
// Paddle reference: https://developer.paddle.com/webhook-reference/risk-dispute-alerts/payment-dispute-created
type PaymentDisputeCreatedAlert struct {
	AlertName        *string `json:"alert_name"`
	AlertID          *string `json:"alert_id"`
	Amount           *string `json:"amount"`
	CheckoutID       *string `json:"checkout_id"`
	Currency         *string `json:"currency"`
	Email            *string `json:"email"`
	EventTime        *string `json:"event_time"`
	FeeUsd           *string `json:"fee_usd"`
	MarketingConsent *int    `json:"marketing_consent"`
	OrderID          *string `json:"order_id"`
	Passthrough      *string `json:"passthrough"`
	Status           *string `json:"status"`
}

// Fired when a dispute/chargeback is closed for a card transaction. This indicates that the dispute/chargeback was contested and won by Paddle.
// Paddle reference: https://developer.paddle.com/webhook-reference/risk-dispute-alerts/payment-dispute-closed
type PaymentDisputeClosedAlert struct {
	AlertName        *string `json:"alert_name"`
	AlertID          *string `json:"alert_id"`
	Amount           *string `json:"amount"`
	CheckoutID       *string `json:"checkout_id"`
	Currency         *string `json:"currency"`
	Email            *string `json:"email"`
	EventTime        *string `json:"event_time"`
	FeeUsd           *string `json:"fee_usd"`
	MarketingConsent *int    `json:"marketing_consent"`
	OrderID          *string `json:"order_id"`
	Passthrough      *string `json:"passthrough"`
	Status           *string `json:"status"`
}

// Fired when a transaction is flagged as high risk.
// Paddle reference: https://developer.paddle.com/webhook-reference/risk-dispute-alerts/high-risk-transaction-created
type HighRiskTransactionCreatedAlert struct {
	AlertName            *string `json:"alert_name"`
	AlertID              *string `json:"alert_id"`
	CaseID               *string `json:"case_id"`
	CheckoutID           *string `json:"checkout_id"`
	CreatedAt            *string `json:"created_at"`
	CustomerEmailAddress *string `json:"customer_email_address"`
	CustomerUserID       *string `json:"customer_user_id"`
	EventTime            *string `json:"event_time"`
	MarketingConsent     *int    `json:"marketing_consent"`
	Passthrough          *string `json:"passthrough"`
	ProductID            *string `json:"product_id"`
	RiskScore            *string `json:"risk_score"`
	Status               *string `json:"status"`
}

// Fired when a flagged transaction is approved or rejected.
// Paddle reference: https://developer.paddle.com/webhook-reference/risk-dispute-alerts/high-risk-transaction-updated
type HighRiskTransactionUpdatedAlert struct {
	AlertName            *string `json:"alert_name"`
	AlertID              *string `json:"alert_id"`
	CaseID               *string `json:"case_id"`
	CheckoutID           *string `json:"checkout_id"`
	CreatedAt            *string `json:"created_at"`
	CustomerEmailAddress *string `json:"customer_email_address"`
	CustomerUserID       *string `json:"customer_user_id"`
	EventTime            *string `json:"event_time"`
	MarketingConsent     *int    `json:"marketing_consent"`
	OrderID              *string `json:"order_id"`
	Passthrough          *string `json:"passthrough"`
	ProductID            *string `json:"product_id"`
	RiskScore            *string `json:"risk_score"`
}

// Fired when a new transfer/payout is created for your account.
// Paddle reference: https://developer.paddle.com/webhook-reference/payout-alerts/transfer-created
type TransferCreatedAlert struct {
	AlertName *string `json:"alert_name"`
	AlertID   *string `json:"alert_id"`
	Amount    *string `json:"amount"`
	Currency  *string `json:"currency"`
	EventTime *string `json:"event_time"`
	PayoutID  *string `json:"payout_id"`
	Status    *string `json:"status"`
}

// Fired when a new transfer/payout is marked as paid for your account.
// Paddle reference: https://developer.paddle.com/webhook-reference/payout-alerts/transfer-paid
type TransferPaidAlert struct {
	AlertName *string `json:"alert_name"`
	AlertID   *string `json:"alert_id"`
	Amount    *string `json:"amount"`
	Currency  *string `json:"currency"`
	EventTime *string `json:"event_time"`
	PayoutID  *string `json:"payout_id"`
	Status    *string `json:"status"`
}

// Fired when a customer opts in to receive marketing communication from you.
// Paddle reference: https://developer.paddle.com/webhook-reference/audience-alerts/new-audience-member
type NewAudienceMemberAlert struct {
	AlertName        *string `json:"alert_name"`
	AlertID          *string `json:"alert_id"`
	CreatedAt        *string `json:"created_at"`
	Email            *string `json:"email"`
	EventTime        *string `json:"event_time"`
	MarketingConsent *int    `json:"marketing_consent"`
	Products         *string `json:"products"`
	Source           *string `json:"source"`
	Subscribed       *string `json:"subscribed"`
	UserID           *string `json:"user_id"`
}

// Fired when the information of an audience member is updated.
// Paddle reference: https://developer.paddle.com/webhook-reference/audience-alerts/update-audience-member
type UpdateAudienceMemberAlert struct {
	AlertName           *string `json:"alert_name"`
	AlertID             *string `json:"alert_id"`
	EventTime           *string `json:"event_time"`
	NewCustomerEmail    *string `json:"new_customer_email"`
	NewMarketingConsent *string `json:"new_marketing_consent"`
	OldCustomerEmail    *string `json:"old_customer_email"`
	OldMarketingConsent *string `json:"old_marketing_consent"`
	Products            *string `json:"products"`
	Source              *string `json:"source"`
	UpdatedAt           *string `json:"updated_at"`
	UserID              *string `json:"user_id"`
}

// Fired when a manual invoice has been successfully paid by a customer.
// Paddle reference: https://developer.paddle.com/webhook-reference/manual-invoicing-alerts/invoice-paid
type InvoicePaidAlert struct {
	AlertName                    *string `json:"alert_name"`
	AlertID                      *string `json:"alert_id"`
	PaymentID                    *string `json:"payment_id"`
	Amount                       *string `json:"amount"`
	SaleGross                    *string `json:"sale_gross"`
	TermDays                     *string `json:"term_days"`
	Status                       *string `json:"status"`
	PurchaseOrderNumber          *string `json:"purchase_order_number"`
	InvoicedAt                   *string `json:"invoiced_at"`
	Currency                     *string `json:"currency"`
	ProductID                    *string `json:"product_id"`
	ProductName                  *string `json:"product_name"`
	ProductAdditionalInformation *string `json:"product_additional_information"`
	CustomerID                   *string `json:"customer_id"`
	CustomerName                 *string `json:"customer_name"`
	Email                        *string `json:"email"`
	CustomerVatNumber            *string `json:"customer_vat_number"`
	CustomerCompanyNumber        *string `json:"customer_company_number"`
	CustomerAddress              *string `json:"customer_address"`
	CustomerCity                 *string `json:"customer_city"`
	CustomerState                *string `json:"customer_state"`
	CustomerZipcode              *string `json:"customer_zipcode"`
	Country                      *string `json:"country"`
	ContractID                   *string `json:"contract_id"`
	ContractStartDate            *string `json:"contract_start_date"`
	ContractEndDate              *string `json:"contract_end_date"`
	Passthrough                  *string `json:"passthrough"`
	DateCreated                  *string `json:"date_created"`
	BalanceCurrency              *string `json:"balance_currency"`
	PaymentTax                   *string `json:"payment_tax"`
	PaymentMethod                *string `json:"payment_method"`
	Fee                          *string `json:"fee"`
	Earnings                     *string `json:"earnings"`
	BalanceEarnings              *string `json:"balance_earnings"`
	BalanceFee                   *string `json:"balance_fee"`
	BalanceTax                   *string `json:"balance_tax"`
	BalanceGross                 *string `json:"balance_gross"`
	DateReconciled               *string `json:"date_reconciled"`
	EventTime                    *string `json:"event_time"`
}

// Fired when a manual invoice has been successfully sent to a customer.
// Paddle reference: https://developer.paddle.com/webhook-reference/manual-invoicing-alerts/invoice-sent
type InvoiceSentAlert struct {
	AlertName                    *string `json:"alert_name"`
	AlertID                      *string `json:"alert_id"`
	PaymentID                    *string `json:"payment_id"`
	Amount                       *string `json:"amount"`
	SaleGross                    *string `json:"sale_gross"`
	TermDays                     *string `json:"term_days"`
	Status                       *string `json:"status"`
	PurchaseOrderNumber          *string `json:"purchase_order_number"`
	InvoicedAt                   *string `json:"invoiced_at"`
	Currency                     *string `json:"currency"`
	ProductID                    *string `json:"product_id"`
	ProductName                  *string `json:"product_name"`
	ProductAdditionalInformation *string `json:"product_additional_information"`
	CustomerID                   *string `json:"customer_id"`
	CustomerName                 *string `json:"customer_name"`
	Email                        *string `json:"email"`
	CustomerVatNumber            *string `json:"customer_vat_number"`
	CustomerCompanyNumber        *string `json:"customer_company_number"`
	CustomerAddress              *string `json:"customer_address"`
	CustomerCity                 *string `json:"customer_city"`
	CustomerState                *string `json:"customer_state"`
	CustomerZipcode              *string `json:"customer_zipcode"`
	Country                      *string `json:"country"`
	ContractID                   *string `json:"contract_id"`
	ContractStartDate            *string `json:"contract_start_date"`
	ContractEndDate              *string `json:"contract_end_date"`
	Passthrough                  *string `json:"passthrough"`
	DateCreated                  *string `json:"date_created"`
	BalanceCurrency              *string `json:"balance_currency"`
	PaymentTax                   *string `json:"payment_tax"`
	PaymentMethod                *string `json:"payment_method"`
	Fee                          *string `json:"fee"`
	Earnings                     *string `json:"earnings"`
	EventTime                    *string `json:"event_time"`
}

// Fired when a manual invoice has exceeded the payment term and is now overdue.
// Paddle reference: https://developer.paddle.com/webhook-reference/manual-invoicing-alerts/invoice-overdue
type InvoiceOverdueAlert struct {
	AlertName                    *string `json:"alert_name"`
	AlertID                      *string `json:"alert_id"`
	PaymentID                    *string `json:"payment_id"`
	Amount                       *string `json:"amount"`
	SaleGross                    *string `json:"sale_gross"`
	TermDays                     *string `json:"term_days"`
	Status                       *string `json:"status"`
	PurchaseOrderNumber          *string `json:"purchase_order_number"`
	InvoicedAt                   *string `json:"invoiced_at"`
	Currency                     *string `json:"currency"`
	ProductID                    *string `json:"product_id"`
	ProductName                  *string `json:"product_name"`
	ProductAdditionalInformation *string `json:"product_additional_information"`
	CustomerID                   *string `json:"customer_id"`
	CustomerName                 *string `json:"customer_name"`
	Email                        *string `json:"email"`
	CustomerVatNumber            *string `json:"customer_vat_number"`
	CustomerCompanyNumber        *string `json:"customer_company_number"`
	CustomerAddress              *string `json:"customer_address"`
	CustomerCity                 *string `json:"customer_city"`
	CustomerState                *string `json:"customer_state"`
	CustomerZipcode              *string `json:"customer_zipcode"`
	Country                      *string `json:"country"`
	ContractID                   *string `json:"contract_id"`
	ContractStartDate            *string `json:"contract_start_date"`
	ContractEndDate              *string `json:"contract_end_date"`
	Passthrough                  *string `json:"passthrough"`
	DateCreated                  *string `json:"date_created"`
	BalanceCurrency              *string `json:"balance_currency"`
	PaymentTax                   *string `json:"payment_tax"`
	PaymentMethod                *string `json:"payment_method"`
	Fee                          *string `json:"fee"`
	Earnings                     *string `json:"earnings"`
	EventTime                    *string `json:"event_time"`
}
