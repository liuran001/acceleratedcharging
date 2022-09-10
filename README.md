# acceleratedcharging
MIUI全机型充电加速(888系列可用)

充电时自动修改温控(实时生效)和强行修改充电电流实现满速充电(跑满120w)
温度墙45度
要设置温度墙请修改模块目录下的wall.ini文件
格式:
45度=450

#go编译命令:
```bash
cd go
go build -ldflags="-s -w" -o charge-current run.go
```
