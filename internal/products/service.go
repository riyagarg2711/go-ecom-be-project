package products

import (
	"context"

	repo "github.com/riyagarg2711/ecom-api-course/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product,error)
}


type svc struct {
	repo repo.Querier

}
func NewService(repo repo.Querier) Service{
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product,error) {
return s.repo.ListProducts(ctx)
	
}