# cmd
- submodule

```
git submodule add git@github.com:longpt233/proto-go-grpc.git ./external/proto
hình như là phải push được submodule lên ->  module chính pull về ->  module chính push lên thì mới nhận
```

- cài proto compiler

```
apt install -y protobuf-compiler
```

- cài protoc gen go

```
long@hello:~/Documents/20221/datn/auth-service$ go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
long@hello:~/Documents/20221/datn/auth-service$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
long@hello:~/Documents/20221/datn/auth-service$ export PATH="$PATH:$(go env GOPATH)/bin"
```

# gen file

- gen file encode/decode message: (the code necessary to encode/decode messages)
```
protoc external/proto/invoicer.proto --go_out=.

```

- gen grpc server (the code necessary to run a gRPC server/client)
```
protoc external/proto/invoicer.proto --go-grpc_out=.
```

# tạo server grpc 

- go get -u google.golang.org/grpc
- 