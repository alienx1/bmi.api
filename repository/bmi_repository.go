package repository

import (
	"context"

	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
)

type BmiRepository interface {
	Insert(ctx context.Context, product entity.Bmi) entity.Bmi
	FindAl(ctx context.Context) []entity.Bmi
}
