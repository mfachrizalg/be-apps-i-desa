package dtos

type AddVillageRequest struct {
	Name string `json:"name" validate:"required"`
}
