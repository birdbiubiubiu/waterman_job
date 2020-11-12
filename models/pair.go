package models

import (
	"gorm.io/gorm"
	"waterman_job/pkg/logging"
)

type Pair struct {
	Model

	Name            string  `json:"name"`
	PlatformId      int     `json:"platform_id"`
	Liquidity       float64 `json:"liquidity"`
	Volume24h       float64 `json:"volume_24h"`
	Volume7d        float64 `json:"volume_7d"`
	Fee24h          float64 `json:"fee_24h"`
	Apy             float64 `json:"apy"`
	LpTotalSupply   string  `json:"lp_total_supply"`
	Token0Address   string  `json:"token0_address"`
	Token1Address   string  `json:"token1_address"`
	Token0Amount    string  `json:"token0_amount"`
	Token1Amount    string  `json:"token1_amount"`
	Token0          string  `json:"token0"`
	Token1          string  `json:"token1"`
	ContractAddress string  `json:"contract_address"`
}

func (Pair) TableName() string {
	return "pair"
}

func AddPair(data map[string]interface{}) error {
	pair := Pair{
		Name:            data["name"].(string),
		PlatformId:      data["platform_id"].(int),
		Liquidity:       data["liquidity"].(float64),
		Volume24h:       data["volume_24h"].(float64),
		Volume7d:        data["volume_7d"].(float64),
		Fee24h:          data["fee_24h"].(float64),
		Apy:             data["apy"].(float64),
		LpTotalSupply:   data["lp_total_supply"].(string),
		Token0Address:   data["token0_address"].(string),
		Token1Address:   data["token1_address"].(string),
		Token0Amount:    data["token0_amount"].(string),
		Token1Amount:    data["token1_amount"].(string),
		Token0:          data["token0"].(string),
		Token1:          data["token1"].(string),
		ContractAddress: data["contract_address"].(string),
	}

	if err := db.Create(&pair).Error; err != nil {
		return err
	}
	return nil
}

func GetPair(maps interface{}) (*Pair, error) {
	var pair Pair
	err := db.Where(maps).Find(&pair).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &pair, err
}

func UpdatePairById(id int, data map[string]interface{}) error {
	if err := db.Model(&Pair{}).Where("id = ? ", id).Updates(data).Error; err != nil {
		logging.Error(err)
		return err
	}
	return nil
}
