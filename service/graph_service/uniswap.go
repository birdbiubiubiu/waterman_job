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

type Variables struct {
	AllPairs      []string `json:"allPairs"`
	LastTimestamp int      `json:"lastTimestamp"`
	First 		  int `json:"first"`
	Offset 		  int `json:"offset"`
}

type UniSwapGraphql struct {
	Variables Variables `json:"variables"`
	Query string `json:"query"`
	Action string `json:"-"`
}

type Mints struct {
	AmountUSD string `json:"amountUSD"`
	ID        string `json:"id"`
	Pair      struct {
		Token0 struct {
			Symbol string `json:"symbol"`
		} `json:"token0"`
		Token1 struct {
			Symbol string `json:"symbol"`
		} `json:"token1"`
	} `json:"pair"`
	Timestamp    string `json:"timestamp"`
	Transaction struct {
		ID        string `json:"id"`
		Timestamp string `json:"timestamp"`
	} `json:"transaction"`
	Amount0 string `json:"amount0"`
	Amount1 string `json:"amount1"`
}

type Result struct {
	Data map[string][]Mints `json:"data"`
}

var url = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"

func (u UniSwapGraphql) Get() {
	jsonStr, err := json.Marshal(u)
	if err != nil {
		logging.Error(err)
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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
	for _, v := range r.Data{
		for _,av := range v {
			t,_ := strconv.Atoi(av.Timestamp)
			a0,_  := strconv.ParseFloat(av.Amount0, 64)
			a1,_  := strconv.ParseFloat(av.Amount1, 64)
			au,_ :=strconv.ParseFloat(av.AmountUSD, 64)
			w := models.Whales{
				Token0: av.Pair.Token0.Symbol,
				Token1: av.Pair.Token1.Symbol,
				AmountUsd: au,
				Amount0: a0,
				Amount1: a1,
				Action: u.Action,
				TransactionId: av.Transaction.ID,
				Timestamp: t,
			}
			slack_service.UniSwapWhaleCh <- &w

			if err := models.AddWhales(&w);err!=nil {
				logging.Error(err)
			}
		}
	}
}

func (u UniSwapGraphql) GetSwaps() {
	jsonStr, err := json.Marshal(u)
	if err != nil {
		logging.Error(err)
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		logging.Error(err)
	}
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
	for _, v := range r.Data{
		for _,av := range v {
			t,_ := strconv.Atoi(av.Timestamp)
			a0,_  := strconv.ParseFloat(av.Amount0, 64)
			a1,_  := strconv.ParseFloat(av.Amount1, 64)
			au,_ :=strconv.ParseFloat(av.AmountUSD, 64)
			w := models.Whales{
				Token0: av.Pair.Token0.Symbol,
				Token1: av.Pair.Token1.Symbol,
				AmountUsd: au,
				Amount0: a0,
				Amount1: a1,
				Action: u.Action,
				TransactionId: av.Transaction.ID,
				Timestamp: t,
			}
			if err := models.AddWhales(&w);err!=nil {
				logging.Error(err)
			}
		}
	}

}