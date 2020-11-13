package models

import (
	"gorm.io/gorm"
)

type PairSnapshot struct {
	Model

	Name       string  `json:"name"`
	PlatformId int     `json:"platform_id"`
	Apy             float64 `json:"apy"`
	LpTotalSupply   string `json:"lp_total_supply"`
	Token0Amount    string `json:"token0_amount"`
	Token1Amount    string `json:"token1_amount"`
}

func (PairSnapshot) TableName() string {
	return "pair_snapshot"
}

func AddPairSnapshot(data map[string]interface{}) error {
	pairSnapshot := PairSnapshot{
		Name:            data["name"].(string),
		PlatformId:      data["platformId"].(int),
		LpTotalSupply:   data["lpTotalSupply"].(string),
		Token0Amount:    data["token0Amount"].(string),
		Token1Amount:    data["token1Amount"].(string),
		Apy: data["apy"].(float64),
	}
	if err := db.Create(&pairSnapshot).Error; err != nil {
		return err
	}
	return nil
}

func GetPairSnapshot(maps interface{}) (*PairSnapshot, error) {
	var ps PairSnapshot
	err := db.Where(maps).Find(&ps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &ps, err
}
