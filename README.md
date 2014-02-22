Stock/Financial data at your fingertips
======================================

[go-finance] Aims to provide an easy to use stock data client api

* Uses Yahoo Finance csv http api

## Install

`go get github.com/TrevorDev/go-finance`

### Setup API (stable)
	Add the following to the top of your .go file:

	import (
		"github.com/TrevorDev/go-finance"
	)

### Basic Usage
	
	out, err := finance.GetStockInfo([]string{"MSFT", "GOOG"},[]string{finance.Last_Trade_Price_Only,finance.Price_Per_Earning_Ratio})
 	if(err!=nil){
 		//This may happen often if yahoo's server is under high volume
 		fmt.Println(err)
 	}else{
 		//outputs map of stock symbols and their specified attributes given above
 		fmt.Println(out)

 		//outputs microsofts current price per earing ratio as a string
 		fmt.Println(out["MSFT"][finance.Price_Per_Earning_Ratio])
 	}

### Possibly supported attributes

	finance.Ask
	finance.Average_Daily_Volume				
	finance.Ask_Size
	finance.Bid
	finance.Ask_Real_time			
	finance.Bid_Real_time
	finance.Book_Value
	finance.Bid_Size
	finance.Change_And_Percent_Change				
	finance.Change
	finance.Commission	
	finance.Change_Real_time			
	finance.After_Hours_Change_Real_time		
	finance.Dividend_Per_Share
	finance.Last_Trade_Date	
	finance.Trade_Date
	finance.Earnings_Per_Share
	finance.Error_Indication_returned_for_symbol_changed_Per_invalid											
	finance.EPS_Estimate_Current_Year
	finance.EPS_Estimate_Next_Year
	finance.EPS_Estimate_Next_Quarter	
	finance.Float_Shares
	finance.Days_Low
	finance.Days_High
	finance.Year_Low
	finance.Year_High
	finance.Holdings_Gain_Percent		
	finance.Annualized_Gain
	finance.Holdings_Gain
	finance.Holdings_Gain_Percent_Real_time					
	finance.Holdings_Gain_Real_time
	finance.More_Info
	finance.Order_Book_Real_time				
	finance.Market_Capitalization
	finance.Market_Cap_Real_time
	finance.EBITDA
	finance.Change_From_Year_Low				
	finance.Percent_Change_From_Year_Low		
	finance.Last_Trade_Real_time_With_Time
	finance.Change_Percent_Real_time
	finance.Last_Trade_Size
	finance.Change_From_Year_High		
	finance.Percebt_Change_From_Year_High		
	finance.Last_Trade_With_Time
	finance.Last_Trade_Price_Only	
	finance.High_Limit
	finance.Low_Limit
	finance.Days_Range
	finance.Days_Range_Real_time			
	finance.Fifty_day_Moving_Average
	finance.Two_Hundred_Day_Moving_Average	
	finance.Change_From_200_day_Moving_Average		
	finance.Percent_Change_From_200_day_Moving_Average		
	finance.Change_From_50_day_Moving_Average
	finance.Percent_Change_From_50_day_Moving_Average		
	finance.Name
	finance.Notes
	finance.Open
	finance.Previous_Close	
	finance.Price_Paid
	finance.Change_in_Percent	
	finance.Price_Per_Sales
	finance.Price_Per_Book
	finance.Ex_Dividend_Date
	finance.Price_Per_Earning_Ratio
	finance.Dividend_Pay_Date		
	finance.Price_Per_Earning_Ratio_Real_time	
	finance.PEG_Ratio
	finance.Price_Per_EPS_Estimate_Current_Year						
	finance.Price_Per_EPS_Estimate_Next_Year
	finance.Symbol
	finance.Shares_Owned		
	finance.Short_Ratio
	finance.Last_Trade_Time	
	finance.Trade_Links
	finance.Ticker_Trend
	finance.One_Year_Target_Price	
	finance.Volume
	finance.Holdings_Value	
	finance.Holdings_Value_Real_time			
	finance.Year_Range
	finance.Days_Value_Change	
	finance.Days_Value_Change_Real_time			
	finance.Stock_Exchange
	finance.Dividend_Yield

## Contributors

- [Trevor Baron](https://github.com/TrevorDev) (author)
