package main

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	dsn := os.Getenv("SENTRY_DSN")

	err := sentry.Init(sentry.ClientOptions{
		Dsn:   dsn,
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	//monitorSchedule := sentry.IntervalSchedule(1, sentry.MonitorScheduleUnitMinute)
	//monitorConfig := &sentry.MonitorConfig{
	//	Schedule:      monitorSchedule,
	//	MaxRuntime:    2,
	//	CheckInMargin: 1,
	//}

	// NOTIFY START
	checkinId := sentry.CaptureCheckIn(
		&sentry.CheckIn{
			MonitorSlug: "scheduler",
			Status:      sentry.CheckInStatusInProgress,
		},
		nil,
	)

	// DO WORK
	log.Println("DOING WORK...")
	time.Sleep(10 * time.Second)

	// NOTIFY COMPLETE THE TASK
	sentry.CaptureCheckIn(
		&sentry.CheckIn{
			ID:          *checkinId,
			MonitorSlug: "scheduler",
			Status:      sentry.CheckInStatusOK,
		},
		nil,
	)

	//sentry.CaptureMessage("TEST CRON")

	defer sentry.Flush(4 * time.Second)
}
