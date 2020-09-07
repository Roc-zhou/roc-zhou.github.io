## 1、Mac下编译Linux, Windows平台的64位可执行程序：

```
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

## 2、Linux下编译Mac, Windows平台的64位可执行程序：

```
$ CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

## 3、Windows下编译Mac, Linux平台的64位可执行程序：

```
$ SET CGO_ENABLED=0 SET GOOS=darwin SET GOARCH=amd64 go build main.go
$ SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build main.go
```

> 如果编译web等工程项目，直接cd到工程目录下直接执行以上命令
> GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows
> GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm