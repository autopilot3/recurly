package recurly

import "encoding/xml"

type Usage struct {
	XMLName           xml.Name   `xml:"usage"`
	AddOnCode         string     `xml:"add_on_code"`
	SubscriptionID    string     `xml:"subscription_id"`
	ID                string     `xml:"id"`
	UnitAmountInCents NullInt    `xml:"unit_amount_in_cents,omitempty"`
	Amount            UnitAmount `xml:"amount,omitempty"`
	UsageType         string     `xml:"usage_type,omitempty"`
	CreatedAt         NullTime   `xml:"created_at,omitempty"`
	ModifiedAt        NullTime   `xml:"modified_at,omitempty"`
	BilledAt          NullTime   `xml:"billed_at,omitempty"`
}
