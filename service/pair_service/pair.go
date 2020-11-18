package pair_service

import (
	"waterman_job/models"
	"waterman_job/pkg/logging"
)

type Pair struct {
	ID            int
	Name          string
	Apy           float64
	PlatformId    int
	LpTotalSupply string
	Token0Amount  string
	Token1Amount  string
}

func (p *Pair) UpdatePair() error {
	pair, err := models.GetPair(p.getQueryMaps())
	if err != nil {
		logging.Error(err)
		return err
	}
	updateErr := models.UpdatePairById(pair.ID, p.getUpdateMaps())
	if updateErr != nil {
		logging.Error(updateErr)
		return updateErr
	}
	snapshotData := map[string]interface{}{
		"name":          p.Name,
		"platformId":    p.PlatformId,
		"lpTotalSupply": p.LpTotalSupply,
		"token0Amount":  p.Token0Amount,
		"token1Amount":  p.Token1Amount,
		"apy":           p.Apy,
	}
	if err := models.AddPairSnapshot(snapshotData); err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

func (p *Pair) GetPair() (*models.Pair, error) {
	pair, err := models.GetPair(p.getQueryMaps())
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	return pair, nil
}

func (p *Pair) getQueryMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if p.PlatformId != 0 {
		maps["platform_id"] = p.PlatformId
	}

	if p.Name != "" {
		maps["name"] = p.Name
	}

	return maps
}

func (p *Pair) getUpdateMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if p.LpTotalSupply != "'" {
		maps["lp_total_supply"] = p.LpTotalSupply
	}

	if p.Token0Amount != "" {
		maps["token0_amount"] = p.Token0Amount
	}

	if p.Token1Amount != "" {
		maps["token1_amount"] = p.Token1Amount
	}

	if p.Apy != 0 {
		maps["apy"] = p.Apy
	}

	return maps
}

func CalculateAPY(proportion float64, token0Name, platform string) (float64, error) {
	if platform == "uniswap" {
		where := map[string]interface{}{
			"name": "UNI",
		}
		symbolPrice, err := models.GetSymbolPrice(where)
		whereToken := map[string]interface{}{
			"name": token0Name,
		}
		uniPrice := symbolPrice.Price
		dailyUniValue := uniPrice * 583333 / 7
		token0Price, err := models.GetSymbolPrice(whereToken)
		if err != nil {
			return 0.0, err
		}
		basePrice := token0Price.Price
		rewardValue := proportion * dailyUniValue / 2
		apy := rewardValue / basePrice * 365
		return apy, nil
	}
	return 0.0, nil
}
