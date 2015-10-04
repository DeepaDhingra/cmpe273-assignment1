# cmpe273-assignment1
Assignment Description:
This repository contains a collection of Go programs and libraries that demonstrate the Virtual Stock Trading System. The system uses the real-time pricing via Yahoo finance API and supports USD currency only. The system will have 2 components: client and server over JSON-RPC.

Server: The trading engine will have JSON-RPC interface for the above features.

Client: The JSON-RPC client will take command line input and send requests to the server.

How  the assignment Works:
The application takes command line arguments:

##	Buying stocks Service:

a) First start server by command:
go run Server.go StockService.go CheckPortfolioService.go Client.go

b) Then start Client.go with command line arguments 
example : go run Client.go GOOG:50%,YHOO:50% 3000

This command runs the invokes the method registered for StockService.go file and generates the output as 
desired in assignment :

“tradeId”: number
“stocks”: string (E.g. “GOOG:100:$500.25”, “YHOO:200:$31.40”)
“unvestedAmount”: float32)


##	Checking your portfolio (loss/gain)

Stock Service generates a TradeId which a user uses to check his/her portfolio.
 To check the portfolio user again gives the command line arguments example:

a) go run Client.go 2106 (e.g. 2106 is the TradeID here, generated in first step)
This will execute the second service CheckPortfolioService.go and generates the response as desired in the assignment

“stocks”: string (E.g. “GOOG:100:+$520.25”, “YHOO:200:-$30.40”)
“currentMarketValue” : float32
“unvestedAmount”: float32


