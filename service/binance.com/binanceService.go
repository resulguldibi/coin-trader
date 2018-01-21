package service

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"resulguldibi/coin-trader/contract/binance.com"
	"strconv"
)


type binanceService struct{
	httpClient *http.Client
	baseUrl string
}

type BinanceServiceFactory struct{}

func (f *BinanceServiceFactory) NewBinanceService(httpClient *http.Client,baseUrl string) *binanceService{
	return &binanceService{httpClient:httpClient,baseUrl: baseUrl}
}

func (s *binanceService) GetTime() *contract.GetTimeResponse{

	resp,err := s.httpClient.Get(fmt.Sprintf("%s/api/v1/time",s.baseUrl))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	response :=&contract.GetTimeResponse{}
	err = json.Unmarshal(body,response)

	if err != nil {
		panic(err)
	}

	return response
}

func (s *binanceService) GetExchangeInfo() *contract.GetExchangeInfoResponse{

	resp,err := s.httpClient.Get(fmt.Sprintf("%s/api/v1/exchangeInfo",s.baseUrl))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	response :=&contract.GetExchangeInfoResponse{}
	err = json.Unmarshal(body,response)

	if err != nil {
		panic(err)
	}

	return response
}

func (s *binanceService) GetDepth(symbol string, limit int32) *contract.GetDepthResponse{


	url := fmt.Sprintf("%s/api/v1/depth",s.baseUrl)

	if symbol != ""{
		url +="?symbol="+symbol
	}

	if limit > 0{
		url +="&limit="+ strconv.Itoa(int(limit))
	}

	resp,err := s.httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	response :=&contract.GetDepthResponse{}
	err = json.Unmarshal(body,response)

	if err != nil {
		panic(err)
	}

	return response
}

func (s *binanceService) GetTrades(symbol string, limit int32) []contract.Trade{


	url := fmt.Sprintf("%s/api/v1/trades",s.baseUrl)

	if symbol != ""{
		url +="?symbol="+symbol
	}

	if limit > 0{
		url +="&limit="+ strconv.Itoa(int(limit))
	}

	resp,err := s.httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	trades := make([]contract.Trade,0,0)

	err = json.Unmarshal(body,&trades)

	if err != nil {
		panic(err)
	}
	return trades
}

func (s *binanceService) GetAggTrades(symbol string, limit int32,fromId int64, startTime int64, endTime int64) []contract.AggTrade{


	url := fmt.Sprintf("%s/api/v1/aggTrades",s.baseUrl)

	if symbol != ""{
		url +="?symbol="+symbol
	}

	if limit > 0{
		url +="&limit="+ strconv.Itoa(int(limit))
	}

	if fromId > 0{
		url +="&fromId="+ strconv.Itoa(int(fromId))
	}

	if startTime > 0{
		url +="&startTime="+ strconv.Itoa(int(startTime))
	}

	if endTime > 0{
		url +="&endTime="+ strconv.Itoa(int(endTime))
	}

	resp,err := s.httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	trades := make([]contract.AggTrade,0,0)

	err = json.Unmarshal(body,&trades)

	if err != nil {
		panic(err)
	}
	return trades
}

func (s *binanceService) GetKlines(symbol string,interval contract.KlineType, limit int32, startTime int64, endTime int64) []contract.KlineInfo{


	url := fmt.Sprintf("%s/api/v1/klines",s.baseUrl)

	if symbol != ""{
		url +="?symbol="+symbol
	}

	if limit > 0{
		url +="&limit="+ strconv.Itoa(int(limit))
	}

	if interval != ""{
		url +="&interval="+ string(interval)
	}

	if startTime > 0{
		url +="&startTime="+ strconv.Itoa(int(startTime))
	}

	if endTime > 0{
		url +="&endTime="+ strconv.Itoa(int(endTime))
	}

	resp,err := s.httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	klines := make([]contract.KlineInfo,0,0)

	data := make([][]interface{},0,0)

	err = json.Unmarshal(body,&data)
	
	if err != nil {
		panic(err)
	}

	if len(data) > 0{
		for _,list := range data{
			fmt.Println("list : ",list)
			klineInfo := contract.KlineInfo{
				OpenTime: int64((list[0].(float64))),
				OpenPrice: list[1].(string),
				HighPrice : list[2].(string),
				LowPrice : list[3].(string),
				ClosePrice :list[4].(string),
				Volume: list[5].(string),
				CloseTime : int64((list[6].(float64))),
				QuoteAssetVolume : list[7].(string),
				TradeCount : int32(list[8].(float64)),
				TakerBuyBaseAssetVolume : list[9].(string),
				TakerBuyQuoteAssetVolume : list[10].(string),
				Ignore : list[11].(string),
			}

			klines = append(klines,klineInfo)
		}
	}


	return klines
}

func (s *binanceService) GetTicker24h(symbol string) interface{}{

	url := fmt.Sprintf("%s/api/v1/ticker/24hr",s.baseUrl)

	if symbol != ""{
		url +="?symbol="+symbol
	}


	resp,err := s.httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var tickers interface{}

	if symbol != ""{
		tickers = contract.TickerInfo{}
	}else{
		tickers = make([]contract.TickerInfo,0,0)
	}

	err = json.Unmarshal(body,&tickers)
	
	if err != nil {
		panic(err)
	}

	return tickers
}