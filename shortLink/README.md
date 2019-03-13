# shortLink

## use
1. 本地hosts文件添加映射：127.0.0.1 www.s.cn

2. 启动服务：
```bash
    go build
    ./shortLink
```

3. http请求获取短链接服务：127.0.0.1:8081/?url=https://www.baidu.com

4. http请求获取的短链接：www.s.cn:8082/A