package pair_service

import (
	"waterman_job/models"
	"waterman_job/pkg/logging"
)

type Pair struct {
	ID            int
	Name          string
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

	return maps
}
