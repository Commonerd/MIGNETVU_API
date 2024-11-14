package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type INetworkValidator interface {
	NetworkValidate(network model.Network) error
}

type networkValidator struct{}

func NewNetworkValidator() INetworkValidator {
	return &networkValidator{}
}

func (tv *networkValidator) NetworkValidate(network model.Network) error {
	return validation.ValidateStruct(&network,
		validation.Field(
			&network.Title,
			validation.Required.Error("name is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		// validation.Field(
		// 	&network.MigrationYear,
		// 	validation.Required.Error("migration year is required"),
		// 	validation.Max(time.Now()).Error("limited max is this year"),
		// ),
		validation.Field(
			&network.Latitude,
			validation.Required.Error("latitude is required"),
			validation.Min(-90.0).Error("latitude must be >= -90"),
			validation.Max(90.0).Error("latitude must be <= 90"),
		),
		validation.Field(
			&network.Longitude,
			validation.Required.Error("longitude is required"),
			validation.Min(-180.0).Error("longitude must be >= -180"),
			validation.Max(180.0).Error("longitude must be <= 180"),
		),
	)
}
