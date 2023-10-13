- đổi số lượng gomaxprocs 

```
runtime.GOMAXPROCS(2)
count := runtime.GOMAXPROCS(0)   // nếu set mỗi như này sẽ không print đúng
fmt.Println(count)

or 

GOMAXPROCS=2 go run main.go
```