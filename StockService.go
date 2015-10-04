package main

import (
    "strconv"
 	"net/http"
	"io/ioutil"
   	"encoding/json"
   	"strings"
   	"time"
   	"math/rand"
   		)

//Structure to fetch URL records
type Urlrecord struct {
	Query struct {
	Results struct {
			Quote struct {
			Ask float32 `json:",string"`
			} `json:"quote"`
		} `json:"results"`
	} `json:"query"`
	}

var stockObjMap[] Stocksinfo
// To hold array of multiple stocks + MAP
type Holder struct {
	Hold []Stocksinfo
}

var AccessHolder Holder
// used to have key value pair of different stocks and their attributes
var M  map[int] Holder 

type GetStockService struct{}

func(h *GetStockService) Say(r *http.Request, args *StocksInput, reply *StocksinfoReply) error {
var index int
var a,b int
var count int
M = make(map[int]Holder)
a =0
count=0
b=0

    var fields[] string

f := func(c rune) bool {
	return c == ',' || c == ':' || c == '%'
    }
fields = strings.FieldsFunc(args.StockSymbolAndPercentage, f)
for index = 0; index < len(fields)-1; index+=2{
count+=1
}
var stockobj[] Stocksinfo = make([]Stocksinfo,count)
for index = 0; index < len(fields)-1; index+=2{
stockobj[a].Name = fields[index]
stockobj[a].TotalBudget = (float32) (args.Budget)
a+=1
}

var indexVar int

for indexVar = 1; indexVar < len(fields); indexVar+=2{
stockobj[b].Percentage = fields[indexVar]
b+=1
}

TotalUnvestedAmount:= make([]float32,count)
TotalStockMessage:= make([]string,count)
var FinalUnvestedAmount float32
FinalUnvestedAmount=0.0
var FinalStockMessage string
FinalUnvestedAmount=0.0
 

for index = 0; index < count; index+=1{
s := []string{}
s = append(s, "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20yahoo.finance.quotes%20where%20symbol%20in%20(")
s = append(s,"'")
s = append(s,stockobj[index].Name)
s = append(s,"'")
s = append(s,")&format=json&diagnostics=true&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys&callback=")

var url string = s[0] + s[1] +s[2] +s[3] + s[4]
record,_:= GetContentDetails(url)
var amount float32
var  Number_of_Shares float32
var UnvestedAmount float32 
var floatPercentage float64
var s1 string
 
s1= stockobj[index].Percentage
floatPercentage,_ =  strconv.ParseFloat(s1, 64)
amount = (float32)(stockobj[index].TotalBudget * (float32)(floatPercentage))/ 100
var Exact_Number_of_Shares int
    
    if record.Query.Results.Quote.Ask > amount {
    		    Exact_Number_of_Shares=0;
         UnvestedAmount = (float32) (amount - (float32(Exact_Number_of_Shares) * record.Query.Results.Quote.Ask))
    } else  {
        
                Number_of_Shares = (float32) (amount / record.Query.Results.Quote.Ask)
        
      
        Exact_Number_of_Shares = int (Number_of_Shares)
        
        UnvestedAmount = (float32) (amount - (float32(Exact_Number_of_Shares) * record.Query.Results.Quote.Ask))
        
        
       }

       TotalUnvestedAmount[index]=(float32) (UnvestedAmount)		
       FinalUnvestedAmount= (float32)(FinalUnvestedAmount+ TotalUnvestedAmount[index])
		

 	   stockobj[index].NumberOfShares = Exact_Number_of_Shares
       stockobj[index].CurrentStockPrice= strconv.FormatFloat((float64)(record.Query.Results.Quote.Ask), 'f', 2, 64)

msg := []string{}
msg = append(msg,stockobj[index].Name)
msg = append(msg,":")
msg = append(msg, strconv.Itoa(Exact_Number_of_Shares))
msg = append(msg,":")
msg = append(msg,"$")
msg = append(msg,strconv.FormatFloat((float64)(record.Query.Results.Quote.Ask), 'f', 2, 64))


		TotalStockMessage[index] = msg[0] + msg[1] +msg[2] +msg[3] + msg[4] + msg[5]	
}

FinalStockMessage = strings.Join(TotalStockMessage, ", ")
stockobj[0].UnvestedAmount = (float32)(FinalUnvestedAmount)

reply.Stocks = FinalStockMessage
reply.UnvestedAmount = (float32)(FinalUnvestedAmount)
reply.TradeId = TradeIdGenerator();

AccessHolder.Hold = stockobj;

stockObjMap =AccessHolder.Hold

M[reply.TradeId] = Holder{stockObjMap}
	
return nil;
}

func getContent(url string) ([]byte, error) {
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

func GetContentDetails(ip string) (*Urlrecord, error) {
content, err := getContent(ip)
if err != nil {
}
	// Fill the record with the data from the JSON
	var record Urlrecord
	err = json.Unmarshal(content, &record)
	if err != nil {
	}
	return &record, err
	}
	
func TradeIdGenerator() int  {

	RandomNumber1 := rand.NewSource(time.Now().UnixNano())
    FinalRandomNumber := rand.New(RandomNumber1)
    RandomId := FinalRandomNumber.Intn(10000)
    return RandomId
    }