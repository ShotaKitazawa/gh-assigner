package googlecalendar

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"

	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
)

const (
	SREOrderStaffScheduleName = "SREO担当:"
)

// CalendarInfrastructure is Infrastructure
type CalendarInfrastructure struct {
	ID      string
	Service *calendar.Service
	Logger  interfaces.Logger
}

func (r CalendarInfrastructure) GetCurrentStaff() (staff string, err error) {
	start := time.Now()
	end := start.Add(time.Duration(1) * time.Second)
	events, err := r.Service.Events.List(r.ID).ShowDeleted(false).SingleEvents(true).TimeMin(start.Format(time.RFC3339)).TimeMax(end.Format(time.RFC3339)).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Unable to retrieve next ten of the user's events: %v", err))
		return
	}
	if len(events.Items) == 0 {
		err = fmt.Errorf("GetCurrentStaff: No upcoming events found.")
		return
	}
	for _, item := range events.Items {
		// Trim Space & Newline Charactor
		summary := item.Summary
		summary = strings.Replace(summary, " ", "", -1)
		summary = strings.Replace(summary, "\n", "", -1)
		summary = strings.Replace(summary, "\r", "", -1)
		summary = strings.Replace(summary, "\r\n", "", -1)

		// Check
		if strings.HasPrefix(item.Summary, SREOrderStaffScheduleName) {
			staff = strings.Split(item.Summary, ":")[1]
			return
		}
	}

	err = fmt.Errorf("GetCurrentStaff: No SREO staff schedule found.")
	return
}
