

## 生成镜像
> docker build -t goblog:v1 .


# 运行示例
> docker run -it -v /Users/alan/goblog/src/goblog:/goprojects/src/goblog -p 8081:8081 goblog:v1

# 访问网站
> http://192.168.99.100:8081/

# 删除无用的镜像 --force 强制删除
> docker rmi  `docker images | grep none | awk '{print $3}'` --force

# 计划

* 增加mysql
* 增加redis
* 编写docker compose

## 参考地址
* https://www.cnblogs.com/foxy/p/9274329.html
* https://birdben.github.io/2017/02/06/Docker/Docker%E5%AE%9E%E6%88%98%EF%BC%88%E4%BA%8C%E5%8D%81%E5%9B%9B%EF%BC%89Docker%E5%AE%89%E8%A3%85Golang%E7%8E%AF%E5%A2%83/
