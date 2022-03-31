package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"net"
	"net/http"
	"os"
	"sync"
)

type program struct {
	log service.Logger
	cfg *service.Config
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {

	router := gin.Default()

	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname")

		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
	wg.Done()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	svcConfig := &service.Config{
		Name:        "fuck",          //在windows下可用net start
		DisplayName: "fuck test",     //显示的名称
		Description: "这是一个service程序", //详情
	}
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			x := s.Install()
			if x != nil {
				fmt.Println("error:", x.Error())
				return
			}
			fmt.Println("服务安装成功")
			return
		} else if os.Args[1] == "uninstall" {
			x := s.Uninstall()
			if x != nil {
				fmt.Println("error:", x.Error())
				return
			}
			fmt.Println("服务卸载成功")
			return
		}
	}
	err = s.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	wg.Wait()
}

func GetFreePorts(count int) ([]int, error) {
	var ports []int
	for i := 0; i < count; i++ {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			return nil, err
		}
		fmt.Println(addr)

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			return nil, err
		}
		defer l.Close()
		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
	}
	return ports, nil
}
