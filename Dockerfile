############################################
# version : wangkeya/golang:v1
# desc : 当前版本安装的golang
############################################

FROM centos:7

MAINTAINER wangkeya <783242528@qq.com>

#golang下载地址 https://studygolang.com/dl
ENV GOLANG_VERSION 1.12.5
ENV GOLANG_DOWNLOAD_URL https://studygolang.com/dl/golang/go1.12.5.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 aea86e3c73495f205929cfebba0d63f1382c8ac59be081b6351681415f4063cf

# install gcc vim wget curl git
# -y means saying yes to all questions
RUN yum install -y gcc vim wget curl git

# 下载并安装golang
RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

# config GOPATH
ENV GOPATH /goprojects

# config GOROOT
ENV GOROOT /usr/local/go
ENV PATH=$PATH:/usr/local/go/bin

# config GOPATH
RUN mkdir -p /goprojects
RUN mkdir -p /goprojects/src
RUN mkdir -p /goprojects/pkg
RUN mkdir -p /goprojects/bin

#创建项目目录
RUN mkdir -p /goprojects/src/goblog

# VOLUME 选项是将本地的目录挂载到容器中　
# 当你运行-v　＜hostdir>:<Containerdir>时要确保目录内容相同否则会出现数据丢失
# 对应关系如下
# /Users/alan/goblog/src/goblog:/goprojects/src/goblog
# 这里挂载的路径是goblog项目的目录
VOLUME ["/goprojects/src/goblog"]

# build the server
WORKDIR /goprojects/src/goblog
# RUN ls /goprojects/src/goblog
# RUN go build -o server.bin main.go
#RUN ls /goprojects

# startup the server
# 执行go_docker.sh脚本，该脚本在goblog项目根目录下，用于打包编译启动golang项目
# 挂载的go_docker.sh必须有可执行权限
# 需要执行chmod +x /Users/alan/goblog/src/goblog/go_docker.sh
CMD ["sh","/goprojects/src/goblog/go_docker.sh"]
#ENTRYPOINT ['/goprojects/src/goblog/go_docker.sh']

