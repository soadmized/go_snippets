package main

import (
	"context"
	"log"

	"mongo-transactions/cmd"
)

func main() {
	ctx := context.Background()

	err := cmd.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
