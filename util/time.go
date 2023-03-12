/*
* File     : util/time.go
* Author   : *
* Mail     : *
* Creation : Sat 11 Mar 2023 10:28:40 PM CST
 */
package util

import (
	"fmt"
	"time"
)

func Test() {
	GetTime()
	GetDate()
	FormatTime()
	CompareTime()
	GetSpecialTime()
	fmt.Println(IsTradingDay())
}

func GetTime() {
	//  10 位时间戳（秒）
	timestamp := time.Now().Unix()
	// 19 位时间戳（纳秒）
	timestampNano := time.Now().UnixNano()
	// 13 位时间戳（毫秒）
	timestampMillisecond := timestampNano / 1e6

	fmt.Printf("时间戳（秒）：%v\n", timestamp)
	fmt.Printf("时间戳（纳秒）：%v\n", timestampNano)
	fmt.Printf("时间戳（毫秒）：%v\n", timestampMillisecond)
	fmt.Printf("时间戳（纳秒转换为秒）：%v\n", timestampNano/1e9)
}

func GetDate() {
	nowTime := time.Now()
	// 获取具体年月日
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	nanosecond := time.Now().Nanosecond()

	fmt.Println(nowTime)
	fmt.Println(year, month, day, hour, minute, second, nanosecond)
	// 如下毫秒 = 0 精确到秒
	currentDate := time.Date(year, month, day, hour, minute, second, 0, time.Local)
	fmt.Println(currentDate)
}

func FormatTime() {
	// 方式一：
	timeFormat := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeFormat)
	// 方式二：先获取当前时间戳
	timestamp := time.Now().Unix()
	// 再格式化时间戳转化为日期
	datetime := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	fmt.Println(datetime)

	datetime = time.Unix(timestamp, 0).Format("2006-01-02 15:04")
	fmt.Println(datetime)

	datetime = "2020-11-26 11:25:39"
	formatTime, _ := time.Parse("2006-01-02 15:04:05", datetime)
	fmt.Println(formatTime)
	// 2020-11-26 11:25:39 +0000 UTC
	// 输出
	// 2020-11-26 11:29:11
	// 2020-11-26 11:29:11
	datetime = "2020-11-26 11:25:39"
	formatTime, _ = time.Parse("2006-01-02 15:04:05", datetime)
	// 再上一步的基础上使用 Unix() 函数
	fmt.Println(formatTime.Unix())
	// 1606389939
}

func CompareTime() {
	time1 := "2021-03-19 09:23:29"
	time2 := "2021-03-20 08:50:29"
	//先把时间字符串格式化成相同的时间类型
	t1, err1 := time.Parse("2006-01-02 15:04:05", time1)
	t2, err2 := time.Parse("2006-01-02 15:04:05", time2)
	if err1 == nil && err2 == nil && t1.Before(t2) {
		fmt.Println("true")
	}
	time3 := "2021-03-20 08:50:29"
	time4 := "2021-03-20 08:50:29"
	//先把时间字符串格式化成相同的时间类型
	t3, _ := time.Parse("2006-01-02 15:04:05", time3)
	t4, _ := time.Parse("2006-01-02 15:04:05", time4)
	if err1 == nil && err2 == nil && t3.Equal(t4) {
		fmt.Println("true")
	}
}

func GetSpecialTime() {
	// today zero time
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	timeStr := addTime.Format("2006-01-02")
	fmt.Println(timeStr) // 2022-04-15

	ts := time.Now().AddDate(0, 0, -1)
	time := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location())
	timeStr = time.Format("2006/01/02")
	fmt.Println(timeStr) // 2022-04-14
}
