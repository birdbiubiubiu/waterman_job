package models

import (
	"gorm.io/gorm"
)

type Whales struct {
	Model

	Token0        string  `json:"token0"`
	Token1        string  `json:"token1"`
	AmountUsd     float64 `json:"amount_usd"`
	Amount0       float64 `json:"amount0"`
	Amount1       float64 `json:"amount1"`
	Action        string  `json:"action"`
	Platform      string  `json:"platform"`
	TransactionId string  `json:"transaction_id"`
	Timestamp     int     `json:"timestamp"`
}

func (Whales) TableName() string {
	return "whales"
}

func AddWhales(w *Whales) error {
	if err := db.Create(w).Error; err != nil {
		return err
	}
	return nil
}

func GetWhales(maps interface{}) (*Whales, error) {
	var w Whales
	err := db.Where(maps).Find(&w).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &w, err
}

func GetLastWhalesByAction(action, platform string) (*Whales, error) {
	var w Whales
	err := db.Limit(1).Order("timestamp desc").Where("action = ? and platform = ?", action, platform).Find(&w).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &w, err
}
