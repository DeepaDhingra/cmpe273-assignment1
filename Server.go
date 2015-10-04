package main

import (
   "github.com/gorilla/rpc"
   "github.com/gorilla/rpc/json"
   "net/http"
   "log"
)  

func init() {
log.Printf("inside server")
   jsonRPC := rpc.NewServer()
   jsonCodec := json.NewCodec()
   jsonRPC.RegisterCodec(jsonCodec, "application/json")
   jsonRPC.RegisterCodec(jsonCodec, "application/json; charset=UTF-8") // For firefox 11 and other browsers which append the charset=UTF-8
  //Registration of both the Services
   jsonRPC.RegisterService(new(GetStockService), "") 
   jsonRPC.RegisterService(new(CheckPortfolioService), "")
    
   http.Handle("/rpc", jsonRPC)
   http.ListenAndServe(":8080", jsonRPC)  
   
    
     
}