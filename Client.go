package main

import (
"encoding/json"
"errors"
"fmt"
"io"
"os"
"math/rand"
"bytes"
"net/http"
"strconv"
)

// This Structure is used to check gain/loss on stocks
type CheckPortfolioResponse struct { 
Stocks string
CurrentMarketValue float32
UnvestedAmounts float32 
}
// Request Structure of Investment
type StocksInput struct { 
StockSymbolAndPercentage string
Budget float32
}
// Intermediate Structure to hold Maps{Key/Value] pair
type Stocksinfo struct { 
Name string
Percentage string
TotalBudget float32
NumberOfShares int
CurrentStockPrice string
UnvestedAmount float32
}
// Final Reply Structure for first time investment
type StocksinfoReply struct { 
TradeId int
Stocks string
UnvestedAmount float32
}
// Request TradeID to view Portfolio
type CheckPortFolioInput struct { 
TradeId int
}
// JSON Client Request Structure
type ClientRequest struct {
Method string `json:"method"`
Params [1]interface{} `json:"params"` 
 Id uint64 `json:"id"`
}
// clientResponse represents a JSON-RPC response returned to a client.
type ClientResponse struct {
Result *json.RawMessage `json:"result"`
Error  interface{}      `json:"error"`
Id     uint64           `json:"id"`
}
// EncodeClientRequest encodes parameters for a JSON-RPC client request.
func EncodeClientRequest(method string, args interface{}) ([]byte, error) {
c := &ClientRequest{
Method: method,
Params: [1]interface{}{args},
Id:     uint64(rand.Int63()),
}
return json.Marshal(c)
}
// DecodeClientResponse decodes the response body of a client request into
// the interface reply.
func DecodeClientResponse(r io.Reader, reply interface{}) error {
var c ClientResponse
if err := json.NewDecoder(r).Decode(&c); err != nil {
return err
}
if c.Error != nil {
return fmt.Errorf("%v", c.Error)
}
if c.Result == nil {
return errors.New("result is null")
}
return json.Unmarshal(*c.Result, reply)
}


func Execute(method string, req, res interface{}) error {
buf, _ := EncodeClientRequest(method, req)
body := bytes.NewBuffer(buf)

r, _ := http.NewRequest("POST", "http://localhost:8080/rpc/", body)
r.Header.Set("Content-Type", "application/json")

client := &http.Client{}
resp, _ := client.Do(r)
	return DecodeClientResponse(resp.Body, res)
}

//MAIN Function to invoke StockService and CheckPortfolio Service
func main() {
var TradeIduser string
var Investment string
var StockInformation string
var Investmentfinal float64
var reply StocksinfoReply
var response CheckPortfolioResponse

   if (len(os.Args) >= 3)  {
    
StockInformation = os.Args[1]
Investment = os.Args[2]
Investmentfinal,_ = strconv.ParseFloat(Investment,64)    
if err := Execute("GetStockService.Say", &StocksInput{StockInformation,(float32)(Investmentfinal)}, &reply); err != nil {
fmt.Println("Error!" , err)
}
fmt.Println("Your current stock investment details")
fmt.Println("tradeId",reply.TradeId)
fmt.Println("stocks",reply.Stocks)
fmt.Println("unvestedAmount",reply.UnvestedAmount) 

} else {

TradeIduser = os.Args[1]
var CheckPortFolioObj CheckPortFolioInput
CheckPortFolioObj.TradeId,_=strconv.Atoi(TradeIduser)
if err := Execute("CheckPortfolioService.Say", &CheckPortFolioObj, &response); err != nil {
fmt.Println("Error!" , err)
}
fmt.Println("Check Portfolio Response")
fmt.Println("stocks",response.Stocks)
fmt.Println("currentMarketValue",response.CurrentMarketValue)
fmt.Println("unvestedAmount",response.UnvestedAmounts)
}
}