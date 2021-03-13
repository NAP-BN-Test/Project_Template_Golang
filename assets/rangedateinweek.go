package assets

import (
	"time"
)

func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func WeekRange(year, week int) (start, end string) {
	start = WeekStart(year, week).Format("2006-01-02")
	end = WeekStart(year, week).AddDate(0, 0, 6).Format("2006-01-02")
	return
}

func WeekYear() (year, week int) {
	tn := time.Now().UTC()
	year, week = tn.ISOWeek()
	ts := time.Now().UTC().Unix()
	tn = time.Unix(ts, 0)
	year, week = tn.ISOWeek()
	return
}
