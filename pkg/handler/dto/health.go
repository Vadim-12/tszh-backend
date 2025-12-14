package dto

type HealthResponseDto struct {
	Status string `json:"status"`
	Db     string `json:"db"`
}
