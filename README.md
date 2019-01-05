# tools
## 基于beego开发的一些小工具 
 ```
 1.clone 项目 --git clone https://github.com/wangle201210/tools.git
 2.数据库文件在sql内,可以直接使用
 3.本地运行项目(默认位8080端口) --bee run 
 4.访问localhost:80080
 ```

## 线上运行
```
1.打包linux版本项目 --CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
2.给可执行文件权限  --chmod +x [可执行文件名字]
3.运行 --nohup ./[可执行文件名字]
4.退出 
  先使用--ps -aux | grep [可执行文件名字] 获取PID
  然后 kill [PID]
  ```
  
## 完成
  ```
  已完成通过身份证号查看个人信息 2019/01/05
  已完成个人所得税计算 2019/01/06
  ```
