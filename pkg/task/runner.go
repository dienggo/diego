package task

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

// Runner to run task scheduler
func Runner(task ITask, ctx context.Context, isBusy bool) {
	taskSchedule := task.Task()
	taskSchedule.isBusy = isBusy
	duration := taskSchedule.duration
	err := durationMultipleEntity(*duration)
	if err != nil {
		log.Error("Error on durationMultipleEntity ", err.Error())
		return
	}
	if durationIsAllNil(*duration) {
		log.Error("Ups on durationIsAllNil == true")
		return
	}

	if taskSchedule.name == "" {
		log.Error("Ups on taskSchedule name can not be empty")
		return
	}

	if taskSchedule.overlaps == false && taskSchedule.isBusy {
		log.Info(taskSchedule.name+" : ", "Overlaps is false and task is busy, waiting task done and will continue on next time")
		return
	}

	run(duration,
		// callback runner
		func() {
			taskSchedule.isBusy = true
			// log.Debugln("Debug Task", taskSchedule.name, "Run on", time.Now().String())
			if taskSchedule.overlaps {
				go taskSchedule.Run(ctx)
			} else {
				taskSchedule.Run(ctx)
			}
			taskSchedule.isBusy = false
		},
		// callback recursive
		func(hour int) {
			// on sleep method
			if duration.SleepDuration != nil && hour == 0 {
				time.Sleep(*duration.SleepDuration)
				Runner(task, ctx, taskSchedule.isBusy)
			} else {
				onUpdatedHour(hour, func() {
					Runner(task, ctx, taskSchedule.isBusy)
				})
			}
		})
}

// run is method switcher to run callback
func run(duration *Duration, callback func(), recursive func(hour int)) {
	hour := 0

	// on sleep method
	if duration.SleepDuration != nil {
		callback()
	}

	// on every day at method
	if duration.EveryDayAt != nil {
		dr := *duration.EveryDayAt
		if dr.TimeAt == time.Now().Hour() {
			hour = dr.TimeAt
			callback()
		}
	}

	// on every week at, day at method
	if duration.EveryWeekAt != nil {
		dr := *duration.EveryWeekAt
		if dr.TimeAt == time.Now().Hour() && dr.DayAt == time.Now().Day() {
			hour = dr.TimeAt
			callback()
		}
	}

	// on every month at, week at, day at method
	if duration.EveryMonthAt != nil {
		dr := *duration.EveryMonthAt
		_, currentWeek := time.Now().ISOWeek()
		if dr.TimeAt == time.Now().Hour() && dr.DayAt == time.Now().Day() && dr.WeekAt == currentWeek {
			hour = dr.TimeAt
			callback()
		}
	}

	// on every year at, month at, week at, day at method
	if duration.EveryYearAt != nil {
		dr := *duration.EveryYearAt
		_, currentWeek := time.Now().ISOWeek()
		if dr.TimeAt == time.Now().Hour() && dr.DayAt == time.Now().Day() && dr.WeekAt == currentWeek && dr.MonthAt == int(time.Now().Month()) {
			hour = dr.TimeAt
			callback()
		}
	}

	// call recursive callback
	recursive(hour)
}

// onUpdatedHour is method to stream updated hour
func onUpdatedHour(hour int, call func()) {
	if hour != time.Now().Hour() {
		call()
	} else {
		time.Sleep(1 * time.Second)
		onUpdatedHour(hour, call)
	}
}
