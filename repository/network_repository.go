package repository

import (
	"encoding/json"
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
	// 모든 네트워크 데이터를 조회
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(networks).Error; err != nil {
		return err
	}

	// 각 네트워크의 connections JSON 데이터를 언마샬링
	for i := range *networks {
		network := &(*networks)[i]
		if len(network.Connections) > 0 {
			// JSON 데이터를 구조체로 언마샬링
			if err := json.Unmarshal(network.Connections, &network.Connections); err != nil {
				return fmt.Errorf("failed to unmarshal connections for network ID %d: %w", network.ID, err)
			}
		}
	}

	return nil
}
func (tr *networkRepository) GetNetworkById(network *model.Network, userId uint, networkId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(network, networkId).Error; err != nil {
		return err
	}

	// connections JSON 데이터를 언마샬링하여 구조체로 변환
	if err := json.Unmarshal(network.Connections, &network.Connections); err != nil {
		return fmt.Errorf("failed to unmarshal connections: %w", err)
	}
	return nil
}


func (tr *networkRepository) CreateNetwork(network *model.Network) error {
	// connections 필드를 JSON으로 직렬화
	connectionsData, err := json.Marshal(network.Connections)
	if err != nil {
		return fmt.Errorf("failed to marshal connections: %w", err)
	}
	
	// JSON으로 직렬화된 connections 데이터를 구조체의 해당 필드에 할당
	network.Connections = connectionsData

	// 네트워크 생성
	if err := tr.db.Create(network).Error; err != nil {
		return err
	}
	return nil
}

func (tr *networkRepository) UpdateNetwork(network *model.Network, userId uint, networkId uint) error {
	// connections를 JSON으로 직렬화
	connectionsData, err := json.Marshal(network.Connections)
	if err != nil {
		return fmt.Errorf("failed to marshal connections: %w", err)
	}
	result := tr.db.Model(network).
		Clauses(clause.Returning{}).
		Where("id = ? AND user_id = ?", networkId, userId).
		Updates(map[string]interface{}{
			"title":       network.Title,
			"type":        network.Type,
			"nationality": network.Nationality,
			"ethnicity": network.Ethnicity,
			"latitude":    network.Latitude,
			"longitude":   network.Longitude,
			"connections": connectionsData, // JSON 데이터를 그대로 업데이트

		})

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
