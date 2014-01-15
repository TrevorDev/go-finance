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
	Ask="a"	
	Average_Daily_Volume="a2"				
	Ask_Size="a5"
	Bid="b"
	Ask_Real_time="b2"			
	Bid_Real_time="b3"
	Book_Value="b4"
	Bid_Size="b6"
	Change_And_Percent_Change="c"					
	Change="c1"
	Commission="c3"	
	Change_Real_time="c6"			
	After_Hours_Change_Real_time="c8"		
	Dividend_Per_Share="d"
	Last_Trade_Date="d1"	
	Trade_Date="d2"
	Earnings_Per_Share="e"	
	Error_Indication_returned_for_symbol_changed_Per_invalid="e1"											
	EPS_Estimate_Current_Year="e7"
	EPS_Estimate_Next_Year="e8"
	EPS_Estimate_Next_Quarter="e9"	
	Float_Shares="f6"
	Days_Low="g"
	Days_High="h"	
	Year_Low="j"
	Year_High="k"
	Holdings_Gain_Percent="g1"		
	Annualized_Gain="g3"
	Holdings_Gain="g4"
	Holdings_Gain_Percent_Real_time="g5"					
	Holdings_Gain_Real_time="g6"
	More_Info="i"
	Order_Book_Real_time="i5"				
	Market_Capitalization="j1"
	Market_Cap_Real_time="j3"
	EBITDA="j4"
	Change_From_Year_Low="j5"				
	Percent_Change_From_Year_Low="j6"		
	Last_Trade_Real_time_With_Time="k1"
	Change_Percent_Real_time="k2"
	Last_Trade_Size="k3"
	Change_From_Year_High="k4"		
	Percebt_Change_From_Year_High="k5"		
	Last_Trade_With_Time="l"
	Last_Trade_Price_Only="l1"	
	High_Limit="l2"
	Low_Limit="l3"
	Days_Range="m"	
	Days_Range_Real_time="m2"			
	Fifty_day_Moving_Average="m3"
	Two_Hundred_Day_Moving_Average="m4"	
	Change_From_200_day_Moving_Average="m5"		
	Percent_Change_From_200_day_Moving_Average="m6"		
	Change_From_50_day_Moving_Average="m7"
	Percent_Change_From_50_day_Moving_Average="m8"		
	Name="n"
	Notes="n4"
	Open="o"
	Previous_Close="p"		
	Price_Paid="p1"
	Change_in_Percent="p2"	
	Price_Per_Sales="p5"
	Price_Per_Book="p6"
	Ex_Dividend_Date="q"	
	Price_Per_Earning_Ratio="r"
	Dividend_Pay_Date="r1"		
	Price_Per_Earning_Ratio_Real_time="r2"	
	PEG_Ratio="r5"
	Price_Per_EPS_Estimate_Current_Year="r6"						
	Price_Per_EPS_Estimate_Next_Year="r7"
	Symbol="s"
	Shares_Owned="s1"		
	Short_Ratio="s7"
	Last_Trade_Time="t1"	
	Trade_Links="t6"
	Ticker_Trend="t7"
	One_Year_Target_Price="t8"	
	Volume="v"
	Holdings_Value="v1"	
	Holdings_Value_Real_time="v7"			
	Year_Range="w"
	Days_Value_Change="w1"	
	Days_Value_Change_Real_time="w4"			
	Stock_Exchange="x"
	Dividend_Yield="y"
)

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

	resp, err := timedHttp.Get("http://download.finance.yahoo.com/d/quotes.csv?s="+strings.Join(symbols,"+")+"&f="+strings.Join(attributes,""))
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

																									
