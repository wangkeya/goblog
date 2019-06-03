# goblog

### 技术栈

* govendor
* gin 
* gorm 
* go-ini
* grpc
* redis

### 安装
```
    go get github.com/kardianos/govendor
    cd project
    govendor init
    govendor fetch github.com/gin-gonic/gin
    #读取配置
    govendor fetch github.com/go-ini/ini
    #数据库
    govendor fetch github.com/jinzhu/gorm
    govendor fetch github.com/jinzhu/gorm/dialects/mysql
    
    #grpc
    govendor fetch google.golang.org/grpc
    govendor fetch golang.org/x/net
    govendor fetch golang.org/x/text
    govendor fetch google.golang.org/genproto
    
    #redis
    govendor fetch github.com/go-redis/redis
    
```

### 启动
```
    cd project
    go run main.go
    
```

### 接口示例

* http://localhost:8080/web/user/getUsers?tagId=1
* http://localhost:8080/web/post/getInfo/1
* http://localhost:8080/web/post/getList?page=1&pageRows=1 分页实现
* http://localhost:8080/web/post/delete/1
* http://localhost:8080/web/tag/getTagById?tagId=1

### 并发测试

``` 
    ab -n 100 -c 100  http://localhost:8080/web/getUsers?tagId=1
```

### git标签
``` 
    #推送
    git tag v1.0.0
    git push origin v1.0.0
    #拉取
    git pull origin v1.0.0
```



### 计划

* react前端
* docker部署
* jwt
* utc时间格式化
* redis


### md写法

### 删除线

~~删除线~~

#### 表格

| name | age |
| :----: | :----:|
| w | w |

#### 超链接

[百度]:http://www.baidu.com 

[百度]


#### 图片

![image](https://github.com/wangkeya/goblog/blob/master/go.png)

<img src="https://github.com/wangkeya/goblog/blob/master/go.png" width = "100" height = "150" div align="left" />



