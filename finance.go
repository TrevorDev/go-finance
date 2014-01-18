package finance

import (
	"net/http"
	"net"
	"time"
	"io/ioutil"
	//"log"
	"strings"
	"errors"
)

//List of constants that may be supported by yahoo finance api
const (
	Ask = "Ask" 	
	Average_Daily_Volume = "Average_Daily_Volume" 				
	Ask_Size = "Ask_Size" 
	Bid = "Bid" 
	Ask_Real_time = "Ask_Real_time" 			
	Bid_Real_time = "Bid_Real_time" 
	Book_Value = "Book_Value" 
	Bid_Size = "Bid_Size" 
	Change_And_Percent_Change = "Change_And_Percent_Change" 					
	Change = "Change" 
	Commission = "Commission" 	
	Change_Real_time = "Change_Real_time" 			
	After_Hours_Change_Real_time = "After_Hours_Change_Real_time" 		
	Dividend_Per_Share = "Dividend_Per_Share" 
	Last_Trade_Date = "Last_Trade_Date" 	
	Trade_Date = "Trade_Date" 
	Earnings_Per_Share = "Earnings_Per_Share" 	
	Error_Indication_returned_for_symbol_changed_Per_invalid = "Error_Indication_returned_for_symbol_changed_Per_invalid" 											
	EPS_Estimate_Current_Year = "EPS_Estimate_Current_Year" 
	EPS_Estimate_Next_Year = "EPS_Estimate_Next_Year" 
	EPS_Estimate_Next_Quarter = "EPS_Estimate_Next_Quarter" 	
	Float_Shares = "Float_Shares" 
	Days_Low = "Days_Low" 
	Days_High = "Days_High" 	
	Year_Low = "Year_Low" 
	Year_High = "Year_High" 
	Holdings_Gain_Percent = "Holdings_Gain_Percent" 		
	Annualized_Gain = "Annualized_Gain" 
	Holdings_Gain = "Holdings_Gain" 
	Holdings_Gain_Percent_Real_time = "Holdings_Gain_Percent_Real_time" 					
	Holdings_Gain_Real_time = "Holdings_Gain_Real_time" 
	More_Info = "More_Info" 
	Order_Book_Real_time = "Order_Book_Real_time" 				
	Market_Capitalization = "Market_Capitalization" 
	Market_Cap_Real_time = "Market_Cap_Real_time" 
	EBITDA = "EBITDA" 
	Change_From_Year_Low = "Change_From_Year_Low" 				
	Percent_Change_From_Year_Low = "Percent_Change_From_Year_Low" 		
	Last_Trade_Real_time_With_Time = "Last_Trade_Real_time_With_Time" 
	Change_Percent_Real_time = "Change_Percent_Real_time" 
	Last_Trade_Size = "Last_Trade_Size" 
	Change_From_Year_High = "Change_From_Year_High" 		
	Percebt_Change_From_Year_High = "Percebt_Change_From_Year_High" 		
	Last_Trade_With_Time = "Last_Trade_With_Time" 
	Last_Trade_Price_Only = "Last_Trade_Price_Only" 	
	High_Limit = "High_Limit" 
	Low_Limit = "Low_Limit" 
	Days_Range = "Days_Range" 	
	Days_Range_Real_time = "Days_Range_Real_time" 			
	Fifty_day_Moving_Average = "Fifty_day_Moving_Average" 
	Two_Hundred_Day_Moving_Average = "Two_Hundred_Day_Moving_Average" 	
	Change_From_200_day_Moving_Average = "Change_From_200_day_Moving_Average" 		
	Percent_Change_From_200_day_Moving_Average = "Percent_Change_From_200_day_Moving_Average" 		
	Change_From_50_day_Moving_Average = "Change_From_50_day_Moving_Average" 
	Percent_Change_From_50_day_Moving_Average = "Percent_Change_From_50_day_Moving_Average" 		
	Name = "Name" 
	Notes = "Notes" 
	Open = "Open" 
	Previous_Close = "Previous_Close" 		
	Price_Paid = "Price_Paid" 
	Change_in_Percent = "Change_in_Percent" 	
	Price_Per_Sales = "Price_Per_Sales" 
	Price_Per_Book = "Price_Per_Book" 
	Ex_Dividend_Date = "Ex_Dividend_Date" 	
	Price_Per_Earning_Ratio = "Price_Per_Earning_Ratio" 
	Dividend_Pay_Date = "Dividend_Pay_Date" 		
	Price_Per_Earning_Ratio_Real_time = "Price_Per_Earning_Ratio_Real_time" 	
	PEG_Ratio = "PEG_Ratio" 
	Price_Per_EPS_Estimate_Current_Year = "Price_Per_EPS_Estimate_Current_Year" 						
	Price_Per_EPS_Estimate_Next_Year = "Price_Per_EPS_Estimate_Next_Year" 
	Symbol = "Symbol" 
	Shares_Owned = "Shares_Owned" 		
	Short_Ratio = "Short_Ratio" 
	Last_Trade_Time = "Last_Trade_Time" 	
	Trade_Links = "Trade_Links" 
	Ticker_Trend = "Ticker_Trend" 
	One_Year_Target_Price = "One_Year_Target_Price" 	
	Volume = "Volume" 
	Holdings_Value = "Holdings_Value" 	
	Holdings_Value_Real_time = "Holdings_Value_Real_time" 			
	Year_Range = "Year_Range" 
	Days_Value_Change = "Days_Value_Change" 	
	Days_Value_Change_Real_time = "Days_Value_Change_Real_time" 			
	Stock_Exchange = "Stock_Exchange" 
	Dividend_Yield = "Dividend_Yield"
)

func associatedQueryValue(s string)string {
		switch s {
			case "Ask":
				return "a"	
			case "Average_Daily_Volume":
				return "a2"				
			case "Ask_Size":
				return "a5"
			case "Bid":
				return "b"
			case "Ask_Real_time":
				return "b2"			
			case "Bid_Real_time":
				return "b3"
			case "Book_Value":
				return "b4"
			case "Bid_Size":
				return "b6"
			case "Change_And_Percent_Change":
				return "c"					
			case "Change":
				return "c1"
			case "Commission":
				return "c3"	
			case "Change_Real_time":
				return "c6"			
			case "After_Hours_Change_Real_time":
				return "c8"		
			case "Dividend_Per_Share":
				return "d"
			case "Last_Trade_Date":
				return "d1"	
			case "Trade_Date":
				return "d2"
			case "Earnings_Per_Share":
				return "e"	
			case "Error_Indication_		returned_for_symbol_changed_Per_invalid":
				return "e1"											
			case "EPS_Estimate_Current_Year":
				return "e7"
			case "EPS_Estimate_Next_Year":
				return "e8"
			case "EPS_Estimate_Next_Quarter":
				return "e9"	
			case "Float_Shares":
				return "f6"
			case "Days_Low":
				return "g"
			case "Days_High":
				return "h"	
			case "Year_Low":
				return "j"
			case "Year_High":
				return "k"
			case "Holdings_Gain_Percent":
				return "g1"		
			case "Annualized_Gain":
				return "g3"
			case "Holdings_Gain":
				return "g4"
			case "Holdings_Gain_Percent_Real_time":
				return "g5"					
			case "Holdings_Gain_Real_time":
				return "g6"
			case "More_Info":
				return "i"
			case "Order_Book_Real_time":
				return "i5"				
			case "Market_Capitalization":
				return "j1"
			case "Market_Cap_Real_time":
				return "j3"
			case "EBITDA":
				return "j4"
			case "Change_From_Year_Low":
				return "j5"				
			case "Percent_Change_From_Year_Low":
				return "j6"		
			case "Last_Trade_Real_time_With_Time":
				return "k1"
			case "Change_Percent_Real_time":
				return "k2"
			case "Last_Trade_Size":
				return "k3"
			case "Change_From_Year_High":
				return "k4"		
			case "Percebt_Change_From_Year_High":
				return "k5"		
			case "Last_Trade_With_Time":
				return "l"
			case "Last_Trade_Price_Only":
				return "l1"	
			case "High_Limit":
				return "l2"
			case "Low_Limit":
				return "l3"
			case "Days_Range":
				return "m"	
			case "Days_Range_Real_time":
				return "m2"			
			case "Fifty_day_Moving_Average":
				return "m3"
			case "Two_Hundred_Day_Moving_Average":
				return "m4"	
			case "Change_From_200_day_Moving_Average":
				return "m5"		
			case "Percent_Change_From_200_day_Moving_Average":
				return "m6"		
			case "Change_From_50_day_Moving_Average":
				return "m7"
			case "Percent_Change_From_50_day_Moving_Average":
				return "m8"		
			case "Name":
				return "n"
			case "Notes":
				return "n4"
			case "Open":
				return "o"
			case "Previous_Close":
				return "p"		
			case "Price_Paid":
				return "p1"
			case "Change_in_Percent":
				return "p2"	
			case "Price_Per_Sales":
				return "p5"
			case "Price_Per_Book":
				return "p6"
			case "Ex_Dividend_Date":
				return "q"	
			case "Price_Per_Earning_Ratio":
				return "r"
			case "Dividend_Pay_Date":
				return "r1"		
			case "Price_Per_Earning_Ratio_Real_time":
				return "r2"	
			case "PEG_Ratio":
				return "r5"
			case "Price_Per_EPS_Estimate_Current_Year":
				return "r6"						
			case "Price_Per_EPS_Estimate_Next_Year":
				return "r7"
			case "Symbol":
				return "s"
			case "Shares_Owned":
				return "s1"		
			case "Short_Ratio":
				return "s7"
			case "Last_Trade_Time":
				return "t1"	
			case "Trade_Links":
				return "t6"
			case "Ticker_Trend":
				return "t7"
			case "One_Year_Target_Price":
				return "t8"	
			case "Volume":
				return "v"
			case "Holdings_Value":
				return "v1"	
			case "Holdings_Value_Real_time":
				return "v7"			
			case "Year_Range":
				return "w"
			case "Days_Value_Change":
				return "w1"	
			case "Days_Value_Change_Real_time":
				return "w4"			
			case "Stock_Exchange":
				return "x"
			case "Dividend_Yield":
				return "y"
			default:
				return "NOTFOUND"
		}
	}

func createassociatedQueryArray(a []string) []string {
	ret := make([]string, len(a))
	for i:=0;i<len(a);i++ {
		ret[i] = associatedQueryValue(a[i])
	}
	return ret
}

// Use for timeouts from yahoo
func dialTimeout(timeoutLen time.Duration) (func(network, addr string) (net.Conn, error)) {
	    timeout := timeoutLen
    return func(network, addr string) (net.Conn, error) {
        return net.DialTimeout(network, addr, timeout)
    }
}

func getHttpClient(timeout time.Duration)http.Client {
	transport := http.Transport{
		Dial: dialTimeout(timeout),
	}
	client := http.Client{
		Transport: &transport,
	}
	return client;
}

func GetStockInfoWithTimeOut(symbols []string, attributes []string, timeout time.Duration) (map[string] map[string] string, error) {
	ret := make(map[string] map[string] string);
	
	timedHttp := getHttpClient(timeout)

	associatedQueryAr := createassociatedQueryArray(attributes)
	resp, err := timedHttp.Get("http://download.finance.yahoo.com/d/quotes.csv?s="+strings.Join(symbols,"+")+"&f="+strings.Join(associatedQueryAr,""))
	if err != nil { return ret, err }
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil { return ret, err }

	symbolRows := strings.Split(string(contents), "\n")
	if(len(symbolRows)-1!=len(symbols)){
		return ret, errors.New("stock: Api request returned invalid number of symbols. Got "+ string(len(symbolRows))+" returned")
	}

	for i:=0;i<len(symbols);i++ {
		ret[symbols[i]] = make(map[string] string)
		symboAttributes := strings.Split(symbolRows[i],",");
		for j,element := range symboAttributes {
			ret[symbols[i]][attributes[j]] = strings.TrimSpace(element)
		}
	}

	return ret, nil
}

func GetStockInfo(symbols []string, attributes []string) (map[string] map[string] string, error) {
	return GetStockInfoWithTimeOut(symbols, attributes, time.Duration(2 * time.Second));
}

																									
