package user

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"mongo-transactions/internal/journal"
)

type Repo interface {
	Get(ctx context.Context, id int32) (*User, error)
	Upsert(ctx context.Context, user User) error
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}

type JournalService interface {
	Record(ctx context.Context, record journal.Record) error
}

type Service struct {
	repo    Repo
	journal JournalService
}

func NewService(repo Repo, journalService JournalService) *Service {
	return &Service{
		repo:    repo,
		journal: journalService,
	}
}

func (s *Service) Get(ctx context.Context, id int32) (*User, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) Upsert(ctx context.Context, user User) error {
	err := s.repo.InTransaction(ctx, func(ctx context.Context) error {
		if err := s.repo.Upsert(ctx, user); err != nil {
			return errors.Wrap(err, "upsert user in service")
		}

		rec := journal.Record{
			UserID:    user.ID,
			Timestamp: time.Now(),
		}

		if err := s.journal.Record(ctx, rec); err != nil {
			return errors.Wrap(err, "store record in service")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "upsert user in transaction")
	}

	return nil
}
