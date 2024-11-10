package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type INetworkRepository interface {
	GetAllNetworks(networks *[]model.Network, userId uint) error
	GetNetworkById(network *model.Network, userId uint, networkId uint) error
	CreateNetwork(network *model.Network) error
	UpdateNetwork(network *model.Network, userId uint, networkId uint) error
	DeleteNetwork(userId uint, networkId uint) error
}

type networkRepository struct {
	db *gorm.DB
}

func NewNetworkRepository(db *gorm.DB) INetworkRepository {
	return &networkRepository{db}
}

func (tr *networkRepository) GetAllNetworks(networks *[]model.Network, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(networks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *networkRepository) GetNetworkById(network *model.Network, userId uint, networkId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(network, networkId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *networkRepository) CreateNetwork(network *model.Network) error {
	if err := tr.db.Create(network).Error; err != nil {
		return err
	}
	return nil
}

func (tr *networkRepository) UpdateNetwork(network *model.Network, userId uint, networkId uint) error {
	result := tr.db.Model(network).Clauses(clause.Returning{}).Where("id=? AND user_id=?", networkId, userId).Update("title", network.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *networkRepository) DeleteNetwork(userId uint, networkId uint) error {
	result := tr.db.Where("id=? AND user_id=?", networkId, userId).Delete(&model.Network{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
