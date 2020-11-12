package main

import (
	"github.com/robfig/cron/v3"
	"waterman_job/jobs/cmc_jobs"
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
	c.AddJob("*/5 * * * ?", cmc_jobs.UpdateSymbolPriceJob{"update symbol price from cmc"})
	c.Start()
	select {}

}