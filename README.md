# reportingtool-go
go语言版reportingtool
## 使用须知
因项目使用 mysql 数据库 需增加依赖  go-sql-driver/mysql
在终端 运行命令go get -u github.com/go-sql-driver/mysql


增加第三方路由器  依赖
在终端 运行命令go get -u github.com/naoina/denco


第三方XML解析  go package
go get github.com/beevik/etree


修改服务器请求地址 在assets/js/core/ReportingTool.js的16行
var serverURL="http://localhost:8080/ReportingTool"