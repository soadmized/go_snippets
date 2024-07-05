package journal

import "context"

type Repo interface {
	Insert(ctx context.Context, rec Record) error
}

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{repo: repo}
}

func (s Service) Record(ctx context.Context, rec Record) error {
	return s.repo.Insert(ctx, rec)
}
