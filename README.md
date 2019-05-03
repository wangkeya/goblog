# goblog

### 技术栈

* govendor v1.0.9
* gin v1.3
* gorm 
* go-ini

### 安装
```
    go get github.com/kardianos/govendor
    cd project/src
    govendor init
    govendor fetch github.com/gin-gonic/gin@v1.3
    govendor fetch github.com/golang/protobuf/proto
    govendor fetch github.com/ugorji/go/codec
    govendor fetch gopkg.in/go-playground/validator.v8
    govendor fetch gopkg.in/yaml.v2
    #读取配置
    govendor fetch github.com/go-ini/ini
    #数据库
    govendor fetch github.com/jinzhu/gorm
    govendor fetch github.com/jinzhu/gorm/dialects/mysql
```

### 启动
```
    cd project/goblog
    go run src/goblog/main.go
    
```

### 访问

* http://localhost:8080/web/getUsers?tagId=1



### 系统特点


