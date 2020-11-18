package graph_service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"waterman_job/models"
	"waterman_job/pkg/logging"
	"waterman_job/service/slack_service"
)

type UniSwapGraphql struct {
	Variables Variables `json:"variables"`
	Query     string    `json:"query"`
	Action    string    `json:"-"`
}

func (u UniSwapGraphql) Get() {
	jsonStr, err := json.Marshal(u)
	if err != nil {
		logging.Error(err)
		return
	}
	req, err := http.NewRequest("POST", UniSwapUrl, bytes.NewBuffer(jsonStr))
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
				Action:        u.Action,
				TransactionId: av.Transaction.ID,
				Platform:      PlatformUniSwap,
				Timestamp:     t,
			}

			if err := models.AddWhales(&w); err != nil {
				logging.Error(w)
				logging.Error(err)
			} else {
				slack_service.SwapWhaleCh <- &w
			}
		}
	}
}
