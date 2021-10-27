# FAQ

1. 编译问题，带有动态外部库，所有会造成本地编译，环境上无法运行。还有，编译的二进制文件带有本地路径。

   [CGO_ENABLED环境变量对Go静态编译机制的影响]: https://johng.cn/cgo-enabled-affect-go-static-compile/
   
   ```bash
   # 推荐使用下面命令解决
   # CGO_ENABLED 是否启用CGO
   # -trimpath 移除编译结果的系统路径
   # -ldflags "-s -w" -w 表示关闭DWARF的调试信息，-s 表示strip -s关闭符号链接表
   export CGO_ENABLED=0 go build -trimpath -ldflags "-s -w"
   ```
   
   ```Go
   // hello.go
   package main
   
   import (
       "github.com/gin-gonic/gin"
       "net/http"
   )
   
   func main() {
       router := gin.Default()
   
       router.GET("/", func(c *gin.Context) {
           c.JSON(http.StatusOK, "Hello, Welcome Gin World!")
       })
       err := router.Run(":8080")
       if err != nil {
           panic(err)
       }
   }
   ```
   
   ```bash
   root@ubuntu:~/GolandProjects/GoExample/daily/example/hello# go build
   root@ubuntu:~/GolandProjects/GoExample/daily/example/hello# ldd hello 
           linux-vdso.so.1 (0x00007ffca27d1000)
           libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007fd777430000)
           libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fd77723e000)
           /lib64/ld-linux-x86-64.so.2 (0x00007fd777467000)
   
   root@ubuntu:~/GolandProjects/GoExample/daily/example/hello# strings hello |grep GoExample
   /root/GolandProjects/GoExample/daily/example/hello/main.go
   ```
   
   ```bash
   1. 编译时候，设置CGO_ENABLED=0，解决动态链接库问题
   root@ubuntu:~/GolandProjects/GoExample/daily/example/hello# export CGO_ENABLED=0
   root@ubuntu:~/GolandProjects/GoExample/daily/example/hello# go build
   root@ubuntu:~/GolandProjects/GoExample/daily/example/hello# ldd hello 
   	not a dynamic executable
   root@ubuntu:~/GolandProjects/GoExample/daily/example/hello# 
   
   2. 编译时候，增加-trimpath参数，解决二进制文件带有本地路径问题
   root@ubuntu:~/GolandProjects/GoExample/daily/example/hello# go build -trimpath
   root@ubuntu:~/GolandProjects/GoExample/daily/example/hello# strings hello |grep GoExample
   ```

