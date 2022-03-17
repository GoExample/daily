package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
)

type srvConfig struct {
	tlsRandSrvConfig *tls.Config
	tlsSrvConfig     *tls.Config
	tlsCliConfig     *tls.Config
	certFile         string
	keyFile          string
	caFile           string
	pubKeySum        string
	inStream         string
	outStream        string
	listenPort       string
	manageKey        string
	comKey           string
	shell            string
	check            int
	retry            int
}

var RELEASE = "true"
var srvCfg srvConfig

func main() {
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))

	fmt.Println(path.Join("a/b", "../../../xyz"))

	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))

	abs, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	var pathName interface{}
	if RELEASE == "true" {
		pathName, _ = Asset("conf/conf.ini")
	} else {
		pathName = abs + "/conf/conf.ini"
	}
	cfg, err := ini.Load(pathName)
	if err != nil {
		fmt.Printf("fail to read file: %v", err)
		os.Exit(1)
	}

	if RELEASE == "true" {
		srvCfg.caFile = cfg.Section("connection").Key("ca").String()
		srvCfg.certFile = cfg.Section("connection").Key("cert").String()
		srvCfg.keyFile = cfg.Section("connection").Key("key").String()
	} else {
		srvCfg.caFile = abs + "/" + cfg.Section("connection").Key("ca").String()
		srvCfg.certFile = abs + "/" + cfg.Section("connection").Key("cert").String()
		srvCfg.keyFile = abs + "/" + cfg.Section("connection").Key("key").String()
	}

	fmt.Println(srvCfg.caFile)
	fmt.Println(srvCfg.certFile)
	fmt.Println(srvCfg.keyFile)

	ca1, _ := ioutil.ReadFile(abs + "/" + srvCfg.caFile)
	ca2, _ := Asset(srvCfg.caFile)

	fmt.Println("==========================")
	fmt.Println(bytes.Equal(ca1, ca2))
	fmt.Println(bytes.Equal(ca1, nil))
	fmt.Println(bytes.Equal(ca2, nil))

	fmt.Println(srvCfg.certFile, srvCfg.keyFile)
	cert1, _ := tls.LoadX509KeyPair(abs+"/"+srvCfg.certFile, abs+"/"+srvCfg.keyFile)

	certContent, _ := Asset(srvCfg.certFile)
	keyContent, _ := Asset(srvCfg.keyFile)
	cert2, _ := tls.X509KeyPair(certContent, keyContent)

	fmt.Println(cert1)
	fmt.Println(cert2)
	fmt.Println(reflect.TypeOf(cert1))
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, "Hello world!") })
	err = r.Run(":8888")
	if err != nil {
		return
	}
}
