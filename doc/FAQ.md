# FAQ

1. 编译问题，带有动态外部库，所有会造成本地编译，环境上无法运行。还有，编译的二进制文件带有本地路径。

   [CGO_ENABLED环境变量对Go静态编译机制的影响]: https://johng.cn/cgo-enabled-affect-go-static-compile/

   ```bash
   # 推荐使用下面命令解决
   # CGO_ENABLED 是否启用CGO
   # -trimpath 移除编译结果的系统路径
   # -ldflags "-s -w" -w 表示关闭DWARF的调试信息，-s 表示strip -s关闭符号链接表
   # -w 去掉调试信息
   # -s 去掉符号表
   # -X 注入变量, 编译时赋值
   
   export CGO_ENABLED=0 go build -trimpath -ldflags "-s -w"
   
   go tool link 查看符号表
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

2. sql语句中and级别优先于or，sqlalchemy中有两种关联表之间的关系。

   ```python
   #方法一、在modes上用ForeignKey进行外键管理
   class DpiUserLogin(db.Model):
       __tablename__ = 'dpiuserlogins'
   
       dpiUserLoginId = db.Column(db.Integer, primary_key=True, autoincrement=True)
       operationLogId = db.Column(db.Integer, db.ForeignKey('operationlogs.operationLogId'))
       
   DpiUserCmdLog.query.join(
       OperationLog
   ).order_by(
       DpiUserCmdLog.timestamp.asc()
   ).all()
   
   # 方法二、
   exp = or_(
           (and_(Ip2Mac.mac == TopDevice.mac, Ip2Mac.ip == TopDevice.ip)),
           (and_(TopDevice.mac == '', Ip2Mac.ip == TopDevice.ip))
       )
   db.session.query(Ip2Mac, TopDevice).outerjoin(TopDevice, exp).filter(Ip2Mac.id > 1).all()
   ```

3. github拉去代码换成token，拉取路径如下：

   ```
   使用方法：
   1)从Settings页面 Personal access tokens 生成唯一的Token
   2 获取token 注意生成了就要保存 有效期过期或者忘记了只能重新生成了
   2) 手动拼接出远程仓库的地址，比如：https://$token@github.com/owner/repo.git
   3）从以上地址克隆或使用git remote add 的方式关联本地仓库，之后都不需要输入用户名和密码信息。
   实例：git remote set-url origin https://$token@github.com/work/base.git
   ```

   
