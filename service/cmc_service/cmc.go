package cmc_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"waterman_job/models"
	"waterman_job/pkg/logging"
)

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
}

type Usd struct {
	Price           float64 `json:"price"`
	Volume24H       float64 `json:"volume_24h"`
	PercentChange1H float64 `json:"percent_change_1h"`
}

type Data struct {
	Name  string         `json:"name"`
	Quote map[string]Usd `json:"quote"`
}

type Result struct {
	Status Status          `json:"status"`
	Data   map[string]Data `json:"data"`
}

func UpdateSymbolPrice(symbols []string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		return err
	}

	q := url.Values{}
	q.Add("convert", "USD")
	querySymbols := strings.Replace(strings.Trim(fmt.Sprint(symbols), "[]"), " ", ",", -1)
	q.Add("symbol", querySymbols)
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "852b04b3-4ec4-40b9-a026-1e727b56dda0")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		logging.Error(err)
		return err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	result := &Result{}
	err = json.Unmarshal([]byte(respBody), result)
	if err != nil {
		logging.Error(err)
		return err
	}

	for k, v := range result.Data {
		updateData := make(map[string]interface{})
		updateData["price"] = v.Quote["USD"].Price
		models.UpdateSymbolPrice(k, updateData)
		snapshotData := map[string]interface{}{
			"name":  k,
			"price": v.Quote["USD"].Price,
		}
		AddSymbolPriceSnapshot(snapshotData)
	}
	return nil
}

func AddSymbolPriceSnapshot(data map[string]interface{}) error {
	s := models.SymbolPriceSnapshot{
		Name:  data["name"].(string),
		Price: data["price"].(float64),
	}
	if err := models.AddSymbolPriceSnapshot(&s); err != nil {
		logging.Error(err)
		return err
	}
	return nil
}
