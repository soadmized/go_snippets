package main

import (
	"time"

	"github.com/google/uuid"
)

func generateLogs(count int) []logDoc {
	logs := make([]logDoc, 0, count)

	for range count {
		logs = append(logs, logDoc{
			EventID:        uuid.New(),
			ProductID:      uuid.New(),
			PromoID:        123,
			Timestamp:      time.Now(),
			CreatedBy:      "John Snow",
			ModType:        modTypeAdd,
			LocalityID:     uuid.New(),
			Price:          10000,
			MaxSoldCount:   1000,
			OrderLimit:     100,
			IsPriority:     true,
			IsUploadToFeed: true,
			Lifetime:       time.Now().Add(time.Hour),
		})
	}

	return logs
}
