数据库配置

```
// 路径 config/config.toml
[database]
    [[database.default]]
    link  = "mysql:YOUR_ACCOUNT:YOUR_PASSWORD@tcp(127.0.0.1:3306)/DATAVASE_NAME"
```



docker 运行 bridgeSafe(误差监控)

跨链桥误差监控逻辑：
每分钟自动查询所有跨链对的资产情况，
主链资产查询质押数额，侧链资产查询totalSupply，并作记录。
后台管理系统可以查看当前误差数额，
并且可以设置关停跨链对的阈值，即误差绝对值超过该值则立即停止该跨链对跨链。
但是关停期间任然会记录所有的跨出质押交易。
一旦将跨链对恢复跨链，之前记录的所有质押会补签跨链，
所以需要手动删除异常期间的异常质押记录后，再恢复跨链对。

```
docker run --name bridge-safe -itd -v $PWD:/code -w /code --restart=always golang:latest go run bridgeSafe.go
```