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

3. GitHub相关问题事项

   ```bash
   # github拉去代码换成token，拉取路径如下：
   # 使用方法：
   # 1)从Settings页面 Personal access tokens 生成唯一的Token
   # 2 获取token 注意生成了就要保存 有效期过期或者忘记了只能重新生成了
   # 2) 手动拼接出远程仓库的地址，比如：https://$token@github.com/owner/repo.git
   # 3）从以上地址克隆或使用git remote add 的方式关联本地仓库，之后都不需要输入用户名和密码信息。
   # 实例：git remote set-url origin https://$token@github.com/work/base.git
   那么这些字母分别代表什么意思呢？
   # 在看 github 上的 PR 回复的时候，往往会出现类似于 LGTM、WIP等一类缩写，作为一个无知少年，还真是一脸蒙蔽。那么这些字母分别代表什么意思呢？通常，我们在 github 上最为常见的是以下这些词：
   # PR：Pull Request，如果给其它项目提交合并代码的请求时，就说会提交一个PR。
   # WIP：Work In Progress，如果你要做一个很大的改动，可以在完成部分的情况下先提交，但说明WIP，方便项目维护人员知道你还在 Work，同时他们可以先审核已经完成的。
   # PTAL：Please Take A Look，请求项目维护人员进行 code review。
   # TBR：To Be Reviewed，提示这些代码要进行审核。
   # TL;DR：Too Long; Didn't Read，太长了，懒得看。
   # LGTM：Looks Good To Me，通常是 code review 的时候回复的，即审核通过的意思。
   # SGTM：Sounds Good To Me，跟 LGTM 同义。
   # AFAIK：As Far As I Know，据我所知。
   # CC：Carbon Copy，抄送。
   ```

   ```bash
   # 列出标签
   git tag
   git tag -l
   git tag --list
   
   # 打附注标签
   git tag -a v0.1.0 -m "init project"
   
   # 打轻量标签
   git tag v1.4-lw
   
   # 查看标签
   git show v0.1.0
   
   # 后期打标签
   git log --pretty=oneline
   git tag -a v1.2 9fceb02
   
   # 单个标签推送到远端
   git push origin <tagname>
   
   # 全部标签推送到远端
   git push origin --tags
   
   # 删除标签
   git tag -d v1.4-lw
   git push origin --delete v1.4-lw

   # 导出最新代码
   git archive -o latest.zip HEAD
   ```

   ```bash
   # 如果我们git clone的下载代码的时候是连接的https://而不是git@git (ssh)的形式，当我们操作git pull/push到远程的时候，总是提示我们输入账号和密码才能操作成功，频繁的输入账号和密码会很麻烦。解决办法：
   # git bash进入你的项目目录，输入：
   git config --global credential.helper store

   # 然后你会在你本地生成一个文本，上边记录你的账号和密码。当然这些你可以不用关心。
   ```
   
4. https://ipwhoisinfo.com/ IP信息特别准确

5. linux ulimit 调优问题

   inux系统默认open files数目为1024, 有时应用程序会报Too many open files的错误，是因为进程打开了太多的文件，导致open files数目不够。这就需要修改参数进行调优。特别是有大量文件访问的应用，如elasticsearch、filebeat等, 更要注意这个问题。
   网上的很多教程，都只是简单的说明如何设置，但很多东西没讲明白，而且有时候设置后却不能生效，关于其验证也比较模糊。这里对使用的一些经验进行整理。

   关于进程打开文件描述数量限制，有三个相关参数
   file-max：系统所有进程一共可以打开的文件数量 ；**这项参数是系统级别的。**系统不一样，默认值不同。
   nr_open：单个进程可分配的最大文件数；**这项参数也是系统级别的。**默认：1048576
   limits：当前shell以及由它启动的进程的资源限制；**这项参数是用户所在shell或其所启动的进程级别的。**默认：1024
   **大概结论，file-max是内核可分配的最大文件数，nr_open是单个进程可分配的最大文件数，在在配置ulimit时，如果要超过1048576，需要先增大nr_open；并且根据需求调整file-max 。**

   ```bash
   root@ubuntu18:~# sysctl  -a | grep 'fs.file-max'
   fs.file-max = 52706963
   root@ubuntu18:~# ulimit -a | grep "open files"
   open files            (-n) 1024
   root@ubuntu18:~# sysctl -a | grep 'fs.nr_open'
   fs.nr_open = 1048576
   ```

   ##### **修改file-max 和 nr_open 参数配置**

   ```bash
   # 临时修改，重启失效
   echo  167772166 > /proc/sys/fs/file-max  或者
   sysctl -w "fs.file-max=167772166"
   
   echo  167772166 > /proc/sys/fs/nr_open或者
   sysctl -w "fs.nr_open=167772166"
   
   # 永久生效
   echo "fs.file-max = 167772166 " >> /etc/sysctl.conf
   sysctl -p # 立即生效
   
   echo "fs.nr_open = 167772166 " >> /etc/sysctl.conf
   sysctl -p # 立即生效
   ```

   **修改ulimit参数配置**

   ```bash
   # 临时修改，重启失效
   ulimit -n 204800
   
   # 永久生效，root 用户和其它（*表示其它）一起生效，单用可能不生效
   * soft nofile 204800
   * hard nofile 204800
   * soft nproc 204800
   * hard nproc 204800
   root soft nofile 204800
   root hard nofile 204800
   root soft nproc 204800
   root hard nproc 204800
   
   # 查看结果，-n是可以打开最大文件描述符的数量。 -u是用户最大可用的进程数。
   ulimit -n
   ulimit -u
   
   # 不生效原因及解决方案，
   1. 检查/etc/pam.d/login、/etc/pam.d/su、/etc/pam.d/sshd必须存在session required pam_limits.so
   2. 检查 /etc/ssh/ssd_config 中，必须存在UsePAM yes
   3. limits.conf 文件中， root和*一起使用，单独root或单独*，可能都不生效
   ```

6. 
