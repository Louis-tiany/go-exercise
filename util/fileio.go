/*
* File     : read_file.go
* Author   : *
* Mail     : *
* Creation : Sat 11 Mar 2023 06:03:29 PM CST
 */

package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"stock/model"
)

func ReadFile() []string {
	filePath := "./file.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic("open file error")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func WriteFile(content []string) {
	filePath := "data.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		panic("open file error!")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range content {
		writer.WriteString(line)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func WriteJsonFile(data []model.Stock) {
	//打开文件
	file, _ := os.OpenFile("./first.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()
	//创建encoder 数据输出到file中
	encoder := json.NewEncoder(file)
	// json
	err := encoder.Encode(data)
	//异常处理
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("编码成功")
}

func ReadJsonFile() {
	file, _ := os.OpenFile("./first.json", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	//创建map，用于接收解码好的数据
	dataMap1 := make(map[string]interface{})
	//创建文件的解码器
	decoder := json.NewDecoder(file)
	//解码文件中的数据，丢入dataMap所在的内存
	err8 := decoder.Decode(&dataMap1)
	if err8 != nil {
		fmt.Println(err8)
		return
	}
	fmt.Println("解码成功", dataMap1)
}
