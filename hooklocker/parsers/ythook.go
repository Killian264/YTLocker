package parsers

import (
	"encoding/xml"
	"time"

	"github.com/Killian264/YTLocker/hooklocker/models"
)

// ParseYTHook parses the xml data sent with YT Subscription webhooks
func ParseYTHook(hookXML string) (models.YTHookPush, error) {
	var hook models.YTHookPush
	err := xml.Unmarshal([]byte(hookXML), &hook)
	if err != nil {
		return models.YTHookPush{}, nil
	}
	hook.Received = time.Now()
	return hook, nil
}
