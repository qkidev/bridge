## 安装
1.配置env

复制 .env.example ,修改正确

2.npm安装

```
docker run -it --rm -v "$PWD":/usr/src/app -w /usr/src/app  node:14-alpine npm install
```

3.导入数据库

4.运行
```
docker-composer up
```