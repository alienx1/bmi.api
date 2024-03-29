package service

import (
	"context"

	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
)

type BmiService interface {
	Calculator(ctx context.Context, model model.BmiCreateModel) model.BmiModel
	FindAll(ctx context.Context) (responses []model.BmiModel)
}
