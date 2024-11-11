package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type INetworkUsecase interface {
	GetAllNetworks(userId uint) ([]model.NetworkResponse, error)
	GetNetworkById(userId uint, networkId uint) (model.NetworkResponse, error)
	CreateNetwork(network model.Network) (model.NetworkResponse, error)
	UpdateNetwork(network model.Network, userId uint, networkId uint) (model.NetworkResponse, error)
	DeleteNetwork(userId uint, networkId uint) error
}

type networkUsecase struct {
	tr repository.INetworkRepository
	tv validator.INetworkValidator
}

func NewNetworkUsecase(tr repository.INetworkRepository, tv validator.INetworkValidator) INetworkUsecase {
	return &networkUsecase{tr, tv}
}

func (tu *networkUsecase) GetAllNetworks(userId uint) ([]model.NetworkResponse, error) {
	networks := []model.Network{}
	if err := tu.tr.GetAllNetworks(&networks, userId); err != nil {
		return nil, err
	}
	resNetworks := []model.NetworkResponse{}
	for _, v := range networks {
		t := model.NetworkResponse{
			ID:        v.ID,
			Title:     v.Title,
			Type:     v.Type,
			Nationality:     v.Nationality,
			Ethnicity: v.Ethnicity,
			Latitude:    v.Latitude,
			Longitude:   v.Longitude,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resNetworks = append(resNetworks, t)
	}
	return resNetworks, nil
}

func (tu *networkUsecase) GetNetworkById(userId uint, networkId uint) (model.NetworkResponse, error) {
	network := model.Network{}
	if err := tu.tr.GetNetworkById(&network, userId, networkId); err != nil {
		return model.NetworkResponse{}, err
	}
	resNetwork := model.NetworkResponse{
		ID:        network.ID,
		Title:     network.Title,
		Type:     network.Type,
		Nationality:     network.Nationality,
		Ethnicity: network.Ethnicity,
		Latitude:    network.Latitude,
		Longitude:   network.Longitude,
		CreatedAt: network.CreatedAt,
		UpdatedAt: network.UpdatedAt,
	}
	return resNetwork, nil
}

func (tu *networkUsecase) CreateNetwork(network model.Network) (model.NetworkResponse, error) {
	if err := tu.tv.NetworkValidate(network); err != nil {
		return model.NetworkResponse{}, err
	}
	if err := tu.tr.CreateNetwork(&network); err != nil {
		return model.NetworkResponse{}, err
	}
	resNetwork := model.NetworkResponse{
		ID:        network.ID,
		Title:     network.Title,
		Type:     network.Type,
		Nationality:     network.Nationality,
		Ethnicity:     network.Ethnicity,
		Latitude:    network.Latitude,
		Longitude:   network.Longitude,
		CreatedAt: network.CreatedAt,
		UpdatedAt: network.UpdatedAt,
	}
	return resNetwork, nil
}

func (tu *networkUsecase) UpdateNetwork(network model.Network, userId uint, networkId uint) (model.NetworkResponse, error) {
	if err := tu.tv.NetworkValidate(network); err != nil {
		return model.NetworkResponse{}, err
	}
	if err := tu.tr.UpdateNetwork(&network, userId, networkId); err != nil {
		return model.NetworkResponse{}, err
	}
	resNetwork := model.NetworkResponse{
		ID:        network.ID,
		Title:     network.Title,
		Type:     network.Type,
		Nationality:     network.Nationality,
		Ethnicity:     network.Ethnicity,
		Latitude:    network.Latitude,
		Longitude:   network.Longitude,
		CreatedAt: network.CreatedAt,
		UpdatedAt: network.UpdatedAt,
	}
	return resNetwork, nil
}

func (tu *networkUsecase) DeleteNetwork(userId uint, networkId uint) error {
	if err := tu.tr.DeleteNetwork(userId, networkId); err != nil {
		return err
	}
	return nil
}
