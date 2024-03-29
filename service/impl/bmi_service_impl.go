package impl

import (
	"context"
	"math"

	"github.com/RizkiMufrizal/gofiber-clean-architecture/common"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/go-redis/redis/v9"
)

func NewBmiServiceImpl(bmiRepository *repository.BmiRepository, cache *redis.Client) service.BmiService {
	return &bmiServiceImpl{BmiRepository: *bmiRepository, Cache: cache}
}

type bmiServiceImpl struct {
	repository.BmiRepository
	Cache *redis.Client
}

func (service *bmiServiceImpl) Calculator(ctx context.Context, bmiModel model.BmiCreateModel) model.BmiModel {
	common.Validate(bmiModel)
	calculate := float64(int((bmiModel.Kg/(math.Pow(bmiModel.M, 2)))*100)) / 100
	var description string
	if calculate < 18.50 {
		description = "น้ำหนักน้อยกว่าเกณฑ์มาตรฐาน"
	} else if calculate >= 18.5 && calculate <= 24.9 {
		description = "น้ำหนักปกติ"
	} else if calculate >= 25.0 && calculate <= 29.9 {
		description = "น้ำหนักเกิน"
	} else {
		description = "อ้วน"
	}
	bmi := entity.Bmi{
		Kg:          bmiModel.Kg,
		M:           bmiModel.M,
		Bmi:         calculate,
		Description: description,
	}
	bmi = service.BmiRepository.Insert(ctx, bmi)
	return model.BmiModel{
		Id:          bmi.Id.String(),
		Kg:          bmiModel.Kg,
		M:           bmiModel.M,
		Bmi:         calculate,
		Description: description,
	}
}
func (service *bmiServiceImpl) FindAll(ctx context.Context) (responses []model.BmiModel) {
	bmis := service.BmiRepository.FindAl(ctx)
	for _, bmi := range bmis {
		responses = append(responses, model.BmiModel{
			Id:          bmi.Id.String(),
			Kg:          bmi.Kg,
			M:           bmi.M,
			Description: bmi.Description,
		})
	}
	if len(bmis) == 0 {
		return []model.BmiModel{}
	}
	return responses
}
