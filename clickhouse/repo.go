package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	addr          = "localhost:9000"
	logsTableName = "my_logs"
	logsDBName    = "my_logs"
	modTypeAdd    = "add"
	modTypeChange = "change"
	modTypeDelete = "delete"
)

type Repo struct {
	db clickhouse.Conn
}

type logDoc struct {
	EventID   uuid.UUID `ch:"event_id"`
	ProductID uuid.UUID `ch:"product_id"`
	Timestamp time.Time `ch:"timestamp"`
	CreatedBy string    `ch:"created_by"`
	Price     int32     `ch:"price"`
	Lifetime  time.Time `ch:"lifetime"`
}

func NewRepo(ctx context.Context) (*Repo, error) {
	db, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{addr},
	})
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return &Repo{db: db}, nil
}

func (r *Repo) InsertBatch(ctx context.Context, logDocs []logDoc) error {
	query := fmt.Sprintf("insert into %s.%s", logsDBName, logsTableName)

	batch, err := r.db.PrepareBatch(ctx, query)
	if err != nil {
		return errors.Wrap(err, "prepare logs batch")
	}

	for _, log := range logDocs {
		if err = batch.AppendStruct(&log); err != nil {
			return errors.Wrap(err, "append log doc to batch")
		}
	}

	if err = batch.Send(); err != nil {
		return errors.Wrap(err, "writing batch")
	}

	return nil
}
