# mac 安装protobuf

> brew install protobuf


go get github.com/golang/protobuf/protoc-gen-go
cd github.com/golang/protobuf/protoc-gen-go
go build
go install
vi /etc/profile 将$GOPATH/bin 加入环境变量
source profile

go get github.com/golang/protobuf/proto
cd github.com/golang/protobuf/proto
go build
go install


## 编译proto

protoc --go_out=plugins=grpc:. *.proto