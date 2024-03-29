package impl

import (
	"context"

	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewBmiRepositoryImpl(DB *gorm.DB) repository.BmiRepository {
	return &bmiRepositoryImpl{DB: DB}
}

type bmiRepositoryImpl struct {
	*gorm.DB
}

func (repository *bmiRepositoryImpl) Insert(ctx context.Context, product entity.Bmi) entity.Bmi {
	product.Id = uuid.New()
	err := repository.DB.WithContext(ctx).Create(&product).Error
	exception.PanicLogging(err)
	return product
}

func (repository *bmiRepositoryImpl) FindAl(ctx context.Context) []entity.Bmi {
	var bmi []entity.Bmi
	repository.DB.WithContext(ctx).Find(&bmi)
	return bmi
}
