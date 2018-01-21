package main

import(
	"resulguldibi/coin-trader/security/encryption"
	"resulguldibi/coin-trader/contract/binance.com"
	"resulguldibi/coin-trader/service/binance.com"
	"fmt"
	"os"
	"net/http"
	"time"
	"encoding/json"
)

func main(){
	factory :=&encryption.Sha256EncryptorFactory{}
	
	sha256Encryptor :=factory.NewSha256Encryptor(os.Getenv("BINANCE_SECRET_KEY"))
	encryptedMessage := sha256Encryptor.Encrypt("asasasasas")

	fmt.Println(encryptedMessage)

	binanceServiceFactory :=service.BinanceServiceFactory{}
	baseUrl := "https://api.binance.com"
	httpClient := &http.Client{
		Timeout : time.Millisecond *1000 * 2,
	}
	binanceService:=binanceServiceFactory.NewBinanceService(httpClient,baseUrl)
	getTimeResponse := binanceService.GetTime()
	

	getExchangeInfoResponse :=binanceService.GetExchangeInfo()

	fmt.Println("getTimeResponse : ",getTimeResponse)


	getExchangeInfoResponseBytes, _ :=json.Marshal(getExchangeInfoResponse)

	fmt.Println("getExchangeInfoResponse : ",string(getExchangeInfoResponseBytes))

	getDepthResponse := binanceService.GetDepth("ETHBTC",5)

	getDepthResponseBytes, _ :=json.Marshal(getDepthResponse)

	fmt.Println("getDepthResponse : ",string(getDepthResponseBytes))

	trades := binanceService.GetTrades("ETHBTC",5)

	tradesBytes, _ :=json.Marshal(trades)

	fmt.Println("getTradesResponse : ",string(tradesBytes))

	aggTrades := binanceService.GetAggTrades("ETHBTC",5,0,0,0)

	aggTradesBytes, _ :=json.Marshal(aggTrades)

	fmt.Println("getAggTradesResponse : ",string(aggTradesBytes))

	klines := binanceService.GetKlines("ETHBTC",contract.M1,1,0,0)

	klinesBytes, _ :=json.Marshal(klines)

	fmt.Println("klinesBytesResponse : ",string(klinesBytes))

	tickers := binanceService.GetTicker24h("")

	tickersBytes, _ :=json.Marshal(tickers)

	fmt.Println("tickersResponse : ",string(tickersBytes))

}