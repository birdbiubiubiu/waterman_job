package graph_jobs

import (
	"fmt"
	"waterman_job/models"
	"waterman_job/service/graph_service"
)

type UniswapGraphql struct {
	Name string
}

func(u UniswapGraphql) Run()  {
	lastRecord, err := models.GetLastWhales()
	lastTimestamp := int(0)

	if err != nil {
		fmt.Println(err)
	}
	if lastRecord != nil {
		lastTimestamp = lastRecord.Timestamp
	}
	graphqlVar := graph_service.Variables{
		AllPairs: []string{"0xbb2b8038a1640196fbe3e38816f3e67cba72d940", "0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc"},
		LastTimestamp: lastTimestamp,
		First: 10,
	}
	ug := graph_service.UniswapGraphql{Variables: graphqlVar, Query: "query($allPairs: [Bytes]!,$lastTimestamp :BigInt!){mints(first:20, orderBy: timestamp, orderDirection: asc,where:{pair_in: $allPairs,timestamp_gt:$lastTimestamp,amountUSD_gt:500000}){pair{token0{symbol},token1{symbol}},amountUSD,id,amount0,amount1,timestamp,transaction{id,timestamp}}}"}
	ug.Get()
}
