package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Services struct {
	Log service.Logger
	Srv *http.Server
	Cfg *service.Config
}

// ExecPath 获取可执行文件的绝对路径
func ExecPath() string {
	file, e := os.Executable()
	if e != nil {
		log.Printf("Executable file path error : %s\n", e.Error())
	}
	path := filepath.Dir(file)
	return path
}

// 获取 service 对象
func getSrv() service.Service {
	File, err := os.Create(ExecPath() + "/http-server.log")
	if err != nil {
		File = os.Stdout
	}
	defer File.Close()

	log.SetOutput(File)

	s := &Services{
		Cfg: &service.Config{
			Name:        "goWeb",
			DisplayName: "goWeb",
			Description: "基于gin的web服务",
		}}
	serv, er := service.New(s, s.Cfg)
	if er != nil {
		log.Printf("Set logger error:%s\n", er.Error())
	}
	s.Log, er = serv.SystemLogger(nil)
	return serv
}

// Start 启动windows服务
func (srv *Services) Start(s service.Service) error {
	if srv.Log != nil {
		srv.Log.Info("Start run http server")
	}
	go srv.StarServer()
	return nil
}

// Stop 停止windows服务
func (srv *Services) Stop(s service.Service) error {
	if srv.Log != nil {
		srv.Log.Info("Start stop http server")
	}
	log.Println("Server exiting")
	return srv.Srv.Shutdown(context.Background())
}

// StarServer 运行gin web服务
func (srv *Services) StarServer() {
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create(ExecPath() + "/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv.Srv = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	srv.Srv.ListenAndServe()
}

func main() {
	s := getSrv()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err := s.Install()
			if err != nil {
				log.Fatalf("Install service error:%s\n", err.Error())
			}
			fmt.Printf("服务已安装")
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				log.Fatalf("Uninstall service error:%s\n", err.Error())
			}
			fmt.Printf("服务已卸载")
		case "start":
			err := s.Start()
			if err != nil {
				log.Fatalf("Start service error:%s\n", err.Error())
			}
			fmt.Printf("服务已启动")
		case "stop":
			err := s.Stop()
			if err != nil {
				log.Fatalf("top service error:%s\n", err.Error())
			}
			fmt.Printf("服务已关闭")
		}
		return
	}
	err := s.Run()
	if err != nil {
		log.Fatalf("Run programe error:%s\n", err.Error())
	}
}
