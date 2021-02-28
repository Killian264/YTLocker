package parsers

import (
	"encoding/xml"
	"time"

	"github.com/Killian264/YTLocker/models"
)

// ParseYTHook parses the xml data sent with YT Subscription webhooks
func ParseYTHook(hookXML string) models.YTHookPush {
	var hook models.YTHookPush
	xml.Unmarshal([]byte(hookXML), &hook)
	hook.Received = time.Now()
	return hook
}
