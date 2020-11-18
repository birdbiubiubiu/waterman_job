package graph_jobs

import (
	"waterman_job/models"
	"waterman_job/pkg/e"
	"waterman_job/pkg/logging"
	"waterman_job/service/graph_service"
)

type UniSwapGraphql struct {
	Action string
}

func (u UniSwapGraphql) Run() {
	lastRecord, err := models.GetLastWhalesByAction(u.Action, graph_service.PlatformUniSwap)
	lastTimestamp := int(0)
	query := ""
	switch u.Action {
	case "mint":
		query = "query($allPairs: [Bytes]!,$lastTimestamp :BigInt!, $first :Int!){mints(first:$first, orderBy: timestamp, orderDirection: asc,where:{pair_in: $allPairs,timestamp_gt:$lastTimestamp,amountUSD_gt:100000}){pair{token0{symbol},token1{symbol}},amountUSD,id,amount0,amount1,timestamp,transaction{id,timestamp}}}"
	case "burn":
		query = "query($allPairs: [Bytes]!,$lastTimestamp :BigInt!, $first :Int!){burns(first:$first, orderBy: timestamp, orderDirection: asc,where:{pair_in: $allPairs,timestamp_gt:$lastTimestamp,amountUSD_gt:100000}){pair{token0{symbol},token1{symbol}},amountUSD,id,amount0,amount1,timestamp,transaction{id,timestamp}}}"
	case "swap":
		query = "query($allPairs: [Bytes]!,$lastTimestamp :BigInt!, $first :Int!){swaps(first:$first, orderBy: timestamp, orderDirection: asc,where:{pair_in: $allPairs,timestamp_gt:$lastTimestamp,amountUSD_gt:100000}){pair{token0{symbol},token1{symbol}},amountUSD,id,amount0,amount1,timestamp,transaction{id,timestamp}}}"
	}

	if err != nil {
		logging.Error(err)
	}

	if lastRecord != nil {
		lastTimestamp = lastRecord.Timestamp
	}
	graphqlVar := graph_service.Variables{
		AllPairs:      []string{e.UNI_WBTC_ETH_CONTRACT_ADDRESS, e.UNI_ETH_DAI_CONTRACT_ADDRESS, e.UNI_ETH_USDT_CONTRACT_ADDRESS, e.UNI_ETH_USDC_CONTRACT_ADDRESS},
		LastTimestamp: lastTimestamp,
		First:         20,
	}

	ug := graph_service.UniSwapGraphql{Variables: graphqlVar, Query: query, Action: u.Action}
	ug.Get()
}
