package main

import (
	"github.com/robfig/cron/v3"
	"waterman_job/jobs/graph_jobs"
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
	c.AddJob("@every 10s", graph_jobs.UniswapGraphql{Name: "xx"})
	//c.AddJob("*/5 * * * ?", cmc_jobs.UpdateSymbolPriceJob{Name :"update symbol price from cmc"})
	//c.AddJob("@every 10s", etherscan_jobs.UniJob{Name:"WBTC-ETH", Token0Name: "BTC"})
	//c.AddJob("@every 10s", etherscan_jobs.UniJob{Name:"ETH-DAI", Token0Name: "ETH"})
	//c.AddJob("@every 10s", etherscan_jobs.UniJob{Name:"ETH-USDT", Token0Name: "ETH"})
	//c.AddJob("@every 10s", etherscan_jobs.UniJob{Name:"ETH-USDC", Token0Name: "ETH"})
	//
	c.Start()
	select {}

}