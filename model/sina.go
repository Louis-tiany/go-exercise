/*
* File     : sina.go
* Author   : *
* Mail     : *
* Creation : Sat 04 Mar 2023 03:55:40 PM CST
 */
package model

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func (DB *DbConn) CreateStockRecord(records []Stock) {
	for i := 0; i < len(records); i++ {
		DB.Db.Create(&records[i])
	}
}

func GetSinaData() []Stock {
	client := &http.Client{}
	url := "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/MoneyFlow.ssl_bkzj_ssggzj"
	req, err := http.NewRequest("GET", url, nil)

	query := req.URL.Query()
	query.Add("page", "1")
	query.Add("num", "500")
	query.Add("sort", "netamount")
	query.Add("asc", "0")
	query.Add("bankuai", "")
	query.Add("shichang", "")
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)

	list := gjson.Parse(bodyStr).Array()
	stocks := []Stock{}

	for _, item := range list {
		symbol := item.Get("symbol").String()
		name := item.Get("name").String()
		trade := item.Get("trade").Float()
		inAmont := item.Get("inamount").Float()
		outAmont := item.Get("outamount").Float()
		ratioAmount := item.Get("ratioamount").Float()
		netAmount := item.Get("netamount").Float()

		stock := Stock{
			Symbol:      symbol,
			Name:        name,
			Trade:       trade,
			InAmont:     inAmont,
			OutAmont:    outAmont,
			NetAmount:   netAmount,
			RatioAmount: ratioAmount,
		}
		stocks = append(stocks, stock)
	}
	return stocks
}
