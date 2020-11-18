package models

type SymbolPriceSnapshot struct {
	Model

	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (SymbolPriceSnapshot) TableName() string {
	return "symbol_price_snapshot"
}

func AddSymbolPriceSnapshot(s *SymbolPriceSnapshot) error {
	if err := db.Create(s).Error; err != nil {
		return err
	}
	return nil
}
