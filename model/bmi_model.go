package model

type BmiModel struct {
	Id          string  `json:"id"`
	Kg          float64 `json:"kg"`
	M           float64 `json:"m"`
	Bmi         float64 `json:"bmi"`
	Description string  `json:"description"`
}

type BmiCreateModel struct {
	Kg float64 `json:"kg"`
	M  float64 `json:"m"`
}
