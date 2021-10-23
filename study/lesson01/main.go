package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func printAddress() error {
	person := Person{Name: "eric", Age: 18}
	p1 := &person

	fmt.Println(person)
	fmt.Println(&person)
	fmt.Println(p1)
	fmt.Println(&p1)

	fmt.Println(person.Name)
	fmt.Println(&person.Name)
	fmt.Println(p1.Age)
	fmt.Println(&p1.Age)

	p2 := Person{"mary", 18}
	fmt.Println(p2)

	p2Bytes, _ := json.Marshal(p2)
	fmt.Println(string(p2Bytes))
	return nil
}

func printMap() error {
	m := make(map[string]int)
	m["float"] = 1
	m["int"] = 2
	m["bool"] = 3
	m["byte"] = 4

	for k, v := range m {
		fmt.Println(k, v)
	}
	return nil
}

func printNewAddress() error {
	p := new(Person)
	person := Person{"Panda", 28}
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", person)
	return nil
}

func printGoroutine() error {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A ", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return nil
}

func randomHose() (chain string, err error) {
	var hosts = make([]string, 0)
	hosts = append(hosts, "192.168.1.100:8443", "192.168.2.100:8443", "192.168.3.100:8443",
		"192.168.4.100:8443", "192.168.5.100:8443", "192.168.6.100:8443", "192.168.7.100:8443",
		"192.168.8.100:8443", "192.168.9.100:8443", "192.168.10.100:8443", "192.168.11.100:8443",
		"192.168.12.100:8443", "192.168.13.100:8443", "192.168.14.100:8443", "192.168.15.100:8443",
		"192.168.16.100:8443", "192.168.17.100:8443", "192.168.18.100:8443", "192.168.19.100:8443",
		"192.168.20.100:8443", "192.168.21.100:8443", "192.168.22.100:8443", "192.168.23.100:8443",
		)

	//message := fmt.Sprint("Gonna work from home...", pick a random reason )
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	message := fmt.Sprint("Gonna work from home...", hosts[rand.Intn(len(hosts))])
	fmt.Println(message)
	return chain, err
}

func main() {
	iniContent, err := Asset("conf_tls/conf.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err := ini.Load(iniContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.Section("node").Key("nodes").SetValue("192.168.1.66:8443")

	b := new(bytes.Buffer)
	w := bufio.NewWriter(b)
	_, err = cfg.WriteTo(w)
	err = w.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b.String())

	err = printAddress()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = printMap()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = printNewAddress()
	if err != nil {
		fmt.Println(err)
		return
	}

	chain, err := randomHose()
	if err != nil {
		return
	}
	fmt.Println(chain)
}
