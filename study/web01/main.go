package main

import (
    "fmt"
    "gopkg.in/ini.v1"
    "io/ioutil"
    "net/http"
    "os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
    buf, _ := ioutil.ReadFile("./hello.html")
    _, err := fmt.Fprintf(w, string(buf))
    if err != nil {
        return
    }
}

func main() {
    ff, _ := Asset("conf/conf.ini")
    //fmt.Println(reflect.TypeOf(ff))
    cfg, err := ini.Load(ff)
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println(cfg.Section("node").Key("nodes").String())

    dir, _ := os.Getwd()
    //fmt.Println(dir)
    path := dir + "/conf/"
    pathName := path + "my.ini"
    cfg, err = ini.Load(pathName)
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }
    b, err := modifyCfg(cfg)
    fmt.Sprintf("%v", b.String())
    /*
       fmt.Println(b.String())
       path = dir + "/conf/my.ini.local"
       f, err := os.Create(path)
       defer f.Close()
       f.WriteString(b.String())
    */

    http.HandleFunc("/hello", sayHello)
    http.ListenAndServe(
        ":9000",
        nil,
    )
    if err != nil {
        fmt.Printf("http serve failed, err:%v\n", err)
        return
    }
    studyChi()
}
