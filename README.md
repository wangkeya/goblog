# goblog

### 技术栈

* gin 
* gorm 
* go-ini
* redis
* mysql
* jwt

### 安装
```
   go mod init 
   go mod tidy
   go mod vendor
   
```

### 启动
```
    cd project
    go run main.go
    
```

### 接口示例

* http://localhost:5200/web/user/getUsers?tagId=1
* http://localhost:5200/web/post/getInfo/1
* http://localhost:5200/web/post/getList?page=1&pageRows=1 分页实现
* http://localhost:5200/web/post/delete/1
* http://localhost:5200/web/tag/getTagById?tagId=1

### 并发测试

``` 
    ab -n 100 -c 100  http://localhost:5200/web/getUsers?tagId=1
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
* utc时间格式化



