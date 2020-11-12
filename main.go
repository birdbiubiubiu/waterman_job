package main

import (
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
	logging.Error(111)
	//c := cron.New()
	//c.AddJob("@every 1s", cmc_jobs.UpdateSymbolPriceJob{"11"})
	//c.Start()
	//
	//time.Sleep(5 * time.Second)
}