package contract

type GetTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

type GetExchangeInfoResponse struct {
	Timezone   string      `json:"timezone"`
	ServerTime int64       `json:"serverTime"`
	RateLimits []RateLimit `json:"rateLimits"`
	Symbols    []Symbol    `json:"symbols"`
}

type GetDepthResponse struct {
	LastUpdateId int64           `json:"lastUpdateId"`
	Bids         [][]interface{} `json:"bids"`
	Asks         [][]interface{} `json:"asks"`
}

type Trade struct {
	Id           int64  `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

type AggTrade struct {
	Id           int64  `json:"a"`
	Price        string `json:"p"`
	Qty          string `json:"q"`
	FirstTradeId int32  `json:"f"`
	LastTradeId  int32  `json:"l"`
	Time         int64  `json:"T"`
	IsBuyerMaker bool   `json:"m"`
	IsBestMatch  bool   `json:"M"`
}

type KlineInfo struct {
	OpenTime                 int64  `json:"openTime"`
	OpenPrice                string `json:"openPrice"`
	HighPrice                string `json:"highPrice"`
	LowPrice                 string `json:"lowPrice"`
	ClosePrice               string `json:"closePrice"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	TradeCount               int32  `json:"tradeCount"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
	Ignore                   string `json:"ignore"`
}

type RateLimit struct {
	RateLimitType RateLimitType         `json:"rateLimitType"`
	Interval      RateLimitIntervalType `json:"interval"`
	Limit         int32                 `json:"limit"`
}

type TickerInfo struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}

type Filter struct {
	FilterType  string `json:"filterType"`
	MinPrice    string `json:"minPrice"`
	MaxPrice    string `json:"maxPrice"`
	TickSize    string `json:"tickSize"`
	MinQty      string `json:"minQty"`
	MaxQty      string `json:"maxQty"`
	StepSize    string `json:"stepSize"`
	MinNotional string `json:"minNotional"`
}

type Symbol struct {
	Symbol             string      `json:"symbol"`
	Status             string      `json:"status"`
	BaseAsset          string      `json:"baseAsset"`
	BaseAssetPrecision int8        `json:"baseAssetPrecision"`
	QuoteAsset         string      `json:"quoteAsset"`
	QuotePrecision     int8        `json:"quotePrecision"`
	IcebergAllowed     bool        `json:"icebergAllowed"`
	OrderTypes         []OrderType `json:"orderTypes"`
	Filters            []Filter    `json:"filters"`
}

type SymbolStatusType string

var (
	PRE_TRADING   SymbolStatusType = "PRE_TRADING"
	TRADING       SymbolStatusType = "TRADING"
	POST_TRADING  SymbolStatusType = "POST_TRADING"
	END_OF_DAY    SymbolStatusType = "END_OF_DAY"
	HALT          SymbolStatusType = "HALT"
	AUCTION_MATCH SymbolStatusType = "AUCTION_MATCH"
	BREAK         SymbolStatusType = "BREAK"
)

type SymbolType string

var (
	SPOT SymbolType = "SPOT"
)

type OrderStatus string

var (
	NEW              OrderStatus = "NEW"
	PARTIALLY_FILLED OrderStatus = "PARTIALLY_FILLED"
	FILLED           OrderStatus = "FILLED"
	CANCELED         OrderStatus = "CANCELED"
	PENDING_CANCEL   OrderStatus = "PENDING_CANCEL"
	REJECTED         OrderStatus = "REJECTED"
	EXPIRED          OrderStatus = "EXPIRED"
)

type OrderType string

var (
	LIMIT             OrderType = "LIMIT"
	MARKET            OrderType = "MARKET"
	STOP_LOSS         OrderType = "STOP_LOSS"
	STOP_LOSS_LIMIT   OrderType = "STOP_LOSS_LIMIT"
	TAKE_PROFIT       OrderType = "TAKE_PROFIT"
	TAKE_PROFIT_LIMIT OrderType = "TAKE_PROFIT_LIMIT"
	LIMIT_MAKER       OrderType = "LIMIT_MAKER"
)

type OrderSideType string

var (
	BUY  OrderSideType = "BUY"
	SELL OrderSideType = "SELL"
)

type TimeInForceType string

var (
	GTC TimeInForceType = "GTC"
	IOC TimeInForceType = "IOC"
	FOK TimeInForceType = "FOK"
)

type KlineType string

var (
	M1  KlineType = "1m"
	M3  KlineType = "3m"
	M5  KlineType = "5m"
	M15 KlineType = "15m"
	M30 KlineType = "30m"
	H1  KlineType = "1h"
	H2  KlineType = "2h"
	H4  KlineType = "4h"
	H6  KlineType = "6h"
	H8  KlineType = "8h"
	H12 KlineType = "12h"
	D1  KlineType = "1d"
	D3  KlineType = "3d"
	W1  KlineType = "1w"
	MM1 KlineType = "1M"
)

type RateLimitType string

var (
	REQUESTS RateLimitType = "REQUESTS"
	ORDERS   RateLimitType = "ORDERS"
)

type RateLimitIntervalType string

var (
	SECOND RateLimitIntervalType = "SECOND"
	MINUTE RateLimitIntervalType = "MINUTE"
	DAY    RateLimitIntervalType = "DAY"
)
