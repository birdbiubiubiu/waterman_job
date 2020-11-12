package etherscan_jobs

import (
	"fmt"
	"waterman_job/pkg/e"
	"waterman_job/pkg/tools"
	"waterman_job/service/etherscan_service"
	"waterman_job/service/pair_service"
)

type UniJob struct {
	Name string
}

func(j UniJob) getParams() map[string]interface{} {
	params := make(map[string]interface{})
	switch j.Name {
		case "WBTC-ETH":
			params["lpContractAddress"] = e.UNI_WBTC_ETH_CONTRACT_ADDRESS
			params["token0Address"]     = e.UNI_WBTC_TOKEN_ADDRESS
			params["token1Address"]     = e.UNI_ETH_TOKEN_ADDRESS
			params["lpDecimals"]	    = int32(18)
			params["token0Decimals"]    = int32(8)
			params["token1Decimals"]    = int32(18)
			params["platformId"]        = 1
			params["apiKey"]			= e.API_KEY2
		case "ETH-DAI":
			params["lpContractAddress"] = e.UNI_ETH_DAI_CONTRACT_ADDRESS
			params["token0Address"]     = e.UNI_ETH_TOKEN_ADDRESS
			params["token1Address"]     = e.UNI_DAI_TOKEN_ADDRESS
			params["lpDecimals"]	    = int32(18)
			params["token0Decimals"]    = int32(18)
			params["token1Decimals"]    = int32(18)
			params["platformId"]        = 1
			params["apiKey"]			= e.API_KEY3
		case "ETH-USDC":
			params["lpContractAddress"] = e.UNI_ETH_USDC_CONTRACT_ADDRESS
			params["token0Address"]     = e.UNI_ETH_TOKEN_ADDRESS
			params["token1Address"]     = e.UNI_USDC_TOKEN_ADDRESS
			params["lpDecimals"]	    = int32(18)
			params["token0Decimals"]    = int32(6)
			params["token1Decimals"]    = int32(18)
			params["platformId"]        = 1
			params["apiKey"]			= e.API_KEY4
	case "ETH-USDT":
			params["lpContractAddress"] = e.UNI_USDT_TOKEN_ADDRESS
			params["token0Address"]     = e.UNI_ETH_TOKEN_ADDRESS
			params["token1Address"]     = e.UNI_USDT_TOKEN_ADDRESS
			params["lpDecimals"]	    = int32(18)
			params["token0Decimals"]    = int32(6)
			params["token1Decimals"]    = int32(18)
			params["platformId"]        = 1
			params["apiKey"]			= e.API_KEY2
	}
	return params
}

func (j UniJob) Run()  {
	fmt.Println(j.Name)
	params := j.getParams()
	supply, err := etherscan_service.GetTokenSupply(params["lpContractAddress"].(string))
	if err != nil {
		return
	}
	supply = tools.BigIntStrToFloatStr(supply, params["lpDecimals"].(int32))

	wBtcBalance, err := etherscan_service.GetTokenBalance(params["token0Address"].(string), params["lpContractAddress"].(string), params["apiKey"].(string))
	if err != nil {
		return
	}
	wBtcBalance = tools.BigIntStrToFloatStr(wBtcBalance, params["token0Decimals"].(int32))

	ethBalance, err := etherscan_service.GetTokenBalance(params["token1Address"].(string), params["lpContractAddress"].(string), params["apiKey"].(string))
	if err != nil {
		return
	}
	ethBalance = tools.BigIntStrToFloatStr(ethBalance, params["token1Decimals"].(int32))

	pair := pair_service.Pair{Name: j.Name, PlatformId: params["platformId"].(int), LpTotalSupply: supply, Token0Amount: wBtcBalance, Token1Amount: ethBalance}

	err = pair.UpdatePair()
	if err != nil {
		return
	}
}