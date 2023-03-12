/*
* File     : common.go
* Author   : *
* Mail     : *
* Creation : Sat 11 Mar 2023 11:50:02 PM CST
 */
package model

type HodayStatus struct {
	Date   string `json:"date"`
	Day    int64  `json:"day"`
	Month  int64  `json:"month"`
	Status int64  `json:"status"`
	Year   int64  `json:"year"`
}
