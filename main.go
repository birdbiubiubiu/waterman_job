package main

import (
	"github.com/robfig/cron/v3"
	"waterman_job/jobs/cmc_jobs"
	"waterman_job/jobs/etherscan_jobs"
	"waterman_job/models"
	"waterman_job/pkg/logging"
	"waterman_job/pkg/setting"
)

func init()  {
	setting.Setup()
	logging.Setup()
	models.Setup()
}

func main()  {
	c := cron.New()
	c.AddJob("*/5 * * * ?", cmc_jobs.UpdateSymbolPriceJob{Name :"update symbol price from cmc"})
	c.AddJob("@every 10s", etherscan_jobs.UniJob{Name:"WBTC-ETH",})
	c.AddJob("@every 10s", etherscan_jobs.UniJob{Name:"ETH-DAI",})
	c.AddJob("@every 10s", etherscan_jobs.UniJob{Name:"ETH-USDT",})
	c.AddJob("@every 10s", etherscan_jobs.UniJob{Name:"ETH-USDC",})

	c.Start()
	select {}

}