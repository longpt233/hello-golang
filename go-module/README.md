1. cd vào thư mục
```
go mod init github.com/longpt233/hello-module
```
- khi đó sẽ tạo ra file go.mod có nội dung như sau 

```
module github.com/longpt233/hello-module

go 1.19
```

2. từ hello mà import greetings 

```
long@hello:~/Documents/vccorp/hello-golang/go-module/hello$ go mod edit -replace github.com/longpt233/greetings=../greetings
long@hello:~/Documents/vccorp/hello-golang/go-module/hello$ go mod tidy
go: found github.com/longpt233/greetings in github.com/longpt233/greetings v0.0.0-00010101000000-000000000000
long@hello:~/Documents/vccorp/hello-golang/go-module/hello$ go run . 
Hi, Gladys. Welcome!
```

3. test 
- The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go).
- You can add the -v flag to get verbose output that lists all of the tests and their results. 

```
long@hello:~/Documents/vccorp/hello-golang/go-module/greetings$ go test
PASS
ok      github.com/longpt233/greetings  0.001s
```

4. install 

```
# sẽ tiến hành tạo bin ở thư mục cài đặt của go
long@hello:~/Documents/vccorp/hello-golang/go-module/hello$ go install

# lấy thư mục cài đặt ở  (go instal directory)
long@hello:~/Documents/vccorp/hello-golang/go-module/hello$ go list -f '{{.Target}}'
/home/long/go/bin/hello

# liệt kê các bin đã cài
long@hello:~/Documents/vccorp/hello-golang/go-module/hello$ ll /home/long/go/bin
total 68432
drwxrwxr-x 2 long long     4096 Thg 10 25 16:16 ./
drwxrwxr-x 4 long long     4096 Thg 6   6  2022 ../
-rwxrwxr-x 1 long long  9330058 Thg 6   6  2022 godef*
-rwxrwxr-x 1 long long  4640145 Thg 6   6  2022 go-outline*
-rwxrwxr-x 1 long long 28005660 Thg 10  3 11:15 gopls*
-rwxrwxr-x 1 long long  1930960 Thg 10 25 16:16 hello*
-rwxrwxr-x 1 long long  8766440 Thg 10 12 12:29 mockgen*
-rwxrwxr-x 1 long long  8780391 Thg 10 21 10:42 protoc-gen-go*
-rwxrwxr-x 1 long long  8593957 Thg 10 21 10:42 protoc-gen-go-grpc*

# những thư mục này là bin chỉ cần ./ là chạy

# sửa đổi thư mục instal 
go env -w GOBIN=/path/to/your/bin
```


# go path 

- go root là cái path chạy go ấy 
```
long@hello:~/Documents/vccorp/hello-golang/go-test$ which go
/usr/local/go/bin/go
```

- go path là nơi mấy cái package tải về

```
long@hello:~/Documents/vccorp/hello-golang/go-module$ go env
...
GOOS="linux"
GOPATH="/home/long/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
...
```

- có thể để prj ở đâu cx được khi get github.com/longpt233/gopkg nó đều dùng được (không nhất thiết phải để proj trong go path)
- nếu nó lỗi k nhận thì tốt nhất k mod init lại 
