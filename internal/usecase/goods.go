package usecase

import (
	"context"
	"fmt"
	"github.com/hexhoc/go-mall-api/internal/entity"
)

type GoodsUseCase struct {
	repo GoodsRepo
}

// New -.
func New(r GoodsRepo) *GoodsUseCase {
	return &GoodsUseCase{
		repo: r,
	}
}

// History - getting translate history from store.
func (uc *GoodsUseCase) FindAll(ctx context.Context) ([]entity.Goods, error) {
	translations, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("GoodsUseCase - History - s.repository.GetHistory: %w", err)
	}

	return translations, nil
}
