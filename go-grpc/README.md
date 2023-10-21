# cmd

```
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc



git submodule add git@github.com:longpt233/proto-go-grpc.git ./external/proto
```

old
```
go mod init auth-service
go get -u github.com/golang/protobuf/protoc-gen-go    # sudo snap install protobuf --classic
go get -u google.golang.org/grpc    

ok 
long@hello:~/Documents/20221/datn/auth-service$ go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
long@hello:~/Documents/20221/datn/auth-service$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
long@hello:~/Documents/20221/datn/auth-service$ export PATH="$PATH:$(go env GOPATH)/bin"
```