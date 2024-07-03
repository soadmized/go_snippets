package main

import (
	"context"
	"log"
)

const (
	logsCount = 1000000
)

func main() {
	ctx := context.Background()

	repo, err := NewRepo(ctx)
	if err != nil {
		log.Print(err)

		return
	}

	logs := generateLogs(logsCount)

	err = repo.InsertBatch(ctx, logs)
	if err != nil {
		log.Print(err)

		return
	}
}
