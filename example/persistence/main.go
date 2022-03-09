package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

var persistPath = "/tmp/StateRecord"

type Info struct {
	Name string
	Age  int
	Sex  string
}

type StateRecord struct {
	Index int
	Date  string
}

func (sr *StateRecord) Load() error {
	data, err := ioutil.ReadFile(persistPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, sr)
	if err != nil {
		return err
	}

	return nil
}

func (sr *StateRecord) Save() error {
	content, err := json.Marshal(sr)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(persistPath, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	datetimeOperate()
	jsonSerialization()
	jsonPersistence()
}

func datetimeOperate() {
	label := "Datetime Operate Example"
	fmt.Println("\n" + strings.Repeat("-", 80/2-len(label)) + label + strings.Repeat("-", 80/2-len(label)))

	// 2006-01-02 15:04:05 据说是 golang 的诞生时间，固定写法，按照612345去记忆
	layout := "2006-01-02 15:04:05"

	timeStr := "2021-05-17 11:00:00"          // 认为服务器时间
	location, _ := time.LoadLocation("Local") // timeStr是北京时间，注意使用Local
	localTime, _ := time.ParseInLocation(layout, timeStr, location)
	fmt.Println(localTime.UTC().Format(layout)) // 转为UTC时间

	compareTimeStr := "2021-05-17 02:00:00" // 认为UTC时间
	location, _ = time.LoadLocation("UTC")  // timeStr是北京时间，注意使用UTC
	compareTime, _ := time.ParseInLocation(layout, compareTimeStr, location)
	fmt.Println(compareTime.Local().Format(layout)) // 转为服务器时间

	fmt.Println(compareTime.Before(localTime)) // Before 会统一时间基准点
	fmt.Println(compareTime.After(localTime))  // After 会统一时间基准点

	resultA, _ := time.Parse(layout, time.Now().Format(layout))
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	resultB, _ := time.Parse(layout, time.Now().Format(layout))

	resultC := resultB.Sub(resultA)
	randomValue := rand.Intn(10)
	if resultC.Seconds() > float64(randomValue) {
		fmt.Printf("resultA = %s, resultB = %s, resultB - resultA = %s\n", resultA, resultB, resultC)
	} else {
		fmt.Printf("resultA = %s, resultB = %s, resultB - resultA = %s\n", resultA, resultB, resultC)
	}
}

func jsonSerialization() {
	label := "Json Serialization Example"
	fmt.Println("\n" + strings.Repeat("-", 80/2-len(label)) + label + strings.Repeat("-", 80/2-len(label)))
	infos := []Info{
		{
			Name: "Sophia",
			Age:  23,
			Sex:  "female",
		},
		{
			Name: "Benjie",
			Age:  24,
			Sex:  "male",
		},
	}
	data, err := json.Marshal(infos)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))

	var info []Info
	if err := json.Unmarshal(data, &info); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(info)
	fmt.Println(reflect.TypeOf(info))
}

func jsonPersistence() {
	label := "Json Persistence Example"
	fmt.Println("\n" + strings.Repeat("-", 80/2-len(label)) + label + strings.Repeat("-", 80/2-len(label)))
	var sr StateRecord
	fmt.Printf("声明 %v\n", sr)

	sr = StateRecord{Index: 1, Date: time.Now().Format("2006-01-02 15:04:05")}
	fmt.Printf("初始化 %v\n", sr)

	err := sr.Save()
	if err != nil {
		return
	}

	err = sr.Load()
	if err != nil {
		return
	}
	fmt.Printf("从文件读取 %v\n", sr)

	sr.Index = 2
	sr.Date = time.Now().Format("2006-01-02 15:04:05")
	err = sr.Save()
	if err != nil {
		return
	}
}
