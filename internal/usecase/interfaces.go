package usecase

import (
	"context"
	"github.com/hexhoc/go-mall-api/internal/entity"
)

type (
	// Goods -.
	Goods interface {
		GetGoodsList(context.Context) ([]entity.Goods, error)
	}

	// GoodsRepo -.
	GoodsRepo interface {
		FindAll(context.Context) ([]entity.Goods, error)
	}
)
