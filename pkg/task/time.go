package task

import (
	"errors"
	"fmt"
	"github.com/dienggo/diego/pkg/helper"
	"strings"
	"time"
)

type Day struct {
	// TimeAt is time value on day [1 - 24] in hour
	TimeAt int
}

type Week struct {
	Day
	// DayAt is range day on week [1 - 7]
	DayAt int
}

type Month struct {
	Week
	// WeekAt is range week on month [1 - 3]
	WeekAt int
}

type Year struct {
	Month
	// MonthAt is range month on year [1 - 12]
	MonthAt int
}

type Duration struct {
	SleepDuration *time.Duration
	EveryDayAt    *Day
	EveryWeekAt   *Week
	EveryMonthAt  *Month
	EveryYearAt   *Year
}

// durationMultipleEntity is method to check all entity has multiple Duration stored
func durationMultipleEntity(d Duration) error {
	maps := helper.StructToMap(d)
	var notNilData []string
	for key, val := range maps {
		if val != nil && fmt.Sprint(val) != `<nil>` {
			notNilData = append(notNilData, key+":"+fmt.Sprint(val))
		}
	}

	if len(notNilData) > 1 {
		text := strings.Join(notNilData, ", ")
		return errors.New("Not nil entities : " + text)
	}
	return nil
}

// durationIsAllNil is method to check all entity
func durationIsAllNil(d Duration) bool {
	return d.SleepDuration == nil && d.EveryDayAt == nil && d.EveryWeekAt == nil && d.EveryYearAt == nil && d.EveryMonthAt == nil
}
