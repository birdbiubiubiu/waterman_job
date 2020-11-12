package etherscan_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"waterman_job/pkg/logging"
)

const (
	URI      = "https://api.etherscan.io"
	API_KEY  = "2X7KGBMYBDF354CGKRDGXB6XVB6UZXCGAE"
	API_KEY2 = "SJN5BSUXRVEMG436AKUDP547IBGV5HWHFB"
	API_KEY3 = "F3RJGNISMQ25PCTXZIKJUAZI266NM5UJPC"
	API_KEY4 = "NJZTAMIQEYZMV7NEQTEDPCKGR5FV4A4EXI"
)

type TokenSupplyResponseResult struct {
	Status, Message string
	Result          string
}

func GetTokenSupply(contractAddress string) (string, error) {
	url := fmt.Sprintf("%s/api?module=%s&action=%s&contractaddress=%s&apikey=%s", URI, "stats", "tokensupply", contractAddress, API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		logging.Error(err)
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	result := &TokenSupplyResponseResult{}
	json.Unmarshal(body, result)
	if result.Status == "1" && result.Message == "OK" {
		return result.Result, nil
	}
	defer resp.Body.Close()
	return "", nil
}

func GetTokenBalance(contractAddress, tokenAddress, apiKey string) (string, error) {
	url := fmt.Sprintf("%s/api?module=%s&action=%s&contractaddress=%s&address=%s&apikey=%s", URI, "account", "tokenbalance", contractAddress, tokenAddress, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		logging.Error(err)
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	result := &TokenSupplyResponseResult{}
	json.Unmarshal(body, result)
	if result.Status == "1" && result.Message == "OK" {
		return result.Result, nil
	}
	defer resp.Body.Close()
	return "", nil
}
