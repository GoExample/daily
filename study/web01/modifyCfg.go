package main

import (
    "bufio"
    "bytes"
    "fmt"
    "gopkg.in/ini.v1"
)

func modifyCfg(cfg *ini.File) (*bytes.Buffer, error) {
    //dir, _ := os.Getwd()
    //path := dir + "/conf/"
    // 典型读取操作，默认分区可以使用空字符串表示
    fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
    fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

    // 我们可以做一些候选值限制的操作
    fmt.Println("Server Protocol:",
        cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
    // 如果读取的值不在候选列表内，则会回退使用提供的默认值
    fmt.Println("Email Protocol:",
        cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

    // 试一试自动类型转换
    fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
    fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

    // 差不多了，修改某个值然后进行保存
    cfg.Section("").Key("app_mode").SetValue("production")
    //cfg.SaveTo(path + "my.ini.local")

    //b := bytes.NewBuffer(make([]byte, 0))
    b := new(bytes.Buffer)
    w := bufio.NewWriter(b)
    _, err := cfg.WriteTo(w)
    if err != nil {
        return nil, err
    }
    err = w.Flush()
    if err != nil {
        return nil, err
    }
    return b, nil
}
