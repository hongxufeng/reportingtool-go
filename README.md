# reportingtool-go
reportingtool go语言版 貌似做得有点多啊 前端js，html，css后端 golang，sql ，xml，json 白菜价！
## 使用须知
因项目使用 mysql 数据库 需增加依赖  go-sql-driver/mysql
在终端 运行命令go get -u github.com/go-sql-driver/mysql


增加第三方路由器  依赖
在终端 运行命令go get -u github.com/naoina/denco


第三方XML解析  go package
go get github.com/beevik/etree


修改服务器请求地址 在assets/js/core/ReportingTool.js的16行
var serverURL="http://localhost:8080/ReportingTool"


## Note
由于涉及前后端交互，如果不涉及账号密码，会造成后端参数解析问题等等
思前向后  决定加入用户登录功能   以求go后端框架  完整  不需要再次改动
并决定用cookie保存信息，实现自动登录

涉及用户名密码  数据库结构问题  建表语句等会之后在代码中贴出

登录页面在https://hongxufeng.github.io/reportingtool-go/pages_login.html
