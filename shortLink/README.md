# shortLink

## Description
go实现短链接服务。

## Requirement
1. go
2. mysql

## Reference
1. Mysql表link
```sql
    CREATE TABLE `link` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
      `url` varchar(200) NOT NULL DEFAULT '' COMMENT '链接',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='链接表'
```

## Usage
1. 本地hosts文件添加映射：127.0.0.1 www.s.cn
2. 启动服务：
```bash
    go build
    ./shortLink
```
3. http请求获取短链接服务：127.0.0.1:8081/?url=https://www.baidu.com
4. http请求获取的短链接：www.s.cn:8082/A