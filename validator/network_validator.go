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
	)
}
