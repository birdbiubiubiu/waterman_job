package graph_service

type Variables struct {
	AllPairs      []string `json:"allPairs"`
	LastTimestamp int      `json:"lastTimestamp"`
	First         int      `json:"first"`
	Offset        int      `json:"offset"`
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
	Timestamp   string `json:"timestamp"`
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

var UniSwapUrl = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"
var SushiSwapUrl = "https://api.thegraph.com/subgraphs/name/zippoxer/sushiswap-subgraph-fork"

var PlatformUniSwap = "uniswap"
var PlatformSushiSwap = "sushiswap"
