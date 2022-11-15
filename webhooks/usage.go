package webhooks

import "github.com/autopilot3/recurly"

const (
	NewUsage = "new_usage_notification"
)

type UsageNotification struct {
	Type    string        `xml:"-"`
	Account Account       `xml:"account"`
	Usage   recurly.Usage `xml:"usage"`
}
