package main 

import (
  "strconv"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strings"
   )

type MedPortfolio struct {
NameofShares string
NumofShares string
CurrentPriceShares string
}

type Urlrecord1 struct {
	Query struct {
	Results struct {
			Quote struct {
			Ask float32 `json:",string"`
			} `json:"quote"`
		} `json:"results"`
	} `json:"query"`
	}

type CheckPortfolioService struct{}

func(h *CheckPortfolioService) Say(r *http.Request, args *CheckPortFolioInput, response *CheckPortfolioResponse) error {

var AccessHolderPortfolio Holder

var stockObjMapPortfolio[] Stocksinfo

AccessHolderPortfolio.Hold = M[args.TradeId].Hold;

stockObjMapPortfolio = AccessHolderPortfolio.Hold

var currentSP float64
var index int
var var1 string 
var TotalUnvestedAmoutPortfolio float32

TotalStocksendMessage:= make([]string,len(stockObjMapPortfolio))
var FinalStocksendMessage string
  
  
for index = 0; index < len(stockObjMapPortfolio); index+=1{


s1 := []string{}

s1 = append(s1, "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20yahoo.finance.quotes%20where%20symbol%20in%20(")
s1 = append(s1,"'")
s1 = append(s1,stockObjMapPortfolio[index].Name)
s1 = append(s1,"'")
s1 = append(s1,")&format=json&diagnostics=true&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys&callback=")

var url string = s1[0] + s1[1] +s1[2] +s1[3] + s1[4]
record1,_:= GetContentDetailscheck(url)
response.UnvestedAmounts=stockObjMapPortfolio[index].UnvestedAmount;
currentSP,_=(strconv.ParseFloat(stockObjMapPortfolio[index].CurrentStockPrice, 64))
var ask float64
ask= (float64)(record1.Query.Results.Quote.Ask)
var x_ask float32
x_ask = (float32) (ask)
var b_currentSP float32;
b_currentSP=(float32) (currentSP)
if(b_currentSP < x_ask) {
var1 = "+"

} else if (b_currentSP > x_ask){ 
var1 = "-"
} else {

var1 = ""

}
msg1 := []string{}
msg1 = append(msg1,stockObjMapPortfolio[index].Name)
msg1 = append(msg1,":")
msg1 = append(msg1, strconv.Itoa(stockObjMapPortfolio[index].NumberOfShares))
msg1 = append(msg1,":")
msg1 = append(msg1,var1)
msg1 = append(msg1,"$")
msg1 = append(msg1,strconv.FormatFloat((float64)(record1.Query.Results.Quote.Ask), 'f', 2, 32))


TotalStocksendMessage[index] = msg1[0] + msg1[1] +msg1[2] +msg1[3] + msg1[4] + msg1[5] + msg1[6]
TotalUnvestedAmoutPortfolio += (float32)(stockObjMapPortfolio[index].NumberOfShares)  * record1.Query.Results.Quote.Ask
	
}

FinalStocksendMessage = strings.Join(TotalStocksendMessage, ", ")

response.Stocks=FinalStocksendMessage;
response.CurrentMarketValue=TotalUnvestedAmoutPortfolio;
response.UnvestedAmounts = stockObjMapPortfolio[0].UnvestedAmount
return nil
}
func getContentcheck(url string) ([]byte, error) {
req, err := http.NewRequest("GET", url, nil)
if err != nil {
return nil, err
}
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
return nil, err
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
return nil, err
}
return body, nil
}

func GetContentDetailscheck(ip string) (*Urlrecord1, error) {
content, err := getContentcheck(ip)
if err != nil {
}
	var record1 Urlrecord1
	err = json.Unmarshal(content, &record1)
	if err != nil {
	}
	return &record1, err
	}
	
          