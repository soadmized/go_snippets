package main

import (
	"time"

	"github.com/google/uuid"
)

func generateLogs(count int) []logDoc {
	logs := make([]logDoc, 0, count)

	for range count {
		logs = append(logs, logDoc{
			EventID:   uuid.New(),
			ProductID: uuid.New(),
			Timestamp: time.Now(),
			CreatedBy: "John Snow",
			Price:     10000,
			Lifetime:  time.Now().Add(time.Hour),
		})
	}

	return logs
}
