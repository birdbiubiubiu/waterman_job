package graph_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"waterman_job/models"
	"waterman_job/pkg/logging"
	"waterman_job/service/slack_service"
)

type SushiSwapGraphql struct {
	Variables Variables `json:"variables"`
	Query     string    `json:"query"`
	Action    string    `json:"-"`
}

func (s SushiSwapGraphql) Get() {
	jsonStr, err := json.Marshal(s)
	if err != nil {
		logging.Error(err)
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", SushiSwapUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	r := &Result{}
	json.Unmarshal(body, &r)
	for _, v := range r.Data {
		for _, av := range v {
			t, _ := strconv.Atoi(av.Timestamp)
			a0, _ := strconv.ParseFloat(av.Amount0, 64)
			a1, _ := strconv.ParseFloat(av.Amount1, 64)
			au, _ := strconv.ParseFloat(av.AmountUSD, 64)
			w := models.Whales{
				Token0:        av.Pair.Token0.Symbol,
				Token1:        av.Pair.Token1.Symbol,
				AmountUsd:     au,
				Amount0:       a0,
				Amount1:       a1,
				Action:        s.Action,
				Platform:      PlatformSushiSwap,
				TransactionId: av.Transaction.ID,
				Timestamp:     t,
			}

			if err := models.AddWhales(&w); err != nil {
				logging.Error(err)
			} else {
				slack_service.SwapWhaleCh <- &w
			}
		}
	}
}
