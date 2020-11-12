package models

import "waterman_job/pkg/logging"

type SymbolPrice struct {
	Model

	Name   string `json:"name"`
	Price float64 `json:"price"`
}

func (SymbolPrice) TableName() string {
	return "symbol_price"
}

func UpdateSymbolPrice(name string, data map[string]interface{}) error {
	if err := db.Model(&SymbolPrice{}).Where("name = ? ", name).Updates(data).Error; err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

