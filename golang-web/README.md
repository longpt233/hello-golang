## step by step

``` bash
go mod init <name module>
```

> sau buoc nay ta se co file go.mod

``` bash
go mod tidy
```

> install, up/downgrade 

``` bash
go run main.go
```

## project structure

```
- internal --> folder where we store application code that should not be reused by other applications
   |- model
      |- a.go
   |- driver
   |- services
- .gitignore
- go.mod
- go.sum 
- main.go
- README.md

```