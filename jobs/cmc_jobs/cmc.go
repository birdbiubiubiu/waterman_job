package cmc_jobs

import (
	"waterman_job/pkg/logging"
	"waterman_job/service/cmc_service"
)

type UpdateSymbolPriceJob struct {
	Name string
}

func (uspj UpdateSymbolPriceJob) Run() {
	logging.Info("execute" + uspj.Name)
	cmc_service.UpdateSymbolPrice([]string{"BTC", "ETH", "BNB", "CAKE", "DOT", "FARM", "SUSHI"})
}
