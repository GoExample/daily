package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)

func InitLogger(flag string) {
	switch flag {
	case "prod":
		file, err := os.OpenFile("/tmp/ip-issue.log",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("Failed to open error log file: %s\n", err)
			os.Exit(-1)
		}
		Trace = log.New(io.MultiWriter(file),
			"TRACE: ", log.LstdFlags|log.Lshortfile)

		Info = log.New(io.MultiWriter(file),
			"INFO: ", log.LstdFlags|log.Lshortfile)

		Warning = log.New(io.MultiWriter(file),
			"WARNING: ", log.LstdFlags|log.Lshortfile)

		Error = log.New(io.MultiWriter(file),
			"ERROR: ", log.LstdFlags|log.Lshortfile)
	case "test":
		Trace = log.New(ioutil.Discard,
			"TRACE: ", log.LstdFlags|log.Lshortfile)
		Info = log.New(ioutil.Discard,
			"INFO: ", log.LstdFlags|log.Lshortfile)
		Warning = log.New(ioutil.Discard,
			"WARNING: ", log.LstdFlags|log.Lshortfile)
		Error = log.New(ioutil.Discard,
			"ERROR: ", log.LstdFlags|log.Lshortfile)
	case "debug":
		Trace = log.New(os.Stdout,
			"TRACE: ", log.LstdFlags|log.Lshortfile)
		Info = log.New(os.Stdout,
			"INFO: ", log.LstdFlags|log.Lshortfile)
		Warning = log.New(os.Stdout,
			"WARNING: ", log.LstdFlags|log.Lshortfile)
		Error = log.New(os.Stdout,
			"ERROR: ", log.LstdFlags|log.Lshortfile)
	default:
		log.Printf("Invalid version flag: %s\n", flag)
		os.Exit(-1)
	}
}
