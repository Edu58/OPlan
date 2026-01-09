package time

import (
	"fmt"
	"time"
)

// Format date and start time to (March 15, 2025 • 9:00 AM - 6:00 PM)
func FormatEventTime(startTime, endTime time.Time) string {
	date := startTime.Format("January 2, 2006")
	start := startTime.Format("3:04 PM")
	end := endTime.Format("3:04 PM")

	return fmt.Sprintf("%s • %s - %s", date, start, end)
}
