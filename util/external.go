package util

import (
	"encoding/json"
	"io"
	"net/http"
	"stock/model"
	"time"
)

func IsTradingDay() bool {
	todayStatus := GetHoliday()
	return todayStatus.Status == 0
}

func GetHoliday() *model.HodayStatus {
	todayStr := time.Now().Format("2006-01-02")
	url := "http://api.haoshenqi.top/holiday" + "?date=" + todayStr
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var hodayStatus []model.HodayStatus
	json.Unmarshal(body, &hodayStatus)
	return &hodayStatus[0]
}
