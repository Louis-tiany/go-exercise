/*
* File     : main.go
* Author   : *
* Mail     : *
* Creation : Sat 04 Mar 2023 10:59:02 PM CST
 */

package main

import (
	"fmt"
	. "stock/model"
	"stock/util"
)

func main() {
	util.Test()
	return
	stocks := GetSinaData()
	db := NewStockDb()
	fmt.Println(db)
	strs := []string{"hello", "world", "hello", "world"}
	fmt.Println(stocks)
	util.WriteFile(strs)
	// fmt.Println(util.ReadFile())

	util.WriteJsonFile(stocks)
	util.ReadJsonFile()

	db.CreateStockRecord(stocks)
}
