# reportingtool-go
reportingtool go语言版 貌似做得有点多啊 前端js，html，css后端 golang，sql ，xml，json 白菜价！
## 使用须知
因项目使用 mysql 数据库 需增加依赖  go-sql-driver/mysql
在终端 运行命令go get -u github.com/go-sql-driver/mysql


增加第三方路由器  依赖
在终端 运行命令go get -u github.com/gorilla/mux   //以前的路由功能太少，不用了


第三方XML解析  go package
go get github.com/beevik/etree


修改服务器请求地址 在assets/js/core/ReportingTool.js的16行
var serverURL="http://localhost:8080/ReportingTool"


## Note
1.由于涉及前后端交互，如果不涉及账号密码，会造成后端参数解析问题等等
思前向后  决定加入用户登录功能   以求go后端框架  完整  不需要再次改动
并决定用cookie保存信息，实现自动登录

涉及用户名密码  数据库结构问题  建表语句等会之后在代码中贴出

登录页面在https://hongxufeng.github.io/reportingtool-go/web/pages_login.html

2.客户端提交 md5(password) 密码（此方法只是简单保护了密码，是可能被查表获取密码的）。服务端数据库通过 md5(salt+md5(password)) 的规则存储密码，该 salt 仅存储在服务端，且在每次存储密码时都随机生成。这样即使被拖库，制作字典的成本也非常高。密码被 md5() 提交到服务端之后，可通过 md5(salt + form['password']) 与数据库密码比对。此方法可以在避免明文存储密码的前提下，实现密码加密提交与验证。这里还有防止 replay 攻击（请求被重新发出一次即可能通过验证）的问题，由服务端颁发并验证一个带有时间戳的可信 token （或一次性的）即可。

3.准备写个彩蛋   用户auth  密码W  可永久登录系统 
