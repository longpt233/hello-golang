# cmd

```
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc
```
- submodule

```
git submodule add git@github.com:longpt233/proto-go-grpc.git ./external/proto
hình như là phải push được submodule lên ->  module chính pull về ->  module chính push lên thì mới nhận
```

```

protoc external/proto/user.proto --go_out=.
protoc external/proto/user.proto --go-grpc_out=.
```

old
``` 
go get -u github.com/golang/protobuf/protoc-gen-go    # sudo snap install protobuf --classic
go get -u google.golang.org/grpc    

ok 
long@hello:~/Documents/20221/datn/auth-service$ go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
long@hello:~/Documents/20221/datn/auth-service$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
long@hello:~/Documents/20221/datn/auth-service$ export PATH="$PATH:$(go env GOPATH)/bin"
```